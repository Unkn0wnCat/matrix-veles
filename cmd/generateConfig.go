/*
 * Copyright © 2022 Kevin Kandlbinder.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// generateConfigCmd is a command to generate a config file
var generateConfigCmd = &cobra.Command{
	Use:   "generateConfig",
	Short: "Generate a config file",
	Long:  `This command generates a config file with all values populated with either the defaults or the contents of the current config with any unpopulated fields set to default.`,
	Run: func(cmd *cobra.Command, args []string) {
		out, err := cmd.PersistentFlags().GetString("out")
		if err != nil {
			fmt.Println("Could not parse --out flag:")
			fmt.Println(err)
			os.Exit(1)
		}

		force, err := cmd.PersistentFlags().GetBool("force")
		if err != nil {
			fmt.Println("Could not parse --force flag:")
			fmt.Println(err)
			os.Exit(1)
		}

		if force {
			err = viper.WriteConfigAs(out)
		}
		if !force {
			err = viper.SafeWriteConfigAs(out)
		}

		if err != nil {
			fmt.Println("Could not write config:")
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Config written to %s!\n", out)
	},
}

func init() {
	rootCmd.AddCommand(generateConfigCmd)

	generateConfigCmd.PersistentFlags().StringP("out", "o", "./config.yaml", "Output path for the config file")
	generateConfigCmd.PersistentFlags().BoolP("force", "f", false, "Force writing the config file (may overwrite existing files)")
}
