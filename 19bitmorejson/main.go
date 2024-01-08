package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type course struct {
	Name     string   `json:"coursename"` // these aliases will replace the actual name in JSON data
	Price    int      `json:"courseprice"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              // this will hide this field from the json data
	Tags     []string `json:"tags,omitempty"` // omitempty doesn't show any nil or null responses
}

func main() {
	fmt.Println("We will learn more about JSON in this video alongside Go")

	// JSONDataFromWeb := []byte(
	// 	`{
	// 			"coursename": "Golang",
	// 			"courseprice": 800,
	// 			"website": "Youtbe",
	// 			"tags": ["golang","backend","cloud"]
	// 		}
	// `)

	// EncodeJSON()
	// DecodeJSONFromWeb(JSONDataFromWeb)
	GetJSONFromURL("https://rating.jenkins.io/rate/result.php")
}

func GetJSONFromURL(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// never forget to close the response body it's your responsibility as a caller
	defer response.Body.Close()

	fmt.Println(response)

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	DummyJSONDataFromWeb := databytes
	fmt.Println(string(DummyJSONDataFromWeb))
}

func EncodeJSON() {
	courses := []course{
		{"Java", 900, "Youtbe", "youtube", []string{"python", "backend"}},
		{"Golang", 800, "Youtbe", "youtube", []string{"golang", "backend", "cloud"}},
		{"Javascript", 100, "Udemy", "udemy", []string{"js", "fullstack"}},
		{"Python", 100, "Youtbe", "youtube", nil},
	}

	// it converts the struct to json data which will be a slice of JSON Objects
	// coursesJSON, err := json.Marshal(courses)

	// MarshalIndent - Beautifies the JSON Data (like JSON data)
	coursesJSON, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", coursesJSON)
}

func DecodeJSONFromWeb(jsonData []byte) {
	// checks if the JSON is valid or Not syntactically
	validDataJSON := json.Valid(jsonData)

	// if the unmarshalledJSON results are stored in anything other than the struct based on which the JSON was created it will return nothing
	var unmarshalledJSON course

	// I tested that by taking the reference of a byte array. an Empty array was printed
	// var unmarshalledJSON []byte

	if validDataJSON {
		// the reference is where the result data will be stored after the JSON gets unmarshalled
		json.Unmarshal(jsonData, &unmarshalledJSON)
		fmt.Printf("%#v\n", unmarshalledJSON)
	} else {
		fmt.Println("JSON Data is invalid")
	}

	// Let's take the case of not knowing how the struct will look if i get some data from the web or a backend I can't access.
	// let's explain this a bit more. we know for a fact that all the keys in a JSON object are strings and values can be anything
	// interface is just like `any` keyword in golang
	var dataFromBackendOrInternet map[string]interface{}

	// unmarshalling the data from internet
	json.Unmarshal(jsonData, &dataFromBackendOrInternet)
	databytes, _ := json.MarshalIndent(dataFromBackendOrInternet, "", "\t")
	fmt.Println(string(databytes))
	// fmt.Printf("%s\n", dataFromBackendOrInternet)
	// for k, v := range dataFromBackendOrInternet {
	// 	fmt.Printf("Key: %v, Value: %v and Value Type: %T\n", k, v, v)
	// }

	// %#v prints how the data was coded
	// fmt.Printf("%#v\n", dataFromBackendOrInternet)

}
