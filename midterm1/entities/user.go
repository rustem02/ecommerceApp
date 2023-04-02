package entities

//структура пользователя

type User struct {
	Id       uint16 `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Pass     string `json:"pass"`
	Confpass string `json:"confpass"`
}

type Product struct {
	Id                                     uint16
	Title, Model, Price, Rating, CommentId string
}
