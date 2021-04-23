package user

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	CreatedAt time.Time
	Status    bool
}

func (this *Usuario) Signup(id int, nombre string, createdAt time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.CreatedAt = createdAt
	this.Status = status
}
