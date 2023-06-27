package processes

import "os"

func Sqlc() error {

	file, err := os.Create("sqlc.yaml")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(`
version: "2"
sql:
- schema: "./db/migration/"
  queries: "./db/query/"
  engine: "postgresql"
  gen:
    go: 
      package: "db"
      out: "./db/sqlc"
      emit_json_tags: true
      emit_prepared_queries: false
      emit_interface: false
      emit_exact_table_names: false
      emit_empty_slices: true
`)
	if err != nil {
		return err
	}

	return nil
}
