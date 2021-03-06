// Copyright 2017 CoreOS, Inc.
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

package packages

import (
	"github.com/coreos/go-semver/semver"

	"github.com/coreos/mantle/kola/cluster"
	"github.com/coreos/mantle/kola/register"
)

func init() {
	register.Register(&register.Test{
		// MinVersion currently set for ipvsadm
		// TODO(ajeddeloh) allow subtests to specify minversion
		Run:         packageTests,
		ClusterSize: 1,
		MinVersion:  semver.Version{Major: 1520},
		Name:        "packages",
	})
}

func packageTests(c cluster.TestCluster) {
	c.Run("sys-cluster/ipvsadm", ipvsadm)
}
