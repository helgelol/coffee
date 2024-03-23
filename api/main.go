package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"k8s.io/apimachinery/pkg/util/json"
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
	id       int
	country  string
	region   string
	producer string
	name     string
	process  string
	flavours string
}

func main() {
	var err error
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/all", getAllBeans)
	// mux.HandleFunc("/bean/{id}", getBeanByID)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Acknowledged")
	})

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getAllBeans(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, country, region, producer, name, process, flavours FROM beans")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	beans := []Bean{}
	for rows.Next() {
		var bean Bean
		if err := rows.Scan(&bean.id, &bean.country, &bean.region, &bean.producer, &bean.name, &bean.process, &bean.flavours); err != nil {
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

	// for _, bean := range beans {
	// 	fmt.Fprintf(w, "ID: %d, Name: %s\n", bean.id, bean.name)
	// }
}

// func getBeanByID(w http.ResponseWriter, r *http.Request, id string) {
// 	// id := r.URL.Path[len("/1"):]
// 	var bean Bean
// 	err := db.QueryRow("SELECT id, name FROM beans WHERE id = ?", id).Scan(&bean.id, &bean.name)
// 	switch {
// 	case err == sql.ErrNoRows:
// 		http.NotFound(w, r)
// 		return
// 	case err != nil:
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	default:
// 		fmt.Fprintf(w, "ID: %d, Name: %s\n", bean.id, bean.name)
// 	}
// }
