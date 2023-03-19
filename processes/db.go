package processes

import "os"

func Db() error {

	err := os.Mkdir("db", 0755)
	if err != nil {
		return err
	}

	err = os.Chdir("db")
	if err != nil {
		return err
	}

	err = os.Mkdir("migration", 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir("query", 0755)
	if err != nil {
		return err
	}

	err = os.Mkdir("sqlc", 0755)
	if err != nil {
		return err
	}

	err = os.Chdir("sqlc")
	if err != nil {
		return err
	}

	file, err := os.Create("store.go")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(`
package db

import "database/sql"

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
	`)
	if err != nil {
		return err
	}

	err = os.Chdir("../..")
	if err != nil {
		return err
	}

	return nil
}
