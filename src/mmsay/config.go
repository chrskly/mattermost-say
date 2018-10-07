
package mmsay

import (
    "fmt"
    "os"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

// Contains application configuration read in from config file
type conf struct {
    WebhookUrl string `yaml:webhook_url:`
    User string `yaml:user:`
    DefaultChannel string `yaml:default_channel:`
}

// Search paths at which to look for config file
var config_paths = []string{"/etc/mmsay.conf", "/usr/local/etc/mmsay.conf"}

// Search for the config file and return its path
func findConfigFile() (string, error) {
    for _, p := range config_paths {
        fmt.Println("Looking for config file at", p)
        if _, err := os.Stat(p); !os.IsNotExist(err) {
            // File found
            return p, nil
        }
    }
    fmt.Println("no config file found")
    return "", fmt.Errorf("no config file found")
}

// Load the config file. Return `conf` struct of its contents.
func (c *conf) LoadConfig() (*conf, error) {
    // Does the config file exist?
    config_path, err := findConfigFile()
    if err != nil {
        return nil, err
    }
    yamlFile, err := ioutil.ReadFile(config_path)
    if err != nil {
        return nil, fmt.Errorf("could not read config file %s: %v", config_path, err)
    }
    fmt.Println("yamlFile ", yamlFile)
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        return nil, fmt.Errorf("failed to parse config file %s: %v", config_path, err)
    }
    empty_config := conf{}
    if c == &empty_config {
        return nil, fmt.Errorf("config file was empty")
    }
    return c, nil
}

