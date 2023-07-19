package db

import (
	"gambit-user/models"
	"gambit-user/secretm"
	"os"
)

var SecretModel models.SecretRDSJson
var err error
func ReadSecret()error{
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}