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

	router.HandleFunc("/login", Login).Methods(http.MethodPost)
	router.HandleFunc("/logout", Logout).Methods(http.MethodPost)
	router.HandleFunc("/getevents", GetEvents).Methods(http.MethodPost)
	router.HandleFunc("/addevent", AddEvent).Methods(http.MethodPost)
	router.HandleFunc("/getfriends", GetFriends).Methods(http.MethodPost)
	router.HandleFunc("/addfriend", AddFriend).Methods(http.MethodPost)

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

	var u model.JsonLoginRequest
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check for valid login
	// tsql := fmt.Sprintf("SELECT user_id FROM users WHERE username='%s' AND password='%s';", u.Username, u.Password)
	tsql := "SELECT user_id FROM users WHERE username='josh6h' AND password='5ecretPassword';"
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

func GetEvents(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ug := r.Header.Get("userguid")
	tsql := fmt.Sprintf("SELECT UserId FROM Sessions where UserGuid='%s' AND IsActive=1;", ug)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	row := db.QueryRowContext(ctx, tsql)
	var uid int
	if err = row.Scan(&uid); err != nil {
		http.Error(w, "No login found", http.StatusUnauthorized)
		return
	}
	// print(uid)
	defer rows.Close()
	tsql = fmt.Sprintf("SELECT events.description, events.title, events.start_datetime, events.end_datetime FROM user_events INNER JOIN events ON user_events.event_id = events.event_id WHERE user_events.user_id=%d", uid)
	rows, err = db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var event_list []model.Event
	for rows.Next() {
		var event model.Event
		rows.Scan(&event.Title, &event.Description, &event.StartDateTime, &event.EndDateTime)
		event_list = append(event_list, event)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	err = json.NewEncoder(w).Encode(event_list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddEvent(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ug := r.Header.Get("userguid")
	tsql := fmt.Sprintf("SELECT UserId FROM Sessions where UserGuid='%s' AND IsActive=1;", ug)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	row := db.QueryRowContext(ctx, tsql)
	var uid int
	if err = row.Scan(&uid); err != nil {
		http.Error(w, "No login found", http.StatusUnauthorized)
		return
	}

	var e model.NewEventRequest
	err = json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tsql = fmt.Sprintf("INSERT INTO events (title, description, start_datetime, end_datetime) VALUES ('%s', '%s', '%s', '%s')", e.NewEvent.Title, e.NewEvent.Description, e.NewEvent.StartDateTime, e.NewEvent.EndDateTime)
	_, err = db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var eid int
	tsql = "SELECT TOP 1 event_id FROM events Order By event_id DESC"
	row = db.QueryRowContext(ctx, tsql)
	row.Scan(&eid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, v := range e.Users {
		tsql = fmt.Sprintf("INSERT INTO user_events VALUES ('%d', '%d')", v, eid)
		rows, err = db.QueryContext(ctx, tsql)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	defer rows.Close()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonGenericResponse{Type: "Success", Message: ""}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetFriends(w http.ResponseWriter, r *http.Request) {
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

	tsql := fmt.Sprintf("SELECT users.user_id, users.first_name, users.last_name, users.username FROM users INNER JOIN friends ON users.user_id = friends.friends_with_id WHERE friends.user_id = %d", u.UserID)
	rows, err := db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var friends []model.User
	for rows.Next() {
		var usr model.User
		rows.Scan(&usr.UserID, &usr.FirstName, &usr.LastName, &usr.Username)
		friends = append(friends, usr)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
	err = json.NewEncoder(w).Encode(friends)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddFriend(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ug := r.Header.Get("userguid")
	tsql := fmt.Sprintf("SELECT UserId FROM Sessions where UserGuid='%s' AND IsActive=1;", ug)
	_, err = db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	row := db.QueryRowContext(ctx, tsql)
	var uid int
	if err = row.Scan(&uid); err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	var friend_id int
	err = json.NewDecoder(r.Body).Decode(&friend_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tsql = fmt.Sprintf("INSERT INTO friends VALUES (%d, %d), (%d, %d)", uid, friend_id, friend_id, uid)
	_, err = db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonGenericResponse{Type: "Success", Message: ""}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
