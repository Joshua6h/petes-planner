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
	EventID       int32  `json:"event_id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	StartDateTime string `json:"start_datetime"`
	EndDateTime   string `json:"end_datetime"`
	GroupID       string `json:"group_id"`
}

type Friendship struct {
	FriendID      int32 `json:"friend_id"`
	UserID        int32 `json:"user_id"`
	FriendsWithID int32 `json:"friends_with_id"`
}
