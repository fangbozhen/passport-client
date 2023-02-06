package passport_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	url = "localhost:3000"
)

func QscLogin(qscid string, password string) (cookie string, err error) {
	client := &http.Client{}
	data := make(map[string]interface{})
	data["qscid"] = qscid
	data["password"] = password
	bytesData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", url+"/qsc/login", bytes.NewReader(bytesData))
	if err != nil {
		return "", err
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	cookies := resp.Cookies()
	return cookies[0].Value, nil
}

func GetProfile(token string) (user *User, err error) {
	type Data struct {
		User    User `json:"User"`
		Logined bool `json:"logined"`
	}
	var val struct {
		Err  string `json:"err"`
		Code int    `json:"code"`
		Data Data   `json:"data"`
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url+"/profile", nil)
	if err != nil {
		return nil, err
	}
	cookie := http.Cookie{
		Name:  "SessionToken",
		Value: token,
	}
	req.AddCookie(&cookie)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
	err = json.Unmarshal(body, &val)
	if err != nil {
		return nil, err
	}
	return &val.Data.User, nil
}
