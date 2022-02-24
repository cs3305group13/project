package env

import (
	"log"

	"github.com/joho/godotenv"
)

// Retrieves env variables from file named `.env`
//
// param:
// path parameter requires relative path from caller to the env file.
//
//
// Env variables are retrieved from `map[string]string` like follows:
//         env := GetEnvironmentVariables("../../.env")
//         hostname := env["HOSTNAME"]
func GetEnvironmentVariables(path string) map[string]string {
	envVariables, err := godotenv.Read(path)
	if err != nil {
		if path != ".env" {
			return GetEnvironmentVariables(path[3:]) // cuts off the starting `../` of the path
		} else {
			log.Fatal("could not find .env file through path: " + path)
		}
	}

	return envVariables
}
