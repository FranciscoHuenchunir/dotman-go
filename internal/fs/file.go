package fs

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(file string) {
	_, err := os.Stat(file)

	if os.IsNotExist(err) {
		file, err := os.Create(file)

		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		fmt.Println("archivo creando en: ", file.Name())
	}
}
