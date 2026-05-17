package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://api.github.com/users/%s/events"

func FetchGithubActivity(username string)([]Event, error){
	url:=fmt.Sprintf(baseURL,username)

	req,err:=http.NewRequest(http.MethodGet,url,nil)
	if err!=nil{
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "github-activity")

	resp, err:= http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("network error: %w", err)
	}

	defer resp.Body.Close()
	
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user '%s' not found", username)
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status: %s", resp.Status)
	}

	var events []Event
	if err:=json.NewDecoder(resp.Body).Decode(&events); err!=nil{
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return events, nil
}