/*
Copyright 2024 The Nho Luong.

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

package kube

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"https://github.com/nholuongut/kind/pkg/log"
)

// TODO(bentheelder): plumb through arch

// directoryBuilder implements Bits for a local docker-ized make / bash build
type directoryBuilder struct {
	tarballPath string
	logger      log.Logger
}

var _ Builder = &directoryBuilder{}

// NewTarballBuilder returns a new Bits backed by the docker-ized build,
// given kubeRoot, the path to the kubernetes source directory
func NewTarballBuilder(logger log.Logger, tarballPath string) (Builder, error) {
	return &directoryBuilder{
		tarballPath: tarballPath,
		logger:      logger,
	}, nil
}

// Build implements Bits.Build
func (b *directoryBuilder) Build() (Bits, error) {
	tmpDir, err := os.MkdirTemp(os.TempDir(), "k8s-tar-extract-")
	if err != nil {
		return nil, fmt.Errorf("error creating temporary directory for tar extraction: %w", err)
	}

	b.logger.V(0).Infof("Extracting %q", b.tarballPath)
	err = extractTarball(b.tarballPath, tmpDir, b.logger)
	if err != nil {
		return nil, fmt.Errorf("error extracting tar file: %w", err)
	}

	binDir := filepath.Join(tmpDir, "kubernetes/server/bin")
	contents, err := os.ReadFile(filepath.Join(tmpDir, "kubernetes/version"))
	if err != nil {
		return nil, err
	}
	sourceVersionRaw := strings.TrimSpace(string(contents))
	return &bits{
		binaryPaths: []string{
			filepath.Join(binDir, "kubeadm"),
			filepath.Join(binDir, "kubelet"),
			filepath.Join(binDir, "kubectl"),
		},
		imagePaths: []string{
			filepath.Join(binDir, "kube-apiserver.tar"),
			filepath.Join(binDir, "kube-controller-manager.tar"),
			filepath.Join(binDir, "kube-scheduler.tar"),
			filepath.Join(binDir, "kube-proxy.tar"),
		},
		version: sourceVersionRaw,
	}, nil
}
