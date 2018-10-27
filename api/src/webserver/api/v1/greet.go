package v1

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	api "webserver/api"
	config "webserver/config"
)

// Request send from the ui
type GreetingRequest struct {
	Name string `json:"name"`
}

type GreetingMetadata struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// Response send by the api
type GreetingResponse struct {
	Message  string           `json:"message"`
	Metadata GreetingMetadata `json:"__metadata"`
}

// Response send by the metadatastore
type GreetingMetadataResponse struct {
	Count int `json:"count"`
}

func HandleGreet(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data GreetingRequest
	err := decoder.Decode(&data)
	if err != nil {
		api.BadRequest(err, w)
		return
	}

	count, err := fetchMetadataCount(data.Name)
	if err != nil {
		api.InternalServerError(err, w)
		return
	}

	responseJson, err := json.Marshal(buildResponse(data.Name, count))
	if err != nil {
		api.InternalServerError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func fetchMetadataCount(name string) (int, error) {
	url := fmt.Sprintf("%s/%s/%s", config.METADATASTORE_URL, "metadata/names", name)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return -1, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return -1, err
	}

	decoder := json.NewDecoder(resp.Body)
	var data GreetingMetadataResponse
	errDecode := decoder.Decode(&data)
	if errDecode != nil {
		return -1, errDecode
	}

	return data.Count, nil
}

var messages = []string{"Hello %s!", "Greeting %s.", "Happy to see you %s :)"}

func buildResponse(name string, count int) *GreetingResponse {
	return &GreetingResponse{
		Message: fmt.Sprintf(messages[rand.Intn(len(messages))], name),
		Metadata: GreetingMetadata{
			Name:  name,
			Count: count,
		},
	}
}
