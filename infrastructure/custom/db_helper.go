package custom

import (
	"fmt"
	"gorm.io/gorm"
)

func Each(closure func(item interface{}) error) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		rows, err := db.Rows()
		if err != nil {
			fmt.Printf("Error fetching rows: %s\n", err.Error())
			return db
		}
		defer rows.Close()

		model := db.Statement.Dest

		for rows.Next() {
			if err := db.ScanRows(rows, model); err != nil {
				fmt.Printf("Error scanning row: %s\n", err.Error())
				continue
			}

			if err := closure(model); err != nil {
				// Handle the error accordingly
				fmt.Printf("Error in closure function: %s\n", err.Error())
				continue
			}
		}

		return db
	}
}
