/*
Copyright 2020 The Kubernetes Authors.

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

package main

import (
	"flag"

	"github.com/kubernetes-csi/csi-driver-nfs/pkg/nfs"
	"k8s.io/klog"
)

var (
	endpoint = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	nodeID   = flag.String("nodeid", "", "node id")
	// todo: Uncomment this flag when support for printing version is added
	// version        = flag.Bool("version", false, "Print the version and exit.")
	// todo: Uncomment this flag when support for exposing metrics is added
	// metricsAddress = flag.String("metrics-address", "0.0.0.0:29644", "export the metrics")
	kubeconfig = flag.String("kubeconfig", "", "Absolute path to the kubeconfig file. Required only when running out of cluster.")
)

func main() {
	flag.Parse()
	// todo: Add code for printing version information

	if *nodeID == "" {
		klog.Warning("NodeID is empty")
	}
	handle()
}

func handle() {
	var perm uint32 = 0755
	driver := nfs.NewNFSdriver(*nodeID, *endpoint, &perm)
	driver.Run()
}
