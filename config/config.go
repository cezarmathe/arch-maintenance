package config

import (
	"errors"
	"os"

	"github.com/cezarmathe/arch-maintenance/logging"
	"github.com/pelletier/go-toml"
)

var (
	// paths for loading the config
	configPaths = [...]string{"", "/etc/arch-maintenance.toml", "/home/cezar/Projects/Go/src/github.com/cezarmathe/arch-maintenance/config/config.toml"}

	// Config contains the configuration
	Config config

	// The logger
	logger *logging.Logger
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
			countries   []string `toml:"countries"`
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
func LoadConfig() error {
	for i := 0; i < len(configPaths); i++ {
		configFile, err := os.Open(configPaths[i])
		if err == nil {
			tomlDecoder := toml.NewDecoder(configFile)
			tomlDecoder.Decode(&Config)
			break
		} else if i == len(configPaths)-1 {
			return errors.New("The package is corrupt. The config file was not found in /etc")
		} else {

		}
	}
	return nil
}

func SetLogger(newLogger *logging.Logger) {
	logger = newLogger
}
