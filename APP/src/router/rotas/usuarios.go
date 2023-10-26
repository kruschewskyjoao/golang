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
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarUsuario,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarUsuario,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarUsuario,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/seguir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.SeguirUsuario,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.PararDeSeguirUsuario,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/seguidores",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarSeguidores,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/seguindo",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarSeguindo,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:              http.MethodPost,
		Funcao:              controllers.AtualizarSenha,
		RequerAtutenticacao: true,
	},
}
