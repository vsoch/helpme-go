/*

  Copyright (c) 2017, Vanessa Sochat
  All rights reserved.
  See LICENSE file in main repository

*/

package control


import "log"
import "net"
import "os"
import "os/signal"
import "syscall"

import "github.com/vsoch/helpme/database"
import "github.com/vsoch/helpme/model"
import "github.com/vsoch/helpme/views"


type Config struct {
    ListenSpec string
    DB database.Config
    UI views.Config
}

func Run(cfg *Config) error {

    log.Printf("Starting, HTTP on: %s\n", cfg.ListenSpec)

    db, err := database.InitDb(cfg.DB)
    if err != nil {
        log.Printf("Error initializing database: %v\n", err)
        return err
    }

    m := model.New(db)

    l, err := net.Listen("tcp", cfg.ListenSpec)
    if err != nil {
        log.Printf("Error creating listener: %v\n", err)
        return err
    }

    views.Start(cfg.UI, m, l)
    waitForSignal()
    return nil
}

func waitForSignal() {
    ch := make(chan os.Signal)
    signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
    s := <-ch
    log.Printf("Got signal: %v, exiting.", s)
}
