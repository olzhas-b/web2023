package interfaces

import (
	"github.com/olzhas-b/social-media/config"
	irepo "github.com/olzhas-b/social-media/internal/interfaces/repository"
)

type ICore interface {
	Config() *config.Config
	SmsConfirmationRepo() irepo.IPosts
}
