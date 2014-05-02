package main

import (
  "encoding/json"
  "io/ioutil"
  "log"
  "os"
  "os/exec"
  "strings"
)

type Config struct {
  Name    string `json:"name"`
  Version string `json:"version"`
  Org     string `json:"org"`
}

func main() {
  var config Config

  curDir, _ := os.Getwd()
  file, e := ioutil.ReadFile("./config.json")

  if e != nil {
    log.Fatalf("Error: %v\n", e)
  }

  if e := json.Unmarshal(file, &config); e != nil {
    log.Fatalf("Error: %v\n", e)
  }

  repoPath := config.Org + "/" + config.Name

  if e := os.MkdirAll("./gopath/src/"+config.Org, 0755); e != nil {
    if !strings.Contains(e.Error(), "exists") {
      log.Fatalf("Error: %v\n", e)
    }
  }

  if e := os.Symlink("../../../..", "gopath/src/"+repoPath); e != nil {
    if !strings.Contains(e.Error(), "exists") {
      log.Fatalf("Error: %v\n", e)
    }
  }

  os.Setenv("GOBIN", curDir+"/bin")
  os.Setenv("GOPATH", curDir+"/gopath")

  log.Println("Getting Dependencies...")

  depsCmd := exec.Command("go", "get")
  depsCmd.Dir = curDir + "/gopath/src/" + repoPath
  depsCmd.Stdout, depsCmd.Stderr = os.Stdout, os.Stderr
  depsCmd.Run()

  log.Printf("Building %s...", config.Name)

  cmd := exec.Command("go", "install", repoPath)
  cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
  cmd.Run()
}
