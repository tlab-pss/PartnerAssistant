package topiccategory

// TopicCategory : サービスのカテゴリを定義
type TopicCategory int

const (
	// Uncategorized : 未分類
	Uncategorized TopicCategory = iota
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
	// PersonalData : 個人情報
	PersonalData
)

func (s TopicCategory) String() string {
	switch s {
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
	case PersonalData:
		return "PersonalData"
	case Uncategorized:
		return "Uncategorized"
	default:
		return "Uncategorized"
	}
}
