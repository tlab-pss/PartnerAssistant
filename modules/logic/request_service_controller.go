package logic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	topiccategory "github.com/sskmy1024/PartnerAssistant/models/category/topic_category"
)

// BranchLogic : カテゴリに応じて実行ロジックを分岐させる
func (r *RequireServiceType) BranchLogic() (*RequireServiceType, error) {
	switch r.TopicCategory {
	case topiccategory.PersonalData:
		// パーソナルデータの更新を呼び出す
		personalData, err := r.PersonalDataValue.InvokePDUpdate()
		if err != nil {
			return nil, err
		}
		r.PersonalDataValue = *personalData
		return r, nil
	case topiccategory.Uncategorized:
		return r, nil
	default:
		res, err := r.RequestService()
		if err != nil {
			return nil, err
		}
		r.ServiceDataValue = *res
		return r, nil
	}
}

// RequestService : レコメンドシステムにリクエストを投げる
func (r *RequireServiceType) RequestService() (*RecommendServiceResType, error) {
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
	rBody = io.TeeReader(rBody, os.Stderr)

	rsRes := &RecommendServiceResType{}

	if err := json.NewDecoder(rBody).Decode(rsRes); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil, err
	}

	return rsRes, nil
}
