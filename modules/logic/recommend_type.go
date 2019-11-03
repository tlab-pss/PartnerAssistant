package logic

// RecommendType : RecommendSystemから返答されるデータの型
type RecommendType struct {
	Success    bool     `json:"success"`
	Text       string   `json:"text"`
	ImagePaths []string `json:"image_paths"`
}
