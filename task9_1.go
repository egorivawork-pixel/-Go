package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const baseURL = "http://localhost:8083"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	return strings.TrimSpace(s)
}

func listUsers() {
	resp, err := http.Get(baseURL + "/users")
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	defer resp.Body.Close()

	var users []User
	_ = json.NewDecoder(resp.Body).Decode(&users)
	fmt.Println("Пользователи:")
	for _, u := range users {
		fmt.Printf("  #%d %s (%d)\n", u.ID, u.Name, u.Age)
	}
}

func getUser() {
	idStr := readLine("id: ")
	id, _ := strconv.Atoi(idStr)

	resp, err := http.Get(fmt.Sprintf("%s/users/%d", baseURL, id))
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println("не найдено")
		return
	}

	var u User
	_ = json.NewDecoder(resp.Body).Decode(&u)
	fmt.Printf("User: #%d %s (%d)\n", u.ID, u.Name, u.Age)
}

func createUser() {
	name := readLine("name: ")
	ageStr := readLine("age: ")
	age, _ := strconv.Atoi(ageStr)

	body, _ := json.Marshal(map[string]any{"name": name, "age": age})
	resp, err := http.Post(baseURL+"/users", "application/json", bytes.NewReader(body))
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	defer resp.Body.Close()

	var u User
	_ = json.NewDecoder(resp.Body).Decode(&u)
	fmt.Println("Создан:", u)
}

func updateUser() {
	idStr := readLine("id: ")
	id, _ := strconv.Atoi(idStr)
	name := readLine("new name: ")
	ageStr := readLine("new age: ")
	age, _ := strconv.Atoi(ageStr)

	body, _ := json.Marshal(map[string]any{"name": name, "age": age})
	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/users/%d", baseURL, id), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println("не найдено")
		return
	}

	var u User
	_ = json.NewDecoder(resp.Body).Decode(&u)
	fmt.Println("Обновлён:", u)
}

func deleteUser() {
	idStr := readLine("id: ")
	id, _ := strconv.Atoi(idStr)

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("%s/users/%d", baseURL, id), nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("ошибка:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		fmt.Println("не найдено")
		return
	}
	if resp.StatusCode == 204 {
		fmt.Println("Удалено.")
		return
	}
	fmt.Println("Статус:", resp.Status)
}

func main() {
	for {
		fmt.Println("\n1) list  2) get  3) create  4) update  5) delete  0) exit")
		choice := readLine("> ")

		switch choice {
		case "1":
			listUsers()
		case "2":
			getUser()
		case "3":
			createUser()
		case "4":
			updateUser()
		case "5":
			deleteUser()
		case "0":
			return
		default:
			fmt.Println("не понял команду")
		}
	}
}
