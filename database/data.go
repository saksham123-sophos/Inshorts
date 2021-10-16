package database

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func PersistCovidData() error {
	initStateCodeMappings()
	dbClient := ConnectToDb()
	req, err := http.NewRequest("GET", "https://data.covid19india.org/v4/min/data.min.json", nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data map[string]interface{}
	err = json.Unmarshal([]byte(body), &data)
	if err != nil {
		return err
	}
	var docs []interface{}
	var india = getObjectForIndia()
	for key := range data {
		value := data[key].(map[string]interface{})
		value["state"] = getStateName(key)
		docs = append(docs, value)
		var last_updated = (value["meta"].(map[string]interface{})["last_updated"].(string))
		var meta = make(map[string]interface{})
		meta["last_updated"] = last_updated
		india["meta"] = meta
		var total = make(map[string]float64)
		total["recovered"] = (value["total"].(map[string]interface{})["recovered"].(float64)) + (india["total"].(map[string]float64)["recovered"])
		total["tested"] = (value["total"].(map[string]interface{})["tested"].(float64)) + (india["total"].(map[string]float64)["tested"])
		total["vaccinated1"] = (value["total"].(map[string]interface{})["vaccinated1"].(float64)) + (india["total"].(map[string]float64)["vaccinated1"])
		total["vaccinated2"] = (value["total"].(map[string]interface{})["vaccinated2"].(float64)) + (india["total"].(map[string]float64)["vaccinated2"])
		total["confirmed"] = (value["total"].(map[string]interface{})["confirmed"].(float64)) + (india["total"].(map[string]float64)["confirmed"])
		total["deceased"] = (value["total"].(map[string]interface{})["deceased"].(float64)) + (india["total"].(map[string]float64)["deceased"])
		india["total"] = total
	}
	docs = append(docs, india)
	err = InsertToDb(dbClient, docs)
	if err != nil {
		return err
	}
	return nil
}
