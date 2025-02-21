package healthcheck

import (
	"errors"
	"fmt"
	"webapp/db"
)

func Check() error {
	// Get the MySQL connection
	con, err := db.GetMySQLConn()
	if err != nil {
		// Handle the error
		fmt.Println(err)
		return errors.New("error connecting to MySQL")
	}
	_, err = con.Exec("INSERT INTO webapp (datetime) VALUES(now())")
	if err != nil {
		fmt.Println(err)
		return errors.New("error inserting into healthcheck")
	}
	// Close the connection
	con.Close()
	return nil
}
