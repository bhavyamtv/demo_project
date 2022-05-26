package models

import (
	"demo_project/config"
	"log"
)

func ContactUs(Data map[string]string) (int, error) {

	stmt, err := config.DB_READ.Prepare(
		`INSERT INTO contact_us (name,email,mobile,msg) VALUES (?,?,?,?)`)
	if err != nil {
		log.Println("Error 1")
		log.Println(err)
		return 0, err
	}
	_, err = stmt.Exec(Data["name"], Data["email"], Data["mobile"], Data["msg"])
	if err != nil {
		log.Println("Error 2")
		log.Println(err)
		return 0, err
	}
	stmt.Close()

	return 1, nil
}
