package homerepo

import(
	// "fmt"

	"unimatch/data/knowledge"
	"unimatch/data/network"
)

type HomeRepository interface {
	LoadKnowledgeBase() (int, error)
}

type homeRepositoryImpl struct {
	knowledgeSource knowledgesource.Knowledge
	api api.API
}

var homeRepository HomeRepository

func GetHomeRepo() HomeRepository {
	if homeRepository == nil {
		homeRepository = homeRepositoryImpl{ knowledgeSource: knowledgesource.GetKnowledgeSource(), api: api.GetApi() }
	}

	return homeRepository
}

func (r homeRepositoryImpl) LoadKnowledgeBase() (int, error) {
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