package logic

import (
	topiccategory "github.com/sskmy1024/PartnerAssistant/models/category/topic_category"
	personaldata "github.com/sskmy1024/PartnerAssistant/models/personal_data"
	rsquery "github.com/sskmy1024/PartnerAssistant/models/request_service_query"
)

// RequireServiceType : サービスリクエストを行うにあたって必要な情報をいれる型
type RequireServiceType struct {
	TopicCategory     topiccategory.TopicCategory     `json:"topic_category"`
	RequireService    bool                            `json:"require_service"`
	PersonalDataValue personaldata.PersonalDataValue  `json:"personal_data_value"`
	ServiceDataValue  rsquery.RequestServiceQueryType `json:"service_data_value"`
	UserSendDataValue LogicPayload                    `json:"user_send_data_value"`
}
