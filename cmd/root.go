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
	"github.com/Unkn0wnCat/matrix-veles/internal/tracer"
	"github.com/spf13/viper"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "matrix-veles",
	Short: "A anti-spam bot for your Matrix chatroom",
	//Long: ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")

	viper.SetDefault("bot.homeserver", "matrix.org")
	viper.SetDefault("bot.homeserver_url", "https://matrix.org")
	viper.SetDefault("bot.username", "")
	viper.SetDefault("bot.password", "")
	viper.SetDefault("bot.accessKey", "")
	viper.SetDefault("bot.mongo.uri", "mongodb://localhost:27017")
	viper.SetDefault("bot.mongo.database", "veles")
	viper.SetDefault("bot.mongo.collection.entries", "entries")
	viper.SetDefault("bot.mongo.collection.lists", "lists")
	viper.SetDefault("bot.mongo.collection.rooms", "rooms")
	viper.SetDefault("bot.mongo.collection.users", "users")
	viper.SetDefault("bot.web.listen", "127.0.0.1:8123")
	viper.SetDefault("bot.web.secret", "hunter2")

	viper.SetDefault("tracing.enable", false)
	viper.SetDefault("tracing.jaeger.endpoint", "http://localhost:14268/api/traces")

	cobra.OnInitialize(loadConfig)
	cobra.OnInitialize(func() {
		if viper.GetBool("tracing.enable") {
			tracer.SetupJaeger()
		}
		if !viper.GetBool("tracing.enable") {
			tracer.SetupDummy()
		}
	})
}

func loadConfig() {
	viper.SetConfigType("yaml")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
