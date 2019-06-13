// Copyright 2015 Light Code Labs, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package caddyhttp

import (
	// plug in the server
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/httpserver"

	// plug in the standard directives
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/basicauth"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/bind"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/browse"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/errors"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/expvar"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/extensions"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/fastcgi"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/gzip"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/header"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/index"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/internalsrv"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/limits"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/log"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/markdown"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/mime"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/pprof"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/proxy"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/push"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/redirect"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/requestid"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/rewrite"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/root"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/status"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/templates"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/timeouts"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/caddyhttp/websocket"
	_ "github.com/go-stargate/stargate/third_party/forked/caddy/onevent"
)
