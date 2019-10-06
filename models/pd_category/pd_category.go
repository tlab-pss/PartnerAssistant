package pdcategory

// PersonalDataCategory : サービスのカテゴリを定義
type PersonalDataCategory int

const (
	// Uncategorized : 未分類
	Uncategorized PersonalDataCategory = iota
	// Name : 名前
	Name
	// Birthday : 誕生日
	Birthday
)

func (s PersonalDataCategory) String() string {
	switch s {
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
