package rest

import (
	"encoding/json"
	"fmt"
	"golangorm/db"
	"golangorm/entity"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func PostPerson(rw http.ResponseWriter, r *http.Request) {
	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)
	if hasError(rw, err, "Internal issue") {
		return
	}

	var person = entity.Person{}
	err = json.Unmarshal(bodyBytes, &person)
	if hasError(rw, err, "Internal issue") {
		return
	}

	db.GetDB().Create(&person)

	fmt.Println(person)
	rw.Write(bodyBytes)
}

func GetPerson(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("id")
	var person entity.Person
	result := db.GetDB().Find(&person, "id=?", name)
	if result.RecordNotFound() {
		http.Error(rw, "No Record Found", http.StatusInternalServerError)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	personData, _ := json.Marshal(person)
	rw.Write(personData)
}

func hasError(rw http.ResponseWriter, err error, message string) bool {
	logger := new(logrus.Entry)
	if err != nil {
		logger.WithError(err).Error(message)
		rw.Write([]byte(message))
		return true
	}
	return false
}

func DeletePerson(rw http.ResponseWriter, r *http.Request) {
	idValue := r.URL.Query().Get("id")
	result := db.GetDB().Delete(&entity.Person{}, "id=?", idValue)
	if result.Error != nil {
		http.Error(rw, "Internal Error. Please try again after some time", http.StatusInternalServerError)
		return
	}

	rw.Write([]byte("Record successfully deleted"))
}
