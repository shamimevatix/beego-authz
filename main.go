// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/casbin/beego-authz/authz"
	"github.com/casbin/beego-authz/auth"
	"github.com/casbin/casbin"
)

func main() {
	// authenticate every request.
	web.InsertFilter("*", web.BeforeRouter, auth.Basic("alice", "123"))

	// authorize every request.
	web.InsertFilter("*", web.BeforeRouter, authz.NewAuthorizer(casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")))

	//beego.Router("*", &TestController{})
	web.Run()
}
