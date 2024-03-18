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
	"time"
)

type NowFunc func() time.Time

type Idempotent interface {
	String() string
}

// InFlight helps to manage in-flight requests.
type InFlight interface {
	// Get returns true when the key already exists.
	Get(entry Idempotent) bool

	// Insert inserts the entry to the current list of in-flight requests.
	// Returns false when the key already exists.
	Insert(entry Idempotent) bool

	// Delete removes the entry from the in-flight entries map.
	// It doesn't return anything, and will do nothing if the specified key doesn't exist.
	Delete(entry Idempotent)
}
