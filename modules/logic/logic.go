package logic

// ExecuteLogic : AIとクライアントを仲介するLogicを実行する
func ExecuteLogic(request string) (*ReplyAIType, error) {
	replyData, err := RequestAI(request)
	if err != nil {
		return &ReplyAIType{
			Message: "AIとの接続に失敗しました",
		}, err
	}

	requestArgs := ConvertRequireServiceType(replyData)

	// レスポンスのパラメタによって動作を分岐させる
	requestArgs.BranchLogic()

	result := &ReplyAIType{
		Message: replyData.ReplyText(),
	}

	return result, nil
}
