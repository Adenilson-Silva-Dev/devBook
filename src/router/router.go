package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Criando funçao gerar que ira retorna um *mux.Router
func Gerar() *mux.Router {

// r retorna uma rota	
	r := mux.NewRouter()

//passando para nossa rotas.Configurar o a nova rota "r"	
	return rotas.Configurar(r)
}
