package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var folder string = os.Getenv("FOLDER")
var port string = ":8080"
var remote string = os.Getenv("REMOTE")
var delay time.Duration

func init() {
	var err error
	delay, err = time.ParseDuration(os.Getenv("DELAY"))
	if err != nil {
		delay = 24 * time.Hour
	}

	flag.StringVar(&folder, "folder", folder, "Folder to store mirror into")
	flag.StringVar(&port, "address", port, "Address to listen on")
	flag.StringVar(&remote, "remote", remote, "Remote to fetch data from")
	flag.DurationVar(&delay, "delay", delay, "Regular interval to run sync at")
	flag.Parse()

	if folder == "" {
		panic("-folder may not be blank")
	}

	if remote == "" {
		panic("-remote may not be blank")
	}
}

func main() {
	go Sync()
	Serve()
}

// Serve serves the http server
func Serve() {
	log.Printf("Accepting connections on port %s", port)
	err := http.ListenAndServe(port, http.FileServer(http.Dir(folder)))
	if err == nil {
		return
	}
	log.Fatal(err)
}

// Sync regularly sync the folder with the provided
func Sync() {
	for {
		RunSyncCommand()
		log.Printf("Waiting %s for next sync invocation", delay)
		time.Sleep(delay)
	}
}

// RunSyncCommand runs a single command to sync
func RunSyncCommand() {
	args := []string{"rsync", "-a", "--delete", remote, folder}
	log.Printf("Running sync command: %v", args)

	command := exec.Command(args[0], args[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		log.Printf("Sync command exited with error: %s", err)
	} else {
		log.Printf("Sync command finished")
	}
}
