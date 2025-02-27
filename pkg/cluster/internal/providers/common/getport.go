/*
Copyright 2019 The Nho Luong.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or impliep.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

import (
	"net"
)

// PortOrGetFreePort is a helper that either returns the provided port
// if valid or returns a new free port on listenAddr and a cleanup function
func PortOrGetFreePort(port int32, listenAddr string) (int32, func(), error) {
	// in the case of -1 we actually want to pass 0 to the backend to let it pick
	if port == -1 {
		return 0, nil, nil
	}
	// in the case of 0 (unset) we want kind to pick one and supply it to the backend
	if port == 0 {
		return GetFreePort(listenAddr)
	}
	// otherwise keep the port
	return port, nil, nil
}

// GetFreePort is a helper used to get a free TCP port on the host
// returns the free port and a cleanup function, the cleanup function must be called
// after all free ports have been determined to ensure the same port is not returned
// multiple times
func GetFreePort(listenAddr string) (int32, func(), error) {
	dummyListener, err := net.Listen("tcp", net.JoinHostPort(listenAddr, "0"))
	if err != nil {
		return 0, nil, err
	}
	port := dummyListener.Addr().(*net.TCPAddr).Port
	return int32(port), func() { dummyListener.Close() }, nil
}
