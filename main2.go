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

  dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+
  	"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
   db, err := sql.Open("postgres", dbinfo)
   if err != nil{
   	panic(err)
   }
 defer db.Close()
 sql := `insert into users(age, email, first_name, last_name) values($1, $2, $3, $4)
 RETURNING id`

 _,err = db.Exec(sql, 20, "mayank@gmail.com", "mayank", "prajapati")
 //fmt.Println("last inserted id: ",id)

 // sql2 := `delete from users where id = $1`
 // _, err = db.Exec(sql2, 2)
 // if err != nil{
 // 	fmt.Println(err)
 // }

//  sql3 := `update users set first_name=$2, last_name=$3 where id= $1`
//  res, err := db.Exec(sql3, 3, "Ridhdhi", "pujara")
// if err != nil{
// 	panic(err)
// }

// count, er := res.RowsAffected()
// if er != nil{
// 	panic(er)
// }
// fmt.Println(count, "rows affected")
// }


// rows, err := db.Query("select * from users")
// if err != nil {
// 	panic(err)
// }

// for rows.Next(){
// 	var id int
// 	var age int
// 	var email string
// 	var first_name string
// 	var last_name string

// 	err = rows.Scan(&id, &age, &first_name, &last_name, &email)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("\tID | AGE | FIRST NAME | LAST NAME | EMAIL ")
// 	fmt.Printf(" \t%3v | %8v | %6v | %6v | %6v\n", id, age, first_name, last_name, email)
// }



}
