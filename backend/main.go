package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/google/uuid"
)

type Todo struct{
   id string
   text string
   created time.Time
}

type TodoService interface {
    add(todo Todo) (string,error)
    delete(id string) (error)
    getAll() ([]Todo,error)
}

type TodoServer struct {
    service TodoService
}

func NewTodoServer(service TodoService) *TodoServer{
    return &TodoServer{service: service}
}

func (t *TodoServer) addTodo(w http.ResponseWriter, r *http.Request){
    text := r.FormValue("text")
    
    id := uuid.NewString()
    t.service.add(Todo{id: id, text: text, created: time.Now()})

    t.getTodo(w,r)
}

func (t *TodoServer) getTodo(w http.ResponseWriter, r *http.Request){
    todos,err := t.service.getAll()

    if err != nil{
        w.WriteHeader(http.StatusNotFound)
    }
    
    w.Header().Set("Content-Type","text/html;charset=utf-8")
    
    var sb strings.Builder 
    
    for _,todo := range todos {
        fmt.Fprintf(&sb,"<li>%s</li>", todo.text)
    }
    
    fmt.Fprint(w,sb.String())
}

func (t *TodoServer) deleteTodo (w http.ResponseWriter, r *http.Request){
    
}

type SliceService struct{
   todos []Todo 
}

func (s *SliceService) add(todo Todo) (string,error){
    s.todos = append(s.todos, todo)
    return todo.id, nil
}

func (s *SliceService) delete(id string) (error){
    index := -1

    for i, t := range s.todos {
       if t.id == id {
           index = i
       } 
    }

    if index == -1{
        return errors.New("item not found")
    }

    s.todos[index] = s.todos[len(s.todos)-1]
    s.todos = s.todos[:len(s.todos)-1]
    return nil
}

func (s *SliceService) getAll() ([]Todo,error){
    if s.todos != nil{
        return s.todos,nil
    }

    return nil,errors.New("slice not initialized")
}

func main()  {
    service := &SliceService{} 
    service.add(Todo{id: "test",text: "hi",created: time.Now()})
    server := NewTodoServer(service)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
       template := `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <script src="https://unpkg.com/htmx.org@1.9.3" integrity="sha384-lVb3Rd/Ca0AxaoZg5sACe8FJKF0tnUgR2Kd7ehUOG5GCcROv5uBIZsOqovBAcWua" crossorigin="anonymous"></script>
  </head>
  <body>
    <ul id="list" hx-get="/todos" hx-trigger="load"></ul>
    <form hx-post="/todos/add" hx-target="#list">
        <label>What Todo</label>
        <input type="text" name="text">
        <button>submit</button>
    </form>
  </body>
</html>
       `
       fmt.Fprint(w,template) 
    })

    http.HandleFunc("/todos",server.getTodo) 
    http.HandleFunc("/todos/delete",server.deleteTodo)
    http.HandleFunc("/todos/add",server.addTodo)

    log.Fatal(http.ListenAndServe(":8080",nil))
}
