package model

type User struct {
	UserID    int32  `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Event struct {
	EventID       int32    `json:"event_id"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	StartDateTime string   `json:"start_datetime"`
	EndDateTime   string   `json:"end_datetime"`
	GroupID       string   `json:"group_id"`
	Friends       []string `json:"friends"`
}

type NewEventRequest struct {
	NewEvent Event `json:"event"`
	Users    []int `json:"users"`
}

type Friendship struct {
	FriendID      int32 `json:"friend_id"`
	UserID        int32 `json:"user_id"`
	FriendsWithID int32 `json:"friends_with_id"`
}

type JsonGenericResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type JsonLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JsonLoginResponse struct {
	Type     string `json:"type"`
	Message  string `json:"message"`
	UserGuid string `json:"userguid"`
}

type JsonAddUserRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type JsonProfileResponse struct {
	Type    string  `json:"type"`
	Message string  `json:"message"`
	Data    Profile `json:"data"`
}

type JsonEmailRequest struct {
	Email string `json:"email"`
}

type Profile struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type Session struct {
	SessionId int    `json:"sessionid"`
	UserId    int    `json:"userid"`
	UserGuid  string `json:"userguid"`
	IsActive  bool   `json:"isactive"`
}
