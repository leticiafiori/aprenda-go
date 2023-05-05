package entities

import "time"

type Account struct {
	id         int
	name       string
	cpf        string
	secret     string  //hash
	balance    float64 //tem que ser maior que 0
	created_at time.Time
}
