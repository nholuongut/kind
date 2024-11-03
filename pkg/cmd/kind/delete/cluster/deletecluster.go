/*
Copyright 2018 The Nho Luong.

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

// Package cluster implements the `delete` command
package cluster

import (
	"github.com/spf13/cobra"

	"https://github.com/nholuongut/kind/pkg/cluster"
	"https://github.com/nholuongut/kind/pkg/cmd"
	"https://github.com/nholuongut/kind/pkg/errors"
	"https://github.com/nholuongut/kind/pkg/log"

	"https://github.com/nholuongut/kind/pkg/internal/cli"
	"https://github.com/nholuongut/kind/pkg/internal/runtime"
)

type flagpole struct {
	Name       string
	Kubeconfig string
}

// NewCommand returns a new cobra.Command for cluster deletion
func NewCommand(logger log.Logger, streams cmd.IOStreams) *cobra.Command {
	flags := &flagpole{}
	cmd := &cobra.Command{
		Args: cobra.NoArgs,
		// TODO(bentheelder): more detailed usage
		Use:   "cluster",
		Short: "Deletes a cluster",
		Long:  "Deletes a resource",
		RunE: func(cmd *cobra.Command, args []string) error {
			cli.OverrideDefaultName(cmd.Flags())
			return deleteCluster(logger, flags)
		},
	}
	cmd.Flags().StringVarP(
		&flags.Name,
		"name",
		"n",
		cluster.DefaultName,
		"the cluster name",
	)
	cmd.Flags().StringVar(
		&flags.Kubeconfig,
		"kubeconfig",
		"",
		"sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config",
	)
	return cmd
}

func deleteCluster(logger log.Logger, flags *flagpole) error {
	provider := cluster.NewProvider(
		cluster.ProviderWithLogger(logger),
		runtime.GetDefault(logger),
	)
	// Delete individual cluster
	logger.V(0).Infof("Deleting cluster %q ...", flags.Name)
	if err := provider.Delete(flags.Name, flags.Kubeconfig); err != nil {
		return errors.Wrapf(err, "failed to delete cluster %q", flags.Name)
	}
	return nil
}
