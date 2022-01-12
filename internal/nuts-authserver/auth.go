package nutsauthserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Options struct {
	BaseURL string
}

type AuthServer struct {
	HttpClient *http.Client
	BaseUrl    string
}

type Jwt struct {
	Active             bool   `json:"active,omitempty"`
	Audience           string `json:"aud,omitempty"`
	ExpirationDateTime int64  `json:"exp,omitempty"`
	IssueDateTime      int64  `json:"iat,omitempty"`
	Issuer             string `json:"iss,omitempty"`
	Claim              string `json:"service,omitempty"`
	Subject            string `json:"sub,omitempty"`
}

func New(options Options) *AuthServer {
	url, err := url.ParseRequestURI(options.BaseURL)
	if err != nil {
		panic(fmt.Errorf("failed on parsing base url: %w", err))
	}
	return &AuthServer{
		HttpClient: &http.Client{},
		BaseUrl:    url.String(),
	}
}

func (a *AuthServer) VerifyAccessToken(jwt string) (string, error) {
	data := url.Values{}
	data.Set("token", jwt)

	r, err := http.NewRequest("POST", a.BaseUrl+"/internal/auth/v1/accesstoken/introspect", strings.NewReader(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := a.HttpClient.Do(r)
	if err != nil {
		return "", fmt.Errorf("failed to verify access token: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to verify access token: service not available")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to verify access token: %w", err)
	}
	var jwtClaim Jwt
	if err := json.Unmarshal(body, &jwtClaim); err != nil {
		return "", fmt.Errorf("failed to verify access token: %w", err)
	}
	if jwtClaim.Active {
		// Add validation for a leeway time of 120 seconds
		var leeway int64 = 120
		if time.Now().Unix()-leeway > jwtClaim.ExpirationDateTime {
			// The token is expired
			return "", errors.New("access token not valid")
		}
		if time.Now().Unix()+leeway < jwtClaim.IssueDateTime {
			// The token was issued in the future
			return "", errors.New("access token not valid")
		}
		return jwtClaim.Claim, nil
	}
	return "", errors.New("access token not valid")
}
