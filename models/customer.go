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
	Point  int64
}

type customers []*Customer

func (ct *Customer) Show(s string,db *sql.DB) (*Customer, error){

	err := db.QueryRow("select top 1 code,name1,billaddress,isnull(cast(sumofmark1 as int),0) as point from pos.dbo.bcar where code = ? ",s).Scan(&ct.Code,&ct.Name,&ct.Address,&ct.Point)
	fmt.Println("print customer "+string(ct.Point))


	if err != nil {
		return nil,err
	}
	return ct, nil

}

func (ct *Customer) All(db *sql.DB) ([]*Customer, error){
	rows,err := db.Query("select top 200 code,name1,billaddress,cast(isnull(sumofmark1,0) as int) as point from pos.dbo.bcar order by sumofmark1 desc ")

	if err != nil {
		return nil,err
	}
	defer rows.Close()
	var cts customers

	for rows.Next() {
		var c = new(Customer)
		rows.Scan(
			&c.Code,
			&c.Name,
			&c.Address,
			&c.Point,
		)

		cts =  append(cts, c)


	}
	return cts,nil
}