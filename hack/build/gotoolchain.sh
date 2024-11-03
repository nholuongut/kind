#!/bin/bash
# Copyright 2020 The Nho Luong.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# script to set GOTOOLCHAIN as needed
# MUST BE RUN FROM THE REPO ROOT DIRECTORY

# read go-version file unless GO_VERSION is set
GO_VERSION="${GO_VERSION:-"$(cat .go-version)"}"

GOTOOLCHAIN="go${GO_VERSION}"
export GOTOOLCHAIN GO_VERSION
