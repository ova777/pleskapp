// Copyright 1999-2023. Plesk International GmbH.

package cmd

import (
	"fmt"
	"github.com/plesk/pleskapp/plesk/internal/config"
	"github.com/plesk/pleskapp/plesk/internal/locales"
	"github.com/spf13/cobra"
)

// Version information
var (
	Commit    string
	BuildTime string
	Version   string
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: locales.L.Get("version.description"),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Client information")
		fmt.Printf("Version:\t%s\nRevision:\t%s\nBuild time:\t%s\n", Version, Commit, BuildTime)

		defaultServerName, err := config.DefaultServer()
		if err == nil {
			server, _ := config.GetServer(defaultServerName)
			platform := "Linux"
			if server.Info.IsWindows {
				platform = "Windows"
			}

			fmt.Println()
			fmt.Println("Server information")
			fmt.Printf("Host:   \t%s\n", server.Host)
			fmt.Printf("Platform:\t%s\n", platform)
			fmt.Printf("Version:\t%s\n", server.Info.Version)
			fmt.Printf("Revision:\t%s\n", server.Info.Revision)
			fmt.Printf("Build date:\t%s\n", server.Info.BuildDate)
		}

		return nil
	},
}
