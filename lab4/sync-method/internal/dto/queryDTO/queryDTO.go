package queryDTO

type FindByTaskDTO struct {
	Name *string `json:"name" binding:"required"`
}
