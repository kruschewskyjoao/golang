package modelos

// DadosAutenticacao contem o token e id do usuario autenticado
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
