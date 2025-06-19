package main

import (
    "encoding/json"
    "net/http"
    "fmt"
	"errors"
	"bytes"
	"strconv"
    
	"io"	
)

type API interface {
	GetFacts() ([]FactDto, error)
	AddFact(Faculty, Career, Aptitude, Skill, Interest string) (int, error)
	DeleteFact(Faculty, Career string) (int, error)
	UpdateFact(Faculty, Career, Aptitude, Skill, Interest string) (int, error)
}

type apiImpl struct {}

type FactDto struct {
	Faculty  string `json:"Faculty"`
	Career   string `json:"Career"`
	Aptitude string `json:"Aptitude"`
	Skill    string `json:"Skill"`
	Interest string `json:"Interest"`
}

var api API

func GetApi() API {
	if api == nil {
		api = apiImpl{}
	}

	return api
}

var localDbSubstituteServer = "http://localhost:8000"

func (a apiImpl) GetFacts() ([]FactDto, error) {
    // response, err := http.Post(fmt.Sprintf("%s/getFatcs", localDbSubstituteServer), "application/json", bytes.NewBuffer(courseJson) )
	response, err := http.Get(fmt.Sprintf("%s/getFacts", localDbSubstituteServer))

    if err != nil {
		fmt.Printf("error en get %v", err)
        return []FactDto{}, err
    }

	// Check the status code
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return []FactDto{}, errors.New(fmt.Sprintf("Unsuccessful response: %d %s\n", response.StatusCode, http.StatusText(response.StatusCode)))
	}

    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
		fmt.Printf("error reading body %v", err)
        return []FactDto{}, err
    }

	var facts []FactDto

	err = json.Unmarshal(body, &facts)

    if err != nil {
		fmt.Printf("error parsing to object %v\n", err)
		return []FactDto{}, err
	}

	return facts, nil

}

func (a apiImpl) AddFact(Faculty, Career, Aptitude, Skill, Interest string) (int, error) {
	// response, err := http.Post(fmt.Sprintf("%s/getFatcs", localDbSubstituteServer), "application/json", bytes.NewBuffer(courseJson) )
	careerFact := FactDto{Faculty, Career, Aptitude, Skill, Interest}

	postJson, err := json.Marshal(careerFact)

	response, err := http.Post(fmt.Sprintf("%s/addFact", localDbSubstituteServer), "application/json", bytes.NewBuffer(postJson))

    if err != nil {
		fmt.Printf("error in post %v", err)
        return -1, err
    }

	// Check the status code
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return -1, errors.New(fmt.Sprintf("Unsuccessful response: %d %s\n", response.StatusCode, http.StatusText(response.StatusCode)))
	}

	defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
		fmt.Printf("error reading body %v", err)
        return -1, err
    }

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)

    if err != nil {
		fmt.Printf("error parsing to object %v\n", err)
		return -1, err
	}

	num, err := strconv.Atoi(fmt.Sprintf("%v",(data["rows"])))

	if err != nil {
		fmt.Printf("error parsing rows %v\n", err)
		return -1, err
	}

	return num, nil
}

func (a apiImpl) DeleteFact(Faculty, Career string) (int, error) {
	return 1, nil
}

func (a apiImpl) UpdateFact(Faculty, Career, Aptitude, Skiclll, Interest string) (int, error) {
	return 1, nil
}

func main() {
	apix := GetApi()

	rows, err := apix.AddFact("ingenieria", "civil", "matematica", "dibujo", "construccion")

	if err != nil {
		fmt.Printf("error en %v", err)
	}

	fmt.Println(rows)

	facts, err := apix.GetFacts()

	if err != nil {
		fmt.Printf("error en %v", err)
	}

    fmt.Println(facts)
}
