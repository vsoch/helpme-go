/*

  Copyright (c) 2017, Vanessa Sochat
  All rights reserved.
  See LICENSE file in main repository

*/

package main

import "flag"
import "net/http"
import "log"
import "github.com/vsoch/helpme/control"


var assetsPath string

func getConfig() *control.Config {

    config := &control.Config{}
    flag.StringVar(&config.ListenSpec, "listen", "localhost:3000", "HTTP listen spec")
    flag.StringVar(&config.DB.ConnectString, "db-connect", "host=/var/run/postgresql dbname=helpme sslmode=disable", "DB Connect String")
    flag.StringVar(&assetsPath, "assets-path", "assets", "Path to assets directory")
    flag.Parse()
    return config
}

func setupHttpAssets(cfg *control.Config) {
    log.Printf("Assets served from %q.", assetsPath)
    cfg.UI.Assets = http.Dir(assetsPath)
}

func main() {
    cfg := getConfig()

    setupHttpAssets(cfg)

    if err := control.Run(cfg); err != nil {
        log.Printf("Error in main(): %v", err)
    }
}
