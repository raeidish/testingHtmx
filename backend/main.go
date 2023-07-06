package main

import (
	"database/sql"
	"log"

    _ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

type TodoServer struct {
    db *sql.DB
} 

type Todo struct{
    id int;
    text string;
}

func (t *TodoServer) getAllTodos(c *gin.Context){
    rows, err := t.db.Query("SELECT * FROM todo") 

    if err != nil{
        //send back 500 or maybe 400
    }

    todos := []Todo{}
    
    for rows.Next(){
       todo := Todo{}
       rows.Scan(&todo.id,&todo.text) 
       todos = append(todos,todo)
    }
    
    
}

func main()  {
    db := openDb("postgres://postgres:docker@localhost:5432?sslmode=disable")
    defer db.Close()

    r := gin.Default()

    r.Group("")
}

func setup(db *sql.DB) *gin.Engine{
    r := gin.Default()
    r.GET("/api/getAllTodos", getAllTodos)

    return r
}

func getAllTodos(c *gin.Context){
    
} 

func openDb(connString string) *sql.DB{
    db,err := sql.Open("postgres", connString)

    if err != nil {
        log.Fatal(err)
    }

    return db
}
