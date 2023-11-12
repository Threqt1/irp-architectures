package adapter

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Threqt1/architecture-go/api-gateway/config"
	"github.com/Threqt1/architecture-go/api-gateway/models"
)

type MatchServiceAdapter struct {
	route string
}

func (msa *MatchServiceAdapter) MatchUser(id string) (models.Match, error) {
	match := models.Match{}

	request, err := http.NewRequest(http.MethodPut, msa.route, nil)
	if err != nil {
		return match, err
	}

	query := request.URL.Query()
	query.Add("id", id)

	request.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return match, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return match, err
	}

	if response.StatusCode != http.StatusOK {
		return match, errors.New("match - " + string(body))
	}

	err = json.Unmarshal(body, &match)
	if err != nil {
		return match, err
	}

	return match, nil
}

func (msa *MatchServiceAdapter) GetMatch(id string) (models.Match, error) {
	match := models.Match{}

	request, err := http.NewRequest(http.MethodGet, msa.route, nil)
	if err != nil {
		return match, err
	}

	query := request.URL.Query()
	query.Add("id", id)

	request.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return match, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return match, err
	}

	if response.StatusCode != http.StatusOK {
		return match, errors.New("match - " + string(body))
	}

	err = json.Unmarshal(body, &match)
	if err != nil {
		return match, err
	}

	return match, nil
}

func CreateMatchServiceAdapter() MatchServiceAdapter {
	return MatchServiceAdapter{route: config.MATCHING_MS_ROUTE}
}
