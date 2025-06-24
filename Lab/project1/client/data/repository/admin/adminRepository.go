package adminrepo

import (
	"fmt"
	"sync"

	"unimatch/data/network"
)

type EditableCareer struct {
	CareerId	int
	Faculty		string
	Career   	string
	Aptitude 	[]string
	Skill    	[]string
	Interest 	[]string
}

type AdminRepository interface {
	GetCareers() ([]*EditableCareer, error)
	UpdateCareer(careerId int, Aptitude, Skill, Interest, PAptitude, PSkill, PInterest []string) (int, error)
	DeleteCareer(CareerId int) (int, error)
	AddCareer(Faculty, Career string, Aptitude, Skill, Interest []string) (int, error)
	DeleteFact(careerId int, Aptitude, Skill, Interest string) (int, error)
}

type adminRepositoryImpl struct {
	api api.API
}

var adminRepo AdminRepository
var once sync.Once

func GetAdminRepository() AdminRepository {
	if adminRepo == nil {
		once.Do(func() {
			adminRepo = adminRepositoryImpl{ api: api.GetApi() }
		})
	}

	return adminRepo
}

func (a adminRepositoryImpl) GetCareers() ([]*EditableCareer, error) {
	apiFacts, err := a.api.GetFacts()

	if err != nil {
		return []*EditableCareer{}, err
	}

	// fmt.Println("apiFacts")
	// fmt.Println(apiFacts)

	careers := make([]*EditableCareer, len(apiFacts))

	for i, fact := range apiFacts {
		careers[i] = &EditableCareer{ 
			CareerId: fact.CareerId,
			Faculty: fact.Faculty,
			Career: fact.Career,
			Aptitude: fact.Aptitude, 
			Skill: fact.Skill, 
			Interest: fact.Interest,
		}
	}

	// fmt.Println("careers")
	// fmt.Println(careers)

	return careers, nil
}

func (a adminRepositoryImpl) UpdateCareer(careerId int, Aptitude, Skill, Interest, PAptitude, PSkill, PInterest []string) (int, error) {
	rows, err := a.api.UpdateFact(careerId, Aptitude, Skill, Interest, PAptitude, PSkill, PInterest)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	fmt.Println(rows)
	return rows, nil
}

func (a adminRepositoryImpl) DeleteCareer(CareerId int) (int, error) {
	rows, err := a.api.DeleteCareer(CareerId)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	fmt.Println(rows)
	return rows, nil
}

func (a adminRepositoryImpl) DeleteFact(careerId int, Aptitude, Skill, Interest string) (int, error) {
	rows, err := a.api.DeleteFact(careerId, Aptitude, Skill, Interest)

	if err != nil {
		return -1, err
	}

	return rows, nil
}

func (a adminRepositoryImpl) AddCareer(Faculty, Career string, Aptitude, Skill, Interest []string) (int, error) {
	careerId, err := a.api.AddCareer(Faculty, Career)

	if err != nil {
		return -1, err
	}

	rows, err := a.api.AddFact(careerId, Aptitude, Skill, Interest)

	if err != nil {
		return -1, err
	}

	return rows, nil
}

// func main() {
// 	admin := GetAdminRepository()

// 	perrito, err := admin.GetCareers()

// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

// 	fmt.Println(perrito)





// 	_, err = admin.UpdateCareer("ingenieria", "sistemas", "logica3", "programacion", "tecnologia")

// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

	

// 	perrito, err = admin.GetCareers()

// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

// 	fmt.Println(perrito)




// 	_, err = admin.DeleteCareer("ingenieria", "sistemas")

// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

	

// 	perrito, err = admin.GetCareers()

// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

// 	fmt.Println(perrito)




	
// 	_, err = admin.AddCareer("ingenieria", "sistemas", "logica", "programacion", "tecnologia")

// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

	

// 	perrito, err = admin.GetCareers()

// 	if err != nil {
// 		fmt.Printf("%v", err)
// 	}

// 	fmt.Println(perrito)
// }