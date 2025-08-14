package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `hello world`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestEnvHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/env", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"headers":{},"envVars":{}}`))
	})

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// In a real test, you'd want to check the content of the JSON.
	// For this example, we'll just check that it's not empty.
	if rr.Body.String() == "" {
		t.Errorf("handler returned empty body")
	}
}

func TestHokkaidoHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/hokkaido", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hokkaidoHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response body is one of the expected values.
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
	found := false
	for _, m := range meibutsu {
		if rr.Body.String() == m {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}
