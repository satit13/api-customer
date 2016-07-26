package controllers
import (
	"net/http"
	"log"
	"github.com/api-customer/models"
	"github.com/api-customer/api"
	"fmt"
	"encoding/json"
	"github.com/gorilla/mux"
//	"strconv"

	//"strconv"
)
func (e *Env) ShowCustomer(w http.ResponseWriter, r *http.Request) {
	log.Println("call GET Customer()")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	v := mux.Vars(r)

	fmt.Println(v["code"])
	code := v["code"]
	u := new(models.Customer)
	u.Code = code

	log.Println("Print u.Code"+ u.Code)

//	u := models.Customer{}
	customer, err := u.Show(code,e.DB)
	rs := api.Response{}
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		rs.Status = "500"
		rs.Message = err.Error()
	} else {
		rs.Status = "200"
		rs.Message = "OK"
		rs.Result = customer
	}
	output, err := json.Marshal(rs)
	if err != nil {
		log.Println("Error json.Marshal:", err)
	}
	fmt.Fprintf(w, "%s", string(output))
}


func (e *Env) CustomerAll(w http.ResponseWriter, r *http.Request) {
	log.Println("Call GET : All Customer()")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	u := models.Customer{}
	customer, err := u.All(e.DB)
	rs := api.Response{}
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		rs.Status = "500"
		rs.Message = err.Error()
	} else {
		rs.Status = "200"
		rs.Message = "OK"
		rs.Result = customer
	}
	output, err := json.Marshal(rs)
	if err != nil {
		log.Println("Error json.Marshal:", err)
	}
	fmt.Fprintf(w, "%s", string(output))
}
