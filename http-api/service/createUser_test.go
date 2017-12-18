package service

import (
	"encoding/json"
	"net/http/httptest"
	"net/http"
	"github.com/unrolled/render"
	"testing"
	"GoAgenda/http-api/entities"
)
// func router(f func(*render.Render) http.HandlerFunc,formatter *render.Render,url string) *mux.Router {
//     router := mux.NewRouter()
//     router.HandleFunc(url, f(formatter)).Methods("POST")
//     return router
// }

func TestCreateUser(t *testing.T) {
	formatter := render.New(render.Options{
        IndentJSON: true,
	})
	tests := []struct{ url string; username string; password string; email string;want struct{UserName string;Password string;Email string}} {
		{"/v1/users?key=1e3576bt&UserName=admin&Password=admin&Email=admin@123.com","admin","admin","admin@123.com",struct{UserName string;Password string;Email string}{"admin","admin","admin@123.com"}},
	}
	s := http.NewServeMux()
	s.Handle("/",createUserHandler(formatter))
	for _, test := range tests {
		func() {
			req, _ := http.NewRequest("POST",test.url,nil)
			rw := httptest.NewRecorder()
			s.ServeHTTP(rw, req)
			var got entities.UserInfo
			json.NewDecoder(rw.Body).Decode(&got)
			if got.UserName != test.want.UserName || got.Password != test.want.Password || got.Email != test.want.Email {
				t.Errorf("%s: got %v, want %v", test.url, got, test.want)
			}
		}()
	}

}