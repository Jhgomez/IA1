package knowledgedao

import (
	"fmt"

	"database/sql"

	"unimatchserver/data/db"
)

type careerDto struct {
	Faculty  string
	Career   string
	Aptitude string
	Skill    string
	Interest string
}

type KnowledgeDao interface {
	GetFacts() ([]careerDto, error)
	AddFact(Faculty, Career, Aptitude, Skill, Interest string) (int64, error)
	DeleteFact(Faculty, Career string) (int64, error)
	UpdateFact(Faculty, Career, Aptitude, Skill, Interest string) (int64, error)
}

type knowledgeDaoImpl struct {
	db db.Knowledgedb
}

var knowledgeDao KnowledgeDao

func GetKnowledgeDao() KnowledgeDao {
	if knowledgeDao == nil {
		knowledgeDao = knowledgeDaoImpl{ db.GetKnowledgeDb() }
	}
	return knowledgeDao
}

func (k knowledgeDaoImpl) GetFacts() ([]careerDto, error) {
	rows, err := k.db.GetConnection().Query("SELECT * FROM proyecto1.careers_knowledge")

	if err != nil {
		fmt.Printf("Error querying the database: %v", err)
		return []careerDto{}, err
	}

	var careers []careerDto
	for rows.Next() {
		var career careerDto
		if err := rows.Scan(&career.Faculty, &career.Career, &career.Aptitude, &career.Skill, &career.Interest); err != nil {
			fmt.Printf("Error scanning row: %v", err)
			return []careerDto{}, err
		}
		careers = append(careers, career)
	}

	// Check for errors after iteration
	if err := rows.Err(); err != nil {
		fmt.Printf("Error during row iteration: %v", err)
		return []careerDto{}, err
	}

    return careers, nil
}

func (k knowledgeDaoImpl) UpdateFact(Faculty, Career, Aptitude, Skill, Interest string) (int64, error) {
	stmt, err := k.db.GetConnection().Prepare("UPDATE proyecto1.careers_knowledge SET Aptitude = @Aptitude, Skill = @Skill, Interest = @Interest WHERE Faculty = @Faculty AND Career = @Career;")
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

func (k knowledgeDaoImpl) AddFact(Faculty, Career, Aptitude, Skill, Interest string) (int64, error) {
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

func (k knowledgeDaoImpl) DeleteFact(Faculty, Career string) (int64, error) {
	stmt, err := k.db.GetConnection().Prepare("DELETE FROM proyecto1.careers_knowledge WHERE Faculty = @Faculty AND Career = @Career;")
    if err != nil {
        return 0, err
    }

    defer stmt.Close()

    // Execute the prepared statement
    result, err := stmt.Exec(
        sql.Named("Faculty", Faculty),
        sql.Named("Career", Career),
    )
    if err != nil {
        return 0, err
    }

    // Get the number of row inserted
    rowDeleted, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

    return rowDeleted, nil
}

// func main() {
// 	repo := GetKnowledgeDao()

// 	rows, err := repo.UpdateFact("ingenieria", "sistemas", "logica2", "programacion", "tecnologia")
// 	// rows, err := repo.DeleteFact("ingenieria", "sistemas")

// 	if err != nil {
// 		fmt.Printf("Error inserting career: %v\n", err)
// 	}

// 	fmt.Printf("%d Rows inserted\n", rows)

// 	json, _ := repo.GetFacts()

// 	fmt.Println(json)
// }