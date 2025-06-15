package domain

type StudentRepository interface {
	suggestCareer(interest, aptitude, skill string) []string
}