/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package e2e_node

import (
	"github.com/golang/glog"
	"os/exec"
)

const (
	busyBoxImage = iota

	hostExecImage
	netExecImage
	nginxImage
	pauseImage

	// Images just used for explicitly testing pulling of images
	pullTestExecHealthz
	pullTestAlpineWithBash
)

var ImageRegistry = map[int]string{
	busyBoxImage:  "gcr.io/google_containers/busybox:1.24",
	hostExecImage: "gcr.io/google_containers/hostexec:1.2",
	netExecImage:  "gcr.io/google_containers/netexec:1.4",
	nginxImage:    "gcr.io/google_containers/nginx:1.7.9",
	pauseImage:    "gcr.io/google_containers/pause-amd64:3.0",
}

// These are used by tests that explicitly test the ability to pull images
var NoPullImagRegistry = map[int]string{
	pullTestAlpineWithBash: "gcr.io/google_containers/alpine-with-bash:1.0",
	pullTestExecHealthz:    "gcr.io/google_containers/exechealthz:1.0",
}

// Pre-fetch all images tests depend on so that we don't fail in an actual test
func PrePullAllImages() error {
	for _, image := range ImageRegistry {
		output, err := exec.Command("docker", "pull", image).CombinedOutput()
		if err != nil {
			glog.Warningf("Could not pre-pull image %s %v output:  %s", image, err, output)
			return err
		}
	}
	return nil
}
