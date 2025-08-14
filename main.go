package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type EnvResponse struct {
	Headers map[string][]string `json:"headers"`
	EnvVars map[string]string   `json:"envVars"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	r.Get("/env", func(w http.ResponseWriter, r *http.Request) {
		headers := make(map[string][]string)
		for k, v := range r.Header {
			headers[k] = v
		}

		envVars := make(map[string]string)
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			if len(pair) == 2 {
				envVars[pair[0]] = pair[1]
			} else {
				envVars[pair[0]] = ""
			}
		}

		response := EnvResponse{
			Headers: headers,
			EnvVars: envVars,
		}

		w.Header().Set("Content-Type", "application/json")
		b, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(b)
	})
	r.Get("/hokkaido", hokkaidoHandler)
	r.Get("/fukuoka", fukuokaHandler)

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func hokkaidoHandler(w http.ResponseWriter, r *http.Request) {
	meibutsu := []string{
		"ジンギスカン",
		"札幌ラーメン",
		"函館ラーメン",
		"旭川ラーメン",
		"スープカレー",
		"石狩鍋",
		"ちゃんちゃん焼き",
		"うに丼",
		"いくら丼",
		"豚丼",
		"ザンギ",
		"白い恋人",
		"ロイズのチョコレート",
		"夕張メロン",
		"花畑牧場の生キャラメル",
		"六花亭のマルセイバターサンド",
		"とうもろこし",
		"じゃがいも",
		"カニ",
		"ホッケの開き",
		"松前漬け",
		"いかめし",
	}
	w.Write([]byte(meibutsu[rand.Intn(len(meibutsu))]))
}

func fukuokaHandler(w http.ResponseWriter, r *http.Request) {
	meibutsu := []string{
		"博多ラーメン",
		"もつ鍋",
		"水炊き",
		"明太子",
		"通りもん",
	}
	w.Write([]byte(meibutsu[rand.Intn(len(meibutsu))]))
}
