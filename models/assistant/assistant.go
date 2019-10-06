package assistant

import (
	pdcategory "main/models/pd_category"
	topiccategory "main/models/topic_category"
)

// WatsonResponseType : WatsonAssistantからのレスポンスの型
// 変換はここでやった -> https://mholt.github.io/json-to-go/
type WatsonResponseType struct {
	StatusCode int `json:"StatusCode"`
	Headers    struct {
		AccessControlAllowHeaders    []string `json:"Access-Control-Allow-Headers"`
		AccessControlAllowMethods    []string `json:"Access-Control-Allow-Methods"`
		AccessControlAllowOrigin     []string `json:"Access-Control-Allow-Origin"`
		AccessControlMaxAge          []string `json:"Access-Control-Max-Age"`
		ContentSecurityPolicy        []string `json:"Content-Security-Policy"`
		ContentType                  []string `json:"Content-Type"`
		Date                         []string `json:"Date"`
		StrictTransportSecurity      []string `json:"Strict-Transport-Security"`
		Vary                         []string `json:"Vary"`
		XBacksideTransport           []string `json:"X-Backside-Transport"`
		XContentTypeOptions          []string `json:"X-Content-Type-Options"`
		XDNSPrefetchControl          []string `json:"X-Dns-Prefetch-Control"`
		XDownloadOptions             []string `json:"X-Download-Options"`
		XDpTransitID                 []string `json:"X-Dp-Transit-Id"`
		XDpWatsonTranID              []string `json:"X-Dp-Watson-Tran-Id"`
		XEdgeconnectMidmileRtt       []string `json:"X-Edgeconnect-Midmile-Rtt"`
		XEdgeconnectOriginMexLatency []string `json:"X-Edgeconnect-Origin-Mex-Latency"`
		XFrameOptions                []string `json:"X-Frame-Options"`
		XGlobalTransactionID         []string `json:"X-Global-Transaction-Id"`
		XXSSProtection               []string `json:"X-Xss-Protection"`
	} `json:"Headers"`
	Result struct {
		Input struct {
			Text string `json:"text"`
		} `json:"input"`
		Intents []struct {
			Intent     string  `json:"intent"`
			Confidence float64 `json:"confidence"`
		} `json:"intents"`
		Entities []EntitiesType `json:"entities"`
		Context  struct {
			ConversationID string `json:"conversation_id"`
			System         struct {
				NodeOutputMap        interface{} `json:"_node_output_map"`
				BranchExited         bool        `json:"branch_exited"`
				BranchExitedReason   string      `json:"branch_exited_reason"`
				DialogRequestCounter int         `json:"dialog_request_counter"`
				DialogStack          []struct {
					DialogNode string `json:"dialog_node"`
				} `json:"dialog_stack"`
				DialogTurnCounter int  `json:"dialog_turn_counter"`
				Initialized       bool `json:"initialized"`
			} `json:"system"`
			RequireService bool   `json:"require_service"`
			TopicCategory  string `json:"topic_category"`
			PdCategory     string `json:"pd_category"`
		} `json:"context"`
		Output struct {
			Generic []struct {
				ResponseType string `json:"response_type"`
				Text         string `json:"text"`
			} `json:"generic"`
			LogMessages  []interface{} `json:"log_messages"`
			NodesVisited []string      `json:"nodes_visited"`
			Text         []string      `json:"text"`
		} `json:"output"`
	} `json:"Result"`
}

// EntitiesType : Entityの型
type EntitiesType struct {
	Entity     string `json:"entity"`
	Location   []int  `json:"location"`
	Value      string `json:"value"`
	Confidence int    `json:"confidence"`
}

// InputText : Watsonに送ったテキストを返す
func (r WatsonResponseType) InputText() string {
	return r.Result.Input.Text
}

// ReplyText : Watson Assistantのテキストを返す
func (r WatsonResponseType) ReplyText() string {
	return r.Result.Output.Generic[0].Text
}

// OriginEntityWords : Watson AssistantのEntityに該当したオリジナルワードの配列を返す
func (r WatsonResponseType) OriginEntityWords() []string {
	words := []string{}
	originEntities := r.Result.Entities

	for _, originEntity := range originEntities {
		words = append(words, originEntity.Value)
	}
	return words
}

// IsRequireService : サービスリクエストTriggerに引っかかっているかどうか
func (r WatsonResponseType) IsRequireService() bool {
	return r.Result.Context.RequireService
}

func (r WatsonResponseType) getTopicCategory() string {
	return r.Result.Context.TopicCategory
}

func (r WatsonResponseType) getPersonalDataCategory() string {
	return r.Result.Context.PdCategory
}

// TopicCategory : 会話内容のカテゴリを返す
func (r WatsonResponseType) TopicCategory() topiccategory.TopicCategory {

	stringType := r.getTopicCategory()
	switch stringType {
	case "COMMERCE":
		return topiccategory.Commerce
	case "GOURMET":
		return topiccategory.Gourmet
	case "WEATHER":
		return topiccategory.Weather
	case "MAP":
		return topiccategory.Map
	case "PERSONALDATA":
		return topiccategory.PersonalData
	default:
		return topiccategory.Uncategorized
	}
}

// PersonalDataCategory : 含まれているパーソナルデータのカテゴリを返す
func (r WatsonResponseType) PersonalDataCategory() pdcategory.PersonalDataCategory {

	stringType := r.getPersonalDataCategory()
	switch stringType {
	case "Name":
		return pdcategory.Name
	default:
		return pdcategory.Uncategorized
	}

}
