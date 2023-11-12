package adapter

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Threqt1/architecture-go/api-gateway/config"
	"github.com/Threqt1/architecture-go/api-gateway/models"
)

type MessagesServiceAdapter struct {
	route string
}

func (msa *MessagesServiceAdapter) SendMessage(user1 string, user2 string, content string) (models.Message, error) {
	message := models.Message{User1ID: user1, User2ID: user2, Message: content}

	marshalled, err := json.Marshal(message)
	if err != nil {
		return message, err
	}

	request, err := http.NewRequest(http.MethodPut, msa.route, bytes.NewReader(marshalled))
	if err != nil {
		return message, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return message, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return message, err
	}

	if response.StatusCode != http.StatusOK {
		return message, errors.New("message - " + string(body))
	}

	err = json.Unmarshal(body, &message)
	if err != nil {
		return message, err
	}

	return message, nil
}

func CreateMessagesServiceAdapter() MessagesServiceAdapter {
	return MessagesServiceAdapter{route: config.MESSAGES_MS_ROUTE}
}
