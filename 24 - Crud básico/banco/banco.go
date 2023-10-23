package banco

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //driver conexão com mysql
)

// Conectar abre conexão com banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		fmt.Println("quebrou no open")
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		fmt.Println("quebrou no ping")
		return nil, erro
	}

	return db, nil
}
