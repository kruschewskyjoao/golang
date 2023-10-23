package enderecos

import "strings"

// verifica se um endereço tem tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoLetraMinuscula, " ")[0]
	enderecoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			enderecoValido = true
		}
	}
	if enderecoValido {
		return strings.Title(primeiraPalavra)
	}
	return "Tipo Inválido"
}
