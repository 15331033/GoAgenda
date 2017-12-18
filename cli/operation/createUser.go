package operation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func CreateUser(key, username, password, email string) User {
	res, err := http.PostForm("http://localhost:8080/v1/users?key="+key,
		url.Values{"UserName": {username}, "Password": {password}, "Email": {email}})
	check(err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	check(err)

	var retJSON User

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
