package JSONmap

type RegisterFile struct {
	Url  string `json:"url" binding:"required"`
	Name string `json:"name" binding:"required"`
	Type uint   `json:"type" binding:"required"`
}
