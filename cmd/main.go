package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	box "github.com/hansbringert/go-quiz/internal"
	log "github.com/sirupsen/logrus"
)

var version = "development"
var commit = "-"
var date = "-"
var builtBy = "-"

var infoVariables = map[string]string{
	"Version:": version,
	"Commit:":  commit,
	"Date:":    date,
	"BuiltBy:": builtBy,
	"Runtime:": runtime.Version(),
}

func printVersion() {
	for key, value := range infoVariables {
		fmt.Printf("%-10s %s\n", key, value)
	}
}

var rootDir = "./static"
var port = ":1959"

func main() {
	printVersion()
	if len(os.Args) > 1 {
		port = fmt.Sprintf(":%s", os.Args[1])
	}
	fileServer := box.FileServer(box.Dir(rootDir))

	// fs := box.MyFileSystem{
	//	Dir: rootDir,
	// }
	// fileServer := http.FileServer(fs)

	http.Handle("/", fileServer)

	// Start server
	fmt.Printf("started server, visit http://localhost%s for quiz\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
