package spotconvenience

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func saveReaderToNewFile(reader io.Reader, fileName string) error {
	newFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file with path %s: %s", fileName, err)
	}
	io.Copy(newFile, reader) // copy the reader to the writer

	newFile.Close() // don't defer since there's nothing in between and defer has a performance cost
	return nil
}

func getLastSplit(str string, delimiter string) string {
	str_split := strings.Split(str, delimiter)
	return str_split[len(str_split)-1]
}

func NiceJsonFormat(object interface{}) string {
	jsonBytes, err := json.MarshalIndent(object, "", "    ")
	if err != nil {
		return ""
	} else {
		return string(jsonBytes)
	}
}
