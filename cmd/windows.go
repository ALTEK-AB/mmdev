//go:build windows

package cmd

import (
	"fmt"
	"github.com/abdfnx/gosh"
	"github.com/spf13/cobra"
	"strings"
)

// windowsCmd represents the windows command
var windowsCmd = &cobra.Command{
	Use:   "windows",
	Short: "This set up the development environment on the Windows Operating System.",
	Long:  `The windows command will setup the tools necessary to react-native development on the Windows OS.`,
	Run: func(cmd *cobra.Command, args []string) {
		// run a command
		fmt.Println(">>>> Updating ExecutionPolicy this Process scope ...")
		gosh.PowershellCommand(`Set-ExecutionPolicy Unrestricted -Scope Process -Force`)

		pErr, pOut, _ := gosh.PowershellOutput(`Get-ExecutionPolicy -Scope Process`)
		if pErr == nil && strings.Contains(pOut, "Unrestricted") {
			fmt.Println(">>>> ExecutionPolicy has been updated to Unrestricted for this Process scope")
		}
	},
}

// Use the -Command parameter to specify the command to execute as a String: powershell -Command "Start-Process powershell -verb RunAs"
func init() {
	rootCmd.AddCommand(windowsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// windowsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// windowsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
