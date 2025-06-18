package knowledgerepo

import (
	"unimatchserver/data/db/dao"
)

type KnowledgeRepo interface {
	GetFacts() (string, error)
	AddFact(Faculty, Career, Aptitude, Skill, Interest string) (int64, error)
	DeleteFact(Faculty, Career string) (int64, error)
	UpdateFact(Faculty, Career, Aptitude, Skill, Interest string) (int64, error)
}

type knowledgeRepoImpl struct {
	dao knowledgedao.KnowledgeDao
}

var knowledgeRepo KnowledgeRepo

func GetKnowledgeRepo() KnowledgeRepo {
	if knowledgeRepo == nil {
		knowledgeRepo = knowledgeRepoImpl { knowledgedao.GetKnowledgeDao() }
	}

	return knowledgeRepo
}

func (r knowledgeRepoImpl) GetFacts() (string, error) {
	return r.dao.GetFacts()
}

func (r knowledgeRepoImpl) AddFact(Faculty, Career, Aptitude, Skill, Interest string) (int64, error) {
	return r.dao.AddFact(Faculty, Career, Aptitude, Skill, Interest)
}

func (r knowledgeRepoImpl) DeleteFact(Faculty, Career string) (int64, error) {
	return r.dao.DeleteFact(Faculty, Career)
}

func (r knowledgeRepoImpl) UpdateFact(Faculty, Career, Aptitude, Skill, Interest string) (int64, error) {
	return r.dao.UpdateFact(Faculty, Career, Aptitude, Skill, Interest)
}

// func main() {
// 	repo := GetKnowledgeRepo()

// 	rows, err := repo.UpdateFact("ingenieria", "sistemas", "logica3", "programacion", "tecnologia")
// 	// rows, err := repo.DeleteFact("ingenieria", "sistemas")

// 	if err != nil {
// 		fmt.Printf("Error inserting career: %v\n", err)
// 	}

// 	fmt.Printf("%d Rows inserted\n", rows)

// 	json, _ := repo.GetFacts()

// 	fmt.Println(json)
// }