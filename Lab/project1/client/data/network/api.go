package main

import (
    "encoding/json"
    "net/http"
    "fmt"
    
	"io"	
)

type API interface {
	GetFacts() ([]FactDto, error)
	AddFact(Faculty, Career, Aptitude, Skill, Interest string) (any, error)
	DeleteFact(Faculty, Career string) (any, error)
	UpdateFact(Faculty, Career, Aptitude, Skill, Interest string) (any, error)
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

var rustServerUrl = "http://localhost:8000"

func (a apiImpl) GetFacts() ([]FactDto, error) {
    // response, err := http.Post(fmt.Sprintf("%s/getFatcs", rustServerUrl), "application/json", bytes.NewBuffer(courseJson) )
	response, err := http.Get(fmt.Sprintf("%s/getFacts", rustServerUrl))

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

	var facts []FactDto

	err = json.Unmarshal(body, &facts)

    if err != nil {
		fmt.Printf("error parsing to object %v\n", err)
		return []FactDto{}, err
	}

	return facts, nil

}

func (a apiImpl) AddFact(Faculty, Career, Aptitude, Skill, Interest string) (any, error) {
	return 1, nil
    // var factDto FactDto

    
    // if courseDataError := c.BindJSON(&factDto); courseDataError != nil {
    //     fmt.Println(courseDataError)
    //     c.String(http.StatusBadRequest, courseDataError.Error())
    //     return
    // }

    // conn, connErr := grpc.NewClient(grpcServerUrl, opts...)

    // if connErr != nil {
    //     log.Fatalf("fail to dial: %v", connErr)
    //     c.String(http.StatusBadRequest, connErr.Error())
    //     return
    // }

    // defer conn.Close()
    
    // client := pb.NewCourseClient(conn)

    // log.Println("gRPC client connected to server", grpcServerUrl)

   

    // response, responseErr := client.PostCourse(ctx, &pb.FactDto{ 
    //     Curso: factDto.Curso,
    //     Facultad: factDto.Facultad,
    //     Carrera: factDto.Carrera,
    //     Region: factDto.Region,
    // })

    // if responseErr != nil {
    //     log.Fatalf("gRPC post failed: %v", responseErr)
    //     c.String(http.StatusBadRequest, responseErr.Error())
    //     return
    // }

    // courseJson, courseJsonErr := json.Marshal(factDto)

    // if courseJsonErr != nil {
    //     log.Fatalf("parsing back to Json failed: %v", courseJsonErr)
    //     c.String(http.StatusBadRequest, courseJsonErr.Error())
    //     return
    // }

    // response, err := http.Post(fmt.Sprintf("%s/course", rustServerUrl), "application/json", bytes.NewBuffer(courseJson) )

    // if err != nil {
    //     log.Fatalf("post to Rust server failed: %v", err)
    //     c.String(http.StatusBadRequest, err.Error())
    //     return
    // }

    // defer response.Body.Close()
    // body, err := io.ReadAll(response.Body)

    // if err != nil {
    //     log.Fatalf("Rust response body failure: %v", err)
    //     c.String(http.StatusBadRequest, err.Error())
    //     return
    // }

    // success := fmt.Sprintf("gRPC server response: %s, Rust REST server response: %s", response.Response, string(body))

    // c.String(http.StatusOK, success)

}

func (a apiImpl) DeleteFact(Faculty, Career string) (any, error) {
	return 1, nil
}

func (a apiImpl) UpdateFact(Faculty, Career, Aptitude, Skiclll, Interest string) (any, error) {
	return 1, nil
}

func main() {
	facts, err := GetApi().GetFacts()

	if err != nil {
		fmt.Printf("error en %v", err)
	}

    fmt.Println(facts)
}
