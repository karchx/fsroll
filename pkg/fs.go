package pkg

import (
	"os"
	"strings"

	log "github.com/gothew/l-og"
)

func check(e error) {
	if e != nil {
		log.Errorf("Error: %v", e)
		panic(e)
	}
}

func ReadFile(path string) []byte {
	file, err := os.ReadFile(path)
	check(err)

	return file
}

func CreateFile(extension string, data []byte) {
	content := string(data)
	parseMapString(&content)
}

func parseMapString(data *string) { //map[string]interface{} {
	*data = strings.Replace(*data, "\n", " ", -1)
	keys, values := parseString(data)
	log.Infof("Keys: %v", keys)
	log.Infof("valuess: %v", values)
}

func parseString(str *string) ([]string, []string) {
	var strForVector string

	for _, char := range *str {
		if char != '=' {
			strForVector += string(char)
		} else {
			strForVector += " " + string(char) + " "
		}
	}
	vector := strings.Split(strForVector, " ")

	return getKeyOrValue(vector)
}

func getKeyOrValue(vector []string) ([]string, []string) {
	var keys []string
	var values []string
	for i, item := range vector {
		if item == "=" {
			key := vector[i-1]
			value := vector[i+1]
			keys = append(keys, key)
			values = append(values, value)
		}
	}

	return keys, values
}
