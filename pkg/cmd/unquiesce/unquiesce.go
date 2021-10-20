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

package unquiesce

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/jibudata/app-hook-operator/pkg/client"
	"github.com/jibudata/app-hook-operator/pkg/cmd"
)

type UnquiesceOptions struct {
	Name string
	//Database    string
}

func NewCommand(client *client.Client) *cobra.Command {

	option := &UnquiesceOptions{}

	c := &cobra.Command{
		Use:   "unquiesce",
		Short: "Unquiesce a Database",
		Long:  "Unquiesce a Database which has been quiesced",
		Run: func(c *cobra.Command, args []string) {
			cmd.CheckError(option.Run(client))
		},
	}

	option.BindFlags(c.Flags())

	return c
}

func (c *UnquiesceOptions) BindFlags(flags *pflag.FlagSet) {
	flags.StringVarP(&c.Name, "name", "n", "", "database configration name")
	//flags.StringVarP(&c.Database, "database", "d", "", "name of the database instance")
}

func (c *UnquiesceOptions) Run(kubeclient *client.Client) error {
	return nil
}
