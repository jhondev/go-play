package ssllabs

import (
	"encoding/json"
	"io"
	"net/http"
)

type AnalyzeResult struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Protocol        string `json:"protocol"`
	IsPublic        bool   `json:"isPublic"`
	Status          string `json:"status"`
	StartTime       int64  `json:"startTime"`
	EngineVersion   string `json:"engineVersion"`
	CriteriaVersion string `json:"criteriaVersion"`
	Endpoints       []struct {
		IPAddress            string `json:"ipAddress"`
		ServerName           string `json:"serverName"`
		StatusMessage        string `json:"statusMessage"`
		StatusDetails        string `json:"statusDetails"`
		StatusDetailsMessage string `json:"statusDetailsMessage"`
		Delegation           int    `json:"delegation"`
	} `json:"endpoints"`
}

func Analyze(url string) (*AnalyzeResult, error) {
	resp, err := http.Get("https://api.ssllabs.com/api/v3/analyze?host=" + url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result AnalyzeResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
