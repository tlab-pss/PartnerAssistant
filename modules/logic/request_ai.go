// Package logic : AIに向けてリクエストを行うモジュール
package logic

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/core"
	assistant "github.com/watson-developer-cloud/go-sdk/assistantv1"

	watsonResType "github.com/sskmy1024/PartnerAssistant/models/assistant"
	personaldata "github.com/sskmy1024/PartnerAssistant/models/personal_data"
	rsquery "github.com/sskmy1024/PartnerAssistant/models/request_service_query"
)

// ReplyAIType : AIのResponseの型を定義する
type ReplyAIType struct {
	Message   string `json:"message"`
	ImagePath string `json:"image_path"`
}

// RequestAI : AIに向けてリクエストを送るところ。今回はWatson Assistantを使用
func RequestAI(reqMessage string) (*watsonResType.WatsonResponseType, error) {
	authenticator := &core.IamAuthenticator{
		ApiKey: os.Getenv("WATSON_IAM_APIKEY"),
	}

	service, err := assistant.NewAssistantV1(&assistant.AssistantV1Options{
		URL:           "https://gateway.watsonplatform.net/assistant/api",
		Version:       "2019-07-10",
		Authenticator: authenticator,
	})

	if err != nil {
		return nil, err
	}

	workspaceID := os.Getenv("watson_workspace_id")
	input := &assistant.MessageInput{}
	input.SetText(core.StringPtr(reqMessage))
	messageOptions := service.NewMessageOptions(workspaceID).
		SetInput(input)

	response, err := service.Message(messageOptions)

	if err != nil {
		return nil, err
	}

	jsonBytes := ([]byte)(response.String())
	replyData := new(watsonResType.WatsonResponseType)

	if err := json.Unmarshal(jsonBytes, replyData); err != nil {
		fmt.Printf("JSON Unmarshal error: %+v \n %+v", err, response.String())
	}

	return replyData, nil
}

// ConvertRequireServiceType : サービス接続用の型に変換する
func ConvertRequireServiceType(replyData *watsonResType.WatsonResponseType) *RequireServiceType {
	return &RequireServiceType{
		TopicCategory:  replyData.TopicCategory(),
		RequireService: replyData.IsRequireService(),
		PersonalDataValue: personaldata.PersonalDataValue{
			Category:    replyData.PersonalDataCategory(),
			BasicValues: replyData.UpdateBasicPersonalData(),
		},
		ServiceDataValue: rsquery.RequestServiceQueryType{
			Keywords: replyData.GetContextKeywords(),
		},
	}
}
