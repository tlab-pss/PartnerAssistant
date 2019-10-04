package assistant

import (
	"main/models/servicetype"
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
			ServiceType    string `json:"service_type"`
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

	for _, originEntitie := range originEntities {
		words = append(words, originEntitie.Value)
	}
	return words
}

// IsRequireService : サービスリクエストTriggerに引っかかっているかどうか
func (r WatsonResponseType) isRequireService() bool {
	return r.Result.Context.RequireService
}

func (r WatsonResponseType) getServiceType() string {
	return r.Result.Context.ServiceType
}

// ServiceType : 要求サービスのカテゴリを返す
func (r WatsonResponseType) ServiceType() servicetype.ServiceType {

	if r.isRequireService() != true {
		return servicetype.NoRequest
	}

	stringType := r.getServiceType()
	switch stringType {
	case "COMMERCE":
		return servicetype.Commerce
	case "GOURMET":
		return servicetype.Gourmet
	case "WEATHER":
		return servicetype.Weather
	case "MAP":
		return servicetype.Map
	default:
		return servicetype.NoRequest
	}
}
