/*
Copyright 2021 The terraform-docs Authors.

Licensed under the MIT license (the "License"); you may not
use this file except in compliance with the License.

You may obtain a copy of the License at the LICENSE file in
the root directory of this source tree.
*/

package tfvars

import (
	"github.com/spf13/cobra"

	"github.com/pulberg/terraform-docs/cmd/tfvars/hcl"
	"github.com/pulberg/terraform-docs/cmd/tfvars/json"
	"github.com/pulberg/terraform-docs/internal/cli"
	"github.com/pulberg/terraform-docs/print"
)

// NewCommand returns a new cobra.Command for 'tfvars' formatter
func NewCommand(runtime *cli.Runtime, config *print.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "tfvars [PATH]",
		Short:       "Generate terraform.tfvars of inputs",
		Annotations: cli.Annotations("tfvars"),
	}

	// subcommands
	cmd.AddCommand(hcl.NewCommand(runtime, config))
	cmd.AddCommand(json.NewCommand(runtime, config))

	return cmd
}
