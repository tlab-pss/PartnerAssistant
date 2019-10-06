// Package logic : AIに向けてリクエストを行うモジュール
package logic

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	assistant "github.com/watson-developer-cloud/go-sdk/assistantv1"

	watsonResType "main/models/assistant"
	personaldata "main/models/personal_data"
)

// ReplyAIType : AIのResqponseの型を定義する
type ReplyAIType struct {
	Message string `json:"message"`
}

// RequestAI : AIに向けてリクエストを送るところ。今回はWatson Assistantを使用
func RequestAI(reqMessage string) (*ReplyAIType, error) {
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
		return &ReplyAIType{}, responseErr
	}

	jsonBytes := ([]byte)(response.String())
	replyData := new(watsonResType.WatsonResponseType)

	if err := json.Unmarshal(jsonBytes, replyData); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
	}

	// fmt.Println(response)

	requestArgs := &RequireServiceType{
		TopicCategory:  replyData.TopicCategory(),
		RequireService: replyData.IsRequireService(),
		PersonalDataValue: personaldata.PersonalDataValue{
			Category:    replyData.PersonalDataCategory(),
			BasicValues: replyData.UpdateBasicPersonalData(),
		},
	}

	// レスポンスのパラメタによって動作を分岐させる
	requestArgs.BranchLogic()

	result := &ReplyAIType{
		Message: replyData.ReplyText(),
	}

	if replyData.IsRequireService() {
		result = result.RequestService()
	}

	return result, nil
}
