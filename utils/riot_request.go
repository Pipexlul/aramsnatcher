package utils

import (
	"net/http"

	"github.com/braycarlson/asol"

	"github.com/pipexlul/aramsnatcher/internal/models"
)

func CreateAndExecuteRequest(
	a *asol.Asol,
	reqType models.RequestType,
	endpoint string,
	data models.RequestData,
) (interface{}, error) {
	var request *http.Request
	var err error

	switch reqType {
	case models.GET:
		request, err = a.NewGetRequest(endpoint)
	case models.POST:
		request, err = a.NewPostRequest(endpoint, data)
	case models.DELETE:
		request, err = a.NewDeleteRequest(endpoint)
	case models.PATCH:
		request, err = a.NewPatchRequest(endpoint, data)
	}

	if err != nil {
		return nil, err
	}

	result, err := a.RiotRequest(request)

	if err != nil {
		return nil, err
	}

	return result, nil
}
