package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	port     = os.Getenv("MYSQL_PORT")
	user     = os.Getenv("MYSQL_USER")
	password = os.Getenv("MYSQL_PASSWORD")
	host     = os.Getenv("MYSQL_HOST")
	dbname   = os.Getenv("MYSQL_DATABASE")
)

var db *sql.DB

type Bean struct {
	ID       int    `json:"id"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	Producer string `json:"producer"`
	Name     string `json:"name"`
	Process  string `json:"process"`
	Flavours string `json:"flavours"`
}

func main() {
	var err error
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := http.NewServeMux()
	router.HandleFunc("GET /all", getAllBeans)
	router.HandleFunc("GET /bean/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		getBeanByID(w, r, id)
	})
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Acknowledged")
	})

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
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
		if err := rows.Scan(&bean.ID, &bean.Country, &bean.Region, &bean.Producer, &bean.Name, &bean.Process, &bean.Flavours); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		beans = append(beans, bean)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(beans)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func getBeanByID(w http.ResponseWriter, r *http.Request, id string) {
	query := fmt.Sprintf("SELECT * FROM beans WHERE id = '%s'", id)
	var bean Bean
	err := db.QueryRow(query).Scan(&bean.ID, &bean.Country, &bean.Region, &bean.Producer, &bean.Name, &bean.Process, &bean.Flavours)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	default:
		jsonData, err := json.Marshal(bean)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}
