package main

import (
	"net/http"
	"os/exec"
	"go_android_server/monkey"
)

func main() {
	monkey.StartMonkeyManager()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/ls", lsHandler)
	http.HandleFunc("/monkey/start", monkey.StartMonkey)
	http.HandleFunc("/monkey/stop", monkey.StopMonkey)
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
		panic(err)
	} else {
		w.Write(dataOutput)
	}
}
