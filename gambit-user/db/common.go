package db

import (
	"database/sql"
	"gambit-user/models"
	"gambit-user/secretm"
	"os"
	"github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB
func ReadSecret()error{
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}