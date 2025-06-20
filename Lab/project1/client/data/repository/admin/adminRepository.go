package adminrepo

import (
	"fmt"
	"sync"

	"unimatch/data/network"
)

type EditableCareer struct {
	Faculty		string
	Career   	string
	Aptitude 	string
	Skill    	string
	Interest 	string
}

type AdminRepository interface {
	GetCareers() ([]EditableCareer, error)
	UpdateCareer(Faculty, Career, Aptitude, Skill, Interest string) (int, error)
	DeleteCareer(Faculty, Career string) (int, error)
	AddCareer(Faculty, Career, Aptitude, Skill, Interest string) (int, error)
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

func (a adminRepositoryImpl) GetCareers() ([]EditableCareer, error) {
	apiFacts, err := a.api.GetFacts()

	if err != nil {
		return []EditableCareer{}, err
	}

	// fmt.Println("apiFacts")
	// fmt.Println(apiFacts)

	careers := make([]EditableCareer, len(apiFacts))

	for i, fact := range apiFacts {
		careers[i] = EditableCareer{ 
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

func (a adminRepositoryImpl) UpdateCareer(Faculty, Career, Aptitude, Skill, Interest string) (int, error) {
	rows, err := a.api.UpdateFact(Faculty, Career, Aptitude, Skill, Interest)

	if err != nil {
		return -1, err
	}

	return rows, nil
}

func (a adminRepositoryImpl) DeleteCareer(Faculty, Career string) (int, error) {
	rows, err := a.api.DeleteFact(Faculty, Career)

	if err != nil {
		return -1, err
	}

	return rows, nil
}

func (a adminRepositoryImpl) AddCareer(Faculty, Career, Aptitude, Skill, Interest string) (int, error) {
	rows, err := a.api.AddFact(Faculty, Career, Aptitude, Skill, Interest)

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