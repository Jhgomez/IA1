package api

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
	AddFact(careerId int, Aptitude, Skill, Interest []string) (int, error)
	AddCareer(Faculty, Career string) (int, error)
	DeleteFact(careerId int, Aptitude, Skill, Interest string) (int, error)
	UpdateFact(careerId int, Aptitude, Skill, Interest, PAptitude, PSkill, PInterest []string) (int, error)
    DeleteCareer(careerId int) (int, error)
}

type apiImpl struct {}

type FactDto struct {
    CareerId int	`json:"CareerId"`
	Faculty  string `json:"Faculty"`
	Career   string `json:"Career"`
	Aptitude []string `json:"Aptitude"`
	Skill    []string `json:"Skill"`
	Interest []string `json:"Interest"`
}

type apiErrorMessage struct {
	Error string `json:"error"`
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

    defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
		fmt.Printf("error reading body %v", err)
        return []FactDto{}, err
    }

	// Check the status code
	err = checkResponseCode(response, body)

	if err != nil {
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

func (a apiImpl) AddCareer(Faculty, Career string) (int, error) {
	// response, err := http.Post(fmt.Sprintf("%s/getFatcs", localDbSubstituteServer), "application/json", bytes.NewBuffer(courseJson) )
	dataJson := make(map[string]interface{})

	dataJson["Faculty"] = Faculty
	dataJson["Career"] = Career

	postJson, err := json.Marshal(dataJson)

	response, err := http.Post(fmt.Sprintf("%s/addCareer", localDbSubstituteServer), "application/json", bytes.NewBuffer(postJson))

    if err != nil {
		fmt.Printf("error in post %v", err)
        return -1, err
    }

	defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
		fmt.Printf("error reading body %v", err)
        return -1, err
    }

	// Check the status code
	err = checkResponseCode(response, body)

	if err != nil {
		return -1, err
	}

	var data map[string]interface{}

	err = json.Unmarshal(body, &data)

    if err != nil {
		fmt.Printf("error parsing to object %v\n", err)
		return -1, err
	}

	num, err := strconv.Atoi(fmt.Sprintf("%v",(data["CareerId"])))

	if err != nil {
		fmt.Printf("error parsing careerId %v\n", err)
		return -1, err
	}

	return num, nil
}

