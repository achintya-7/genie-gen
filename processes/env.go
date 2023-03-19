package processes

import "os"

func Env() error {
	file, err := os.Create("app.env")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(`
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@localhost:5432/DB_NAME?sslmode=disable
HTTP_SERVER_ADDRESS=0.0.0.0:8080
	`)

	return nil
}
