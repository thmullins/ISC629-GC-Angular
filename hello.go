package main

import (
	"net/http" //handles http requests
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore" //access to Datastore NoSQL DB
)

func main() { //starts execution
	http.HandleFunc("/", helloWorldHandler) //calls handler function for hello world
	http.HandleFunc("/addWords", addWordsHandler) //calls handler function to write to DB
	http.HandleFunc("/getWords", getWordsHandler) //calls handler function to read from DB
	appengine.Main()
}  

//Hello World
//here we print a message to the browser
func helloWorldHandler (w http.ResponseWriter, r *http.Request){ //writes back to browser
	fmt.Fprint(w, "Hello. My GAE go app!")
}

//Write to DB
//here we construct the words to add to the DataStore DB
type Word struct{
	Item string
	Meaning string
}

var words = []Word{
	{"Mobile", "city"},
	{"Alabama", "state"},	
}


func addWordsHandler (w http.ResponseWriter, r *http.Request){ //handler function for adding words
	c := appengine.NewContext(r)

	for _, oneWord := range words{
		datastore.Put(c, datastore.NewIncompleteKey(c,"Word", nil), &oneWord) //call put to write to DB
	}	
}

//Query DB
func getWordsHandler (w http.ResponseWriter, r *http.Request){ //handler function for getting words
	c := appengine.NewContext(r)

	q := datastore.NewQuery("Word")
	var getWords []Word
	q.GetAll(c, &getWords)

	fmt.Fprintf(w, "<html><body>%v</body></html>", getWords)
}