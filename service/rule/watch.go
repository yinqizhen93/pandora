package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os/exec"
	"plugin"
)

func watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher failed: ", err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("%s %s\n", event.Name, event.Op)
				if event.Name == "data_struct.go" {
					cmd := exec.Command("/bin/sh", "build_go_plugin.sh")
					outPut, _ := cmd.Output()
					fmt.Println(outPut)
				}
				if event.Name == "data_struct.so" {
					reBuildDataStruct()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("data_struct.go")
	err = watcher.Add("data_struct.so")
	if err != nil {
		log.Fatal("Add failed:", err)
	}
	<-done
}

func reBuildDataStruct() {
	p, err := plugin.Open("data_struct.so")
	if err != nil {
		panic(err)
	}
	Plug = p
	mr, err := p.Lookup("MR")
	if err != nil {
		panic(err)
	}
	fmt.Println("newbuild", mr)
}
