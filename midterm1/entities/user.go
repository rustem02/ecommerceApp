package entities

//структура пользователя

type User struct {
	Id                                    uint16
	Name, Email, Username, Pass, Confpass string
}