func (a apiImpl) AddFact(careerId int, Aptitude, Skill, Interest []string) (int, error) {
	// response, err := http.Post(fmt.Sprintf("%s/getFatcs", localDbSubstituteServer), "application/json", bytes.NewBuffer(courseJson) )
	dataJson := make(map[string]interface{})

	dataJson["CareerId"] = careerId
	dataJson["Aptitude"] = Aptitude
	dataJson["Skill"] = Skill
	dataJson["Interest"] = Interest

	postJson, err := json.Marshal(dataJson)

	response, err := http.Post(fmt.Sprintf("%s/addFact", localDbSubstituteServer), "application/json", bytes.NewBuffer(postJson))

    if err != nil {
		fmt.Printf("error in post %v", err)
        return -1, err
    }

	defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
		fmt.Printf("error reading body %v", err)
        return -1, err
    }

	// Check the status code
	err = checkResponseCode(response, body)

	if err != nil {
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

func (a apiImpl) DeleteFact(careerId int, Aptitude, Skill, Interest string) (int, error) {
	// response, err := http.Post(fmt.Sprintf("%s/getFatcs", localDbSubstituteServer), "application/json", bytes.NewBuffer(courseJson) )
	dataJson := make(map[string]interface{})

	dataJson["CareerId"] = careerId
	dataJson["Aptitude"] = Aptitude
	dataJson["Skill"] = Skill
	dataJson["Interest"] = Interest
	// data["List_Of_SOmething"] = []string{"Math"}  this is how to define a parse an object to a Json without an actual struct

	postJson, err := json.Marshal(dataJson)

	response, err := http.Post(fmt.Sprintf("%s/deleteFact", localDbSubstituteServer), "application/json", bytes.NewBuffer(postJson))

    if err != nil {
		fmt.Printf("error in post %v", err)
        return -1, err
    }

	defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
		fmt.Printf("error reading body %v", err)
        return -1, err
    }

	// Check the status code
	err = checkResponseCode(response, body)

	if err != nil {
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

func (a apiImpl) UpdateFact(careerId int, Aptitude, Skill, Interest, PAptitude, PSkill, PInterest []string) (int, error) {
	// response, err := http.Post(fmt.Sprintf("%s/getFatcs", localDbSubstituteServer), "application/json", bytes.NewBuffer(courseJson) )
	dataJson := make(map[string]interface{})

	dataJson["CareerId"] = careerId
	dataJson["Aptitude"] = Aptitude
	dataJson["Skill"] = Skill
	dataJson["Interest"] = Interest
	dataJson["PAptitude"] = PAptitude
	dataJson["PSkill"] = PSkill
	dataJson["PInterest"] = PInterest

	postJson, err := json.Marshal(dataJson)

	response, err := http.Post(fmt.Sprintf("%s/updateFact", localDbSubstituteServer), "application/json", bytes.NewBuffer(postJson))

    if err != nil {
		fmt.Printf("error in post %v", err)
        return -1, err
    }

	defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
		fmt.Printf("error reading body %v", err)
        return -1, err
    }

	// Check the status code
	err = checkResponseCode(response, body)

	if err != nil {
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

func (a apiImpl) DeleteCareer(careerId int) (int, error) {
	// response, err := http.Post(fmt.Sprintf("%s/getFatcs", localDbSubstituteServer), "application/json", bytes.NewBuffer(courseJson) )
	dataJson := make(map[string]interface{})

	dataJson["CareerId"] = careerId
	// data["List_Of_SOmething"] = []string{"Math"}  this is how to define a parse an object to a Json without an actual struct

	postJson, err := json.Marshal(dataJson)

	response, err := http.Post(fmt.Sprintf("%s/deleteCareer", localDbSubstituteServer), "application/json", bytes.NewBuffer(postJson))

    if err != nil {
		fmt.Printf("error in post %v", err)
        return -1, err
    }

	defer response.Body.Close()

    body, err := io.ReadAll(response.Body)

    if err != nil {
		fmt.Printf("error reading body %v", err)
        return -1, err
    }

	// Check the status code
	err = checkResponseCode(response, body)

	if err != nil {
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

func checkResponseCode(response *http.Response, body []byte) error {
	// Check the status code
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		var errorMsg apiErrorMessage

		err := json.Unmarshal(body, &errorMsg)

		if err != nil {
			fmt.Printf("error parsing to error message %v\n", err)
			return err
		}

		errorString := fmt.Sprintf("Unsuccessful response: error code %d %s %s\n", response.StatusCode, http.StatusText(response.StatusCode), errorMsg.Error)

		return errors.New(errorString)
	} else {
		return nil
	}
}

// func main() {
// 	apix := GetApi()

// 	facts, err := apix.GetFacts()

// 	if err != nil {
// 		fmt.Printf("error en %v", err)
// 	}

//     fmt.Println(facts)




// 	rows, err := apix.UpdateFact("ingenieria", "civil", "gay4", "dibujo", "construccion")

// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

// 	fmt.Println(rows)



// 	facts, err = apix.GetFacts()

// 	if err != nil {
// 		fmt.Printf("error en %v", err)
// 	}

//     fmt.Println(facts)



// 	rows, err = apix.DeleteFact("ingenieria", "civil")

// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

// 	fmt.Println(rows)



// 	facts, err = apix.GetFacts()

// 	if err != nil {
// 		fmt.Printf("error en %v", err)
// 	}

//     fmt.Println(facts)



// 	rows, err = apix.AddFact("ingenieria", "civil", "matematica", "dibujo", "construccion")

// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

// 	fmt.Println(rows)




// 	facts, err = apix.GetFacts()

// 	if err != nil {
// 		fmt.Printf("error en %v", err)
// 	}

//     fmt.Println(facts)
// }


// // Generic map function
// func Map[T any, U any](input []T, transform func(T) U) []U {
// 	output := make([]U, len(input))
// 	for i, v := range input {
// 		output[i] = transform(v)
// 	}
// 	return output
// }