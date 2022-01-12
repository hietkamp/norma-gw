package wallet

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
)

type Options struct {
	BaseURL string
}

type Wallet struct {
	HttpClient *http.Client
	BaseUrl    string
}

func New(options Options) *Wallet {
	url, err := url.ParseRequestURI(options.BaseURL)
	if err != nil {
		panic(fmt.Errorf("failed on parsing base url: %w", err))
	}
	return &Wallet{
		HttpClient: &http.Client{},
		BaseUrl:    url.String(),
	}
}

func (w *Wallet) VerifyVerifiableCredential(vcBytes []byte) error {
	var vc VerifiableCredential
	json.Unmarshal(vcBytes, &vc)

	log.Debug().Msgf("resolution: %s", fmt.Sprintf("%s/internal/vcr/v1/vc/%s", w.BaseUrl, url.QueryEscape(vc.Id)))
	r, err := http.NewRequest("GET", fmt.Sprintf("%s/internal/vcr/v1/vc/%s", w.BaseUrl, url.QueryEscape(vc.Id)), nil)
	if err != nil {
		return fmt.Errorf("failed to verify credential: %w", err)
	}
	r.Header.Add("Accept", "application/vc+json")
	resp, err := w.HttpClient.Do(r)
	if err != nil {
		return fmt.Errorf("failed to verify credential: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to verify credential: service not available")
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to verify credential: %w", err)
	}
	resolution := map[string]interface{}{}
	json.Unmarshal(bodyBytes, &resolution)
	log.Debug().Msgf("resolution: %s", string(bodyBytes))
	switch resolution["currentStatus"] {
	case "trusted":
		return nil
	default:
		return errors.New("verifiable credential not valid")
	}
}
