// Copyright © 2018 Drew J. Sonne <drew.sonne@gmail.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/getlunaform/lunaform/client/stacks"
)

var tfStackGetIdFlag string

// tfStackGetCmd represents the tfStackGet command
var tfStackGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a stack by its id",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		stack, err := gocdClient.Stacks.GetStack(
			stacks.NewGetStackParams().WithID(tfStackGetIdFlag),
			authHandler,
		)
		if err == nil {
			handleOutput(cmd, stack.Payload, useHal, err)
		} else if err1, ok := err.(*stacks.GetStackNotFound); ok {
			handleOutput(cmd, err1.Payload, useHal, nil)
		} else if err1, ok := err.(*stacks.GetStackInternalServerError); ok {
			handleOutput(cmd, err1.Payload, useHal, nil)
		} else {
			handleOutput(cmd, nil, useHal, err)
		}

	},
}

func init() {
	tfStackGetCmd.Flags().StringVar(&tfStackGetIdFlag, "id", "",
		"ID of the terraform module in lunaform")
	tfStackGetCmd.MarkFlagRequired("id")
}
