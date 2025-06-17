package main

import (
	"encoding/json"
	"fmt"

	"database/sql"

	"unimatch/server/data/db"
)

type career struct {
	Faculty  string `json:"Faculty"`
	Career   string `json:"Career"`
	Aptitude string `json:"Aptitude"`
	Skill    string `json:"Skill"`
	Interest string `json:"Interest"`
}

type KnowledgeRepository interface {
	GetFacts() (string, error)
	AddFact(Faculty, Career, Aptitude, Skill, Interest string) (int64, error)
	DeleteFact(Faculty, Career string)
}

type knowledgeRepositoryImpl struct {
	db db.Knowledgedb
}

var knowledgeRepository KnowledgeRepository

func GetKnowledgeRepository() KnowledgeRepository {
	if knowledgeRepository == nil {
		knowledgeRepository = knowledgeRepositoryImpl{ db.GetKnowledgeDb() }
	}
	return knowledgeRepository
}

func (k knowledgeRepositoryImpl) GetFacts() (string, error) {
	rows, err := k.db.GetConnection().Query("SELECT * FROM proyecto1.careers_knowledge")

	if err != nil {
		fmt.Printf("Error querying the database: %v", err)
		return "", err
	}

	var careers []career
	for rows.Next() {
		var career career
		if err := rows.Scan(&career.Faculty, &career.Career, &career.Aptitude, &career.Skill, &career.Interest); err != nil {
			fmt.Printf("Error scanning row: %v", err)
			return "", err
		}
		careers = append(careers, career)
	}

	// Check for errors after iteration
	if err := rows.Err(); err != nil {
		fmt.Printf("Error during row iteration: %v", err)
		return "", err
	}

	// Convert the slice of structs to JSON
	jsonData, err := json.Marshal(careers)
	if err != nil {
		fmt.Printf("Error marshalling data to JSON: %v", err)
		return "", err
	}

	// Print the JSON data
	return string(jsonData), nil
}

func (k knowledgeRepositoryImpl) AddFact(Faculty, Career, Aptitude, Skill, Interest string) (int64, error) {
	stmt, err := k.db.GetConnection().Prepare("INSERT INTO proyecto1.careers_knowledge(Faculty, Career, Aptitude, Skill, Interest) VALUES (@Faculty, @Career, @Aptitude, @Skill, @Interest);")
    if err != nil {
        return 0, err
    }

    defer stmt.Close()

    // Execute the prepared statement
    result, err := stmt.Exec(
        sql.Named("Faculty", Faculty),
        sql.Named("Career", Career),
        sql.Named("Aptitude", Aptitude),
		sql.Named("Skill", Skill),
		sql.Named("Interest", Interest),
    )
    if err != nil {
        return 0, err
    }

    // Get the number of row inserted
    rowInserted, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

    return rowInserted, nil
}

func (k knowledgeRepositoryImpl) DeleteFact(Faculty, Career string) {

}

func main() {
	repo := GetKnowledgeRepository()

	rows, err := repo.AddFact("ingenieria", "sistemas", "logica", "programacion", "tecnologia")

	if err != nil {
		fmt.Printf("Error inserting career: %v\n", err)
	}

	fmt.Printf("%d Rows inserted\n", rows)

	json, _ := repo.GetFacts()

	fmt.Println(json)
}