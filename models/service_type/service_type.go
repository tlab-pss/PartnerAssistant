package servicetype

// ServiceType : サービスのカテゴリを定義
type ServiceType int

const (
	// NoRequest : リクエストなし
	NoRequest ServiceType = iota
	// Commerce : コマース
	Commerce
	// Gourmet : グルメ
	Gourmet
	// Weather : 天気
	Weather
	// Map : マップ
	Map
	// Mail : メール
	Mail
	// Music : ミュージック
	Music
	// Message : メッセージ
	Message
	// Search : 検索
	Search
	// Translation : 翻訳
	Translation
	// News : ニュース
	News
)

func (s ServiceType) String() string {
	switch s {
	case NoRequest:
		return ""
	case Gourmet:
		return "Gourmet"
	case Weather:
		return "Weather"
	case Map:
		return "Map"
	case Mail:
		return "Mail"
	case Music:
		return "Music"
	case Message:
		return "Message"
	case Search:
		return "Search"
	case Translation:
		return "Translation"
	case News:
		return "News"
	default:
		return ""
	}
}
