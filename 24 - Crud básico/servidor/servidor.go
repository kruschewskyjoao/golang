package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()
	// PREPARE STATEMENT -- evita ataques sql injection
	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?,?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement"))
		return
	}
	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o ID inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso! Id: %d", idInserido)))
}

// tras todos os usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com banco de dados"))
		return
	}
	defer db.Close()
	linhas, erro := db.Query("SELECT * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao buscar usuarios"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario"))
			return
		}
		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuarios para JSON"))
		return
	}
}

// tras um usuario especifico
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuario no banco"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear usuario"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario para JSON"))
		return
	}
}

// altera os dados de um usuario no banco
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}
	defer statement.Close()
	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar o usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// deleta usuario do banco
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com banco de dados"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
		return
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao deletar usuario"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
