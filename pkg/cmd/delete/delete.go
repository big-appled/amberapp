/*
Copyright 2021.

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

package delete

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/jibudata/app-hook-operator/pkg/client"
	"github.com/jibudata/app-hook-operator/pkg/cmd"
)

type DeleteOptions struct {
	Name string
}

func NewCommand(client *client.Client) *cobra.Command {

	option := &DeleteOptions{}

	c := &cobra.Command{
		Use:   "delete",
		Short: "Delete a Database configuration",
		Long:  "Delete a Database configraution",
		Run: func(c *cobra.Command, args []string) {
			cmd.CheckError(option.Run(client))
		},
	}

	option.BindFlags(c.Flags())

	return c
}

func (c *DeleteOptions) BindFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&c.Name, "name", "n", "", "database configration name")
}

func (c *DeleteOptions) Run(kubeclient *client.Client) error {
	return nil
}
