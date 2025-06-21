package knowledgedao

import (
	"fmt"

	"database/sql"

	"unimatchserver/data/db"
)

type careerDto struct {
    CareerId int
	Faculty  string
	Career   string
	Aptitudes []string
	Skills    []string
	Interests []string
}

type career struct {
    CareerId int
    Faculty string
    Career string
}

type KnowledgeDao interface {
	GetFacts() ([]careerDto, error)
	AddFact(careerId int, Aptitude, Skill, Interest []string) (int64, error)
    AddCareer(Faculty, Career string) (int64, error)
	DeleteFact(careerId int, Aptitude, Skill, Interest string) (int64, error)
	UpdateFact(careerId int, Aptitude, Skill, Interest, PAptitude, PSkill, PInterest []string) (int64, error)
    DeleteCareer(careerId int) (int64, error)
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
    db := k.db.GetConnection()
	defer db.Close()

	// Step 1: Get all careers
	rows, err := db.Query(`
		SELECT CareerId, Faculty, Career
		FROM proyecto1.careers
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch careers: %w", err)
	}
	defer rows.Close()

	careerMap := make(map[int]*careerDto)
	var careerList []careerDto

	for rows.Next() {
		var c careerDto
		if err := rows.Scan(&c.CareerId, &c.Faculty, &c.Career); err != nil {
			return nil, err
		}
		careerMap[c.CareerId] = &c
		careerList = append(careerList, c)
	}

	// Step 2: Get all aptitudes and map them
	rows, err = db.Query(`
		SELECT CareerId, Aptitude FROM proyecto1.aptitude
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch aptitudes: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var careerId int
		var aptitude string
		if err := rows.Scan(&careerId, &aptitude); err != nil {
			return nil, err
		}
		if c, ok := careerMap[careerId]; ok {
			c.Aptitudes = append(c.Aptitudes, aptitude)
		}
	}

	// Step 3: Get all skills and map them
	rows, err = db.Query(`
		SELECT CareerId, Skill FROM proyecto1.skill
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch skills: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var careerId int
		var skill string
		if err := rows.Scan(&careerId, &skill); err != nil {
			return nil, err
		}
		if c, ok := careerMap[careerId]; ok {
			c.Skills = append(c.Skills, skill)
		}
	}

	// Step 4: Get all interests and map them
	rows, err = db.Query(`
		SELECT CareerId, Interest FROM proyecto1.interest
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch interests: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var careerId int
		var interest string
		if err := rows.Scan(&careerId, &interest); err != nil {
			return nil, err
		}
		if c, ok := careerMap[careerId]; ok {
			c.Interests = append(c.Interests, interest)
		}
	}

	// Optional: refresh careerList to point to updated values
	for i, c := range careerList {
		if updated, ok := careerMap[c.CareerId]; ok {
			careerList[i] = *updated
		}
	}

	return careerList, nil
}

