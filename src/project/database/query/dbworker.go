package query

import (
	"fmt"
	"project/database/models"
	"project/database/pgconnectror"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
)

func GETUSER(Id uuid.UUID) *user.User {

	dbInstance := pgconnect.GetInstance()
	if !dbInstance.Connected {
		fmt.Println("error db not open")
		return nil
	}

	u := user.User{}
	return &u
}

func GET() []user.User {

	dbInstance := pgconnect.GetInstance()
	if !dbInstance.Connected {
		fmt.Println("error db not open")
		return nil
	}

	var usersList = []user.User{}
	rows, err := dbInstance.Db.Query("SELECT id,name,lastname,age,birthdate FROM users")
	if err != nil {
		fmt.Println(err)
		return usersList
	} else {
		for rows.Next() {
			var id uuid.UUID
			var name string
			var lastname string
			var age int
			var birthdate int64
			err = rows.Scan(&id, &name, &lastname, &age, &birthdate)
			if err != nil {
				fmt.Println(err)
				break
			}

			unw := user.User{}
			unw.Name = name
			unw.Age = age
			unw.Birthdate = birthdate
			unw.Id = id
			unw.Lastname = lastname

			usersList = append(usersList, unw)
		}
	}

	return usersList
}

func PUT(Id uuid.UUID) user.User {
	u := user.User{}
	return u
}

func DELETE(Id uuid.UUID) bool {

	dbInstance := pgconnect.GetInstance()
	if !dbInstance.Connected {
		fmt.Println("error db not open")
		return false
	}
	stmt, err := dbInstance.Db.Prepare("delete from users where id=$1")
	if err != nil {
		fmt.Println(err)
		return false
	}

	res, err := stmt.Exec(Id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	affect, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(affect, "rows changed")

	return true
}

func POST(u user.User) bool {

	id, err := uuid.NewUUID()

	if err != nil {
		fmt.Println("cannot gen new uuid")
		return false
	}
	u.Id = id
	dbInstance := pgconnect.GetInstance()
	if dbInstance == nil {
		fmt.Println("error db not open")
		return false
	}

	var lastInsertId uuid.UUID
	err = dbInstance.Db.QueryRow("INSERT INTO public.users(id,name,lastname,age,birthdate) VALUES($1,$2,$3,$4,$5) returning id;", u.Id, u.Name, u.Lastname, u.Age, u.Birthdate).Scan(&lastInsertId)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println("last inserted id =", lastInsertId)
		return true
	}
}
