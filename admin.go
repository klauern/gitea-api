package gitea

import (
	"github.com/dghubble/sling"
)

type AdminService struct {
	sling *sling.Sling
}

func NewAdminService(sling *sling.Sling) *AdminService {
	return &AdminService{
		sling: sling,
	}
}
