package service

import (
	"reflect"
	"encoding/json"
	"bytes"
	"net/http/httptest"
	"net/http"
	"testing"
	"github.com/unrolled/render"
)

func TestGetKeyHandler(t *testing.T) {
	var list = []string{"user","admin"}
	formatter := render.New(render.Options{
        IndentJSON: true,
    })
	tests := []struct { url string; want struct{ Key string `json:"key"`; Permissions []string `json:"permissions"` }} {
		{"/v1/user/getkey?username=root&password=pass", struct{ Key string `json:"key"`; Permissions []string `json:"permissions"`}{"1e3576bt",list}},
		{"/v1/user/getkey?username=admin&password=admin",struct{ Key string `json:"key"`; Permissions []string `json:"permissions"` }{"",nil}},
	}
	s := http.NewServeMux()
	s.Handle("/",getKeyHandler(formatter))
	for _, test := range tests {
		func() {
			req, _ := http.NewRequest("GET",test.url,nil)
			rw := httptest.NewRecorder()
			rw.Body = new(bytes.Buffer)
			s.ServeHTTP(rw, req)
			var got struct{ Key string `json:"key"`; Permissions []string `json:"permissions"` }
			json.NewDecoder(rw.Body).Decode(&got)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("%s: got %v, want %v", test.url, got, test.want)
			}
		}()
	}
}