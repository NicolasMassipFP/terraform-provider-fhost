package smc

import (
	neturl "net/url"
	"sync"
)

type LockRequest struct {
	lock sync.RWMutex
}

// LockRequests manages locks for concurrent requests
type LockRequests struct {
	requestLocksByBaseUrl map[string]*LockRequest
}

// Singleton instance of LockRequests
var instance = &LockRequests{
	requestLocksByBaseUrl: make(map[string]*LockRequest),
}

var internalLock sync.Mutex

func GetLockRequests() *LockRequests {
	return instance
}

func (mgr *LockRequests) get(lockId string) *LockRequest {

	internalLock.Lock()
	defer internalLock.Unlock()

	lock := mgr.requestLocksByBaseUrl[lockId]
	if lock == nil {
		lock = &LockRequest{}
		mgr.requestLocksByBaseUrl[lockId] = lock
	}
	return lock
}

func GetLockId(url string) string {
	parsedURL, err := neturl.Parse(url)
	if err != nil {
		// TODO no tracing in this version - to be efficient we need
		// the tf.logs here
		return "DEFAULT_LOCK_FROM_URL"
	}
	return parsedURL.Scheme + "://" + parsedURL.Host
}

func (mgr *LockRequests) Enter(lockId string) {
	mgr.get(lockId).lock.Lock()
}

func (mgr *LockRequests) Leave(lockId string) {
	mgr.get(lockId).lock.Unlock()
}
