package geocoding

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/insights/task/database"
)

func GetStateFromGPSCoordinates(lat string, lon string) (map[string]interface{}, error) {
	dbClient := database.ConnectToDb()
	// req, err := http.NewRequest("GET", "https://apis.mapmyindia.com/advancedmaps/v1/35a6af2a96e683ae34f1d1458f37f8e5/rev_geocode?lat="+lat+"&lng="+lon, nil)
	req, err := http.NewRequest("GET", "http://api.positionstack.com/v1/reverse?access_key=6a12a0519fe51115c148199939b82e62&query="+lat+","+lon, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	results := response["data"].([]interface{})
	var jsonDocuments []interface{}
	for _, result := range results {
		result := result.(map[string]interface{})
		state := result["region"].(string)
		cursor := database.FindDocumentFromState(dbClient, state)
		ctx, _ := database.GetDbContext()
		var docs []map[string]interface{}
		if err := cursor.All(ctx, &docs); err != nil {
			panic(err)
		}
		for _, doc := range docs {
			delete(doc, "districts")
			delete(doc, "delta")
			delete(doc, "delta21_14")
			delete(doc, "delta7")
			jsonDocuments = append(jsonDocuments, doc)
		}
		break
	}
	response = make(map[string]interface{})
	response["results"] = jsonDocuments
	return response, nil
}
