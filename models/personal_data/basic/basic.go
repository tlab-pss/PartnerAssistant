package basicpersonaldata

// BasicPersonalData : ユーザの基本情報
type BasicPersonalData struct {
	ID       string `json:"ID"`
	Name     string `json:"Name"`
	Birthday string `json:"Birthday"`
	Gender   Gender `json:"Gender"`
	Mail     string `json:"Mail"`
	CreateAt string `json:"CreateAt"`
}

// Gender : 性別
type Gender int

const (
	// Male : 0
	Male Gender = iota
	// Female : 1
	Female
	// Other : 2
	Other
)

// BasicPDColumn : ベーシックデータのカラム
type BasicPDColumn int

const (
	// Uncategorized : 未分類
	Uncategorized BasicPDColumn = iota
	// ID : uuid
	ID
	// Name : 名前
	Name
	// Birthday : 誕生日
	Birthday
)

func (c BasicPDColumn) String() string {
	switch c {
	case ID:
		return "ID"
	case Name:
		return "Name"
	case Birthday:
		return "Birthday"
	case Uncategorized:
		return "Uncategorized"
	default:
		return "Uncategorized"
	}
}

// UpdateBasicPersonalData : Basicパーソナルデータを更新時に必要なパラメタの型
type UpdateBasicPersonalData struct {
	Column BasicPDColumn `json:"column"`
	Value  string        `json:"value"`
}
