package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/eylulkadioglu/Music/appconfig"
	"github.com/eylulkadioglu/Music/models"
	"github.com/eylulkadioglu/Music/salt"
	_ "github.com/lib/pq"
)

var dB *sql.DB

func InitDB() {
	config := appconfig.ReadConfig()
	db, err := sql.Open("postgres", config.DbDSN)
	if err != nil {
		fmt.Println("Cannot connect to database!: ", err)
		os.Exit(-1)
	}

	dB = db
}

func GetArtists() ([]models.Artist, error) {

	artists := []models.Artist{}

	sql := "SELECT * FROM artist"
	rows, err := dB.Query(sql)
	if err != nil {
		return artists, err
	}
	defer rows.Close()

	var artist models.Artist
	for rows.Next() {
		rows.Scan(&artist.Id, &artist.Name)
		artists = append(artists, artist)
	}

	return artists, nil
}

func CreateArtist(artist models.Artist) error {
	sql := fmt.Sprintf("INSERT INTO artist(artist_name) VALUES('%s')", artist.Name)

	_, err := dB.Query(sql)
	if err != nil {
		return err
	}

	return nil
}

func CheckLogin(loginRequest models.User) (bool, models.User) {
	var dbUser models.User
	var id int

	saltedPasswordFromUser := fmt.Sprintf("%s###%s", salt.GetSalt(), loginRequest.Password)

	sql := fmt.Sprintf("SELECT * FROM users WHERE user_email='%s' AND user_password=encode(digest('%s', 'sha256'), 'hex')", loginRequest.Email, saltedPasswordFromUser)
	err := dB.QueryRow(sql).Scan(&id, &dbUser.Email, &dbUser.Password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return false, dbUser
	}

	return true, dbUser
}

func DeleteArtist(artist models.Artist) error {
	sql := fmt.Sprintf("DELETE FROM artist WHERE artist_id = '%d'", artist.Id)

	_, err := dB.Query(sql)
	if err != nil {
		return err
	}

	return nil
}

func CreateUser(user models.User) error {
	saltedPassword := fmt.Sprintf("%s###%s", salt.GetSalt(), user.Password)
	sql := fmt.Sprintf("INSERT INTO users(user_email, user_password) VALUES('%s', encode(digest('%s', 'sha256'), 'hex'))", user.Email, saltedPassword)

	_, err := dB.Query(sql)
	if err != nil {
		return err
	}

	return nil
}

func CheckUser(user models.User) error {
	userEmail := user.Email
	sql := fmt.Sprintf("SELECT * FROM users WHERE user_email='%s'", userEmail)

	_, err := dB.Query(sql)
	if err != nil {
		return err
	}

	return nil
}

func CreatePasswordCode(user models.User, code string) error {
	DeletePasswordCode(user)

	sql := fmt.Sprintf("INSERT INTO password_codes VALUES((SELECT user_id FROM users WHERE user_email='%s'), '%s' )", user.Email, code)

	_, err := dB.Query(sql)
	if err != nil {
		return err
	}

	return nil
}

func DeletePasswordCode(user models.User) error {
	sql := fmt.Sprintf("DELETE FROM password_codes WHERE user_id=(SELECT user_id FROM users WHERE user_email='%s')", user.Email)

	_, err := dB.Query(sql)
	if err != nil {
		return err
	}

	return nil
}
func CheckCode(user models.User, code string) error {
	sql := fmt.Sprintf("SELECT code FROM password_codes WHERE user_id=(SELECT user_id FROM users WHERE user_email=%s)", user.Email)

	var dbCode string
	err := dB.QueryRow(sql).Scan(&dbCode)
	if err != nil {
		return err
	}

	if dbCode != code {
		return errors.New("provided code is not correct")
	}

	return nil
}
func ChangePasswordWithCode(user models.User) error {
	saltedPassword := fmt.Sprintf("%s###%s", salt.GetSalt(), user.Password)
	//sql := fmt.Sprintf("INSERT INTO users(user_email, user_password) VALUES('%s', encode(digest('%s', 'sha256'), 'hex'))", user.Email, saltedPassword)

	sql := fmt.Sprintf("UPDATE users SET user_password=encode(digest('%s', 'sha256'), 'hex') WHERE user_email='%s'", saltedPassword, user.Email)

	_, err := dB.Query(sql)
	if err != nil {
		return err
	}

	DeletePasswordCode(user)
	return nil
}
