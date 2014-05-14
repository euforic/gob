package main

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
)

func Create(config *Config) {
  configFile, _ := json.MarshalIndent(&config, "", "  ")
  err := ioutil.WriteFile("config.json", configFile, 0755)

  if err != nil {
    panic(err)
  }

  fmt.Println("\nProject config.json created.\n")
}