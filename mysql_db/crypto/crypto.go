package crypto

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"io"
	"log"
	"strings"

	"github.com/cs3305/group13_2022/project/mysql_db"
	"github.com/cs3305/group13_2022/project/mysql_db/find"

	"golang.org/x/crypto/scrypt"
)

// Function inserts user under the 3 columns `username`, `hash`, `salt` if they don't already exist.
func AddUser(tx *sql.Tx, tableName, username, password string) bool {
	userFound := find.FindRowByValue(tx, tableName, "username", username, "username")
	
	if userFound != "" {
        return false
	}

	username = strings.ToLower(username)	// --- IMPORTANT ---

	hash, salt := MakeHashPassword(password)  

	const query = `INSERT INTO user_credentials
	               VALUES (DEFAULT, ?, ?, ?);
				   `

	// Adding user.
	if _, err := tx.Exec(query, username, hash, salt); err != nil {  
        log.Printf("Error %s when adding user to DB\n", err)
        return false
	}

	// No errors encountered.. user added.
	return true
}

// Checks if provided password under the given name matches the stored hashed password.
func CredentialsMatch(DB *mysql_db.DB, tableName, usernameColumnName, username, password string) bool {
	db := mysql_db.EstablishConnection(DB)
	tx := mysql_db.NewTransaction(db)
	defer tx.Rollback()
	defer db.Close()

	userSalt := find.FindRowByValue(tx, tableName, usernameColumnName, username, "salt")
	userHash := find.FindRowByValue(tx, tableName, usernameColumnName, username, "hash")


	hash := HashPassword(password, userSalt)

	if userHash != hash {
		return false
	}

	return true
}

// Function hashes the provided password with the salt.
func HashPassword(password, salt string) string {
	PW_HASH_BYTES := 64

	byteArraySalt, _ := hex.DecodeString(salt)

	hash, err := scrypt.Key([]byte(password), byteArraySalt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hash)
}

// Function hashes the provided password with a random generated salt
func MakeHashPassword(password string) (string, string) {
	PW_SALT_BYTES := 32
	PW_HASH_BYTES := 64

	salt := make([]byte, PW_SALT_BYTES)

	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		log.Fatal(err)
	}

	// The recommended parameters for interactive logins as of 2017 are N=32768,
	// r=8 and p=1. The parameters N, r, and p should be increased as memory
	// latency and CPU parallelism increases.
	hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hash), hex.EncodeToString(salt)
}