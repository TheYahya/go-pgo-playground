package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type Data struct {
	Quiz struct {
		Sport struct {
			Q1 struct {
				Question string   `json:"question"`
				Options  []string `json:"options"`
				Answer   string   `json:"answer"`
			} `json:"q1"`
		} `json:"sport"`
		Maths struct {
			Q1 struct {
				Question string   `json:"question"`
				Options  []string `json:"options"`
				Answer   string   `json:"answer"`
			} `json:"q1"`
			Q2 struct {
				Question string   `json:"question"`
				Options  []string `json:"options"`
				Answer   string   `json:"answer"`
			} `json:"q2"`
		} `json:"maths"`
	} `json:"quiz"`
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	file, _ := os.ReadFile("./data.json")
	data := Data{}
	json.Unmarshal(file, &data)
	d, _ := json.Marshal(data)
	w.Write(d)
}
