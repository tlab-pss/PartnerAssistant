// Package logic : AIに向けてリクエストを行うモジュール
package logic

import (
	"errors"

	topiccategory "main/models/category/topic_category"
	personaldata "main/models/personal_data"
)

// RequireServiceType : サービスリクエストを行うにあたって必要な情報をいれる型
type RequireServiceType struct {
	TopicCategory     topiccategory.TopicCategory    `json:"topic_category"`
	RequireService    bool                           `json:"require_service"`
	PersonalDataValue personaldata.PersonalDataValue `json:"personal_data_value"`
}

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
	}
	return nil, errors.New("Unknown topic category")
}
