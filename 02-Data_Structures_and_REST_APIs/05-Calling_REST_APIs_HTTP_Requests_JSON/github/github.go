package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Given a github user login, return name and number of public repos

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	fmt.Println(UserInfo(ctx, "dancoe77"))
}
func demo() {
	resp, err := http.Get("https://api.github.com/users/dancoe77")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("ERROR: bad status - %s\n", resp.Status)
		return
	}

	ctype := resp.Header.Get("Content-Type")
	fmt.Println("content-type:", ctype)

	// io.Copy(os.Stdout, resp.Body)
	var reply struct {
		Name     string
		NumRepos int `json:"public_repos"`
		//Public_Repos int
	}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&reply); err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println(reply.Name, reply.NumRepos)
}

// UserInfo return name and number of public repos from GitHub API.
func UserInfo(ctx context.Context, login string) (string, int, error) {
	url := "https://api.github.com/users/" + login
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", 0, err
	}
	// resp, err := http.Get("https://api.github.com/users/dancoe77")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("%q - bad status: %s\n", url, resp.Status)
	}
	return parseResponse(resp.Body)
}

func parseResponse(r io.Reader) (string, int, error) {
	var reply struct {
		Name     string
		NumRepos int `json:"public_repos"`
		//Public_Repos int
	}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&reply); err != nil {
		return "", 0, err
	}
	return reply.Name, reply.NumRepos, nil
}

/*
JSON <-> Go

Types
string <-> string
true/false <-> gool
number <-> float64, float32, int, int8 ... int64, uint, uint8 ...
array <-> []T, []any
object <-> map[string]any, struct

encoding/json API
JSON -> []byte -> Go: Unmarshal
Go -> []byte -> JSON: Marshal
JSON -> io.Reader -> Go: Decoder
Go -> io.Writer -> JSON: Encoder
*/
