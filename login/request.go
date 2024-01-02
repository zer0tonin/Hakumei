package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

type errorResponse struct {
	Err string `json:"err"`
}

type loginRequest struct {
	target string

	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (l *loginRequest) do() tea.Msg {
	url := l.target + "/api/login"

	requestBody, err := json.Marshal(l)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var responsePayload errorResponse
		if err = json.NewDecoder(resp.Body).Decode(&responsePayload); err != nil {
			return err
		}
		return fmt.Errorf(responsePayload.Err)
	}

	var responsePayload loginResponse
	if err = json.NewDecoder(resp.Body).Decode(&responsePayload); err != nil {
		return err
	}
	return responsePayload
}
