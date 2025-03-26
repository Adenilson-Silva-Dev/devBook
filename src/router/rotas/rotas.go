package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(w http.ResponseWriter, r *http.Request)
	RequerAutenticacao bool
}

// Criando funcao configurar, para configurar as rotas

func Configurar(r *mux.Router) *mux.Router {
	//vou criar uma variavel rotas para receber as variaveis rotas que estao dentro do pacote rotas/usuarios

	rotas := rotasUsuarios

	//agora irei iterar em cada rota e passar para o r.HandleFunc

	for _, rota := range rotas {

		// passando os valores das rotas para r.HandleFunc
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	// retornano o r com todas as rotas
	return r

}
