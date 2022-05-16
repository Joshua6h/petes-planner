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
	gomail "gopkg.in/gomail.v2"
)

var db *sql.DB

//var server = "DESKTOP-0C0FDTP" // Josh
var server = "localhost" // Peter
var port = 1433          // Josh
// var port = 49678 // Peter
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

	router.HandleFunc("/getevents", mwCheck(GetEvents)).Methods(http.MethodPost)
	router.HandleFunc("/addevent", mwCheck(AddEvent)).Methods(http.MethodPost)
	router.HandleFunc("/getfriends", mwCheck(GetFriends)).Methods(http.MethodPost)
	router.HandleFunc("/addfriend", mwCheck(AddFriend)).Methods(http.MethodPost)
	router.HandleFunc("/removefriend", mwCheck(RemoveFriend)).Methods(http.MethodPost)
	router.HandleFunc("/login", Login).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/logout", mwCheck(Logout)).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/getprofile", mwCheck(GetProfile)).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/adduser", AddUser).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/sendemail", SendMessage).Methods(http.MethodPost, http.MethodOptions)

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

func GetEvents(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ug := r.Header.Get("Authorization")
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
	defer rows.Close()
	tsql = fmt.Sprintf("SELECT events.event_id, events.title, events.description, events.start_datetime, events.end_datetime FROM user_events INNER JOIN events ON user_events.event_id = events.event_id WHERE user_events.user_id=%d", uid)
	rows, err = db.QueryContext(ctx, tsql)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var event_list []model.Event
	for rows.Next() {
		var event model.Event
		rows.Scan(&event.EventID, &event.Title, &event.Description, &event.StartDateTime, &event.EndDateTime)
		event_list = append(event_list, event)
	}

	for event := range event_list {
		tsql = fmt.Sprintf("SELECT users.first_name, users.last_name FROM user_events INNER JOIN users ON user_events.user_id = users.user_id where user_events.event_id = %d", event_list[event].EventID)
		rows, err = db.QueryContext(ctx, tsql)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var first_name string
			var last_name string
			rows.Scan(&first_name, &last_name)
			name := first_name + " " + last_name
			event_list[event].Friends = append(event_list[event].Friends, name)
		}
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
	resp := model.JsonProfileResponse{Type: "Success", Message: "Retrieved profile data", Data: p}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var u model.JsonAddUserRequest
	err = json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tsql := fmt.Sprintf("INSERT INTO users VALUES ('%s', '%s', '%s', '%s', '%s');", u.Username, u.Password, u.Email, u.FirstName, u.LastName)
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
	resp := model.JsonGenericResponse{Message: "User created", Type: "Success"}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var e model.JsonEmailRequest
	err = json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tsql := fmt.Sprintf("SELECT password FROM users WHERE email='%s';", e.Email)
	row := db.QueryRowContext(ctx, tsql)
	var pass string
	if err = row.Scan(&pass); err != nil {
		http.Error(w, "Password not found", http.StatusInternalServerError)
		return
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", "cooltestemail23@gmail.com")
	msg.SetHeader("To", e.Email)
	msg.SetHeader("Subject", "Pete's Planner Password")
	msg.SetBody("text/html", fmt.Sprintf("<b>Here is your password</b><div>%s</div>", pass))
	//msg.Attach("/home/User/cat.jpg")

	n := gomail.NewDialer("smtp.gmail.com", 587, "cooltestemail23@gmail.com", "PassTest1!")

	// Send the email
	err = n.DialAndSend(msg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := model.JsonGenericResponse{Message: "Email Sent", Type: "Success"}
	err = json.NewEncoder(w).Encode(resp)
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

	ug := r.Header.Get("Authorization")
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

	var e model.Event
	err = json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	e.Friends = append(e.Friends, fmt.Sprintf("%d", uid))
	tsql = fmt.Sprintf("INSERT INTO events (title, description, start_datetime, end_datetime) VALUES ('%s', '%s', '%s', '%s')", e.Title, e.Description, e.StartDateTime, e.EndDateTime)
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
	for _, v := range e.Friends {
		tsql = fmt.Sprintf("INSERT INTO user_events VALUES (%s, %d)", v, eid)
		rows, err = db.QueryContext(ctx, tsql)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			print(7)
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
		print(8)
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
	ug := r.Header.Get("Authorization")
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
	defer rows.Close()

	// var u model.User
	// err = json.NewDecoder(r.Body).Decode(&u)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	tsql = fmt.Sprintf("SELECT users.user_id, users.first_name, users.last_name, users.username FROM users INNER JOIN friends ON users.user_id = friends.friends_with_id WHERE friends.user_id = %d", uid)
	rows, err = db.QueryContext(ctx, tsql)
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

	ug := r.Header.Get("Authorization")
	tsql := fmt.Sprintf("SELECT UserId FROM Sessions WHERE UserGuid='%s' AND IsActive=1;", ug)
	row := db.QueryRowContext(ctx, tsql)
	var uid int
	if err = row.Scan(&uid); err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	var friend model.JsonFriendRequest
	err = json.NewDecoder(r.Body).Decode(&friend)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tsql = fmt.Sprintf("SELECT user_id FROM users WHERE username='%s';", friend.Username)
	row = db.QueryRowContext(ctx, tsql)
	var friend_id int
	if err = row.Scan(&friend_id); err != nil {
		http.Error(w, "Cannot find friend", http.StatusInternalServerError)
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
	resp := model.JsonGenericResponse{Type: "Success", Message: "Added friend"}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func RemoveFriend(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	err := db.PingContext(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ug := r.Header.Get("Authorization")
	tsql := fmt.Sprintf("SELECT UserId FROM Sessions WHERE UserGuid='%s' AND IsActive=1;", ug)
	row := db.QueryRowContext(ctx, tsql)
	var uid int
	if err = row.Scan(&uid); err != nil {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	var friend model.JsonFriendRequest
	err = json.NewDecoder(r.Body).Decode(&friend)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tsql = fmt.Sprintf("SELECT user_id FROM users WHERE username='%s';", friend.Username)
	row = db.QueryRowContext(ctx, tsql)
	var friend_id int
	if err = row.Scan(&friend_id); err != nil {
		http.Error(w, "Cannot find friend", http.StatusInternalServerError)
		return
	}

	tsql = fmt.Sprintf("DELETE FROM friends WHERE (user_id=%d AND friends_with_id=%d) OR (user_id=%d AND friends_with_id=%d);", uid, friend_id, friend_id, uid)
	res, err := db.ExecContext(ctx, tsql)
	if err != nil {
		http.Error(w, "Cannot remove friend", http.StatusInternalServerError)
		return
	}
	count, err := res.RowsAffected()
	if err != nil || count != 1 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := model.JsonGenericResponse{Type: "Success", Message: "Removed friend"}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
