package discgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ApiConfig struct {
	APIBaseURL    string
	ApplicationID string
	PublicKey     string
	ClientID      string
	ClientSecret  string
	BotToken      string
}
type ApiClient struct {
	discordConfig *ApiConfig
	client        HTTPClient
	logger        Logger
}

type ApiClientOption func(*ApiClient)

func WithApiConfig(cfg *ApiConfig) ApiClientOption {
	return func(d *ApiClient) {
		d.discordConfig = cfg
	}
}

func WithHttpClient(c HTTPClient) ApiClientOption {
	return func(d *ApiClient) {
		d.client = c
	}
}

func WithClientLogger(l Logger) ApiClientOption {
	return func(dc *ApiClient) {
		// TODO make this a logger interface
		dc.logger = l
	}
}

func NewApiClient(options ...ApiClientOption) *ApiClient {
	d := ApiClient{}

	for _, opt := range options {
		opt(&d)
	}

	if d.client == nil {
		d.client = http.DefaultClient
	}

	if d.discordConfig.APIBaseURL == "" {
		d.discordConfig.APIBaseURL = "https://discord.com/api"
	}

	if d.logger == nil {
		d.logger = &noopLogger{}
	}

	return &d
}

// Do adds headers before calling ApiClient.client.Do
func (d *ApiClient) Do(req *http.Request) (*http.Response, error) {

	req.Header.Set("Authorization", "Bot "+d.discordConfig.BotToken)

	return d.client.Do(req)
}

func (d *ApiClient) Close() {
	// nothing to do here right now
}

func (d *ApiClient) Get(ctx context.Context, endpoint string, t interface{}) (*http.Response, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", d.discordConfig.APIBaseURL, endpoint), nil)
	if err != nil {
		return nil, fmt.Errorf(ErrUnableToCreateRequest, err)
	}

	resp, err := d.Do(req)
	if err != nil {
		return nil, fmt.Errorf(ErrUnableToCompleteRequest, err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, fmt.Errorf(ErrUnableToParseResponse, err)
	}

	// reset the response body for later reading
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	err = json.Unmarshal(data, t)
	if err != nil {
		return resp, fmt.Errorf("unable to unmarshal response: %w", err)
	}

	return resp, nil
}

func (d *ApiClient) Post(ctx context.Context, endpoint string, payload interface{}, t interface{}) (*http.Response, error) {

	var body io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal payload: %w", err)
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("%s%s", d.discordConfig.APIBaseURL, endpoint), body)
	if err != nil {
		return nil, fmt.Errorf(ErrUnableToCreateRequest, err)
	}

	resp, err := d.Do(req)
	if err != nil {
		return nil, fmt.Errorf(ErrUnableToCompleteRequest, err)
	}

	err = errorFromResponse(resp)
	if err != nil {
		return resp, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, fmt.Errorf(ErrUnableToParseResponse, err)
	}

	// reset the response body for later reading
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	err = json.Unmarshal(data, t)
	if err != nil {
		return resp, fmt.Errorf("unable to unmarshal response: %w", err)
	}

	return resp, nil
}

func (d *ApiClient) Patch(ctx context.Context, endpoint string, payload interface{}, t interface{}) (*http.Response, error) {
	var body io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal payload: %w", err)
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, fmt.Sprintf("%s%s", d.discordConfig.APIBaseURL, endpoint), body)
	if err != nil {
		return nil, fmt.Errorf(ErrUnableToCreateRequest, err)
	}

	resp, err := d.Do(req)
	if err != nil {
		return nil, fmt.Errorf(ErrUnableToCompleteRequest, err)
	}

	// if we have t, unmarshal the body into t
	if t != nil {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp, fmt.Errorf(ErrUnableToParseResponse, err)
		}

		// reset the response body for later reading
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		err = json.Unmarshal(data, t)
		if err != nil {
			return resp, fmt.Errorf("unable to unmarshal response: %w", err)
		}
	}

	return resp, nil
}

func (d *ApiClient) Delete(ctx context.Context, endpoint string, payload interface{}, t interface{}) (*http.Response, error) {
	var body io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal payload: %w", err)
		}
		body = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf("%s%s", d.discordConfig.APIBaseURL, endpoint), body)
	if err != nil {
		return nil, fmt.Errorf(ErrUnableToCreateRequest, err)
	}

	resp, err := d.Do(req)
	if err != nil {
		return nil, fmt.Errorf(ErrUnableToCompleteRequest, err)
	}

	// if we have t, unmarshal the body into t
	if t != nil {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp, fmt.Errorf(ErrUnableToParseResponse, err)
		}

		// reset the response body for later reading
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		err = json.Unmarshal(data, t)
		if err != nil {
			return resp, fmt.Errorf("unable to unmarshal response: %w", err)
		}
	}

	return resp, nil
}

func errorFromResponse(resp *http.Response) error {
	statusCode := resp.StatusCode

	if statusCode < 200 || statusCode >= 300 {
		var errMessage string
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			errMessage = fmt.Errorf("unable to read error body: %w", err).Error()
		} else {
			var errorResponse ErrorResponse
			err := json.Unmarshal(data, &errorResponse)
			if err != nil {
				errMessage = string(data)
			} else {
				errMessage = errorResponse.Message
			}

			// reset the response body for later reading
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		}
		return fmt.Errorf("unexpected status: %d - %s", statusCode, errMessage)
	}

	return nil
}