func (k knowledgeDaoImpl) UpdateFact(careerId int, Aptitude, Skill, Interest, PAptitude, PSkill, PInterest []string) (int64, error) {
    conn := k.db.GetConnection()
    defer conn.Close()

    for i, aptitude := range Aptitude {
        stmt, err := conn.Prepare("UPDATE proyecto1.aptitude SET Aptitude = @Aptitude, WHERE CareerId = @CareerId AND Aptitude = @PAptitude;")
        if err != nil {
            return 0, err
        }

        defer stmt.Close()

        // Execute the prepared statement
        result, err := stmt.Exec(
            sql.Named("CareerId", careerId),
            sql.Named("Aptitude", aptitude),
            sql.Named("PAptitude", PAptitude[i]),
        )

        if err != nil {
            return 0, err
        }

        // Get the number of row inserted
        _, err = result.RowsAffected()
        if err != nil {
            return 0, err
        }
    }

    for i, skill := range Skill {
        stmt, err := conn.Prepare("UPDATE proyecto1.skill SET Skill = @Skill, WHERE CareerId = @CareerId AND Skill = @PSkill;")
        if err != nil {
            return 0, err
        }

        defer stmt.Close()

        // Execute the prepared statement
        result, err := stmt.Exec(
            sql.Named("CareerId", careerId),
            sql.Named("Skill", skill),
            sql.Named("PSkill", PSkill[i]),
        )

        if err != nil {
            return 0, err
        }

        // Get the number of row inserted
        _, err = result.RowsAffected()
        if err != nil {
            return 0, err
        }
    }

    for i, interest := range Interest {
        stmt, err := conn.Prepare("UPDATE proyecto1.interest SET Interest = @Interest, WHERE CareerId = @CareerId AND Interest = @PInterest;")
        if err != nil {
            return 0, err
        }

        defer stmt.Close()

        // Execute the prepared statement
        result, err := stmt.Exec(
            sql.Named("CareerId", careerId),
            sql.Named("Interest", interest),
            sql.Named("PInterest", PInterest[i]),
        )

        if err != nil {
            return 0, err
        }

        // Get the number of row inserted
        _, err = result.RowsAffected()
        if err != nil {
            return 0, err
        }
    }

    return 1, nil
}

func (k knowledgeDaoImpl) AddFact(careerId int, Aptitude, Skill, Interest []string) (int64, error) {

    conn := k.db.GetConnection()
    defer conn.Close()

    for aptitude := range Aptitude {
        stmt, err := conn.Prepare("INSERT INTO proyecto1.aptitude(CareerId, Aptitude) VALUES (@CareerId, @Aptitude);")
        if err != nil {
            return 0, err
        }

        defer stmt.Close()

        // Execute the prepared statement
        result, err := stmt.Exec(
            sql.Named("CareerId", careerId),
            sql.Named("Aptitude", aptitude),
        )
        if err != nil {
            return 0, err
        }

        // Get the number of row inserted
        _, err = result.RowsAffected()
        if err != nil {
            return 0, err
        }
    }

    for skill := range Skill {
        stmt, err := conn.Prepare("INSERT INTO proyecto1.aptitude(CareerId, Skill) VALUES (@CareerId, @Skill);")
        if err != nil {
            return 0, err
        }

        defer stmt.Close()

        // Execute the prepared statement
        result, err := stmt.Exec(
            sql.Named("CareerId", careerId),
            sql.Named("Skill", skill),
        )
        if err != nil {
            return 0, err
        }

        // Get the number of row inserted
        _, err = result.RowsAffected()
        if err != nil {
            return 0, err
        }
    }


    for interest := range Interest {
        stmt, err := conn.Prepare("INSERT INTO proyecto1.aptitude(CareerId, Interest) VALUES (@CareerId, @Interest);")
        if err != nil {
            return 0, err
        }

        defer stmt.Close()

        // Execute the prepared statement
        result, err := stmt.Exec(
            sql.Named("CareerId", careerId),
            sql.Named("Interest", interest),
        )
        if err != nil {
            return 0, err
        }

        // Get the number of row inserted
        _, err = result.RowsAffected()
        if err != nil {
            return 0, err
        }
    }

    return 1, nil
}

