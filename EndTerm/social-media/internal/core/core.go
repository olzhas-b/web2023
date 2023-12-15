package services

import (
	"context"
	"database/sql"

	"github.com/olzhas-b/social-media/config"
	irepo "github.com/olzhas-b/social-media/internal/interfaces/repository"
	"github.com/olzhas-b/social-media/internal/repository/posts"
)

type Core struct {
	ctx     context.Context
	cfg     *config.Config
	sqlDB   *sql.DB
	isDebug bool

	postsRepo irepo.IPosts
}

func New(
	ctx context.Context,
	cfg *config.Config,
	sqlDB *sql.DB,
) *Core {
	core := &Core{
		ctx:   ctx,
		cfg:   cfg,
		sqlDB: sqlDB,
	}

	return core.InitRepositories()
}

func (s *Core) InitRepositories() *Core {
	return s.
		InitPosts()
}

func (s *Core) InitPosts() *Core {
	s.postsRepo = posts.New(s.sqlDB, s.ctx)
	return s
}

func (s *Core) Posts() irepo.IPosts {
	return s.postsRepo
}

func (s *Core) Config() *config.Config {
	return s.cfg
}
