package yalogapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type httpRequest struct {
	Method string
	URI    string
	Token  string
	Params url.Values
}

func doRequest(httpRequest httpRequest) ([]byte, error) {
	request, err := http.NewRequest(httpRequest.Method, httpRequest.URI, nil)

	errorMessage := fmt.Sprintf("Yandex api GET request failed. Uri %s", httpRequest.URI)

	if err != nil {
		return nil, errors.Wrapf(err, errorMessage)
	}

	request.Header.Add(
		"Authorization",
		fmt.Sprintf("OAuth %s", httpRequest.Token),
	)

	request.URL.RawQuery = httpRequest.Params.Encode()

	client := &http.Client{}
	response, err := client.Do(request)

	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)

	if err != nil {
		return nil, errors.Wrapf(err, errorMessage)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, errors.Wrapf(err, errorMessage)
	}

	return body, nil
}