func (k knowledgeDaoImpl) DeleteCareer(careerId int) (int64, error) {
    conn := k.db.GetConnection()
    conn.Close()

    stmt, err := conn.Prepare("DELETE FROM proyecto1.aptitude WHERE CareerId = @CareerId;")
    if err != nil {
        return 0, err
    }

    defer stmt.Close()

    result, err := stmt.Exec(sql.Named("CareerId", careerId))

    if err != nil {
        return -1, err
    }

    // Get the number of row inserted
    _, err = result.RowsAffected()

    if err != nil {
        return -1, err
    }

    stmt, err = conn.Prepare("DELETE FROM proyecto1.skill WHERE CareerId = @CareerId;")
    if err != nil {
        return 0, err
    }

    defer stmt.Close()

    result, err = stmt.Exec(sql.Named("CareerId", careerId))

    if err != nil {
        return -1, err
    }

    // Get the number of row inserted
    _, err = result.RowsAffected()

    if err != nil {
        return -1, err
    }


    stmt, err = conn.Prepare("DELETE FROM proyecto1.interest WHERE CareerId = @CareerId;")
    if err != nil {
        return 0, err
    }

    defer stmt.Close()

    result, err = stmt.Exec(sql.Named("CareerId", careerId))

    if err != nil {
        return -1, err
    }

    // Get the number of row inserted
    _, err = result.RowsAffected()

    if err != nil {
        return -1, err
    }

    return 1, nil

    stmt, err = conn.Prepare("DELETE FROM proyecto1.careers WHERE CareerId = @CareerId;")
    if err != nil {
        return 0, err
    }

    defer stmt.Close()

    result, err = stmt.Exec(sql.Named("CareerId", careerId))

    if err != nil {
        return -1, err
    }

    // Get the number of row inserted
    _, err = result.RowsAffected()

    if err != nil {
        return -1, err
    }

    return 1, nil
}

func (k knowledgeDaoImpl) DeleteFact(careerId int, Aptitude, Skill, Interest string) (int64, error) {

    conn := k.db.GetConnection()
    conn.Close()

    if Aptitude != "" {
        stmt, err := conn.Prepare("DELETE FROM proyecto1.aptitude WHERE CareerId = @CareerId AND Aptitude = @Aptitude;")
        if err != nil {
            return 0, err
        }

        defer stmt.Close()

        result, err := stmt.Exec(
            sql.Named("CareerId", careerId),
            sql.Named("Aptitude", Aptitude),
        )

        if err != nil {
            return -1, err
        }

        // Get the number of row inserted
        _, err = result.RowsAffected()

        if err != nil {
            return -1, err
        }   
    }

    if Skill != "" {
        stmt, err := conn.Prepare("DELETE FROM proyecto1.skill WHERE CareerId = @CareerId AND Skill = @Skill;")
        if err != nil {
            return 0, err
        }

        defer stmt.Close()

        result, err := stmt.Exec(
            sql.Named("CareerId", careerId),
            sql.Named("Skill", Skill),
        )

        if err != nil {
            return -1, err
        }

        // Get the number of row inserted
        _, err = result.RowsAffected()

        if err != nil {
            return -1, err
        }
    }

    if Interest != "" {
        stmt, err := conn.Prepare("DELETE FROM proyecto1.interest WHERE CareerId = @CareerId AND Interest = @Interest;")
        if err != nil {
            return 0, err
        }

        defer stmt.Close()

        result, err := stmt.Exec(
            sql.Named("CareerId", careerId),
            sql.Named("Interest", Interest),
        )

        if err != nil {
            return -1, err
        }

        // Get the number of row inserted
        _, err = result.RowsAffected()

        if err != nil {
            return -1, err
        }
    }

    return 1, nil
}

func (k knowledgeDaoImpl) AddCareer(Faculty, Career string) (int64, error) {
    conn := k.db.GetConnection()
    stmt, err := conn.Prepare("INSERT INTO proyecto1.careers(Faculty, Career) VALUES (@Faculty, @Career);")
    if err != nil {
        return -1, err
    }

    defer conn.Close()
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
    rowInserted, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

    return rowInserted, nil
}

// func main() {
// 	fmt.Println(GetKnowledgeDao().GetFacts())

// 	// rows, err := repo.UpdateFact("ingenieria", "sistemas", "logica", "programacion", "tecnologia")
// 	// // rows, err := repo.DeleteFact("ingenieria", "sistemas")

// 	// if err != nil {
// 	// 	fmt.Printf("Error inserting career: %v\n", err)
// 	// }

// 	// fmt.Printf("%d Rows inserted\n", rows)

// 	// json, _ := repo.GetFacts()

// 	// fmt.Println(json)
// }