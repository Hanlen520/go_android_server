package monkey

import (
	"os/exec"
	"fmt"
	"net/http"
)

// 这里将bool改成struct就可以进行高度定制
var monkeyStart = make(chan bool, 1)
var monkeyDone = make(chan bool, 1)

func StartMonkeyManager() {
	go func() {
		monkeyCmd := exec.Command(
			"/system/bin/sh",
			"/system/bin/monkey", "1000")

		for {
			select {
			case <-monkeyStart:
				if monkeyCmd.Process == nil {
					monkeyCmd.Start()
				}
				fmt.Println("monkey is running.")
			case <-monkeyDone:
				if monkeyCmd.Process != nil {
					monkeyCmd.Process.Kill()
				}
				fmt.Println("monkey dead.")
				monkeyCmd.Process = nil
			}
		}
	}()
}

func StartMonkey(w http.ResponseWriter, r *http.Request) {
	monkeyStart <- true
	w.Write([]byte("Start Monkey now! :)"))
}

func StopMonkey(w http.ResponseWriter, r *http.Request) {
	monkeyDone <- true
	w.Write([]byte("Stop Monkey now! :)"))
}
