package connections

import (
	"errors"
	"fmt"
	"os"
	"webApp/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

//===========================================  DB connection and table creation   =============================================

func ConnectTodb() {
	dsn := os.Getenv("Database")

	fmt.Println("dsn", os.Getenv("Database"))
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("error detected")
		return
	} else {
		fmt.Println("db connected ")
	}
	err = Db.AutoMigrate(
		models.User{},
	)
	if err != nil {
		fmt.Println("err= ", err)
		return
	} else {
		fmt.Println("auto migrtion complited")
	}
}

//===========================================  adding NEW USER and Insertion TO DB   ============================================

func AddNewUser(Firstname, Emailuser, passworduser string) error {
	var user models.User

	//check existing account or not

	Query := `SELECT * FROM users WHERE email = ?`
	err := Db.Raw(Query, Emailuser).Scan(&user).Error
	if err != nil {
		fmt.Println("find a error at Raw ")
		return err
	}
	if user.ID != 0 {
		fmt.Println("user exist")
		return errors.New("user already exists")
	}
	// hashing the password

	HashPassword, err := bcrypt.GenerateFromPassword([]byte(passworduser), 5)
	if err != nil {
		fmt.Println("error")
		return errors.New("hashing faild")
	}

	//if else insert into db
	if user.ID == 0 {
		s := ""
		o := "user__"
		b := true

		Query_ := `INSERT INTO users(name,Email,Password,Status,Ownstatus,Block)
		VALUES ($1,$2,$3,$4,$5,$6)`
		err = Db.Exec(Query_, Firstname, Emailuser, HashPassword, s, o, b).Error //err
		if err != nil {
			return errors.New("error:culdnt insert data")
		}
	}
	fmt.Printf("\n user added successfully\n")
	return nil
}

// ===	IF WANT ID===== (	Query_ := `INSERT INTO users(name,Email,Password)
// 		VALUES ($1,$2,$3) RETURNING ID`
// 		err = db.Raw(Query_, Firstname, Emailuser, passworduser).Scan(&id).Error //err
// 		if err != nil {
// 			return errors.New("error::culdnt insert data")
// 		}
// 	}
// 	// return success
// 	fmt.Printf("\n user added success , user id : %v \n", id)
// 	return nil  )

//===========================================  find user for LOGIN  ===============================================

func FindUser(email_, password string) (bool, int, string, error) {
	var user models.User

	Query := `SELECT * FROM users WHERE email= ?`
	if err = Db.Raw(Query, email_).Scan(&user).Error; err != nil {
		fmt.Println("error detected")
		return false, 0, "", nil
	}

	if user.ID != 0 {
		fmt.Println("get email")
		//verify email and password
		if email_ == user.Email {
			err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
			if err != nil {
				return false, 0, "", nil
			}
			////////==================

			return bool(user.Block), int(user.ID), user.Ownstatus, errors.New(user.Name)
		}
	}
	return false, int(user.ID), "", nil
}
