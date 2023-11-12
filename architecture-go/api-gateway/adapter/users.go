package adapter

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Threqt1/architecture-go/api-gateway/config"
	"github.com/Threqt1/architecture-go/api-gateway/models"
)

type UserServiceAdapter struct {
	route string
}

func (usa *UserServiceAdapter) GetUser(id string) (models.User, error) {
	user := models.User{}

	request, err := http.NewRequest(http.MethodGet, usa.route, nil)
	if err != nil {
		return user, err
	}

	query := request.URL.Query()
	query.Add("id", id)

	request.URL.RawQuery = query.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return user, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return user, err
	}

	if response.StatusCode != http.StatusOK {
		return user, errors.New("user - " + string(body))
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ua *UserServiceAdapter) CreateUser() (models.User, error) {
	user := models.User{}

	request, err := http.NewRequest(http.MethodPut, ua.route, nil)
	if err != nil {
		return user, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return user, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return user, err
	}

	if response.StatusCode != http.StatusOK {
		return user, errors.New("user - " + string(body))
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUserServiceAdapter() UserServiceAdapter {
	return UserServiceAdapter{route: config.USER_MS_ROUTE}
}
