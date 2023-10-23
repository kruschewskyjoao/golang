// teste unitario
package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

// TestXxxxxXxxx
// go test -v     vai ser mais completo, com mais informações
// go test --cover  vai mostrar a cobertura dos testes
// go test --coverprofile cobertura.txt e depois go tool cover --html=cobertura.txt   vai gerar um html dizendo o que não foi testado (se existir)
func TestTipoDeEndereco(t *testing.T) {
	// \/para rodar testes em paralelo. mas tem que por isso em vários testes
	t.Parallel()

	cenariosDeTeste := []cenarioTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Dos Imigrandes", "Rodovia"},
		{"Praça das Rodas", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

	// tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
	//
	//	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	//		t.Error("O tipo de recebido é diferente do esperado!")
	//	}
}
