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

//var server = "DESKTOP-0C0FDTP" // Josh
var server = "localhost" // Peter
//var port = 1433 // Josh
var port = 49678 // Peter
var user = "apiuser"
var password = "Api2022!"
var database = "petes_planner"

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

	router.HandleFunc("/login", Login).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/logout", mwCheck(Logout)).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/getprofile", mwCheck(GetProfile)).Methods(http.MethodPost, http.MethodOptions)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func mwCheck(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validateUser(r) {
			http.Error(w, "Unauthorized", http.StatusForbidden)
		} else {
			f(w, r)
		}
	}
}

func validateUser(r *http.Request) bool {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		return false
	}
	ug := r.Header.Get("Authorization")
	if ug == "" {
		return false
	}

	tsql := fmt.Sprintf("SELECT SessionId FROM Sessions WHERE UserGuid='%s' AND IsActive=1;", ug)
	row := db.QueryRowContext(ctx, tsql)
	if err != nil {
		return false
	}
	var sid int
	if err = row.Scan(&sid); err != nil {
		return false
	}
	return true
}

func Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var u model.JsonLoginRequest
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check for valid login
	tsql := fmt.Sprintf("SELECT user_id FROM users WHERE username='%s' AND password='%s';", u.Username, u.Password)
	// tsql := "SELECT user_id FROM users WHERE username='josh6h' AND password='5ecretPassword';"
	row := db.QueryRowContext(ctx, tsql)
	var uid int
	if err = row.Scan(&uid); err != nil {
		http.Error(w, "Incorrect login information", http.StatusUnauthorized)
		return
	}

	tsql = fmt.Sprintf("SELECT SessionId FROM Sessions WHERE UserId=%d AND IsActive=1;", uid)
	row = db.QueryRowContext(ctx, tsql)
	var sid int
	if err = row.Scan(&sid); err != nil && err != sql.ErrNoRows {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if sid != 0 {
		http.Error(w, "You are already logged on", http.StatusUnauthorized)
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
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")

	w.WriteHeader(http.StatusOK)
	resp := model.JsonLoginResponse{Message: "Logged In", Type: "Success", UserGuid: guid.String()}
	err = json.NewEncoder(w).Encode(resp)
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

	ug := r.Header.Get("Authorization")

	tsql := fmt.Sprintf("UPDATE Sessions SET IsActive=0 WHERE UserGuid='%s'", ug)
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

func GetProfile(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ug := r.Header.Get("Authorization")

	tsql := fmt.Sprintf(`SELECT u.username, u.first_name, u.last_name, u.email
						FROM users u
						INNER JOIN Sessions s
						ON u.user_id=s.UserId
						WHERE s.UserGuid='%s';`, ug)
	row := db.QueryRowContext(ctx, tsql)
	var p model.Profile
	if err = row.Scan(&p.Username, &p.FirstName, &p.LastName, &p.Email); err != nil {
		http.Error(w, "Cannot find profile", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")

	w.WriteHeader(http.StatusOK)
	resp := model.JsonProfileResponse{Type: "Success", Data: p}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
