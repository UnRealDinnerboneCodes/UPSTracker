package main

import (
	"encoding/json"
	"github.com/mdlayher/apcupsd"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello There Again")
	http.HandleFunc("/unreal", unreal)
	http.HandleFunc("/beef", beef)
	http.ListenAndServe(":8090", nil)

}

func unreal(w http.ResponseWriter, req *http.Request) {
	log.Println("Got Request")
	c, err := apcupsd.Dial("tcp", "10.0.0.251:3551")
	if err != nil {
		writeData(Data{"Unknown", 0.0, 0.0, 0.0}, w)
		return
	}
	status, err := c.Status()
	if err != nil {
		writeData(Data{"Error", 0.0, 0.0, 0.0}, w)
		return
	}
	usesage := (float64(status.NominalPower) * (0.1 * status.LoadPercent)) / 10
	writeData(Data{status.Status, status.LoadPercent, status.NominalPower, usesage}, w)
}

func beef(w http.ResponseWriter, req *http.Request) {
	log.Println("Got Request")
	c, err := apcupsd.Dial("tcp", "10.0.0.224:3551")
	if err != nil {
		writeData(Data{"Unknown", 0.0, 0.0, 0.0}, w)
		return
	}
	status, err := c.Status()
	if err != nil {
		writeData(Data{"Error", 0.0, 0.0, 0.0}, w)
		return
	}
	usesage := (float64(status.NominalPower) * (0.1 * status.LoadPercent)) / 10
	writeData(Data{status.Status, status.LoadPercent, status.NominalPower, usesage}, w)
}

func writeData(d Data, w http.ResponseWriter) {
	log.Print("Request: ")
	log.Println(d)
	out, err := json.Marshal(Return{d})
	if err != nil {
		panic(err)
	}
	w.Write(out)
}

type Return struct {
	Data Data `json:"data"`
}

type Data struct {
	Status       string  `json:"status"`
	LoadPercent  float64 `json:"loadPercent"`
	NominalPower int     `json:"nominalPower"`
	CurrentUsage float64 `json:"currentUsage"`
}
