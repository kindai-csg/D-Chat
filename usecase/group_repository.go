package usecase

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type GroupRepository interface {
	Create(domain.Group) (domain.Group, error)
}
