package structs

import(
	// "time"
)

type UserReg struct {

	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	// Img string `json:"image"`
	// ChatHistory []string `json:"chat_history"`
	// TotalConversations int `json:"total_conversations"`
	// Status string `json:"status"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`

}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
