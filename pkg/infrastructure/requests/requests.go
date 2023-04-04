package requests

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marcelofelixsalgado/financial-commons/api/context"
)

// Make a request to another backend (upstream)
func MakeUpstreamRequest(ctx echo.Context, method, url string, data []byte, authenticated bool) (*http.Response, error) {

	request, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	if authenticated {
		// Get the access token from the cookie and set the upstream request header
		customContext := ctx.(*context.CustomContext)
		accessToken := fmt.Sprintf("Bearer %s", customContext.Get("client"))
		request.Header.Add("Authorization", accessToken)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
