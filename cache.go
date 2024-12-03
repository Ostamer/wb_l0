package main

import (
	"sync"
)

var (
	cache     = make(map[string]Order)
	cacheLock sync.RWMutex
)

func initCache() {
	loadCacheFromDB()
}
