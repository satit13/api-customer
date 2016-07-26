package main

import (

	"fmt"
//	"database/sql"
	"flag"
	"log"
	"net/http"
	"github.com/api-customer/models"
	"github.com/api-customer/controllers"
	"github.com/gorilla/mux"
)
var debug = flag.Bool("debug", false, "enable debugging")
var password = flag.String("password", "[ibdkifu", "the database password")
var port *int = flag.Int("port", 1433, "the database port")
var server = flag.String("server", "olaf", "the database server")
var user = flag.String("user", "sa", "the database user")


func main(){
	flag.Parse() // parse the command line args

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
	}
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	fmt.Println(connString)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}

	db, err := models.NewDB(connString)


	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	c := &controllers.Env{DB: db}
	defer db.Close()



	fmt.Println("main.go api ")


	r := SetupRouter(c)

	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)

}

func SetupRouter(c *controllers.Env) *mux.Router{
	r := mux.NewRouter().StrictSlash(true)

	s := r.PathPrefix("/api/v1/customer").Subrouter()
	s.HandleFunc("/",c.CustomerAll ).Methods("GET")
	s.HandleFunc("/{code}",c.ShowCustomer ).Methods("GET")

	return r

}