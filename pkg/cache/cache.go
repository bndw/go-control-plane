// Copyright 2017 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Package cache defines a configuration cache for the server.
package cache

import (
	"github.com/envoyproxy/go-control-plane/api"
	"github.com/golang/protobuf/proto"
)

// ConfigWatcher requests watches for configuration resources by a node, last
// applied version identifier, and resource names hint. The watch should send
// the responses when they are ready. The watch can be cancelled by the
// consumer, in effect terminating the watch for the request.
// ConfigWatcher implementation must be thread-safe.
type ConfigWatcher interface {
	// WatchEndpoints returns a watch for endpoints.
	WatchEndpoints(*api.Node, string, []string) Watch

	// WatchClusters returns a watch for clusters.
	WatchClusters(*api.Node, string, []string) Watch

	// WatchRoutes returns a watch for routes.
	WatchRoutes(*api.Node, string, []string) Watch

	// WatchListeners returns a watch for listeners.
	WatchListeners(*api.Node, string, []string) Watch
}

// Cache is a generic config cache with a watcher.
type Cache interface {
	ConfigWatcher

	// SetResource sets a response for a node group and a type.
	SetResource(Key, ResponseType, []proto.Message)
}

// Watch is a dedicated stream of configuration resources produced by the
// configuration cache and consumed by the xDS server.
type Watch struct {
	// Value channel produces one or more resources, once they are available.
	// Multiple values can be returned ahead of the consumer requesting a newer
	// watch. If the channel is closed prior to cancellation of the watch, an
	// unrecoverable error has occurred in the producer, and the consumer should
	// close the corresponding stream.
	Value chan Response

	stop func()
}

// Cancel cancels the watch. Watch must be cancelled to release resources in the cache.
// Cancel can be called multiple times.
func (watch Watch) Cancel() {
	if watch.stop != nil {
		watch.stop()
		watch.stop = nil
	}
}

// Response is a pre-serialized response for an assumed configuration type.
type Response struct {
	// Version of the resources as tracked by the cache for the given type.
	// Proxy responds with this version as an acknowledgement.
	Version string

	// Resources to be included in the response.
	Resources []proto.Message

	// Canary bit to control Envoy config transition.
	Canary bool
}

// ResponseType is an enumeration of cache response types.
type ResponseType int

// xDS response types
const (
	EndpointResponse ResponseType = iota
	ClusterResponse
	RouteResponse
	ListenerResponse
)

// Key is the node group identifier
type Key string

// NodeGroup aggregates configuration resources by a hash of the node.
type NodeGroup interface {
	// Hash returns a string identifier for the proxy nodes.
	// Must be a thread-safe function.
	Hash(*api.Node) Key
}
