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

	query := fmt.Sprintf(`
    INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('%v' , '%v', '%v')`,
    sign.UserEmail,
    sign.UserUUID,
    tools.FechaMysql(),
)

	fmt.Println(query)
	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println("Error al insertar en la base de datos")
		return err
	}
	return nil
}