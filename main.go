package main
import (
        "fmt"
        "database/sql"
      _ "github.com/lib/pq"
       )

const (
	 host = "localhost"
	 port = 5432
   user = "postgres"
   password = "test123"
   dbname = "postgres"
)

func main() {

  psqlInfo := fmt.Sprintf("host=%s port=%d user = %s "+
  "password =%s dbname=%s sslmode=disable",
 	host, port ,user, password, dbname)

 db, err := sql.Open("postgres", psqlInfo)
 if err != nil{
 	fmt.Println(">>>>>>>", err)
 }
 defer db.Close()

sqlStatement := `INSERT INTO users (age, email, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING id`

  id := 0
  err = db.QueryRow(sqlStatement, 30, "aniket@gmail.com", "aniket", "amin").Scan(&id)
  if err != nil {
    panic(err)
  }
  fmt.Println("New record ID is:", id)
}


