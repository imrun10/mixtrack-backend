package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mixtrack",
	Short: "mixtrack is a CLI tool for managing your music projects",
	Long:  `mixtrack is a CLI tool for managing your music projects.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var childCmd = &cobra.Command{
	Use:   "start",
	Short: "begin a project",
	Run: func(cmd *cobra.Command, args []string) {
		err := core.Initialise(cmd, args)
		if err != nil {
			cmd.PrintErrf("Error: %v\n", err)
			os.Exit(1)
		}
		cmd.Println("Project initialised successfully!")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mixTrack.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	childCmd.Flags().StringP("name", "n", "", "Name of the project")
	rootCmd.AddCommand(childCmd)
}
