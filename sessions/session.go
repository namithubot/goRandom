//Have created a dummy server to test sessions library : sanad

package main

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("secret_key"))

func MyHandler(w http.ResponseWriter, r *http.Request) {
	//Get a session. Get() will always return a session, even if empty.
	// We are ignoring errors
	session, _ := store.Get(r, "session0")
	//set some session values
	session.Values["foo"] = "bar"
	session.Values["namit"] = "kumar"
	//save the session value
	session.Save(r, w)
}
func showHandler(w http.ResponseWriter, r *http.Request) {
	//Fetching session0 created earlier
	session, _ := store.Get(r, "session0")
	//printing value of the key "namit" stored in the cookie during previous session
	for key, value := range session.Values {
		fmt.Fprintf(w, "\nThe Values for %s is = %s", key, value)
	}
}
func main() {
	http.HandleFunc("/add", MyHandler)
	http.HandleFunc("/show", showHandler)
	fmt.Println("Server listening on port 8080")
	//Used when we are not using gorilla's routing library to prevent memory leak
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux)
}
