package secretm

import (
	"encoding/json"
	"fmt"
	"gambit-user/awsgo"
	"gambit-user/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)



func GetSecret(nombreSecret string)(models.SecretRDSJson, error){

	var datosSecret models.SecretRDSJson
	var err error
	fmt.Println(" > GetSecret"+nombreSecret)

	svc:= secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println("Error al obtener el secret manager: "+err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("Secret obtenido: "+nombreSecret)
	return datosSecret, nil
}