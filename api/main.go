package main

import "encoding/json"
import "fmt"
import "math/rand"
import "time"

type GreetingMetadata struct {
	Name 	string		`json:"name"`
	Count 	int			`json:"count"`
}

type GreetingResponse struct {
	Message 		string				`json:"message"`
	Metadata 	GreetingMetadata		`json:"__metadata"`
}

var messages = []string { "Hello %s!", "Greeting %s.", "Happy to see you %s :)" }

func main() {
	rand.Seed(time.Now().Unix())
	response := &GreetingResponse {
		Message: 	fmt.Sprintf(messages[rand.Intn(len(messages))], "toto"),
		Metadata: 	GreetingMetadata {
			Name:  "toto",
			Count: 2,
		},
	}
	responseJson, _ := json.Marshal(response)
	fmt.Print(string(responseJson))
}