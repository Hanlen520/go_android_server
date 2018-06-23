package main

import (
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ls", lsHandler)
	http.ListenAndServe("0.0.0.0:80", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello From Android :)"))
}

func lsHandler(w http.ResponseWriter, r *http.Request) {
	lsCmd := exec.Command("ls")
	dataOutput, err := lsCmd.Output()
	if err != nil {
		w.Write([]byte("Something error in server :("))
	} else {
		w.Write(dataOutput)
	}
}
