package funcs
import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"

	"structs"
	"db"
	"helpers"

	"github.com/zenazn/goji/web"
)

// func setupResponse(w http.ResponseWriter, req *http.Request) {
// 	(w).Header().Set("Access-Control-Allow-Origin", "*")
//     (w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
//     (w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

func Registration(c web.C, w http.ResponseWriter, r *http.Request){
	// setupResponse(w,r)
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

func Login(c web.C, w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading login body->", err)
	}

	var l structs.Login
	var res = make(map[string]string)

	err = json.Unmarshal(body, &l)
	if err != nil {
		fmt.Println("error unmarshaling login body", err)
	}

	Conn := db.DBops(helpers.Database, helpers.UsersCollection)
	err = Conn.Find(&structs.Login{
		Email : l.Email,
	}).One(&res)
	fmt.Println("res--->", res)
	
	if err != nil {
		fmt.Println("error fetching data ->", err)
	}else{
		fmt.Println("Fetched")
	}
}



