package filedb

import (
	"fmt"
	"itemkart/models"
	"os"
)

type FileDB struct{}

func (fd *FileDB) Create(contact *models.Customer) (interface{}, error) {
	var f *os.File

	f, err := os.Open("contacts.fdb")
	if err != nil {
		f, err = os.Create("contacts.fdb")
		if err != nil {
			return nil, err
		}
	} else {
		defer f.Close()
	}
	return f.WriteString(fmt.Sprintln(contact))
}
