package lib

import (
	"automation/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type LoginRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Session struct {
	API string
	Username string
	Password string
	Token    string
}

type LoginResponse struct {
	Token string `json:"token"`
}

func ToJSONString(payload interface{}) string {
	js, err := json.Marshal(payload)
	if err != nil {
		return ""
	}
	return string(js)
}

func PrettyStruct(data interface{}) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("\n")
	fmt.Printf("%s\n", string(val))
}

func RestClient(url, token, methodType string, payload string) (string, error) {
	var resp *http.Response
	req, err := http.NewRequest(methodType, url, strings.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Add("X-Custom-Header", "myvalue")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Authorization", token)
	client := &http.Client{Timeout: time.Minute*3 + time.Second*40}
	resp, err = client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error occurred while hitting request : %s", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	text := string(body)
	return text, err
}

func (s *Session)Login()error{
	URL:=types.LogInAPI
	loginPayload := ToJSONString(LoginRequest{Username: s.Username, Password: s.Password})
	rawResponse, err := RestClient(URL, "", "POST", loginPayload)
	fmt.Print("\nLogin url : ", URL)
	fmt.Print("\nUsername : ", s.Username)
	fmt.Print("\nPassword : ", s.Password)
	if err != nil {
		fmt.Print("Error while logging in ...")
		fmt.Print("Response:", rawResponse)
		return err
	}
	var loginResponse LoginResponse
	err = json.Unmarshal([]byte(rawResponse), &loginResponse)
	if err != nil {
		fmt.Print("Error while logging in ...")
		return err
	}
	fmt.Print("\nResponse : ")
	PrettyStruct(loginResponse)
	s.Token=loginResponse.Token
	return nil
}

func (s *Session)ListCats()error{
	URL:=types.ListKittyAPI
	rawResponse, err := RestClient(URL, s.Token, "GET", "")
	if err != nil {
		fmt.Print("\nError while listing cats ...")
		fmt.Print("\nResponse:", rawResponse)
		return err
	}
	var Response types.ListKittyResponse
	err = json.Unmarshal([]byte(rawResponse), &Response)
	if err != nil {
		fmt.Print("\nError while unmarshaling ...")
		return err
	}
	PrettyStruct(Response)
	return nil
}