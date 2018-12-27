package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/pelletier/go-toml"
)

var (
	// paths for loading the config
	configPaths = [...]string{"", "/etc/arch-maintenance.toml", "/home/cezar/Projects/Go/src/github.com/cezarmathe/arch-maintenance/config/config_example.toml"}

	// Config contains the configuration
	Config config

	// The logger
	configLogger *logrus.Entry
)

// Config is a simple container for the needed variables for running the maintenance
type config struct {
	Backup struct {
		Dotfiles struct {
			Path          string `toml:"path"`
			Remote        string `toml:"remote"`
			CommitMessage string `toml:"commit_message"`
		} `toml:"dotfiles"`

		// Etc struct {} `toml:"etc"`
		Packages struct {
			BackupPath    string   `toml:"backup_path"`
			Formats       []string `toml:"formats"`
			Archive       bool     `toml:"archive"`
			Remote        string   `toml:"remote"`
			CommitMessage string   `toml:"commit_message"`
		} `toml:"packages"`
	} `toml:"backup"`

	Update struct {
		Mirrors struct {
			Countries   []string `toml:"countries"`
			MirrorCount int      `toml:"mirror_count"`
			AgeOrCount  bool     `toml:"age_or_count"`
			Protocol    string   `toml:"protocol"`
			Save        bool     `toml:"save"`
		} `toml:"mirrors"`

		Packages struct {
			Native bool `toml:"native"`
			Aur    bool `toml:"aur"`
			Yay    struct {
				ShowDiffs    bool `toml:"show_diffs"`
				CleanBuild   bool `toml:"clean_build"`
				EditPkgbuild bool `toml:"edit_pkgbuild"`
				Upgrade      bool `toml:"upgrade"`
			} `toml:"yay"`
		} `toml:"packages"`
	} `toml:"update"`

	Maintenance struct {
		Systemd struct {
			CheckFailed bool `toml:"check_failed"`
		} `toml:"systemd"`

		Symlinks struct {
			Fix bool `toml:"bool"`
		} `toml:"symlinks"`

		Pacman struct {
			RemoveOrphans bool `toml:"remove_orphans"`
		} `toml:"pacman"`

		Cache struct {
			DeletePacmanCache bool     `toml:"delete_pacman_cache"`
			UserCaches        []string `toml:"user_caches"`
		} `toml:"cache"`
	} `toml:"maintenance"`

	// Report struct {} `toml:"report"`
}

// Load the XDG_CONFIG_DIR path if possible, otherwise use the default ~/.config
func init() {
	xdgConfigDir := os.Getenv("XDG_CONFIG_DIR")
	if xdgConfigDir != "" {
		configPaths[0] = xdgConfigDir + "arch-maintenance.toml"
	} else {
		configPaths[0] = os.Getenv("HOME") + "/.config/arch-maintenance.toml"
	}
}

// LoadConfig loads the configuration from the system and reports an error
// if there is one
func LoadConfig() {
	configLogger.Info("loading the configuration")
	for i := 0; i < len(configPaths); i++ {
		configLogger.Debug("trying to load the configuration from \"" + configPaths[i] +"\"")
		configFile, err := os.Open(configPaths[i])
		if err == nil {
			configLogger.Debug("opened the configuration file from \"" + configPaths[i] + "\"")
			tomlDecoder := toml.NewDecoder(configFile)
			err = tomlDecoder.Decode(&Config)
			if err != nil {
				configLogger.Warn(err)
			}
			configLogger.Info("loaded the configuration from \"" +configPaths[i] + "\"")
			return
		} else if i == len(configPaths)-1 {
			configLogger.Fatal("no configuration file found")
		} else {
			configLogger.Warn(err)
		}
	}
}

func SetLogger(newLogger *logrus.Entry) {
	configLogger = newLogger
}
