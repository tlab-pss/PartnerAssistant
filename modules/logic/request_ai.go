// Package logic : AIに向けてリクエストを行うモジュール
package logic

import (
	"fmt"
	"os"
	"encoding/json"

	"github.com/IBM/go-sdk-core/core"
	assistant "github.com/watson-developer-cloud/go-sdk/assistantv1"

	assistantType "main/models/assistant"
)

// RequestAI : AIに向けてリクエストを送るところ。今回はWatson Assistantを使用
func RequestAI(reqMessage string) interface{} {
	// Instantiate the Watson Assistant service
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("watson_iam_apikey"),
	}

	service, serviceErr := assistant.NewAssistantV1(&assistant.AssistantV1Options{
		URL:           "https://gateway.watsonplatform.net/assistant/api",
		Version:       "2019-07-10",
		Authenticator: authenticator,
	})

	// Check successful instantiation
	if serviceErr != nil {
		panic(serviceErr)
	}

	workspaceID := os.Getenv("watson_workspace_id")

	input := &assistant.MessageInput{}
	input.SetText(core.StringPtr(reqMessage))

	messageOptions := service.NewMessageOptions(workspaceID).
		SetInput(input)

	// Call the Message method with no specified context
	response, responseErr := service.Message(messageOptions)

	// Check successful call
	if responseErr != nil {
		panic(responseErr)
	}
	
	jsonBytes := ([]byte)(response.String())
	data := new(assistantType.ResponseType)

	if err := json.Unmarshal(jsonBytes, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}
	
	return data.Result.Output.Generic[0].Text
}
