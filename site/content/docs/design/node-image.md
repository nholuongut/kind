---
title: "Node Image"
menu:
  main:
    parent: "design"
    identifier: "node-image"
---

> **NOTE**: This may not completely cover the current implementation.

The ["node" image][node image] is a Docker image for running
nested containers, systemd, and Kubernetes components.

This image is built on top of the ["base" image][base image].

Logic for building the ["node" image][node image] can be found in [`pkg/build`][build package],
and it can be built with `kind build node-image` respectively.

## Design

Other than the requirements that this image inherits from the "base" image, which
provides most of the tools statically needed for a Kubernetes deployment
(eg `systemd`), variants of this image have the following properties:

- `/kind/images/` contains various `*.tar` files which are 
[Docker image archives][docker image archives],
these images will be loaded by the cluster tooling prior to running `kubeadm`

- `kubeadm`, `kubectl`, `kubelet` are in the path

- A [systemd service][systemd service] is enabled for `kubelet`, and is
configured to not fail on swap being enabled. (we must do the latter because 
swap is inherited from the host and we don't want to force users to disable swap 
before using `kind`)

- `/kind/version` is a regular text file containing the `gitVersion` of the
installed Kubernetes build

These properties are used by the [cluster][cluster package] tooling to boot
each "node" container with [kubeadm][kubeadm].

[node image]: https://https://github.com/nholuongut/kind/images/node
[base image]: /docs/design/base-image
[build package]: https://https://github.com/nholuongut/kind/pkg/build
[cluster package]: https://https://github.com/nholuongut/kind/pkg/cluster
[docker image archives]: https://docs.docker.com/engine/reference/commandline/save/
[systemd service]: https://www.freedesktop.org/software/systemd/man/systemd.service.html
[kubeadm]: https://kubernetes.io/docs/reference/setup-tools/kubeadm/kubeadm/
