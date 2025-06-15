package main

import (
	"os"
	"sync"

	"github.com/mndrix/golog"
)

type Knowledge interface {
	GetMachine() golog.Machine
	LoadMachine(url string) golog.Machine
	SuggestCareers() []string
}

type KnowledgeImpl struct {}

var machine golog.Machine
var once sync.Once

func (k KnowledgeImpl) GetMachine() golog.Machine {
	once.Do(func() {
		knowledgeBase, err := os.ReadFile("./knowledgeBase.pl")

		if err != nil {
			panic(err)
		}

		machine = golog.NewMachine().Consult(string(knowledgeBase))
	})

	return machine
}

func (k KnowledgeImpl) SuggestCareers() []string {
	return []string{}
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
// 	tryIt(KnowledgeImpl{})
// //  ptr = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(unsafe.Sizeof(arr[0]))))
// //  fmt.Println(*ptr) // output: 2
// }

