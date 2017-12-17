package service

import (
	"encoding/json"
	"net/http/httptest"
	"bytes"
	"net/http"
	"testing"
	"github.com/unrolled/render"
	"github.com/caijh23/GoAgenda/http-api/entities"
)

func TestUpdateUser(t *testing.T) {
	formatter := render.New(render.Options{
        IndentJSON: true,
	})
	tests := []struct{ url string; username string; password string; email string;want struct{UserName string;Password string;Email string}} {
		{"/v1/users/0","newName","123456","newName@123.com",struct{UserName string;Password string;Email string}{"newName","123456","newName@123.com"}},
	}
	s := http.NewServeMux()
	for _, test := range tests {
		func() {
			var jsonStr = []byte(`{"UserName":"newName","Password":"newpass","Email":"newemail"}`)
			jsonStr = bytes.Replace(jsonStr,[]byte("newName"),[]byte(test.username),-1)
			jsonStr = bytes.Replace(jsonStr,[]byte("newpass"),[]byte(test.password),-1)
			jsonStr = bytes.Replace(jsonStr,[]byte("newemail"),[]byte(test.email),-1)
			req, _ := http.NewRequest(http.MethodPut,test.url,bytes.NewBuffer(jsonStr))
			rw := httptest.NewRecorder()
			rw.Body = new(bytes.Buffer)
			s.HandleFunc("/",updateUserHandler(formatter))
			s.ServeHTTP(rw,req)
			var got entities.UserInfo
			json.NewDecoder(rw.Body).Decode(&got)
			if got.UserName != test.want.UserName || got.Password != test.want.Password || got.Email != test.want.Email {
				t.Errorf("%s: got %v, want %v", test.url, got, test.want)
			}
		}()
	}
}