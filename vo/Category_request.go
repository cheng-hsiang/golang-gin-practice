package vo

type CreateCatrgoryRequest struct {
	Name string `json:"name" binding:"required"`
}
