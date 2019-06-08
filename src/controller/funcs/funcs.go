package funcs
import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"

	"structs"
	"db"
	"helpers"
)

func Registration(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error in reading registration details ->", err)
	} 
	var reg structs.UserReg

	err = json.Unmarshal(body, &reg)
	if err != nil {
		fmt.Println("Error unmarshaling reg body ->", err)
	}
	Conn := db.DBops(helpers.Database, helpers.UsersCollection)
	err = Conn.Insert(&structs.UserReg{
		FirstName : reg.FirstName,
		LastName  : reg.LastName, 
		Emain     : reg.Emain,
		Password  : reg.Password,
	})
	if err != nil {
		fmt.Println("Error inserting data ->", err)
	}else{
		fmt.Println("Inserted")
	}
}

