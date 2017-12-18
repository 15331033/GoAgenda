package operation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type User struct {
	UID      int `json:"uid"`
	UserName string
	Password string
	Email    string
	Created  *time.Time
}

func ListAllUser(key string) []User {
	res, err := http.Get("http://localhost:8080/v1/users?key="+key)
	check(err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	check(err)

	var retJSON []User

	//w, err := json.Marshal(retJSON)
	//check(err)

	//fmt.Println(string(w), retJSON)

	//fmt.Println(string(body))

	if err := json.Unmarshal(body, &retJSON); err != nil {
		panic(err)
	}

	fmt.Println(retJSON)
	return retJSON
}
