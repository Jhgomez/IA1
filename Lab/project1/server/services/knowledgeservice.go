package knowledgeservice

import (
    "net/http"
    "fmt"

    "github.com/gin-gonic/gin"

	"unimatchserver/data/repository/knowledge"
)

type KnowledgeService interface {
    GetFacts(c *gin.Context)
    AddFact(c *gin.Context)
    DeleteFact(c *gin.Context)
    UpdateFact(c *gin.Context)
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
	Faculty  string `json:"Faculty"`
	Career   string `json:"Career"`
	Aptitude string `json:"Aptitude"`
	Skill    string `json:"Skill"`
	Interest string `json:"Interest"`
}

type career struct {
	Faculty  string `json:"Faculty"`
	Career   string `json:"Career"`
}

func (s knowledgeServiceImpl) GetFacts(c *gin.Context) {
    facts, err := s.repo.GetFacts()

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
        return
    }

    c.JSON(http.StatusOK, facts)
}

func (s knowledgeServiceImpl) AddFact(c *gin.Context) {
    var career careerFact

    if careerDataError := c.BindJSON(&career); careerDataError != nil {
        fmt.Println(careerDataError)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", careerDataError.Error())})
        return
    }

    rows, err := s.repo.AddFact(career.Faculty, career.Career, career.Aptitude, career.Skill, career.Interest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
        return
	}

    c.JSON(http.StatusOK, gin.H{"rows": rows })

    // c.String(http.StatusOK, fmt.Sprintf("fact added, rows inserted: %d", rows))
}

func (s knowledgeServiceImpl) DeleteFact(c *gin.Context) {
    var career career

    if careerDataError := c.BindJSON(&career); careerDataError != nil {
        fmt.Println(careerDataError)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", careerDataError.Error())})
        return
    }

    rows, err := s.repo.DeleteFact(career.Faculty, career.Career)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
        return
	}

    // c.String(http.StatusOK, fmt.Sprintf("fact deleted, rows deleted: %d", rows))
    c.JSON(http.StatusOK, gin.H{"rows": rows })
}

func (s knowledgeServiceImpl) UpdateFact(c *gin.Context) {
    var career careerFact

    if careerDataError := c.BindJSON(&career); careerDataError != nil {
        fmt.Println(careerDataError)
        c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", careerDataError.Error())})
        return
    }

    rows, err := s.repo.UpdateFact(career.Faculty, career.Career, career.Aptitude, career.Skill, career.Interest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%v", err)})
        return
	}

    // c.String(http.StatusOK, fmt.Sprintf("fact added, rows inserted: %d", rows))
    c.JSON(http.StatusOK, gin.H{"rows": rows })
}
