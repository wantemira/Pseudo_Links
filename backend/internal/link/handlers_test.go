package link

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

const (
	testLink = "http://gitpub.som/cin-sonic/cin?ysclid=1234567890"

	url       = "http://localhost:8080"
	urlCreate = url + "/link/create"
	urlGet    = url + "/link/get"
	urlDelete = url + "/link/delete"
)

func TestCreate(t *testing.T) {
	jsonBody := struct {
		OriginLink string `json:"origin_link" binding:"required"`
	}{
		OriginLink: testLink,
	}
	jsonData, err := json.Marshal(jsonBody)
	if err != nil {
		t.Errorf("TestCreate: error with marshal json body: %v", err)
	}
	response, err := http.Post(urlCreate, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf("TestCreate: error with send POST-request: %v", err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			t.Errorf("TestCreate: error with close %v", err)
		}
	}()
	if response.StatusCode != http.StatusCreated {
		t.Errorf("TestCreate: status code: %d (want %d)", response.StatusCode, http.StatusOK)
	}
}

func TestGet(t *testing.T) {
	url := fmt.Sprintf("%s?origin_link=%s", urlGet, testLink)
	response, err := http.Get(url)
	if err != nil {
		t.Errorf("TestGet: error with send GET-request: %v", err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			t.Errorf("TestCreate: error with close %v", err)
		}
	}()
	if response.StatusCode != http.StatusOK {
		t.Errorf("TestGet: status code: %d (want %d)", response.StatusCode, http.StatusOK)
	}
}

func TestDelete(t *testing.T) {
	url := fmt.Sprintf("%s?origin_link=%s", urlDelete, testLink)
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Errorf("TestDelete: error with create new request: %v", err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("TestDelete: error with send DELETE-request: %v", err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			t.Errorf("TestCreate: error with close %v", err)
		}
	}()
	if response.StatusCode != http.StatusOK {
		t.Errorf("TestDelete: status code: %d (want %d)", response.StatusCode, http.StatusOK)
	}
}
