package main

import (
	"encoding/json"
	"fmt"
	"gitlab.hepsiburada.com/campaign/rabbitmq-shovel-retry/config"
	"io/ioutil"
	"os"
)

func GetQueueList() []string {
	jsonFile, err := os.Open("queuelist.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened queuelist.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var arr []string

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &arr)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return arr
}

func GetConfig(env string) (*config.ConfigModel, error) {
	fileName := fmt.Sprintf("config.%s.json", env)
	filePath := fmt.Sprintf("config/%s", fileName)

	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Successfully Opened %s", fileName))
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var model config.ConfigModel

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &model)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &model, nil
}
