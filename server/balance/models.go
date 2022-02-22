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
	UserIdTo   uint      `json:"UserIdTo"`
	UserIdFrom uint      `json:"UserIdFrom"`
	Time       time.Time `json:"time"` //utc
}
