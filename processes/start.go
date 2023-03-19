package processes

import (
	"fmt"
	"os"
	"os/exec"
)

func Start(pkgName string, projectName string) error {

	err := os.Mkdir(pkgName, 0755)
	if err != nil {
		return err
	}

	err = os.Chdir(pkgName)
	if err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", projectName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	file, err := os.Create("main.go")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(`
package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	} else {
		fmt.Println("ENV variables loaded")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	} else {
		fmt.Println("Connected to DB")
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("Cannot create server", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}
	`)
	if err != nil {
		return err
	}

	return nil
}
