package service

import (
	"time"
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/unrolled/render"
	"strings"
	"GoAgenda/http-api/entities"
)

func getKeyHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
			vals := req.URL.Query()
			username, _ := vals["username"]
			password, _ := vals["password"]
			if strings.Join(username, "") == "root" && strings.Join(password, "") == "pass" {
				var list = []string{"user","admin"}
				formatter.JSON(w, http.StatusOK, struct{ Key string `json:"key"`; Permissions []string `json:"permissions"` }{"1e3576bt",list})
			} else {
				formatter.JSON(w, 404, struct{ Key string `json:"key"`; Permissions []string `json:"permissions"` }{"",nil})
			}
    }
}
func listAllUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
			ulist := entities.UserInfoService.FindAllUser()
			formatter.JSON(w, http.StatusOK, ulist)
	}
}
func createUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Println(req.Form)
		if len(req.Form["UserName"][0]) == 0  {
				formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
				return
		}
		u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["UserName"][0]})
		u.Password = req.Form["Password"][0]
		u.Email = req.Form["Email"][0]
		fmt.Println(u)
		uCreate := entities.UserInfoService.CreatUser(u)
		formatter.JSON(w, http.StatusOK, uCreate)
	}
}
func updateUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		defer req.Body.Close()
		body, _ := ioutil.ReadAll(req.Body)
		fmt.Println(string(body))
		var retJSON entities.UserInfo
		if err := json.Unmarshal(body, &retJSON); err != nil {
			panic(err)
		}
		intId, _ := strconv.Atoi(id)
		retJSON.UID = intId
		t := time.Now()
		retJSON.Created = &t
		fmt.Println(retJSON)
		uUpdate := entities.UserInfoService.UpdateUser(&retJSON)
		formatter.JSON(w, http.StatusOK, uUpdate)
	}
}
func deleteUserHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		intId, _ := strconv.Atoi(id)
		err := entities.UserInfoService.DeleteUser(intId)
		if err != nil {
			http.Error(w, "404 not found", 404)
		}
		formatter.JSON(w,http.StatusOK,nil)
	}
}