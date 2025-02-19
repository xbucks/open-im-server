// Copyright © 2023 OpenIM. All rights reserved.
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

package main

import (
	"fmt"
	"os"

	"github.com/openimsdk/open-im-server/v3/internal/push"
	"github.com/openimsdk/open-im-server/v3/pkg/common/cmd"
	"github.com/openimsdk/open-im-server/v3/pkg/common/config"
)

func main() {
	pushCmd := cmd.NewRpcCmd(cmd.RpcPushServer)
	pushCmd.AddPortFlag()
	pushCmd.AddPrometheusPortFlag()
	if err := pushCmd.Exec(); err != nil {
		panic(err.Error())
	}
	if err := pushCmd.StartSvr(config.Config.RpcRegisterName.OpenImPushName, push.Start); err != nil {
		fmt.Fprintf(os.Stderr, "\n\nexit -1: \n%+v\n\n", err)
		os.Exit(-1)
	}
}
