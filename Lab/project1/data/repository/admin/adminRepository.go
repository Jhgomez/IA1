package admin

import (
	"fmt"
	"sync"

 	"unimatch/data/knowledge"
)

type AdminRepository interface {
	LoadMachine(url string) []string
}

type adminRepositoryImpl struct {
	knowledgeSource KnowledgeSource.Knowledge
}

var adminRepo AdminRepository
var once sync.Once

func GetStudentRepository() AdminRepository {
	if adminRepo == nil {
		once.Do(func() {
			adminRepo = adminRepositoryImpl{ knowledgeSource: KnowledgeSource.GetKnowledgeSource() }
		})
	}

	return adminRepo
}

func (s adminRepositoryImpl) LoadMachine(url string) []string {
	return s.knowledgeSource.LoadMachine(url)
}
