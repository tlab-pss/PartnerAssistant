package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	topiccategory "github.com/sskmy1024/PartnerAssistant/models/category/topic_category"
	recommend "github.com/sskmy1024/PartnerAssistant/models/recommend"
)

// BranchLogic : カテゴリに応じて実行ロジックを分岐させる
func (r *RequireServiceType) BranchLogic() (*recommend.RecommendResultResponseType, error) {
	result := &recommend.RecommendResultResponseType{}
	switch r.TopicCategory {
	case topiccategory.PersonalData:
		// パーソナルデータの更新を呼び出す
		personalData, err := r.PersonalDataValue.InvokePDUpdate()
		if err != nil {
			return nil, err
		}
		r.PersonalDataValue = *personalData
		result.ResponseData = recommend.RecommendResultType{
			Success: true,
			Text:    "あなたのデータをアップデートしたにゃ",
		}
		return result, nil
	case topiccategory.Uncategorized:
		return result, nil
	default:
		res, err := r.RequestService()
		if err != nil {
			return nil, err
		}
		result = res
		return result, nil
	}
}

// RequestService : レコメンドシステムにリクエストを投げる
func (r *RequireServiceType) RequestService() (*recommend.RecommendResultResponseType, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", "http://rs:8080/api/recommend", bytes.NewBuffer(jsonBytes))
	if err != nil {
		fmt.Printf("pd error, cannot create http request")
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("pd error! cannot exec http request")
		return nil, err
	}
	defer resp.Body.Close()

	var rBody io.Reader = resp.Body
	// rBody = io.TeeReader(rBody, os.Stderr)

	response := &recommend.RecommendResultResponseType{}

	if err := json.NewDecoder(rBody).Decode(response); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil, err
	}

	return response, nil
}
