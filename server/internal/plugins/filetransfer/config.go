package filetransfer

import (
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Enabled         bool
	RootDir         string
	RefreshInterval time.Duration
}

func (Config) Init(cmd *cobra.Command) error {
	cmd.PersistentFlags().Bool("filetransfer.enabled", false, "whether file transfer is enabled")
	if err := viper.BindPFlag("filetransfer.enabled", cmd.PersistentFlags().Lookup("filetransfer.enabled")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("filetransfer.dir", "/home/neko/Downloads", "root directory for file transfer")
	if err := viper.BindPFlag("filetransfer.dir", cmd.PersistentFlags().Lookup("filetransfer.dir")); err != nil {
		return err
	}

	cmd.PersistentFlags().Duration("filetransfer.refresh_interval", 30*time.Second, "interval to refresh file list")
	if err := viper.BindPFlag("filetransfer.refresh_interval", cmd.PersistentFlags().Lookup("filetransfer.refresh_interval")); err != nil {
		return err
	}

	// v2 config

	cmd.PersistentFlags().Bool("file_transfer_enabled", false, "enable file transfer feature")
	if err := viper.BindPFlag("file_transfer_enabled", cmd.PersistentFlags().Lookup("file_transfer_enabled")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("file_transfer_path", "", "path to use for file transfer")
	if err := viper.BindPFlag("file_transfer_path", cmd.PersistentFlags().Lookup("file_transfer_path")); err != nil {
		return err
	}

	return nil
}

func (s *Config) Set() {
	s.Enabled = viper.GetBool("filetransfer.enabled")
	rootDir := viper.GetString("filetransfer.dir")
	s.RootDir = filepath.Clean(rootDir)
	s.RefreshInterval = viper.GetDuration("filetransfer.refresh_interval")

	// v2 config

	if viper.IsSet("file_transfer_enabled") {
		s.Enabled = viper.GetBool("file_transfer_enabled")
	}
	if viper.IsSet("file_transfer_path") {
		rootDir = viper.GetString("file_transfer_path")
		s.RootDir = filepath.Clean(rootDir)
	}
}
