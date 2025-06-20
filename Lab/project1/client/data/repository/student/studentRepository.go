package studentrepo

import (
	"fmt"
	"sync"

 	"unimatch/data/knowledge"
	// "unimatch/data/repository/home"
)

type Career struct {
	Faculty  		string
	Career   		string
	Aptitude 		string
	Skill    		string
	Interest 		string
	Compatibility	float32
}

type StudentRepository interface {
	SuggestCareers(Aptitude1, Skill1, Interest1, Aptitude2, Skill2, Interest3 string) map[string]Career
}

type studentRepositoryImpl struct {
	knowledgeSource knowledgesource.Knowledge
}

var studentRepo StudentRepository
var once sync.Once

func GetStudentRepository() StudentRepository {
	if studentRepo == nil {
		once.Do(func() {
			studentRepo = studentRepositoryImpl{ knowledgeSource: knowledgesource.GetKnowledgeSource() }
		})
	}

	return studentRepo
}

func (s studentRepositoryImpl) SuggestCareers(Aptitude1, Skill1, Interest1, Aptitude2, Skill2, Interest2 string) map[string]Career {
	suggestedCareers := make(map[string]Career)

	// 33% compatibility
	oneThirdCombinations := [][]string {
		{Aptitude1, "Skill", "Interest"},
		{Aptitude2, "Skill", "Interest"},
		{Skill1, "Skill", "Interest"},
		{Skill2, "Skill", "Interest"},
		{Interest1, "Skill", "Interest"},
		{Interest2, "Skill", "Interest"},

		{"Aptitude", Aptitude1, "Interest"},
		{"Aptitude", Aptitude2, "Interest"},
		{"Aptitude", Skill1, "Interest"},
		{"Aptitude", Skill2, "Interest"},
		{"Aptitude", Interest1, "Interest"},
		{"Aptitude", Interest2, "Interest"},

		{"Aptitude", "Skill", Aptitude1},
		{"Aptitude", "Skill", Aptitude2},
		{"Aptitude", "Skill", Skill1},
		{"Aptitude", "Skill", Skill2},
		{"Aptitude", "Skill", Interest1},
		{"Aptitude", "Skill", Interest2},
	}

	
	for _, suggestion := range s.getSuggestions(oneThirdCombinations, 33.33) {
		suggestedCareers[suggestion.Faculty + suggestion.Career] = suggestion
	}

	// 66.66% compatible
	twoThirdCombinations := [][]string {
		{Aptitude1, "Skill", Interest1},
		{Aptitude2, "Skill", Interest1},
		{Aptitude2, "Skill", Interest2},
		{Aptitude1, "Skill", Interest2},

		{Interest1, "Skill", Aptitude1},
		{Interest1, "Skill", Aptitude2},
		{Interest2, "Skill", Aptitude2},
		{Interest2, "Skill", Aptitude1},
		{Interest2, "Skill", Aptitude1},

		{"Aptitude", Skill1, Interest1},
		{"Aptitude", Skill2, Interest1},
		{"Aptitude", Skill2, Interest2},
		{"Aptitude", Skill1, Interest2},

		{"Aptitude", Interest1, Skill1},
		{"Aptitude", Interest1, Skill2},
		{"Aptitude", Interest2, Skill2},
		{"Aptitude", Interest2, Skill1},

		{Aptitude1, Skill1, "Interest"},
		{Aptitude2, Skill1, "Interest"},
		{Aptitude2, Skill2, "Interest"},
		{Aptitude1, Skill2, "Interest"},

		{Skill1, Aptitude1, "Interest"},
		{Skill2, Aptitude1, "Interest"},
		{Skill2, Aptitude2, "Interest"},
		{Skill1, Aptitude2, "Interest"},
	}

	for _, suggestion := range s.getSuggestions(twoThirdCombinations, 66.67) {
		suggestedCareers[suggestion.Faculty + suggestion.Career] = suggestion
	}

	// 100% compatible careers
	combinations := [][]string{
		{Aptitude1, Skill1, Interest1}, 
		{Aptitude1, Skill1, Interest2},
		{Aptitude1, Skill2, Interest1}, 
		{Aptitude1, Skill2, Interest2}, 
		{Aptitude2, Skill2, Interest2}, 
		{Aptitude2, Skill2, Interest1}, 
		{Aptitude2, Skill1, Interest2}, 
		{Aptitude2, Skill1, Interest1},

		{Aptitude1, Interest1, Skill1}, 
		{Aptitude1, Interest2, Skill1},
		{Aptitude1, Interest1, Skill2}, 
		{Aptitude1, Interest2, Skill2}, 
		{Aptitude2, Interest2, Skill2}, 
		{Aptitude2, Interest1, Skill2}, 
		{Aptitude2, Interest2, Skill1}, 
		{Aptitude2, Interest1, Skill1},

		{Skill1, Interest1, Aptitude1}, 
		{Skill1, Interest2, Aptitude1},
		{Skill2, Interest1, Aptitude1}, 
		{Skill2, Interest2, Aptitude1}, 
		{Skill2, Interest2, Aptitude2}, 
		{Skill2, Interest1, Aptitude2}, 
		{Skill1, Interest2, Aptitude2}, 
		{Skill1, Interest1, Aptitude2},

		{Interest1, Aptitude1, Skill1}, 
		{Interest2, Aptitude1, Skill1},
		{Interest1, Aptitude1, Skill2}, 
		{Interest2, Aptitude1, Skill2}, 
		{Interest2, Aptitude2, Skill2}, 
		{Interest1, Aptitude2, Skill2}, 
		{Interest2, Aptitude2, Skill1}, 
		{Interest1, Aptitude2, Skill1},

		{Interest1, Skill1, Aptitude1}, 
		{Interest2, Skill1, Aptitude1},
		{Interest1, Skill2, Aptitude1}, 
		{Interest2, Skill2, Aptitude1}, 
		{Interest2, Skill2, Aptitude2}, 
		{Interest2, Skill1, Aptitude2}, 
		{Interest1, Skill2, Aptitude2}, 
		{Interest1, Skill1, Aptitude2},

		{Skill1, Aptitude1, Interest1}, 
		{Skill1, Aptitude1, Interest2},
		{Skill2, Aptitude1, Interest1}, 
		{Skill2, Aptitude1, Interest2}, 
		{Skill2, Aptitude2, Interest2}, 
		{Skill2, Aptitude2, Interest1}, 
		{Skill1, Aptitude2, Interest2}, 
		{Skill1, Aptitude2, Interest1},
	}

	for _, suggestion := range s.getSuggestions(combinations, 100) {
		suggestedCareers[suggestion.Faculty + suggestion.Career] = suggestion
	}

	return suggestedCareers
}

func (s studentRepositoryImpl) getSuggestions(queriesData [][]string, compatibility float32) []Career {
	careers := []Career{}
	for _, queryData := range queriesData {

		suggestions := s.knowledgeSource.SuggestCareers(queryData[0], queryData[1], queryData[2])

		for _, suggestion := range suggestions {
			careers = append(careers, Career{Faculty: suggestion.Faculty, Career: suggestion.Career, Aptitude: suggestion.Aptitude, Skill: suggestion.Skill, Interest: suggestion.Interest, Compatibility: compatibility})
		}
	}
	
	return careers
}

// func main() {
// 	homerepo.GetHomeRepo().LoadKnowledgeBase()
// 	fmt.Println(GetStudentRepository().SuggestCareers("matematica", "dibujo", "construccion", "k", "k", "k"))
// }
