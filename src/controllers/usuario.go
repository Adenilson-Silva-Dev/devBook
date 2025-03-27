package controllers

import (
	"api/src/database/db"
	"api/src/model"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// CriarUsuario CRIANDO AS ROTAS DE USUARIOS
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	// se nao estiver dado erro, iremos jogar o corpoRequest para dentro so Struct de Usuario
	var usuario model.Usuarios
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}

	// agora vou abrir a conexão com o banco de dados
	db, erro := db.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}

	//Criando o repositorio e passamos o banco para dentro do  repository.NovoRepositorioDeUsuario(db)
	repositorio := repository.NovoRepositorioDeUsuario(db)

	//passando o metodo Criar() e passando o usuario Criar(usuario)
	usuarioID,erro:= repositorio.Criar(usuario)
	if erro != nil{
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Id inserido: %d", usuarioID)))

}
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar todos os usuários"))
	return
}
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
	return
}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
	return
}
func DeletandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
	return
}
