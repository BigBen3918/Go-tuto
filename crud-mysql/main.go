package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
)

type Employee struct {
	Id   int
	Name string
	City string
}

func dbConn() (db *sql.DB) {
	dbDriver := viper.GetString("database.driver")
	dbUser := viper.GetString("database.username")
	dbPass := viper.GetString("database.password")
	dbName := viper.GetString("database.dbName")

	db, err := sql.Open(
		dbDriver,
		dbUser+":"+dbPass+"@/"+dbName,
	)

	if err != nil {
		panic(err.Error())
	}

	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")

	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}
	res := []Employee{}

	for selDB.Next() {
		var id int
		var name, city string

		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city

		res = append(res, emp)
	}

	tmpl.ExecuteTemplate(w, "Index", res)

	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM Emplyee WHERE id=?", nId)

	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}

	for selDB.Next() {
		var id int
		var name, city string

		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city
	}

	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")

	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)

	if err != nil {
		panic(err.Error())
	}

	emp := Employee{}

	for selDB.Next() {
		var id int
		var name, city string

		err = selDB.Scan(&id, &name, &city)
		if err != nil {
			panic(err.Error())
		}

		emp.Id = id
		emp.Name = name
		emp.City = city
	}

	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		insForm, err := db.Prepare("INSERT INTO Employee(name, city) VALUE (?, ?)")

		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(name, city)
		log.Println("INSERT: Name: " + name + " | City: " + city)
	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")

		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")

	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	delForm.Exec(emp)
	log.Println("DELETE")

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)

	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insrt", Insert)
	http.HandleFunc("/delete", Delete)

	err := viper.ReadInConfig()

	if err != nil {
		panic(err.Error())
	}

	log.Println("Server started on PORT 8000")
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
