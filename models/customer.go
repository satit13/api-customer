package models
import (
	"database/sql"
	"fmt"
)
//import "database/sql"


type Customer struct {
	Code string `json:"code`
	Name string `json:"name"`
	Address   string `json:"address"`
}

type customers []*Customer

func (ct *Customer) Show(s string,db *sql.DB) (*Customer, error){

	err := db.QueryRow("select top 1 code,name1,billaddress from bcnp.dbo.bcar where code = ? ",s).Scan(&ct.Code,&ct.Name,&ct.Address)
	fmt.Println("print customer "+ct.Code)


	if err != nil {
		return nil,err
	}
	return ct, nil

}