package service

import (
	"encoding/json"
	"net/http/httptest"
	"bytes"
	"net/http"
	"testing"
	"github.com/unrolled/render"
)

func TestDeleteUser(t *testing.T) {
	formatter := render.New(render.Options{
        IndentJSON: true,
	})
	tests := []struct{ url string;want string } {
		{"/v1/users/1",""},
	}
	s := http.NewServeMux()
	s.HandleFunc("/",deleteUserHandler(formatter))
	for _, test := range tests {
		func() {
			req, _ := http.NewRequest(http.MethodDelete,test.url,nil)
			rw := httptest.NewRecorder()
			rw.Body = new(bytes.Buffer)
			s.ServeHTTP(rw,req)
			var got string
			json.NewDecoder(rw.Body).Decode(&got)
			if  got != test.want {
				t.Errorf("%s: got %v, want %v", test.url, got, test.want)
			}
		}()
	}
}