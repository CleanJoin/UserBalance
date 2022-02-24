package balance

import "time"

type UserModel struct {
	ID           uint    `json:"id"`
	Username     string  `json:"username"`
	PasswordHash string  `json:"password"`
	Money        float64 `json:"money"`
}

type TransactionsModel struct {
	ID         uint      `json:"id"`
	UserIdTo   uint      `json:"useridto"`
	Money      float64   `json:"money"`
	UserIdFrom uint      `json:"useridfrom"`
	Time       time.Time `json:"time"` //utc
}
