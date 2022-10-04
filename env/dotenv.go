package env

import "github.com/joho/godotenv"

type MyEnv map[string]string

func NewEnv() (MyEnv, error) {
	myEnv, err := godotenv.Read()
	if err != nil {
		return myEnv, err
	} else {
		return myEnv, nil
	}
}
