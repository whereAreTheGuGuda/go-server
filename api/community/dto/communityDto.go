package dto

type CommunityDto struct {
	Context string `json:"context" binding:"required,min=5,max=300"`
}
