package intercom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"intercom/internal/config"
)

type ClientError struct {
	message string
}

func (e ClientError) Error() string {
	return e.message
}

type ContactResponse struct {
	ID        string
	Email     string
	Companies any
}

type Client interface {
	ListContacts() ([]ContactResponse, error)
}

var _ Client = (*client)(nil)

type client struct {
}

func NewClient() *client {
	return &client{}
}

func (c *client) ListContacts() ([]ContactResponse, error) {
	req, err := http.NewRequest(http.MethodGet, config.IntercomAPIBaseURL+"/contacts", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for key, value := range map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", config.IntercomAccessToken),
		"Accept":        "application/json",
	} {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		var apiError ApiErrorReponse
		err = json.NewDecoder(resp.Body).Decode(&apiError)
		if err != nil {
			return nil, fmt.Errorf("can't unmarshal intercom error response to type: %T", ApiErrorReponse{})
		}

		errorMsg := make([]string, len(apiError.Errors))
		for i, e := range apiError.Errors {
			errorMsg[i] = fmt.Sprintf("%s - %s", e.Code, e.Message)
		}
		return nil, ClientError{message: fmt.Sprintf(
			"intercom api responded with status code %d: %s",
			resp.StatusCode,
			strings.Join(errorMsg, "; "),
		)}
	}

	var result ContactList
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	var contacts []ContactResponse
	for _, u := range result.Data {
		contacts = append(contacts, ContactResponse{
			ID:        u.ID,
			Email:     u.Email,
			Companies: u.Companies,
		})
	}

	return contacts, nil
}
