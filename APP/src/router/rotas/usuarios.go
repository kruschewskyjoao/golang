package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                 "/usuarios",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarUsuario,
		RequerAtutenticacao: false,
	},
	{
		URI:                 "/usuarios",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarUsuarios,
		RequerAtutenticacao: false,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarUsuario,
		RequerAtutenticacao: false,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarUsuario,
		RequerAtutenticacao: false,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarUsuario,
		RequerAtutenticacao: false,
	},
}
