package main

import (
	"fmt"
	"os"
)

func main() {
	p := fmt.Println
	pf := fmt.Printf
	file, err := os.Open("myFile.txt")
	if err != nil {
		p(err)
		os.Exit(1)
	}
	data := make([]byte, 100)
	c, err := file.Read(data)
	if err != nil {
		p(err)
		os.Exit(1)
	}
	pf("read: %d bytes: %q\n", c, data[:c])

	// Manejo de enviroments

	env := os.Getenv("OS")
	p(env)

	os.Setenv("MY_ENV", "MY VALUE") // Si se pasa una que exista la remplaza
	p(os.Getenv("MY_ENV"))

	dir, _ := os.Getwd()
	p(dir)

	os.Setenv("ENV_EXIST", "")

	env1 := os.Getenv("ENV_EXIST")
	env2 := os.Getenv("ENV_NOT_EXIST")
	p(env1)
	p(env2)

	env3, ok := os.LookupEnv("ENV_EXIST")
	p(env3, " ", ok)

	env4, ok := os.LookupEnv("ENV_NOT_EXIST")
	p(env4, " ", ok)

	os.Setenv("DB_USERNAME", "nahuel")
	os.Setenv("DB_PASSWORD", "sarasa")
	os.Setenv("D8_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "27018")
	os.Setenv("DB_NAME", "users")

	dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@$DB_HOST:$DB_PORT/$DB_NAME")
	p(dbURL)

}
