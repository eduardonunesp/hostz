package cmd

import (
	"fmt"
	"strings"

	"github.com/eduardonunesp/hostz/internals/generator"
	"github.com/eduardonunesp/hostz/internals/parser"
	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Manage profiles",
}

var createProfileCmd = &cobra.Command{
	Use:   "create <name>",
	Short: "Creates a new profile",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("Profile name needed\n")
			cmd.Usage()
			return
		}

		profileGenerator := generator.NewProfileGenerator()
		if err := profileGenerator.CreateProfileFileFromName(args[0]); err != nil {
			fmt.Printf("Failed to generate profile %+v\n", err)
			return
		}
		fmt.Printf("Profile named %s created with sucess", args[0])
	},
}

var copyProfileFromCmd = &cobra.Command{
	Use:   "copy <profile name> <path for the hosts file>",
	Short: "Ceates a profile based the hosts file passed",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("Profile path needed\n")
			cmd.Usage()
			return
		}

		hostsParser := parser.NewHostsParser()
		bs, err := hostsParser.ReadHostsFile(args[1])

		if err != nil {
			fmt.Printf("Failed to read hosts file %+v\n", err)
			return
		}

		profileGenerator := generator.NewProfileGenerator()
		if err := profileGenerator.CreateProfileFromHostMap(args[0], hostsParser.ParseHosts(bs)); err != nil {
			fmt.Printf("Failed to generate profile %+v\n", err)
			return
		}
		fmt.Printf("Profile named %s created with sucess", args[0])
	},
}

var listProfilesCmd = &cobra.Command{
	Use:   "list",
	Short: "List profiles available",
	Run: func(cmd *cobra.Command, args []string) {
		profileParser := parser.NewProfileParser()
		profileNames, err := profileParser.GetProfileNames()

		if err != nil {
			fmt.Printf("Failed to get profile names %+v\n", err)
			return
		}
		fmt.Printf("Profiles available: %s\n", strings.Join(profileNames, ", "))
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)

	profileCmd.AddCommand(createProfileCmd)
	createProfileCmd.SetUsageTemplate("profile create <name>")

	profileCmd.AddCommand(copyProfileFromCmd)
	copyProfileFromCmd.SetUsageTemplate("profile copy <profile name> <path to hosts file>")

	profileCmd.AddCommand(listProfilesCmd)
	listProfilesCmd.SetUsageTemplate("list")
}
