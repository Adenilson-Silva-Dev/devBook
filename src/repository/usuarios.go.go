package repository

import (
	"api/src/model"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuario(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

//Criar insere um usuário na banco
func (repositorio Usuarios) Criar(usuario model.Usuarios) (uint64, error) {

	statemant, erro := repositorio.db.Prepare("INSERT INTO usuarios(nome, nick, email, senha) values(?,?,?,?)")

	if erro != nil{
		return 0, nil
	}
	defer statemant.Close()
	resultado, erro:= statemant.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil{
		return 0, nil
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil{
		return 0, nil
	}

	return uint64(ultimoIDInserido), nil
	

}