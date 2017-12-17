package operation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DeleteUser(id string) {
	res, err := deleteRequest("http://localhost:8080/v1/users/"+id)
	check(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	check(err)

	var retJSON User

	//w, err := json.Marshal(retJSON)
	//check(err)

	if err := json.Unmarshal(body, &retJSON); err != nil {
		panic(err)
	}

	fmt.Println(retJSON)
}
func deleteRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	check(err)
	res, err := client.Do(req)
	check(err)
	return res, err
}
