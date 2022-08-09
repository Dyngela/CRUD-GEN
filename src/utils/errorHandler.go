package utils

import (
	"log"
	"net/http"
)

func CheckForError(error error, message string) {
	if error != nil {
		log.Println(message)
		log.Println(error)
	}
}

func CheckForQueryError(err error, message string) (int, error) {
	if err != nil {
		log.Println(message)
		log.Println(err)
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, error(nil)
}
