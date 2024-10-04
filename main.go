package main

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"
)

type SystemStats struct {
	IP     string   `json:"ip"`
	PS     []string `json:"ps"`
	DF     string   `json:"df"`
	Uptime string   `json:"uptime"`
}

func runShell(c string, args ...string) string {
	out, _ := exec.Command(c, args...).Output()
	return string(out)
}

func getSystemStats(w http.ResponseWriter, r *http.Request) {
	stats := SystemStats{
		IP:     strings.Split(runShell("hostname", "-I"), " ")[0],
		PS:     strings.Split(runShell("ps", "-ax"), "\n")[1:],
		DF:     strings.Split(runShell("df", "-h", "/"), "\n")[1],
		Uptime: runShell("uptime", "-p"),
	}
	jsonStats, _ := json.MarshalIndent(stats, "", "    ")
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonStats)
}

func main() {
	http.HandleFunc("/", getSystemStats)
	_ = http.ListenAndServe("0.0.0.0:8200", nil)
}
