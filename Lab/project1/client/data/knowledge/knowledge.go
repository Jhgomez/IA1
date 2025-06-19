package Knowledgesource

import (
	// "os"
	"sync"
	"fmt"

	"github.com/mndrix/golog"
)

type careerDto struct {
	Faculty string
	Career string
}

type CareerFactDto struct {
	Faculty string
	Career string
	Aptitude string
	Skill string
	Interest string
}

type Knowledge interface {
	LoadCareerFacts(facts []CareerFactDto)
	SuggestCareers(interest, aptitude, skill string) []careerDto
	StartNewMachine()
	LoadRule(rule string)
}

type knowledgeImpl struct {
	Machine golog.Machine
}

var once sync.Once
var knowledge Knowledge

func GetKnowledgeSource() Knowledge {
	if knowledge == nil {
		once.Do(func() {
			knowledge = &knowledgeImpl{ Machine: golog.NewMachine() }
		})
	}

	return knowledge
}

func (k *knowledgeImpl) SuggestCareers(aptitude, skill, interest string) []careerDto {
	query := fmt.Sprintf("career(Faculty, Career, %s, %s, %s).", aptitude, skill, interest)
	results := []careerDto{}

	solutions := k.Machine.ProveAll(query)

	for _, solution := range solutions {
		faculty := solution.ByName_("Faculty").String()
		career := solution.ByName_("Career").String()

		results = append(results, careerDto{Faculty: faculty, Career: career})
	}

	return results
}

func (k *knowledgeImpl) StartNewMachine() {
	k.Machine = golog.NewMachine()
}

func (k *knowledgeImpl) LoadCareerFacts(facts []CareerFactDto) {
	// knowledgeBase, err := os.ReadFile(url)

	// 	if err != nil {
	// 		panic(err)
	// 	}
	for _, fact := range facts {
		k.Machine = k.Machine.Consult(fmt.Sprintf("career(%s, %s, %s, %s, %s).", fact.Faculty, fact.Career, fact.Aptitude, fact.Skill, fact.Interest))	
	}
}

func (k *knowledgeImpl) LoadRule(rule string) {
	k.Machine = k.Machine.Consult(rule)
}


// func main() {
// 	source1 := GetKnowledgeSource()

// 	source1.LoadCareerFacts([]CareerFactDto{CareerFactDto{"a", "b", "c", "d", "e"}})

// 	// source1.LoadRule("suggested_career(Faculty, Career, Aptitude, Skill, Interest) :- career(Faculty, Career, Aptitude, Skill, Interest).")

// 	fmt.Println(source1.SuggestCareers("c", "d", "e"))

// //  ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(unsafe.Sizeof(arr[0]))))
// //  fmt.Println(*ptr) // output: 2

// 	source1.StartNewMachine()

// 	source1.LoadCareerFacts([]CareerFactDto{CareerFactDto{"f", "g", "h", "i", "j"}})


// 	// source1.LoadRule("suggested_career(Faculty, Career, Aptitude, Skill, Interest) :- career(Faculty, Career, Aptitude, Skill, Interest).")

// 	fmt.Println(source1.SuggestCareers("h", "_", "_"))





	// source := golog.NewMachine()

	// source = source.Consult("career(a, b, c, d, e).")
	// source = source.Consult("suggested_career(Faculty, Career, Aptitude, Skill, Interest) :- career(Faculty, Career, Aptitude, Skill, Interest).")

	// results := []careerDto{}

	// solutions := source.ProveAll("career(Faculty, Career, c, d, e).")

	// for _, solution := range solutions {
	// 	faculty := solution.ByName_("Faculty").String()
	// 	career := solution.ByName_("Career").String()

	// 	results = append(results, careerDto{Faculty: faculty, Career: career})
	// }

	// fmt.Println(results)
// }

