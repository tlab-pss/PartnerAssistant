package recommend

// RecommendResultResponseType : レコメンドシステムを介した返答を格納する型
type RecommendResultResponseType struct {
	UUID         string              `json:"uuid"`
	ServiceName  string              `json:"service_name"`
	ResponseData RecommendResultType `json:"response_data"`
}
