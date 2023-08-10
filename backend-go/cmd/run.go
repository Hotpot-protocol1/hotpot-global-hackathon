/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"github.com/Hotpot-protocol1/hotpot-global/app"
	"github.com/Hotpot-protocol1/hotpot-global/db"
	"github.com/Hotpot-protocol1/hotpot-global/db/migrator"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		connectionString, err := cfg.DB.GetConnectionString()
		if err != nil {
			return err
		}

		db := db.New(connectionString)
		_, err = migrator.MigrateDB("up", db.DB(), migrator.Migrations.Migrate)
		if err != nil {
			return err
		}

		if err := app.Start(cfg, db); err != nil {
			panic(errors.Wrap(err, "failed to start api"))
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
