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

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rBErr := tx.Rollback(); rBErr != nil {
			return errors.New("rollback error")
		}
		return err
	}

	return tx.Commit()
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
