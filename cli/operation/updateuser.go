package operation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type UpdateMessage struct {
	UserName string
	Password string
	Email string
}

func UpdateUser(id, username, password, email string) User {
	updateMessage := UpdateMessage{UserName:username, Password:password, Email:email}
	jsonStr,_ := json.Marshal(updateMessage)
	//var jsonStr = []byte(`{"UserName":"newName","Password":"123456","Email":"newName@123.com"}`)
	res, err := putRequest("http://localhost:8080/v1/users/"+id, bytes.NewBuffer(jsonStr))
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
func putRequest(url string, data io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, data)
	check(err)
	res, err := client.Do(req)
	check(err)
	return res, err

}
