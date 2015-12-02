package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Packager struct{}

func (Packager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		goarch string
		goos   string
		pack   string
	)

	hosts := strings.Split(r.URL.Host, ".")
	if len(hosts) != len(strings.Split("darwin.386.binloader.xyz", ".")) {
		goarch = r.FormValue("arch")
		goos = r.FormValue("os")
	} else {
		goarch = hosts[0]
		goos = hosts[1]
	}

	pack = strings.TrimLeft(r.URL.Path, "/")

	var out bytes.Buffer

	// go get
	cmd := exec.Command("go", "get", pack)
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(w, "error `go get`ing:", err, out.String())
		return
	}

	name := fmt.Sprintf("/tmp/binary_%v", time.Now().UnixNano())

	cmd = exec.Command("env", "GOOS="+goos, "GOARCH="+goarch, "go", "build", "-o="+name, pack)

	fmt.Println("args:", cmd.Args)
	fmt.Println("env:", cmd.Env)

	fmt.Printf("compiling with os: %v and arch: %v, to file: %v\n", goos, goarch, name)

	out = bytes.Buffer{}

	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(w, "error `go build`ing:", err, out.String())
		return
	}

	f, err := os.Open(name)
	if err != nil {
		fmt.Fprintln(w, "error opening file:", err)
		return
	}

	io.Copy(w, f)
}
