package KnowledgeSource

import (
	"os"
	"sync"
	"fmt"
	"log"

	"github.com/mndrix/golog"
)

type Knowledge interface {
	GetMachine() golog.Machine
	LoadMachine(url string) golog.Machine
	SuggestCareers(interest, aptitude, skill string) []string
}

type KnowledgeImpl struct {}

var machine golog.Machine
var once sync.Once

func init() {
	once.Do(func() {
		knowledgeBase, err := os.ReadFile("./knowledgeBase.pl")

		if err != nil {
			panic(err)
		}

		machine = golog.NewMachine().Consult(string(knowledgeBase))
	})

	log.Println("Initialization")
}

func (k KnowledgeImpl) GetMachine() golog.Machine {
	if machine == nil {
		once.Do(func() {
			knowledgeBase, err := os.ReadFile("./knowledgeBase.pl")

			if err != nil {
				panic(err)
			}

			machine = golog.NewMachine().Consult(string(knowledgeBase))
		})
	}

	return machine
}

func (k KnowledgeImpl) SuggestCareers(aptitude, skill, interest string) []string {
	query := fmt.Sprintf("career(Faculty, Career, %s, %s, %s).", aptitude, skill, interest)
	results := []string{}

	solutions := machine.ProveAll(query)

	for _, solution := range solutions {
		faculty := solution.ByName_("Faculty").String()
		career := solution.ByName_("Career").String()

		results = append(results, fmt.Sprintf("%s, %s", faculty, career))
	}

	return results
}

func (k KnowledgeImpl) LoadMachine(url string) golog.Machine {
	knowledgeBase, err := os.ReadFile("./knowledgeBase.pl")

		if err != nil {
			panic(err)
		}

	machine = golog.NewMachine().Consult(string(knowledgeBase))

	return machine
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

