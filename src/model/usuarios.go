package model

import "time"

// Iremos criar um struct deu usuarios e aprovaitar e ja colocar os campos JSON

//^omitempty faz com que o campo não apareça no JSON se estiver vazio

type Usuarios struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json: nick, omitempty`
	Email    string    `json:"email, omitempty"`
	CriadoEm time.Time `json:"criadoEm, omitempty"`
}
