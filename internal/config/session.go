package config

import (
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Session struct {
	ImplicitHosting bool
	APIToken        string

	CookieName       string
	CookieExpiration time.Time
	CookieSecure     bool
}

func (Session) Init(cmd *cobra.Command) error {
	cmd.PersistentFlags().Bool("session.implicit_hosting", true, "allow implicit control switching")
	if err := viper.BindPFlag("session.implicit_hosting", cmd.PersistentFlags().Lookup("session.implicit_hosting")); err != nil {
		return err
	}

	cmd.PersistentFlags().String("session.api_token", "", "API token for interacting with external services")
	if err := viper.BindPFlag("session.api_token", cmd.PersistentFlags().Lookup("session.api_token")); err != nil {
		return err
	}

	// cookie
	cmd.PersistentFlags().String("session.cookie.name", "NEKO_SESSION", "name of the cookie that holds token")
	if err := viper.BindPFlag("session.cookie.name", cmd.PersistentFlags().Lookup("session.cookie.name")); err != nil {
		return err
	}

	cmd.PersistentFlags().Int("session.cookie.expiration", 365*24, "expiration of the cookie in hours")
	if err := viper.BindPFlag("session.cookie.expiration", cmd.PersistentFlags().Lookup("session.cookie.expiration")); err != nil {
		return err
	}

	cmd.PersistentFlags().Bool("session.cookie.secure", true, "use secure cookies")
	if err := viper.BindPFlag("session.cookie.secure", cmd.PersistentFlags().Lookup("session.cookie.secure")); err != nil {
		return err
	}

	return nil
}

func (s *Session) Set() {
	s.ImplicitHosting = viper.GetBool("session.implicit_hosting")
	s.APIToken = viper.GetString("session.api_token")

	s.CookieName = viper.GetString("session.cookie.name")
	s.CookieExpiration = time.Now().Add(time.Duration(viper.GetInt("session.cookie.expiration")) * time.Hour)
	s.CookieSecure = viper.GetBool("session.cookie.secure")
}
