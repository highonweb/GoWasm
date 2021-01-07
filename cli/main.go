package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	init := flag.Bool("init", false, "initialise a go wasm file")
	help := flag.Bool("help", false, "Seeing this help message")
	compile := flag.Bool("compile", false, "Compile")
	serve := flag.Bool("serve", false, "serve the wasm application")
	flag.Parse()
	if *help {
		flag.PrintDefaults()
	}
	if *init {
		err := DownloadFile("https://srv-store3.gofile.io/download/tQjPNe/wasm_exec.js", "wasm_exec.js")
		err = DownloadFile("https://srv-store2.gofile.io/download/eRTuDf/index.html", "index.html")

		if err != nil {
			fmt.Println(err)
		}

	}
	if *compile {
		os.Setenv("GOOS", "js")
		os.Setenv("GOARCH", "wasm")
		_, err := exec.Command("go", "build", "-o", "main.wasm").Output()
		if err != nil {
			fmt.Println(err)
		}
	}
	if *serve {
		open("http://localhost:8080")
		err := http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))

		if err != nil {
			fmt.Println(err)
		}
	}
}
func DownloadFile(url string, filepath string) error {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
func open(url string) error {
	var cmd string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	return exec.Command(cmd, url).Start()
}
