package user

type User struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

type ResponseSuccessTest struct {
	ResponseCode string `json:"response_code"`
	Message      string `json:"message"`
	MessageErr   string `json:"message_err"`
	ExternalID   string `json:"external_id"`
}
