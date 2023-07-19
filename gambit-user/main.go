package main

import (
	"context"
	"errors"
	"fmt"
	"gambit-user/awsgo"
	"gambit-user/models"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)
func main(){
	lambda.Start(LambdaHandler)
	
	
}

func LambdaHandler(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation)(x interface {}, err error){

	awsgo.InicializaAws()

	//Valida que se hayan pasado los parametros
	if !ValidateParams() {
		fmt.Println("No se pasaron los parametros, deben enviar secret manager")
		err = errors.New("No se pasaron los parametros, deben enviar secret manager")
		return event, err
	}
	var datos models.SignUp

}

func ValidateParams()bool{
	var traeParamaetro bool
	_,traeParamaetro = os.LookupEnv("SecretName")
	return traeParamaetro
}

