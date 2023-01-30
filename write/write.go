package write

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/miajio/libx/errors"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GenerateFolder(folder string) error {
	if !FileExists(folder) {
		if err := os.MkdirAll(folder, 0666); err != nil {
			return err
		}
	}
	return nil
}

func handleFileAndFolder(out string) (string, string, error) {
	out = strings.ReplaceAll(out, "\\", "/")
	oa := strings.Split(out, "/")
	ob := make([]string, 0)

	for _, o := range oa {
		if o != "" {
			ob = append(ob, o)
		}
	}

	if len(ob) == 0 {
		return "", "", errors.Newf("file name is empty")
	}

	folder := strings.Join(ob[:len(ob)-1], "/")

	if folder == "" {
		folder, _ = os.Getwd()
	}
	return folder, ob[len(ob)-1], nil
}

// WriterIO
func WriteIO(out string, rd io.Reader) (int64, error) {
	folder, fileName, err := handleFileAndFolder(out)
	if err != nil {
		return 0, err
	}

	if err = GenerateFolder(folder); err != nil {
		return 0, err
	}

	file, err := os.Create(folder + "/" + fileName)
	if err != nil {
		return 0, err
	}

	writer := bufio.NewWriter(file)
	return io.Copy(writer, rd)
}

// WriteByte
func WriteByte(out string, bt []byte) error {
	folder, fileName, err := handleFileAndFolder(out)
	if err != nil {
		return err
	}

	if err = GenerateFolder(folder); err != nil {
		return err
	}
	return ioutil.WriteFile(folder+"/"+fileName, bt, 0666)
}
