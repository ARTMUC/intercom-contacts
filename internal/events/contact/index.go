package contact

import (
	"errors"
	"net/http"

	"intercom/internal/handler"
	"intercom/internal/models/response"
	"intercom/pkg/intercom"
)

type IndexHandler struct {
	intercomClient intercom.Client
}

func NewIndexHandler(intercomClient intercom.Client) *IndexHandler {
	return &IndexHandler{intercomClient: intercomClient}
}

func (c *IndexHandler) Method() string {
	return http.MethodGet
}

func (c *IndexHandler) Route() string {
	return "/contact"
}

func (c *IndexHandler) Handle(w http.ResponseWriter, r *http.Request) handler.Response[response.Pagination[Contact]] {
	responseBuilder := handler.NewResponseBuilder[response.Pagination[Contact]]()

	contacts, err := c.intercomClient.ListContacts()
	if err != nil {
		switch {
		case errors.Is(err, &intercom.ClientError{}):
			var clientError intercom.ClientError
			errors.As(err, &clientError)
			return responseBuilder.SetInternalError(err).SetError(clientError).SetStatus(500).Build()
		default:
			return responseBuilder.SetInternalError(err).SetStatus(500).Build()
		}
	}

	responseData := response.NewPaginationReponseWithMap[intercom.ContactResponse, Contact](
		contacts,
		0,
		func(contact intercom.ContactResponse) Contact { return Contact(contact) },
	)

	return responseBuilder.SetData(responseData).SetStatus(200).Build()
}
