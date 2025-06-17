package KnowledgeSource

import (
	"os"
	"sync"
	"fmt"

	"github.com/mndrix/golog"
)

type Knowledge interface {
	LoadMachine(url string)
	SuggestCareers(interest, aptitude, skill string) []string
}

type knowledgeImpl struct {
	machine golog.Machine
}

var once sync.Once
var knowledge Knowledge

func GetKnowledgeSource() Knowledge {
	if knowledge == nil {
		once.Do(func() {
			knowledgeBase, err := os.ReadFile("../knowledge/knowledgeBase.pl")

			if err != nil {
				panic(err)
			}

			knowledge = knowledgeImpl{ machine: golog.NewMachine().Consult(string(knowledgeBase)) }
		})
	}

	return knowledge
}

func (k knowledgeImpl) SuggestCareers(aptitude, skill, interest string) []string {
	query := fmt.Sprintf("career(Faculty, Career, %s, %s, %s).", aptitude, skill, interest)
	results := []string{}

	solutions := k.machine.ProveAll(query)

	for _, solution := range solutions {
		faculty := solution.ByName_("Faculty").String()
		career := solution.ByName_("Career").String()

		results = append(results, fmt.Sprintf("%s, %s", faculty, career))
	}

	return results
}

func (k knowledgeImpl) LoadMachine(url string) {
	knowledgeBase, err := os.ReadFile(url)

		if err != nil {
			panic(err)
		}

	k.machine = golog.NewMachine().Consult(string(knowledgeBase))
}

// func tryIt(k Knowledge) {
// 	k.GetMachine()
// }

// func main() {
// 	fmt.Println(KnowledgeImpl{}.SuggestCareers("matematica", "dibujo", "construccion"))
// 	// tryIt(KnowledgeImpl{})

// //  ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(unsafe.Sizeof(arr[0]))))
// //  fmt.Println(*ptr) // output: 2
// }

