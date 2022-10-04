package main

import (
	"encoding/json"
	"fmt"
	"github.com/mdlayher/apcupsd"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.Println("Hello There")
	resp, err := http.Get("http://10.0.0.251:9595")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
	http.HandleFunc("/info", hello)
	http.ListenAndServe(":8090", nil)

}

func hello(w http.ResponseWriter, req *http.Request) {
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

func writeData(d Data, w http.ResponseWriter) {
	log.Print("Request: ")
	log.Println(d)
	out, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	w.Write(out)
}

type Data struct {
	Status       string  `json:"status"`
	LoadPercent  float64 `json:"loadPercent"`
	NominalPower int     `json:"nominalPower"`
	CurrentUsage float64 `json:"currentUsage"`
}
