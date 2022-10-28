//go:build windows

package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// windowsCmd represents the windows command
var windowsCmd = &cobra.Command{
	Use:   "windows",
	Short: "This set up the development environment on the Windows Operating System.",
	Long:  `The windows command will setup the tools necessary to react-native development on the Windows OS.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Avinash - I'm inside the windows function")

		out, err := exec.Command("Set-ExecutionPolicy Unrestricted -Scope Process -Force;iex (New-Object System.Net.WebClient).DownloadString('https://aka.ms/rnw-deps.ps1')").Output()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out))
	},
}

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
