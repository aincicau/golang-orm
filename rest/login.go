package rest

import (
	"encoding/json"
	"fmt"
	"golangorm/entity"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func Login(rw http.ResponseWriter, r *http.Request) {
	body := r.Body
	bodyBytes, _ := ioutil.ReadAll(body)
	var credentials entity.Credential
	json.Unmarshal(bodyBytes, &credentials)

	session, _ := store.Get(r, "cookie-name")
	session.ID = credentials.ID
	session.Values["userID"] = credentials.ID
	session.Values["authenticated"] = true
	fmt.Fprintln(rw, "Logged in")
	session.Save(r, rw)
}

func Welcome(rw http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	auth, ok := session.Values["authenticated"].(bool)

	if !ok || !auth {
		http.Error(rw, "Access forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintln(rw, "Welcome")
}

func Logout(rw http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok {
		fmt.Println("couldn't convert type")
		return
	}

	if auth {
		session.Values["authenticated"] = false
		session.Save(r, rw)
	}
}
