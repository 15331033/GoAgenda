package service

import (
	//"fmt"
    "net/http"
    //"strconv"

    "GoAgenda/http-api/entities"

    "github.com/unrolled/render"
)

func postUserInfoHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        if len(req.Form["username"][0]) == 0 {
            formatter.JSON(w, http.StatusBadRequest, struct{ ErrorIndo string }{"Bad Input!"})
            return
        }
        u := entities.NewUserInfo(entities.UserInfo{UserName: req.Form["username"][0]})
        u.Password = req.Form["password"][0]
        entities.UserInfoService.CreatUser(u)
        formatter.JSON(w, http.StatusOK, u)
    }
}
// func listAllUserHandler(formatter *render.Render) http.HandlerFunc {
//     return func(w http.ResponseWriter, req *http.Request) {
//         ulist := entities.UserInfoService.FindAllUser()
//         formatter.JSON(w, http.StatusOK, ulist)
//     }
// }