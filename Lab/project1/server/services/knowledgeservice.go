package knowledgeservice

import (
    "net/http"
    "fmt"

    "github.com/gin-gonic/gin"

	"unimatchserver/data/repository/knowledge"
)

type KnowledgeService interface {
    GetFacts(c *gin.Context) // gets all careers
    AddFact(c *gin.Context)   // "careerFact", add facts(aptitude, skill, interest), should be used in combination with "AddCareer" if adding a new career
    AddCareer(c *gin.Context) // Adds a career without any aptitude, skill nor interest
    DeleteFact(c *gin.Context) // "deletFact" Deletes a single aptitude and/or skill and/or interest
    UpdateFact(c *gin.Context)   // "updateCareerFact" updates a career's aptitude, skill and interest but it needs a list of new aptitude, skill nor interest and the old aptitude, skill nor interest
    DeleteCareer(c *gin.Context) // Deletes a career along with all aptitude, skill nor interest
}

type knowledgeServiceImpl struct {
    repo knowledgerepo.KnowledgeRepo
}

var knowledgeService KnowledgeService

func GetKnowledgeService() KnowledgeService {
    if knowledgeService == nil {
        knowledgeService = knowledgeServiceImpl{ knowledgerepo.GetKnowledgeRepo() }
    }

    return knowledgeService
}

type careerFact struct {
    CareerId int    `json:"CareerId"`
	Aptitude []string `json:"Aptitude"`
	Skill    []string `json:"Skill"`
	Interest []string `json:"Interest"`
}

type deletFact struct {
    CareerId int    `json:"CareerId"`
	Aptitude string `json:"Aptitude"`
	Skill    string `json:"Skill"`
	Interest string `json:"Interest"`
}

type updateCareerFact struct {
    CareerId  int `json:"CareerId"`
	Aptitude []string `json:"Aptitude"`
	Skill    []string `json:"Skill"`
	Interest []string `json:"Interest"`
    PAptitude []string `json:"PAptitude"`
    PSkill      []string `json:"PSkill"`
    PInterest   []string `json:"PInterest"`
}

type deleteCareer struct {
    CareerId  int `json:"CareerId"`
}

type careerFaculty struct {
    Faculty  string `json:"Faculty"`
    Career  string `json:"Career"`
}

func (s knowledgeServiceImpl) GetFacts(c *gin.Context) {
    facts, err := s.repo.GetFacts()

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
        return
    }

    fmt.Println(facts)
    
    c.JSON(http.StatusOK, facts)
}

func (s knowledgeServiceImpl) AddFact(c *gin.Context) {
    var career careerFact

    if careerDataError := c.BindJSON(&career); careerDataError != nil {
        fmt.Println(careerDataError)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", careerDataError.Error())})
        return
    }

    rows, err := s.repo.AddFact(career.CareerId, career.Aptitude, career.Skill, career.Interest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
        return
	}

    c.JSON(http.StatusOK, gin.H{"rows": rows })

    // c.String(http.StatusOK, fmt.Sprintf("fact added, rows inserted: %d", rows))
}


func (s knowledgeServiceImpl) AddCareer(c *gin.Context) {
    var career careerFaculty

    if careerDataError := c.BindJSON(&career); careerDataError != nil {
        fmt.Println(careerDataError)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", careerDataError.Error())})
        return
    }

    rows, err := s.repo.AddCareer(career.Faculty, career.Career)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
        return
	}

    c.JSON(http.StatusOK, gin.H{"CareerId": rows })

    // c.String(http.StatusOK, fmt.Sprintf("fact added, rows inserted: %d", rows))
}

func (s knowledgeServiceImpl) DeleteFact(c *gin.Context) {
    var career deletFact

    if careerDataError := c.BindJSON(&career); careerDataError != nil {
        fmt.Println(careerDataError)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", careerDataError.Error())})
        return
    }

    rows, err := s.repo.DeleteFact(career.CareerId, career.Aptitude, career.Skill, career.Interest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
        return
	}

    // c.String(http.StatusOK, fmt.Sprintf("fact deleted, rows deleted: %d", rows))
    c.JSON(http.StatusOK, gin.H{"rows": rows })
}

func (s knowledgeServiceImpl) UpdateFact(c *gin.Context) {
    var career updateCareerFact

    if careerDataError := c.BindJSON(&career); careerDataError != nil {
        fmt.Println(careerDataError)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", careerDataError.Error())})
        return
    }

    rows, err := s.repo.UpdateFact(career.CareerId, career.Aptitude, career.Skill, career.Interest, career.PAptitude, career.PSkill, career.PInterest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
        return
	}

    // c.String(http.StatusOK, fmt.Sprintf("fact added, rows inserted: %d", rows))
    c.JSON(http.StatusOK, gin.H{"rows": rows })
}

func (s knowledgeServiceImpl) DeleteCareer(c *gin.Context) {
    var career deleteCareer

    if careerDataError := c.BindJSON(&career); careerDataError != nil {
        fmt.Println(careerDataError)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", careerDataError.Error())})
        return
    }

    rows, err := s.repo.DeleteCareer(career.CareerId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
        return
	}

    // c.String(http.StatusOK, fmt.Sprintf("fact added, rows inserted: %d", rows))
    c.JSON(http.StatusOK, gin.H{"rows": rows })
}