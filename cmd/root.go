package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

type Config struct {
	HomeDir struct {
		Symlinks []struct {
			Name        string `json:"name"`
			LinkName    string `json:"link-name"`
			LinkAddress string `json:"link-address"`
		} `json:"symlinks"`
		TrackedDotfiles []string `json:"tracked-dotfiles"`
		Repositories    struct {
			Local  string `json:"local"`
			Remote string `json:"remote"`
		} `json:"repositories"`
	} `json:"homedir"`
	Packages struct {
		Native         string `json:"native"`
		NativeExtended string `json:"native-extended"`
		Foreign        string `json:"foreign"`
	} `json:"packages"`
}

var (
	version string = "v0.1.0"

	AppConfig Config

	Verbose bool
)

var rootCmd = &cobra.Command{
	Use:   "arch-maintenance",
	Short: "Arch-Maintenance is a simple utility for maintaining your Arch Linux system.",
	Long:  "Arch-Maintenance is a simple utility for maintaining your Arch Linux system.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "enable verbose")
}

func initConfig() {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filePath := home + "/.maintenance/config.json"

	configFile, err := os.Open(filePath)
	defer configFile.Close()
	if err != nil {
		fmt.Println("Error loading the config file")
		return
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&AppConfig)

}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
