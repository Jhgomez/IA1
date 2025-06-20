package knowledgerepo

import(
	// "fmt"

	"unimatch/data/knowledge"
	"unimatch/data/network"
)

type KnowledgeRepository interface {
	LoadKnowledgeBase() (int, error)
}

type knowledgeRepositoryImpl struct {
	knowledgeSource knowledgesource.Knowledge
	api api.API
}

var homeRepository KnowledgeRepository

func GetKnowledgeRepo() KnowledgeRepository {
	if homeRepository == nil {
		homeRepository = knowledgeRepositoryImpl{ knowledgeSource: knowledgesource.GetKnowledgeSource(), api: api.GetApi() }
	}

	return homeRepository
}

func (r knowledgeRepositoryImpl) LoadKnowledgeBase() (int, error) {
	apiFacts, err := r.api.GetFacts()

	if err != nil {
		return -1, err
	}

	// fmt.Println("apiFacts")
	// fmt.Println(apiFacts)

	careerFacts := make([]knowledgesource.CareerFactDto, len(apiFacts))

	for i, fact := range apiFacts {
		careerFacts[i] = knowledgesource.CareerFactDto{ 
			Faculty: fact.Faculty,
			Career: fact.Career,
			Aptitude: fact.Aptitude, 
			Skill: fact.Skill, 
			Interest: fact.Interest,
		}
	}

	// fmt.Println("careerFacts")
	// fmt.Println(careerFacts)

	r.knowledgeSource.LoadCareerFacts(careerFacts)

	return len(careerFacts), nil
}

// func main() {
// 	GetHomeRepo().LoadKnowledgeBase()
// }