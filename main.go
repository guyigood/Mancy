package main;

import (
	"github.com/fsnotify/fsnotify"
	"flag"
	"log"
	"time"
)

// Package variables
var (
	// Local dir name witch will be watched
	localDir = flag.String("localDir", "./", "Set the local directory witch will be watched")

	// Local dir name witch will be watched
	remoteDir = flag.String("remoteDir", "./", "Set the remote directory that accept the changes")

	// Global chan variables
	// file_watcher will write the chan and file_handle will read the chan
	// create file
	fileCreateEvent = make(chan string)

	// write
	fileWriteEvent = make(chan string)

	// remove
	fileRemoveEvent = make(chan string)

	// rename
	fileRenameEvent = make(chan string)

	// chmod
	fileChmodEvent = make(chan string)

	// watchMainJob chan
	watcherHandlerDone = make(chan bool)

	// fileHandleMainJob chan
	fileHandlerDone = make(chan bool)

	// timeout for watcher event
	fileHandleTimeOut = time.Second * 4
)

// Init
func init() {
	// Reset log format
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {

	watch, _ := fsnotify.NewWatcher()

	w := Watch{
		watch: watch,
	}

	// Watch the local directory
	go func() {
		w.watchDir(*localDir)
		watcherHandlerDone <- true
	}()

	// handle the file events
	go func() {
		fileHandler()
		fileHandlerDone <- true
	}()

	<-watcherHandlerDone
	<-fileHandlerDone

}