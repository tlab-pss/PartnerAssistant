package personaldata

import (
	"errors"
	"fmt"

	basicpd "github.com/sskmy1024/PartnerAssistant/models/personal_data/basic"
)

// PersonalDataCategory : ベーシックデータのカテゴリを定義
type PersonalDataCategory int

const (
	// Uncategorized : 未確定
	Uncategorized PersonalDataCategory = iota
	// Basic : ベーシック
	Basic
	// Location : 位置情報
	Location
)

func (s PersonalDataCategory) String() string {
	switch s {
	case Basic:
		return "Basic"
	case Location:
		return "Location"
	default:
		return "Uncategorized"
	}
}

// PersonalDataValue : パーソナルデータ更新用の型
type PersonalDataValue struct {
	Category    PersonalDataCategory            `json:"category"`
	BasicValues basicpd.UpdateBasicPersonalData `json:"basic_values"`
}

func (p *PersonalDataValue) getBasicValues() (*basicpd.UpdateBasicPersonalData, error) {
	if p.BasicValues.Column != basicpd.Uncategorized {
		return nil, errors.New("not have to update basic values")
	}
	return &p.BasicValues, nil
}

// InvokePDUpdate : パーソナルデータのアップデーターを呼び出す
func (p *PersonalDataValue) InvokePDUpdate() (*PersonalDataValue, error) {

	if p.Category == Uncategorized {
		fmt.Println("PersonalData Category is not categorized")
		return nil, errors.New("PersonalData Category is not categorized")
	}

	switch p.Category {
	case Basic:
		pd := new(basicpd.BasicPersonalData)
		pd, err := pd.Fetch()
		if err != nil {
			fmt.Println("cannot fetch personal data")
			return nil, err
		}

		updateValue, err := p.getBasicValues()
		if err != nil {
			return nil, err
		}

		// BasicPersonalDataの更新
		err = pd.Update(updateValue)
		if err != nil {
			return nil, err
		}

		return p, nil
	}
	return nil, errors.New("Unknown category")
}
