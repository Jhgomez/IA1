package repository

import (
	"fmt"
	"sync"

 	"unimatch/data/knowledge"
)

type StudentRepository interface {
	SuggestCareers(interest, aptitude, skill string) []string
}

type studentRepositoryImpl struct {
	knowledgeSource KnowledgeSource.Knowledge
}

var studentRepo StudentRepository
var once sync.Once

func GetStudentRepository() StudentRepository {
	if studentRepo == nil {
		once.Do(func() {
			studentRepo = studentRepositoryImpl{ knowledgeSource: KnowledgeSource.GetKnowledgeSource() }
		})
	}

	return studentRepo
}

func (s studentRepositoryImpl) SuggestCareers(interest, aptitude, skill string) []string {
	return s.knowledgeSource.SuggestCareers(interest, aptitude, skill)
}
