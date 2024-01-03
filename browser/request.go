package browser

import (
	"encoding/json"
	"fmt"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

type errorResponse struct {
	Err string `json:"err"`
}

type browseRequest struct {
	target string
	path   string
	search string
}

type FileInfos struct {
	IsDir bool   `json:"isDir"`
	Size  int64  `json:"size"`
	Path  string `json:"path"`
}

type browseResponse struct {
	FileInfos []FileInfos `json:"fileInfos"`
}

func (b *browseRequest) do() tea.Msg {
	url := b.target + "/api/browse/" + b.path

	if b.search != "" {
		url = url + "?search=" + b.search
	}

	resp, err := http.Get(url)
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

	var responsePayload browseResponse
	if err = json.NewDecoder(resp.Body).Decode(&responsePayload); err != nil {
		return err
	}
	return responsePayload
}
