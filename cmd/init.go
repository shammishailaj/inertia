// Copyright © 2017 UBC Launch Pad team@ubclaunchpad.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Config represents the current projects configuration.
type Config struct {
	CurrentRemoteName string     `json:"name"`
	CurrentRemoteVPS  *RemoteVPS `json:"remote"`
	DaemonAPIToken    string     `json:"token"`
	io.Writer         `json:"-"`
}

var (
	configFolderName = ".inertia"
	configFileName   = "config.json"
	noInertiaRemote  = "No inertia remote"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize an inertia project in this repository",
	Long: `Initialize an inertia project in this GitHub repository.
There must be a local git repository in order for initialization
to succeed.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := InitializeInertiaProject()
		if err != nil {
			log.Fatal(err)
		}
		println("A .inertia folder has been created to store Inertia configuration.")
		println("It is recommended that you DO NOT commit this folder in source")
		println("control since it will be used to store sensitive information.")
		println("\nYou can now use 'inertia remote add' to connect your remote")
		println("VPS instance.")
	},
}

func init() {
	RootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// InitializeInertiaProject creates the inertia config folder and
// returns an error if we're not in a git project.
func InitializeInertiaProject() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	err = checkForGit(cwd)
	if err != nil {
		return err
	}
	err = createConfigDirectory()
	if err != nil {
		return err
	}

	return nil
}

// createConfigDirectory returns an error if the config directory
// already exists (the project is already initialized).
func createConfigDirectory() error {
	configDirPath, err := getProjectConfigFolderPath()
	if err != nil {
		return err
	}

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	_, dirErr := os.Stat(configDirPath)
	_, fileErr := os.Stat(configFilePath)

	// Check if everything already exists.
	if os.IsExist(dirErr) && os.IsExist(fileErr) {
		return errors.New("inertia already properly configured in this folder")
	}

	// Something doesn't exist.
	if os.IsNotExist(dirErr) {
		os.Mkdir(configDirPath, os.ModePerm)
	}

	// Directory exists. Make sure JSON exists.
	if os.IsNotExist(fileErr) {
		config := Config{
			CurrentRemoteName: noInertiaRemote,
			CurrentRemoteVPS:  &RemoteVPS{},
			DaemonAPIToken:    "",
		}

		path, err := getConfigFilePath()
		if err != nil {
			return err
		}
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		writer, err := os.OpenFile(configFilePath, os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}
		config.Writer = writer
		defer f.Close()
		config.Write()
	}

	return nil
}

// GetProjectConfigFromDisk returns the current project's configuration.
// If an .inertia folder is not found, it returns an error.
func getProjectConfigFromDisk() (*Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	raw, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("config file doesnt exist, try inertia init")
		}
		return nil, err
	}

	var result Config
	err = json.Unmarshal(raw, &result)
	if err != nil {
		return nil, err
	}

	// Add writer to object for writing/testing.
	result.Writer, err = os.OpenFile(configFilePath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &result, err
}

// Write writes configuration to JSON file in .inertia folder.
func (config *Config) Write() (int, error) {
	inertiaJSON, err := json.Marshal(config)
	if err != nil {
		return -1, err
	}

	return config.Writer.Write(inertiaJSON)
}

// getProjectConfigFolderPath gets the absolute location of the project
// configuration folder.
func getProjectConfigFolderPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(cwd, configFolderName), nil
}

// getConfigFilePath returns the absolute path of the config JSON
// file.
func getConfigFilePath() (string, error) {
	path, err := getProjectConfigFolderPath()
	if err != nil {
		return "", err
	}
	return filepath.Join(path, configFileName), nil
}

// GetConfigFile returns a config file descriptor for R/W.
func getConfigFile() (*os.File, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	return os.OpenFile(path, os.O_RDWR, os.ModePerm)
}