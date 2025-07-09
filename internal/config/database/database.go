package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	 _ "github.com/lib/pq"
)

var (
	DBConnection *sql.DB
)

func ConnectDatabase() error {

	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Erro ao buscar variáveis de ambiente: ", err.Error())
	}

	var (
		db_host = os.Getenv("DB_HOST")
		db_port = os.Getenv("DB_PORT")
		db_user = os.Getenv("DB_USER")
		db_password = os.Getenv("DB_PASSWORD")
		db_name = os.Getenv("DB_NAME")
		db_sslMode = os.Getenv("DB_SSL_MODE")
	)

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
									db_host, 
									db_port, 
									db_user, 
									db_password, 
									db_name,
									db_sslMode)

	DBConnection, err = sql.Open("postgres", stringConnection)
	if err != nil {
		return fmt.Errorf("Erro ao se conectar no banco de dados: ", err.Error())
	}

	err = DBConnection.Ping()
	if err != nil {
		return fmt.Errorf("Erro ao testar conexão com o banco de dados: ", err.Error())
	}

	return nil

}

func GetConnectionDatabase() *sql.DB {

	return DBConnection
	
}