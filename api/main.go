package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("POSTGRES_HOSTNAME")
	port     = 5432
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
)

var db *sql.DB

type Bean struct {
	ID   int
	Name string
}

func main() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", getAllBeans)
	http.HandleFunc("/1", getBeanByID)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getAllBeans(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM beans")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	beans := []Bean{}
	for rows.Next() {
		var bean Bean
		if err := rows.Scan(&bean.ID, &bean.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		beans = append(beans, bean)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, bean := range beans {
		fmt.Fprintf(w, "ID: %d, Name: %s\n", bean.ID, bean.Name)
	}
}

func getBeanByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/1"):]
	var bean Bean
	err := db.QueryRow("SELECT id, name FROM beans WHERE id = $1", id).Scan(&bean.ID, &bean.Name)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	default:
		fmt.Fprintf(w, "ID: %d, Name: %s\n", bean.ID, bean.Name)
	}
}
