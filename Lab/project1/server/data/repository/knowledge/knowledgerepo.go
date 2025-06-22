package knowledgerepo

import (
	"unimatchserver/data/db/dao"
)

type Career struct {
	CareerId int
	Faculty  string
	Career   string
	Aptitude []string
	Skill    []string
	Interest []string
}

type KnowledgeRepo interface {
	GetFacts() ([]Career, error)
	AddFact(careerId int, Aptitude, Skill, Interest []string) (int64, error)
	DeleteFact(careerId int, Aptitude, Skill, Interest string) (int64, error)
	UpdateFact(careerId int, Aptitude, Skill, Interest, PAptitude, PSkill, PInterest []string) (int64, error)
	AddCareer(Faculty, Career string) (int64, error)
    DeleteCareer(careerId int) (int64, error)
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

func (r knowledgeRepoImpl) GetFacts() ([]Career, error) {
	careersDto, err := r.dao.GetFacts()

	if err != nil {
		return []Career{}, err
	}

	careers := make([]Career, len(careersDto))

	for i, career := range careersDto {
		careers[i] = Career{
			CareerId:	career.CareerId,
			Faculty:	career.Faculty,
			Career:		career.Career,
			Aptitude:	career.Aptitudes,
			Skill:		career.Skills,
			Interest:	career.Interests,
		}
	}

	// // Convert the slice of structs to JSON
	// jsonData, err := json.Marshal(facts)
	// if err != nil {
	// 	fmt.Printf("Error marshalling data to JSON: %v", err)
	// 	return []byte{}, err
	// }

	// // Print the JSON data
	// return jsonData, nil

	return careers, nil
}

func (r knowledgeRepoImpl) AddFact(careerId int, Aptitude, Skill, Interest []string) (int64, error) {
	return r.dao.AddFact(careerId, Aptitude, Skill, Interest)
}

func (r knowledgeRepoImpl) DeleteFact(careerId int, Aptitude, Skill, Interest string) (int64, error) {
	return r.dao.DeleteFact(careerId, Aptitude, Skill, Interest)
}

func (r knowledgeRepoImpl) UpdateFact(careerId int, Aptitude, Skill, Interest, PAptitude, PSkill, PInterest []string) (int64, error) {
	return r.dao.UpdateFact(careerId, Aptitude, Skill, Interest, PAptitude, PSkill, PInterest)
}

func (r knowledgeRepoImpl) AddCareer(Faculty, Career string) (int64, error) {
	return r.dao.AddCareer(Faculty, Career)
}

func (r knowledgeRepoImpl) DeleteCareer(careerId int) (int64, error) {
	return r.dao.DeleteCareer(careerId)
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