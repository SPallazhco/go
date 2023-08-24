package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	base = "https://reqres.in/api"
)

func main() {
	p := fmt.Println
	req, err := GetReqExample("https://reqres.in/api/users/2")
	if err != nil {
		log.Fatal(err)
	}
	p(string(req))

	p()

	req1, err := Get("https://reqres.in/api/users/2")
	if err != nil {
		log.Fatal(err)
	}
	p(string(req1))

	p()

	user, err := GetUser("2")
	if err != nil {
		log.Fatal(err)
	}
	p("ID: ", user.ID)
	p("Email: ", user.Email)
	p("FirstName: ", user.FirstName)
	p("LastName: ", user.LastName)
	p()

	user, err = GetUser("2")
	p(err)

	u, err := CreateUser("Sergio", "engineer")
	if err != nil {
		log.Fatal(err)
	}
	p()

	p(u)
	p("ID: ", u.ID)
	p("Name: ", u.Name)
	p("Job: ", u.Job)
	p("Date: ", u.CreateAt)

	v, _ := json.Marshal(u)
	p()
	p(string(v))
}

func GetReqExample(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// Esta manera es mas correcta
func Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{} //Estoy creando un puntero
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}

type Data struct {
	User User `json:"data"`
}
type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserCreated struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Job      string     `json:"job,omitempty"` //Lo omite si viene en cero
	Age      string     `json:"age,omitempty"`
	CreateAt *time.Time `json:"createAt"` // Puntero para cuando no llege un valor sea cero
	// CreateAt time.Time `json:"createAt"`
}

func GetUser(userID string) (*User, error) {
	req1, err := Get(fmt.Sprintf("%s/users/%s", base, userID))
	if err != nil {
		return nil, err
	}
	var dataResponse Data

	if err := json.Unmarshal(req1, &dataResponse); err != nil {
		return nil, err
	}

	return &dataResponse.User, nil
}

func Post(url string, data interface{}) ([]byte, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	//Crear un nuevo request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("statuscode %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	responseBoby, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBoby, nil
}

func CreateUser(name, job string) (*UserCreated, error) {
	user := &UserCreated{
		Name: name,
		Job:  job,
	}
	req1, err := Post(fmt.Sprintf("%s/users", base), user)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(req1, user); err != nil {
		return nil, err
	}
	return user, nil
}
