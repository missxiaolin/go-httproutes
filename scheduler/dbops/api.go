package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ReadVideoDeletetionRecord(count int) ([]string, error) {
	stmIns, err := dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT")

	var ids []string

	
	if err != nil {
		return  ids, err
	}

	rows, err := stmIns.Query(count)

	if err != nil {
		log.Printf("ReadVideoDeletetionRecord error: %v", err)
		return ids, err
	}

	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	defer stmIns.Close()

	return  ids, nil
}
