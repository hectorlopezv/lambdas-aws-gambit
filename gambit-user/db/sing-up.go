package db

import (
	"fmt"
	"gambit-user/models"
	"gambit-user/tools"
)

func SignUp(sign models.SignUp)error{
	fmt.Println("Insertando en la base de datos")

	err := DbConnect()
	if err != nil{
		fmt.Println("Error al conectar a la base de datos")
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sign.UserEmail + "', '" + sign.UserUUID + "', '"+tools.FechaMysql()+"')"
	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println("Error al insertar en la base de datos")
		return err
	}
	return nil
}