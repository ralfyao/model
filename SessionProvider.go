package model

import (
	"fmt"
	"sync"
)

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}
type SessionProvider interface {
	SessionInit(sessionId string) (Session, error)
	SessionRead(sessionId string) (Session, error)
	SessionDestroy(sessionId string) error
	GarbageCollector(maxLifeTime int64)
}

var providers = make(map[string]SessionProvider)

func RegisterProvider(name string, provider SessionProvider) {
	if provider == nil {
		panic("session:Register Provider is null")
	}
	if _, p := providers[name]; p {
		panic("session: Register Provider is existed")
	}
	providers[name] = provider
}

type SessionManager struct {
	cookieName  string
	lock        sync.Mutex
	provider    SessionProvider
	maxLifeTime int64
}

func NewSessionManager(providerName, cookieName string, maxLifeTime int64) (*SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session:unknown provider %q(forgotten to import?)", providerName)
	}
	return &SessionManager{
		cookieName:  cookieName,
		maxLifeTime: maxLifeTime,
		provider:    provider,
	}, nil
}
