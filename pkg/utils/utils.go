package utils

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strings"

	"github.com/karchx/envtoyaml/pkg/models"
)

var FilesExtensions []models.FilesType = []models.FilesType{
	{
		Extension: "yaml",
		Output:    ".yaml",
	},
}

func CheckExtension(extension string) (string, error) {
	var output string
	for _, item := range FilesExtensions {
		if item.Extension == extension {
			output = item.Output
		}
	}

	if output == "" {
		return output, errors.New("Extension not found")
	}
	return output, nil
}

func IgnoreComments(file io.Reader) ([]byte, error) {
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

// ParseBytesToString for performance to create string from slice bytes
func ParseBytesToString(data []byte) string {
	bytesBuffer := bytes.NewBuffer(data)

	return bytesBuffer.String()
}
