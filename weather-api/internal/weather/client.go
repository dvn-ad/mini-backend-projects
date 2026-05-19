package weather

import(
	"context"
	"net/http"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type APIResponse struct{
	Days []struct {
		Date        string  `json:"datetimeStr"`
		TempMaxC    float64 `json:"tempmax"`
		TempMinC    float64 `json:"tempmin"`
		Conditions  string  `json:"conditions"`
	} `json:"days"`
}
	

type Client struct {
	apiKey string
	resty  *resty.Client
}

func NewClient(apiKey string) *Client {
	r := resty.New()
	r.SetTimeout(10 * time.Second)
	r.SetBaseURL("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline")

	return &Client{
		apiKey: apiKey,
		resty:  r,
	}
}

func (c *Client) FetchWeather(ctx context.Context, city string) (*APIResponse, error) {
	var result APIResponse

	resp, err := c.resty.R().
		SetContext(ctx).
		SetPathParams(map[string]string{
			"location":   city,
			"start_date": time.Now().Format("2006-01-02"),
		}).
		SetQueryParams(map[string]string{
			"unitGroup":   "metric",
			"include":     "current",
			"contentType": "json",
			"key":         c.apiKey,
		}).
		SetResult(&result).
		Get("/{location}/{start_date}")

	if err != nil {
		return nil, fmt.Errorf("resty request failed: %w", err)
	}

	if resp.StatusCode() == http.StatusBadRequest {
		return nil, fmt.Errorf("invalid city or location query code")
	}
	
	if resp.IsError() {
		return nil, fmt.Errorf("external weather API returned status: %d", resp.StatusCode())
	}

	return &result, nil
}

// const weatherEndpoint:="https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/{location}/{start_date}?unitGroup={unit_group}&include={include}&contentType={content_type}&key={api_key}"