package knowledgesource

import (
	// "os"
	"sync"
	// "fmt"

	"github.com/mndrix/golog"
	"github.com/mndrix/golog/term"
)

type careerDto struct {
	Faculty string
	Career string
}

type CareerFactDto struct {
	Faculty string
	Career string
	Aptitude []string
	Skill []string
	Interest []string
}

type Knowledge interface {
	LoadCareerFact(fact string)
	SuggestCareers(query string) []term.Bindings
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

func (k *knowledgeImpl) SuggestCareers(query string) []term.Bindings {
	return k.Machine.ProveAll(query)
}

func (k *knowledgeImpl) StartNewMachine() {
	k.Machine = golog.NewMachine()
}

func (k *knowledgeImpl) LoadCareerFact(fact string) {
	k.Machine = k.Machine.Consult(fact)
}

func (k *knowledgeImpl) LoadRule(rule string) {
	k.Machine = k.Machine.Consult(rule)
}

// func main() {
// // 	source1 := GetKnowledgeSource()

// // 	source1.LoadCareerFacts([]CareerFactDto{CareerFactDto{"a", "b", "c", "d", "e"}})

// // 	// source1.LoadRule("suggested_career(Faculty, Career, Aptitude, Skill, Interest) :- career(Faculty, Career, Aptitude, Skill, Interest).")

// // 	fmt.Println(source1.SuggestCareers("c", "d", "e"))

// // //  ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(unsafe.Sizeof(arr[0]))))
// // //  fmt.Println(*ptr) // output: 2

// // 	source1.StartNewMachine()

// // 	source1.LoadCareerFacts([]CareerFactDto{CareerFactDto{"f", "g", "h", "i", "j"}})


// // 	// source1.LoadRule("suggested_career(Faculty, Career, Aptitude, Skill, Interest) :- career(Faculty, Career, Aptitude, Skill, Interest).")

// // 	fmt.Println(source1.SuggestCareers("h", "_", "_"))





// 	source := golog.NewMachine()

// 	source = source.Consult(prologProgram)
	
// 	query := `match_all_careers(
//     ['Problem Solving', 'Creativity'],
//     ['Programming', 'Communication'],
//     ['Technology', 'Design'],
//     Faculty, Career,
//     A, S, I,
//     AT, ST, IT).`



// 	// results := []careerDto{}

// 	solutions := source.ProveAll(query)

// 	for _, solution := range solutions {
// 	// 	faculty := solution.ByName_("Faculty").String()
// 	// 	career := solution.ByName_("Career").String()

// 	// 	results = append(results, careerDto{Faculty: faculty, Career: career})

// 		// bindings := solution.Bindings()
//     fmt.Printf("Faculty: %v\n", solution.ByName_("Faculty"))
//     fmt.Printf("Career: %v\n", solution.ByName_("Career"))
//     fmt.Printf("Aptitude Matches:  %v/%v\n\n", solution.ByName_("A"), solution.ByName_("AT"))
//     fmt.Printf("Skill Matches:  %v/%v\n\n", solution.ByName_("S"), solution.ByName_("ST"))
//     fmt.Printf("Interest Matches: %v/%v\n\n", solution.ByName_("I"), solution.ByName_("IT"))
//     fmt.Println("------------------------")
// 	}

// 	// fmt.Println(results)
// }

