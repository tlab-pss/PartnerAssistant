package updatepersonaldata

import (
	pdcategory "main/models/category/personal_data_category"
)

// UpdatePDType : パーソナルデータアップデート用の型
type UpdatePDType struct {
	Category pdcategory.PersonalDataCategory `json:"category"`
}
