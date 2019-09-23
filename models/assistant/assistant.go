package assistant

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
		Entities []interface{} `json:"entities"`
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
