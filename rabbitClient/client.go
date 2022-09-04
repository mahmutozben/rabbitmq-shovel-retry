package rabbitClient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/mahmutozben/rabbitmq-shovel-retry/config"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	RabbitClient interface {
		MoveErrorToQueue(queueName string) (bool, error)
		GetQueueInfo(queueName string) *RabbitMqQueueInfo
	}
	rabbitClient struct {
		rabbitConfig config.RabbitMqConfig
	}
)

func NewRabbitClient(rabbitConfig config.RabbitMqConfig) (instance RabbitClient) {
	instance = &rabbitClient{
		rabbitConfig: rabbitConfig,
	}
	return
}


func (s rabbitClient) MoveErrorToQueue(queueName string) (bool, error) {
	errorQueue := fmt.Sprintf("%s_error",queueName)
	url := fmt.Sprintf("%s/api/parameters/shovel/%s/Move%s0from%s0%s", s.rabbitConfig.Server , "%2F","%2", "%2",errorQueue)
	request := createRequest(queueName)
	json, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	client := http.Client{Timeout: 10 * time.Second}
	client.Timeout = time.Second * 10
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(json))
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(s.rabbitConfig.User, s.rabbitConfig.Password)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	statusCode := res.StatusCode
	defer res.Body.Close()

	return IsSuccessStatusCode(statusCode), nil
}

func (s rabbitClient) GetQueueInfo(queueName string) *RabbitMqQueueInfo {
	url := fmt.Sprintf("%s/api/queues/%s/%s", s.rabbitConfig.Server, "%2F", queueName)
	client := http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	req.SetBasicAuth(s.rabbitConfig.User, s.rabbitConfig.Password)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer req.Body.Close()
	model := RabbitMqQueueInfo{}
	if res.StatusCode == http.StatusOK {
		json.Unmarshal(body, &model)
	}
	return &model
}

func getToken(rabbitConfig config.RabbitMqConfig) string {
	userpass := fmt.Sprintf("%s.%s",rabbitConfig.User,rabbitConfig.Password)
	return EncodeB64(userpass)
}

func EncodeB64(message string) (retour string) {
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(base64Text, []byte(message))
	return string(base64Text)
}

func IsSuccessStatusCode(statusCode int) bool {
	if statusCode >= 200 && statusCode <= 299 {
		return true
	}
	return false
}