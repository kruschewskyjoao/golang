package modelos

// DadosAutenticacao contem o id e token do usuario
type DadosAutenticacao struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
