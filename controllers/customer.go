package controllers
import (
	"net/http"
	"log"
	"github.com/api-customer/models"
	"github.com/api-customer/api"
	"fmt"
	"encoding/json"
)
func (e *Env) ShowCustomer(w http.ResponseWriter, r *http.Request) {
	log.Println("call GET Customer()")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	u := models.Customer{}
	customer, err := u.Show("41054",e.DB)
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
