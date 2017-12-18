package entity

import (
	"unsafe"
	"reflect"
	//"io/ioutil"
	"fmt"
	//"runtime"
	"encoding/json"
	//"os"
)
type storage struct{
	m_userList []User
	m_meetingList []Meeting
}
type jsonData struct {
	UserList string
	MeetingList string
}
var m_dirty bool
var sStorage storage

func GetStorage() *storage{
	return &sStorage
}
func init() {
	sStorage.readFromFile()
	//runtime.SetFinalizer(&Storage,Storage.writeToFile())
}

func (s *storage) MarshalJSON() ([]byte, error) {
	var userlist string
	var meetinglist string
	if s.m_userList == nil {
		userlist = "null"
	} else {
		
	}
	if s.m_meetingList == nil {
		meetinglist = "null"
	} else {
		m, _ := json.Marshal(s.m_meetingList)
		meetinglist = string(m)
	}
	jsonResult := fmt.Sprintf(`{"UserList":%s, "MeetingList":%s}`, userlist, meetinglist)
	return []byte(jsonResult),nil
    // return json.Marshal(jsonData{
	// 	stringUser(s.m_userList),
	// 	stringMeeting(s.m_meetingList),
    // })
}
func stringUser(b []User) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}
func stringMeeting(b []Meeting) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}
func sliceUser(s string) (u []User) {
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&u))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = 2 * pstring.Len
	return
}
func sliceMeeting(s string) (u []Meeting) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&u))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = 2 * pstring.Len
	return
}
func (s *storage) UnmarshalJSON(b []byte) error {
    temp := &jsonData{}

    if err := json.Unmarshal(b, &temp); err != nil {
        return err
	}
	s.m_userList = sliceUser(temp.UserList)
	s.m_meetingList = sliceMeeting(temp.MeetingList)

    return nil
}
func (s *storage) CreatUser(t_user User) {
	s.m_userList = append(s.m_userList,t_user)
	// n := len(s.m_userList)
	// c := cap(s.m_userList)
	// newSlice := make([]User,n,c + 1)
	// copy(newSlice,s.m_userList)
	// newSlice = append(newSlice,t_user)
	// s.m_userList = newSlice
	m_dirty = true
}
func (s *storage) QueryUser(filter func(t_user User) bool) []User{
	var temp []User
	for _,value := range s.m_userList {
		if filter(value) {
			temp = append(temp,value)
		}
	}
	return temp
}
func (s *storage) UpdateUser(filter func(t_user User) bool, switcher func(t_user *User)) int {
	count := 0
	for i := 0; i < len(s.m_userList); i += 1 {
		if filter(s.m_userList[i]) {
			switcher(&s.m_userList[i])
			count++
			m_dirty = true
		}
	}
	/*for _,value := range s.m_userList {
		if filter(value) {
			switcher(value)
			count++
			s.m_dirty = true
		}
	}*/
	return count
}
func (s *storage) DeleteUser(filter func(t_user User) bool) int{
	count := 0
	for i := len(s.m_userList) - 1;i >= 0;i-- {
		if filter(s.m_userList[i]) {
			s.m_userList = append(s.m_userList[:i],s.m_userList[i + 1:]...)
			count++
			m_dirty = true
		}
	}
	return count
}
func (s *storage) CreatMeeting(t_meeting Meeting) {
	s.m_meetingList = append(s.m_meetingList,t_meeting)
	m_dirty = true
}
func (s *storage) QueryMeeting(filter func(t_meeting Meeting) bool) []Meeting{
	var temp []Meeting
	for _,value := range s.m_meetingList {
		if filter(value) {
			temp = append(temp,value)
		}
	}
	return temp
}
func (s *storage) UpdateMeeting(filter func(t_meeting Meeting) bool, switcher func(t_meeting *Meeting)) int {
	count := 0
	for i := 0; i < len(s.m_meetingList); i += 1 {
		if filter(s.m_meetingList[i]) {
			switcher(&s.m_meetingList[i])
			count++
			m_dirty = true
		}
	}
	return count
}
func (s *storage) DeleteMeeting(filter func(t_meeting Meeting) bool) {
	count := 0
	for i := len(s.m_meetingList) - 1;i >= 0;i-- {
		if (filter(s.m_meetingList[i])) {
			s.m_meetingList = append(s.m_meetingList[:i],s.m_meetingList[i+1:]...)
			count++
			m_dirty = true
		}
	}
}
func (s *storage) readFromFile() bool {
	/*inputStream, inputErr := os.Open("storage.json")
	if inputErr != nil {
		fmt.Println("openFileError : ",inputErr)
		return false
	}
	defer inputStream.Close()
	str, readErr := ioutil.ReadAll(inputStream)
	if readErr != nil {
		fmt.Println("readFileErr : ",readErr)
		return false
	}
	inputString := string(str)
	if inputString == "" {
		return true;
	}
	err := s.UnmarshalJSON([]byte(inputString))
	if err != nil {
		fmt.Println("read : JsonTranslateError : ",err)
		return false
	}*/
	return true
}
func (s *storage) writeToFile() bool {
	/*jsonObj, err := s.MarshalJSON()
	if err != nil {
		fmt.Println("write : JsonTranslateError : ",err)
		return false
	}
	outputStream, outputErr := os.OpenFile("storage.json",os.O_WRONLY|os.O_CREATE, 0666)
	if outputErr != nil {
		fmt.Println("openFileErr : ",outputErr)
		return false
	}
	defer outputStream.Close()
	writeErr := ioutil.WriteFile("storage.json",jsonObj,0x644)
	if writeErr != nil {
		fmt.Println("writeFileErr : ",writeErr)
		return false
	}*/
	return true
}
func (s *storage) Sync() bool {
	if m_dirty {
		s.writeToFile();
		m_dirty = false
		return true
	}
	return false
}