package main

import (
	"net/http"
	"strings"
	"os/exec"
	"os"
)

const (
	defaultLocation = "/Dropbox/1Password.agilekeychain/"
	url = "http://localhost:8080/1Password.html"
	envName = "ONEPASSWORD_DIR"
)
var (
	server = http.NewServeMux()
	dir string
)

func main() {
	if len(os.Getenv(envName)) > 0 {
		dir = os.Getenv(envName)
	}else {
		dir = os.Getenv("HOME")+defaultLocation;
	}
	server.Handle("/", http.FileServer(http.Dir(dir)))
	http.HandleFunc("/", hanlder)
	exec.Command("xdg-open", url).Start()
	http.ListenAndServe(":8080", nil)
}

func hanlder(w http.ResponseWriter, r *http.Request) {
	if (strings.Contains(r.URL.Path, ".")) {
		server.ServeHTTP(w, r)
	}
}
