package filelisting

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("path must stats " + "with " + prefix)
	}
	path := request.URL.Path[len(prefix):]
	file, err := os.Open(path)
	err = errors.New("错误")
	err.Error()
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	//fmt.Println(all)
	writer.Write(all)
	return nil
}
