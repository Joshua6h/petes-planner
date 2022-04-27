package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	model "go-api/api/models"
	"log"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var db *sql.DB
var server = "DESKTOP-0C0FDTP"
var port = 1433
var user = "apiuser"
var password = "Api2022!"
var database = "WebStuff"

func main() {
	// Establish db connection
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error connecting: ", err.Error())
	}

	// Handle api requests
	router := mux.NewRouter()

	router.HandleFunc("/login", Login).Methods(http.MethodPost)
	router.HandleFunc("/logout", Logout).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var u model.User
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check for valid login
	tsql := fmt.Sprintf("SELECT UserId FROM Users WHERE Email='%s' AND Password='%s';", u.Email, u.Password)
	row := db.QueryRowContext(ctx, tsql)
	var uid int
	if err = row.Scan(&uid); err != nil {
		http.Error(w, "No login found", http.StatusUnauthorized)
		return
	}

	// Log user in by creating a session
	guid := uuid.New()
	tsql = fmt.Sprintf("INSERT INTO Sessions VALUES (%d, '%s', 1)", uid, guid)
	res, err := db.ExecContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count, err := res.RowsAffected()
	if err != nil || count != 1 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonLoginResponse{Message: "Logged In", Type: "Success", UserGuid: guid.String()}
	err = json.NewEncoder(w).Encode(resp)

	err = json.NewEncoder(w).Encode("(:")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var s model.Session
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if s.UserGuid == "" {
		http.Error(w, "No userguid provided", http.StatusBadRequest)
		return
	}

	tsql := fmt.Sprintf("UPDATE Sessions SET IsActive=0 WHERE UserGuid='%s'", s.UserGuid)
	res, err := db.ExecContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count, err := res.RowsAffected()
	if err != nil || count != 1 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonGenericResponse{Message: "Logged out", Type: "Success"}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
