package read

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"

	"github.com/miajio/libx/errors"
)

var (
	FileTypeMap = map[FileType]string{
		XML:  "XML",
		JSON: "JSON",
		TEXT: "TEXT",
	}
)

type FileType int

const (
	XML FileType = iota
	JSON
	TEXT
)

func fileTypeDoesNotExist(fileType FileType) error {
	return errors.Newf("This file type does not exist, %s", fileType)
}

func unmarshal(fileType FileType, result []byte, obj any, err error) error {
	if err != nil {
		return err
	}
	_, bl := FileTypeMap[fileType]
	if !bl {
		return fileTypeDoesNotExist(fileType)
	}
	switch fileType {
	case XML:
		return xml.Unmarshal(result, obj)
	case JSON:
		return json.Unmarshal(result, obj)
	default:
		return nil
	}
}

func middle(fileType FileType, obj any, result []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	switch fileType {
	case TEXT:
		return result, err
	default:
		err := unmarshal(fileType, result, obj, err)
		return result, err
	}
}

// ReadFile
func ReadFile(fileType FileType, path string, obj any) ([]byte, error) {
	result, err := ioutil.ReadFile(path)
	return middle(fileType, obj, result, err)
}

// ReadIO
func ReadIO(fileType FileType, rd io.Reader, obj any) ([]byte, error) {
	result, err := ioutil.ReadAll(rd)
	return middle(fileType, obj, result, err)
}

// ReadXml
func ReadXml(path string, obj any) error {
	_, err := ReadFile(XML, path, obj)
	return err
}

// ReadJson
func ReadJson(path string, obj any) error {
	_, err := ReadFile(JSON, path, obj)
	return err
}

// ReadText
func ReadText(path string) (string, error) {
	result, err := ReadFile(TEXT, path, nil)
	return string(result), err
}
