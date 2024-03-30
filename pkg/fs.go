package pkg

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"regexp"
	"strings"

	log "github.com/gothew/l-og"
	"gopkg.in/yaml.v3"
)

var filesExtensions []FilesType = []FilesType{
	{
		Extension: "yaml",
		Output:    ".yaml",
	},
}

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

	data, err := ignoreComments(file)
	log.Infof(string(data))
	check(err)

	return data
}

func CreateFile(extension string, data []byte) {
	content := string(data)
	mapContent := parseMapString(&content)

	writeContent, err := yaml.Marshal(&mapContent)
	check(err)

	extensionFile, err := checkExtension(extension)
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

func checkExtension(extension string) (string, error) {
	var output string
	for _, item := range filesExtensions {
		if item.Extension == extension {
			output = item.Output
		}
	}

	if output == "" {
		return output, errors.New("Extension not found")
	}
	return output, nil
}

func ignoreComments(file io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(file)

	var output bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, "#") {
			output.WriteString(line + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}
