package operation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Key struct {
	MyKey       string `json:"key"`
	Permissions []string
}

func GetKey() Key {
	res, err := http.Get("http://localhost:8080/v1/user/getkey?username=root&password=pass")
	check(err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	check(err)

	retJSON := Key{"w", []string{"w", "w"}}

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
