package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/rs/cors"
)

type MidiRequest struct {
	Fn         string
	NotesTime  []int
	NotesValue []int
}

type LoginRequest struct {
	User     string
	Password string
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var midireq MidiRequest
	err = json.Unmarshal(body, &midireq)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	// TRATAR DA AUTENTICACAO E DOCKER CONTAINERING
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		w.WriteHeader(404)
		return
	}

	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		w.WriteHeader(404)
		return
	}

	tokenString := splitToken[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if _, err := os.Stat(("./static/" + midireq.Fn)); err == nil {
			fmt.Printf("File already exists not generating again.\n")
		} else {

			//cmd := exec.Command("mkdir", "static")
			cmd := exec.Command("cmd.exe", "/c", "mkdir", "static")
			err = cmd.Run()
			if err != nil {
				fmt.Println(err)
			}

			fileNameTime := time.Now().UTC().Unix()
			fileName := strconv.FormatInt(fileNameTime, 10) + ".json"

			file, err := json.MarshalIndent(midireq, "", " ")
			if err != nil {
				fmt.Println(err)
			}
			fileNamePath := fmt.Sprintf("./static/%s", fileName)
			err = ioutil.WriteFile(fileNamePath, file, 0644)

			cmd2 := exec.Command("python", "midi.py", fileNamePath, midireq.Fn)
			log.Printf("Running command and waiting for it to finish...")
			out, err := cmd2.Output()
			if err != nil {
				log.Printf("There was an error executing python midi.py: %s", err)
			} else {
				log.Printf("Command output: %v", string(out))
			}

			e := os.Remove(fileNamePath)
			if e != nil {
				log.Fatal(e)
			}

		}

	} else {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	return

}

func handleLogin(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var loginReq LoginRequest
	err = json.Unmarshal(body, &loginReq)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	if loginReq.User == os.Getenv("SERVER_USER") && loginReq.Password == os.Getenv("SERVER_PW") {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user": loginReq.User,
			"nbf":  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"user": loginReq.User, "token": tokenString})
	} else {
		w.WriteHeader(404)
		return
	}

}

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("ALLOWED_ORIGINS")},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})
	fs := http.FileServer(http.Dir("./static"))

	mux := http.NewServeMux()
	mux.HandleFunc("/midi", handleRequest)
	mux.Handle("/", fs)
	mux.HandleFunc("/login", handleLogin)

	handler := c.Handler(mux)
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), handler))
}
