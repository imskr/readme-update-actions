package helpers

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// create new error for empty env
var ErrEnvEmpty = errors.New("getenv: Environment variable empty")

// Get env string from user through actions env
// if no env variables provided throw error
// requires key string
// returns string, error
func GetEnvString(key string) (string, error) {
	env_string := os.Getenv(key)
	if env_string == "" {
		return env_string, ErrEnvEmpty
	}
	return env_string, nil
}

// Get env string from user through actions env and convert to integer
// if no env variables provided throw error
// requires key string
// returns int, error
func GetEnvInteger(key string) (int, error) {
	strEnv, err := GetEnvString(key)
	if err != nil {
		return 0, err
	}

	intEnv, err := strconv.Atoi(strEnv)
	if err != nil {
		return 0, err
	}
	return intEnv, nil
}

// Read a file and replace strings
// requires path of file, resultString
// returns error
func ReplaceFile(path string, resultString string) error {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "<!-- BLOG-LIST-START -->") {
			lines[i] = resultString
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(path, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
