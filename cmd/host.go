package cmd

import (
	"fmt"

	"github.com/eduardonunesp/hostz/internals/generator"
	"github.com/eduardonunesp/hostz/internals/parser"
	"github.com/spf13/cobra"
)

var hostCmd = &cobra.Command{
	Use:   "host",
	Short: "Hosts file commands",
}

var generateFromProfileCmd = &cobra.Command{
	Use:   "generate <profile>",
	Short: "Generate hosts file output from a profile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("Profile name needed\n")
			cmd.Usage()
			return
		}

		profileParser := parser.NewProfileParser()
		profileNames, err := profileParser.GetProfileNames()

		if err != nil {
			fmt.Printf("Failed to get profile names %+v\n", err)
			return
		}

		for _, name := range profileNames {
			if name == args[0] {
				hostsGenerator := generator.NewHostsGenerator()
				output, err := hostsGenerator.BuildHostsFromProfileName(args[0])

				if err != nil {
					fmt.Printf("Failed to get profile data %+v\n", err)
					return
				}

				fmt.Print(output)
				return
			}
		}

		fmt.Printf("Profile %s not find\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(hostCmd)

	hostCmd.AddCommand(generateFromProfileCmd)
	generateFromProfileCmd.SetUsageTemplate("host <profile name>")
}
