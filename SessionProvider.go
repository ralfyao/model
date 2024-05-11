<<<<<<< HEAD
package model

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
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
func (manager *SessionManager) GetSessionId() string {
	b := make([]byte, 32)
	if _, error := io.ReadFull(rand.Reader, b); error != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (manager *SessionManager) SessionBegin(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sessionId := manager.GetSessionId()
		session, _ = manager.provider.SessionInit(sessionId)
		cookie := http.Cookie{
			Name:     manager.cookieName,
			Value:    url.QueryEscape(sessionId),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(manager.maxLifeTime),
		}
		http.SetCookie(w, &cookie)
	} else {
		sessionId, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sessionId)
	}
	return session
}

func (manager *SessionManager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.SessionDestroy(cookie.Value)
	expiredTime := time.Now()
	newCookie := http.Cookie{
		Name:     manager.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiredTime,
		MaxAge:   -1,
	}
	http.SetCookie(w, &newCookie)
}

func (manager *SessionManager) GarbageCollector() {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	manager.provider.GarbageCollector(manager.maxLifeTime)
	time.AfterFunc(time.Duration(manager.maxLifeTime), func() {
		manager.GarbageCollector()
	})
}
=======
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
>>>>>>> c20ffbcc6c6e31518839cfbedd2688702db76620
