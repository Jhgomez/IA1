package studentrepo

import (
	"fmt"
	"sync"
	"strings"
	"strconv"
	"maps"
	"slices"

 	"unimatch/data/knowledge"
	// "unimatch/data/repository/knowledge"
	// "unimatch/data/repository/admin"
)

type Career struct {
	Faculty  		string
	Career   		string
	AptitudeMatch 	float32
	SkillMatch    	float32
	InterestMatch 	float32
	TotalMatch      float32
}

type StudentRepository interface {
	SuggestCareers(Aptitude, Skill, Interest []string) []*Career
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

func (s studentRepositoryImpl) SuggestCareers(Aptitude, Skill, Interest []string) []*Career {
	careers := make(map[string]*Career)

	// A stands for Aptitudes
	// AT stands for Aptitudes Total
	// S stands for Skills
	// ST stands for Skills Total
	// I stands for Interests
	// IT stands for Interests Total
	
	queryString := fmt.Sprintf(
			"match_all_careers(%s, %s, %s, Faculty, Career, A, S, I, AT, ST, IT).",
			s.formatPrologList(Aptitude),
			s.formatPrologList(Skill), 
			s.formatPrologList(Interest),
		)
	
	// fmt.Printf(queryString)
	suggestions := s.knowledgeSource.SuggestCareers(queryString)

	for _, solution := range suggestions {
		faculty := solution.ByName_("Faculty").String()
		career := solution.ByName_("Career").String()

		aptitudeCount, _ := strconv.ParseFloat(solution.ByName_("A").String(), 32)
		skillCount, _ := strconv.ParseFloat(solution.ByName_("S").String(), 32)
		interestCount, _ := strconv.ParseFloat(solution.ByName_("I").String(), 32)

		skillTotalCount, _ := strconv.ParseFloat(solution.ByName_("ST").String(), 32)
		interestTotalCount, _ := strconv.ParseFloat(solution.ByName_("IT").String(), 32)
		aptitudeTotalCount, _ := strconv.ParseFloat(solution.ByName_("AT").String(), 32)

		AptitudeMatch := float32(aptitudeCount / aptitudeTotalCount)
		SkillMatch := float32(skillCount / skillTotalCount)
		InterestMatch := float32(interestCount / interestTotalCount)

		match, exists := careers[faculty+career]

		if exists {
			if match.AptitudeMatch < AptitudeMatch {
				match.AptitudeMatch = AptitudeMatch
			}

			if match.SkillMatch < SkillMatch {
				match.SkillMatch = SkillMatch
			}

			if match.InterestMatch < InterestMatch {
				match.InterestMatch = InterestMatch
			}
		} else {
			careers[faculty+career] = &Career{
				Faculty: faculty,
				Career: career,
				AptitudeMatch: AptitudeMatch,
				SkillMatch: SkillMatch,
				InterestMatch: InterestMatch,
				TotalMatch: (AptitudeMatch + SkillMatch + InterestMatch) / 300,
			}
		}

		// bindings := solution.Bindings()
		// fmt.Printf("Faculty: %v\n", solution.ByName_("Faculty"))
		// fmt.Printf("Career: %v\n", solution.ByName_("Career"))
		// fmt.Printf("Aptitude Matches:  %s/%s\n", solution.ByName_("A"), solution.ByName_("AT"))
		// fmt.Printf("Skill Matches:  %s/%s\n", solution.ByName_("S"), solution.ByName_("ST"))
		// fmt.Printf("Interest Matches: %s/%s\n", solution.ByName_("I"), solution.ByName_("IT"))
		// fmt.Println("------------------------")
	}

	return slices.Collect(maps.Values(careers))
}

func(r studentRepositoryImpl) formatPrologList(items []string) string {
	var quoted []string

	for _, item := range items {
		str := fmt.Sprintf("%v", item) // safely convert anything to string
		quoted = append(quoted, fmt.Sprintf("'%s'", str))
	}

	return "[" + strings.Join(quoted, ", ") + "]"
}

// func main() {
	
// 	// repoAdmin := adminrepo.GetAdminRepository()

// 	// repoAdmin.AddCareer("a", "b", "c", "d", "f")
// 	// repoAdmin.AddCareer("a", "b", "f", "g", "h")
// 	// repoAdmin.AddCareer("a", "b", "i", "j", "k")

// 	repo := knowledgerepo.GetKnowledgeRepo()

// 	repo.LoadKnowledgeBase()

// 	// println(repoAdmin.GetCareers())

// 	srepo := GetStudentRepository()

// 	fmt.Printf("%v", )

// 	for key, val := range srepo.SuggestCareers([]string{"analisis"}, []string{"laboratorio"}, []string{"procesos industriales"}) {
// 		fmt.Printf("Key: %s, Score: %v\n", key, val)
// 	}
// }
