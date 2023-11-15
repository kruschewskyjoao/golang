package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasPublicacoes = []Rota{
	{
		URI:                 "/publicacoes",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CriarPublicacao,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/publicacoes",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacoes,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacao,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}",
		Metodo:              http.MethodPut,
		Funcao:              controllers.AtualizarPublicacao,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}",
		Metodo:              http.MethodDelete,
		Funcao:              controllers.DeletarPublicacao,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/usuarios/{usuarioId}/publicacoes",
		Metodo:              http.MethodGet,
		Funcao:              controllers.BuscarPublicacoesPorUsuario,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}/curtir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.CurtirPublicacao,
		RequerAtutenticacao: true,
	},
	{
		URI:                 "/publicacoes/{publicacaoId}/descurtir",
		Metodo:              http.MethodPost,
		Funcao:              controllers.DescurtirPublicacao,
		RequerAtutenticacao: true,
	},
}
