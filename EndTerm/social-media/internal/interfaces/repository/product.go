package interfaces

import (
	"github.com/olzhas-b/social-media/internal/models"
)

type IPosts interface {
	List(searchText string, userID uint64) (products []models.PostDTO, err error)
	ByID(productID uint64) (productModel models.Posts, err error)
	Add(productModel models.Posts) (err error)
	Remove(productID uint64) (err error)
}
