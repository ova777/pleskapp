// Copyright 1999-2023. Plesk International GmbH.

package cmd

import (
	"github.com/plesk/pleskapp/plesk/internal/locales"
	"github.com/spf13/cobra"
)

var DatabasesCmd = &cobra.Command{Use: "databases", Short: locales.L.Get("database.description")}
