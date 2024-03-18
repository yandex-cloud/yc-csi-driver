/*
Copyright 2024 YANDEX LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package inflight

import (
	"sync"
	"time"
)

const (
	DefaultEntryTTL = 10 * time.Minute
)

type InFlightTTL struct {
	lock     *sync.Mutex
	inFlight map[string]time.Time
	now      NowFunc
	ttl      time.Duration
}

func NewWithTTL() *InFlightTTL {
	return &InFlightTTL{
		lock:     new(sync.Mutex),
		inFlight: make(map[string]time.Time),
		now:      time.Now,
		ttl:      DefaultEntryTTL,
	}
}

func (db *InFlightTTL) Get(entry Idempotent) bool {
	hash := entry.String()
	t, ok := db.inFlight[hash]

	if !ok {
		return false
	}

	if db.now().Sub(t) >= db.ttl {
		db.Delete(entry)
		return false
	}

	return true
}

func (db *InFlightTTL) Insert(entry Idempotent) bool {
	db.lock.Lock()
	defer db.lock.Unlock()

	hash := entry.String()

	_, ok := db.inFlight[hash]
	if ok {
		return false
	}

	db.inFlight[hash] = db.now()

	return true
}

func (db *InFlightTTL) Delete(entry Idempotent) {
	db.lock.Lock()
	defer db.lock.Unlock()

	delete(db.inFlight, entry.String())
}
