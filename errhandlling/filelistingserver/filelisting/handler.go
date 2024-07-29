package filelisting

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	gopath := os.Getenv("GOPATH")
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(filepath.Join(gopath, "SRC", path))
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}
