package logic

// LogicPayload : Logicに渡すデータ群
type LogicPayload struct {
	ImagePath   string `json:"image_path"`
	UserMessage string `json:"user_message"`
}

// ExecuteLogic : AIとクライアントを仲介するLogicを実行する
func (payload LogicPayload) ExecuteLogic() (*ReplyAIType, error) {
	var requestArgs RequireServiceType
	var replyMessage string

	if payload.UserMessage != "" {
		replyData, err := RequestAI(payload.UserMessage)
		if err != nil {
			return &ReplyAIType{
				Message: "AIとの接続に失敗しました",
			}, err
		}

		requestArgs = *ConvertRequireServiceType(replyData)
		replyMessage = replyData.ReplyText()
	}

	requestArgs.UserSendDataValue = payload

	var result ReplyAIType

	// Note : レスポンスのパラメタによって動作を分岐させる
	response, err := requestArgs.BranchLogic()
	if err != nil {
		result.Message = replyMessage
	}

	result.Message = response.ResponseData.Text

	return &result, nil
}
