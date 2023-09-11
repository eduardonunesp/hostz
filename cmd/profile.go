package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"
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
		if err := profileGenerator.CreateProfileFromHostList(args[0], hostsParser.ParseHosts(bs)); err != nil {
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

func isRoot() bool {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf("[isRoot] Unable to get current user: %s", err)
	}
	return currentUser.Username == "root"
}

var useFromProfileCmd = &cobra.Command{
	Use:   "use <profile>",
	Short: "Configures the hostsfile with the selected profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Usage()
			return fmt.Errorf("profile name needed")
		}

		profileParser := parser.NewProfileParser()
		profileNames, err := profileParser.GetProfileNames()

		if err != nil {
			return fmt.Errorf("failed to get profile names %s", err)
		}

		for _, name := range profileNames {
			if name == args[0] {
				if !isRoot() {
					return fmt.Errorf("you need to be root to use this command")
				}

				hostsGenerator := generator.NewHostsGenerator()
				output, err := hostsGenerator.BuildHostsFromProfileName(args[0])

				if err != nil {
					return fmt.Errorf("Failed to get profile data %s", err)
				}

				err = os.WriteFile("/etc/hosts", []byte(output), 0644)

				if err != nil {
					return fmt.Errorf("Failed to write to hostsfile %s", err)
				}

				return nil
			}
		}

		fmt.Printf("Profile %s not find\n", args[0])
		return nil
	},
}

var printProfileCmd = &cobra.Command{
	Use:   "print <profile>",
	Short: "Print to terminal the profile selected",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			cmd.Usage()
			return fmt.Errorf("profile name needed")
		}

		profileParser := parser.NewProfileParser()
		profileNames, err := profileParser.GetProfileNames()

		if err != nil {
			return fmt.Errorf("failed to get profile names %s", err)
		}

		for _, name := range profileNames {
			if name == args[0] {
				hostsGenerator := generator.NewHostsGenerator()
				output, err := hostsGenerator.BuildHostsFromProfileName(args[0])

				if err != nil {
					return fmt.Errorf("Failed to get profile data %s", err)
				}

				fmt.Println(output)

				return nil
			}
		}

		return nil
	},
}

var currentProfileCmd = &cobra.Command{
	Use:   "current",
	Short: "Get the current profile used",
	Run: func(cmd *cobra.Command, args []string) {
		hostParser := parser.NewHostsParser()
		bs, err := hostParser.ReadHostsFile("/etc/hosts")

		if err != nil {
			fmt.Printf("Failed to read hosts file %+v\n", err)
			return
		}

		result := hostParser.ParseProfile(bs)

		fmt.Printf("Profiles available: %s\n", result)
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

	profileCmd.AddCommand(useFromProfileCmd)
	useFromProfileCmd.SetUsageTemplate("use <profile name>")

	profileCmd.AddCommand(currentProfileCmd)
	currentProfileCmd.SetUsageTemplate("current")

	profileCmd.AddCommand(printProfileCmd)
	currentProfileCmd.SetUsageTemplate("print <profile name>")
}
