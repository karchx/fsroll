package fs

import (
	"os"
	"regexp"
	"strings"

	log "github.com/gothew/l-og"
	"github.com/karchx/envtoyaml/pkg/utils"
	"gopkg.in/yaml.v3"
)

func check(e error) {
	if e != nil {
		log.Error(e)
		panic(e)
	}
}

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	check(err)

	defer file.Close()

	data, err := utils.IgnoreComments(file)
	check(err)

	return data
}

func CreateFile(extension string, data []byte) {
	content := utils.ParseBytesToString(data)
	mapContent := parseMapString(&content)

	writeContent, err := yaml.Marshal(&mapContent)
	check(err)

	extensionFile, err := utils.CheckExtension(extension)
	check(err)
	fileName := "output" + extensionFile
	os.WriteFile(fileName, writeContent, 0644)
}

func parseMapString(data *string) map[string]interface{} {
	re := regexp.MustCompile(`\r?\n`)
	*data = re.ReplaceAllString(*data, " ")
	keys, values := parseString(data)
	keysValuesMap := make(map[string]interface{})

	for i := 0; i < len(keys); i++ {
		keysValuesMap[keys[i]] = values[i]
	}

	return keysValuesMap
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
