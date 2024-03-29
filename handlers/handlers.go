package handler 

import (
	"golang-projects/boltdb_image_server/boltdb"
	"log"
	"os"
	"path/filepath"
	"strconv"
	// "io/ioutil" // ioutil.Readfile
	"github.com/lithammer/shortuuid"
	"golang-projects/boltdb_image_server/aws"
)

var db store.Store
var storePath string

const (
	GET = "get"
	PUT = "new"
	DEL = "delete"
)

func testAws3() {

  a := &aws3.Aws3{}
  a.Bucket = "tassel-app-user-image" 
  a.Region = "us-east-1" 
  name := shortuuid.New()
  // TODO read file
  log.Printf(name)


}

func setup() {

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	storePath = filepath.Join(dir, "/store.db")
	// log.Println(storePath)

	err = db.Init(storePath)
	if err != nil {
		log.Fatal(err)
	}

	// TODO run echo webapp
	todo := os.Args[1]
	if todo == "" {
		log.Fatal("pls enter a command")
	}
	log.Print("command: ", string(todo))

	switch todo {
	case GET:
		getTodos()
	case PUT:
		putTodo(string(todo))
	case DEL:
		{
			// get os.Args[2]
			id, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
			if err := db.DeleteTask(storePath, id); err != nil {
				log.Print(err)
			}
		}
	default:
		// get all todos
		db.AllTasks(storePath)
	}

}

func getTodos() {

	tasks, err := db.AllTasks(storePath)
	// error handling
	if err != nil {
		return
	}
	for _, task := range tasks {
		log.Printf("%d : %s \n", task.Key, task.Value)
	}
}

func putTodo(t string) {

	// stdin argument (./boltdb_todoapp "first todo")
	//task := strings.Join(t, " ")
	id, err := db.CreateTask(storePath, t)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Dope! added %s to your task list with id %d \n", t, id)

}
