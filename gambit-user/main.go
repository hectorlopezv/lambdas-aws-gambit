package main

import (
	"context"
	"errors"
	"fmt"
	"gambit-user/awsgo"
	"gambit-user/db"
	"gambit-user/models"
	"os"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)
func main(){
	lambda.Start(LambdaHandler)
}

func LambdaHandler(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation)(  events.CognitoEventUserPoolsPostConfirmation, error){

	awsgo.InicializaAws()
	var err error

	//Valida que se hayan pasado los parametros
	if !ValidateParams() {
		fmt.Println("No se pasaron los parametros, deben enviar secret manager")
		err = errors.New("no se pasaron los parametros, deben enviar secret manager")
		return event, err
	}
	var datos models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = "+ datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("sub = "+ datos.UserUUID)
		}
	}

	err = db.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer el secret manager")
		return event, err
	}
	 err = db.SignUp(datos)
	return event, err

}

func ValidateParams()bool{
	var traeParamaetro bool
	_,traeParamaetro = os.LookupEnv("SecretName")
	return traeParamaetro
}

