package logic

import (
	topiccategory "main/models/category/topic_category"
	personaldata "main/models/personal_data"
)

// RequireServiceType : サービスリクエストを行うにあたって必要な情報をいれる型
type RequireServiceType struct {
	TopicCategory     topiccategory.TopicCategory    `json:"topic_category"`
	RequireService    bool                           `json:"require_service"`
	PersonalDataValue personaldata.PersonalDataValue `json:"personal_data_value"`
	ServiceDataValue  interface{}                    `json:"service_data_value"`
}

// RecommendServiceResType : レコメンドシステムを介したサービスの返答を格納する型
type RecommendServiceResType struct {
	UUID        string `json:"uuid"`
	ServiceName string `json:"service_name"`
	// todo : プラグインの結果を格納する型を作って、合わせる
	ResponseData interface{} `json:"response_data"`
}
