package recommend

// RecommendResultType : RecommendSystemから返答されるデータの型
type RecommendResultType struct {
	Success    bool     `json:"success"`
	Text       string   `json:"text"`
	ImagePaths []string `json:"image_paths"`
}
