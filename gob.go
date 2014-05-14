package main

import (
  "fmt"
  "os"
)

var Version = "0.2.0"

type Config struct {
  Name    string `json:"name"`
  Version string `json:"version"`
  Org     string `json:"org"`
}


func main() {
  args := os.Args[1:]

  if len(args) > 0 {
    switch args[0] {
    case "init":
      config := Config{}

      fmt.Print("name: ")
      fmt.Scan(&config.Name)

      fmt.Print("version: ")
      fmt.Scan(&config.Version)

      fmt.Print("org: ")
      fmt.Scan(&config.Org)

      Create(&config)
    case "build":
      Build()
    }
  } else {
    fmt.Printf(`
      Gob v%v - Simple go build helper

      Usage: gob <command> [options]

      Commands:

       init     Create project config.json
       build    Fetch dependencies and build project
    `+"\n", Version)
  }
}
