package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	type stepikData struct {
		GlobalId int64 `json:"global_id"`
	}
	type rawData struct {
		inside []stepikData
	}
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Print("error while opening file...")
		return
	}
	log.Print("Successfully opened your file!")
	defer jsonFile.Close()
	//
	var info rawData
	data, err2 := ioutil.ReadAll(jsonFile)
	if err2 != nil {
		log.Print("error while reading file...")
	}
	log.Print("Successfully read your file!")

	unmarshaledData := json.Unmarshal(data, &info.inside)
	if unmarshaledData != nil {
		log.Print(unmarshaledData)
	}
	log.Print("Successfully decoded your file!")
	var result int64
	for i := 0; i < len(info.inside); i++ {
		result += info.inside[i].GlobalId
	}
	log.Print(result)
}
