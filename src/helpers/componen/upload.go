package componen

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

func UploadData(file string, path string, validator string) (string, error) {

	data := strings.Split(file, ",")
	if len(data) != 2 {
		return file, nil
	}
	cektypeFile := strings.Split(data[0], "/")
	fmt.Println(cektypeFile)
	ty := strings.Split(cektypeFile[1], ";")
	nameRandom := uuid.New()
	fileDataName := nameRandom.String() + "-" + time.Now().Month().String() + "." + ty[0]
	filename := path + "/" + fileDataName

	dec, err := base64.StdEncoding.DecodeString(data[1])
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}

	if _, err := f.Write(dec); err != nil {
		return "", err

	}
	if err := f.Sync(); err != nil {
		return "", err
	}

	mtype, _ := mimetype.DetectFile(filename)
	fmt.Println(mtype.String(), mtype.Extension())

	cekValidator := ArrayValidate(mtype.Extension(), validator)
	if !cekValidator {
		return "", errors.New("type file yang di upload salah silah kan unggah file dengan type " + validator)
	}

	return "/" + filename, nil
}
func ArrayValidate(condisi string, arr string) bool {
	fmt.Println(arr)
	ar := strings.Split(arr, ",")
	for _, b := range ar {
		if b == condisi {
			return true
		}
	}
	return false
}
