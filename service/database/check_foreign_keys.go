package database

import "fmt"

func (db *appdbimpl) CheckForeignKeys() error {
	query := "PRAGMA foreign_keys;"
	row := db.c.QueryRow(query)

	var result int
	if err := row.Scan(&result); err != nil {
		fmt.Printf("%v", err)
		return err
	}

	fmt.Print(result)
	return nil
}
