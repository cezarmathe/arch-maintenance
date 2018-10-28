package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	Native        bool
	NativeExended bool
	Foreign       bool
	PackageList   bool

	packagesCmd = &cobra.Command{
		Use:   "packages",
		Short: "Interact with the system packages feature",
		Run:   PackagesCmd,
	}
)

func init() {
	defer rootCmd.AddCommand(packagesCmd)

	packagesCmd.Flags().BoolVarP(&Native, "native", "n", false, "Operate on the native packages")
	packagesCmd.Flags().BoolVarP(&NativeExended, "native-extended", "", false, "Operate on the native-extended packages")
	packagesCmd.Flags().BoolVarP(&Foreign, "foreign", "f", false, "Operate on the foreign packages")
	packagesCmd.Flags().BoolVarP(&PackageList, "pkg-list", "", false, "Create a package list")
}

func PackagesCmd(cmd *cobra.Command, args []string) {

	if PackageList {
		if Native {
			cmdNative := exec.Command("pacman", "-Qent", ">", AppConfig.Packages.Native)
			err1 := cmdNative.Run()
			if err1 != null {
				os.Stderr.WriteString(err1.Error())
				panic(err1)
			}

			fmt.Println("Created a list of all explicitly installed native packages.")
		}
	}
}
