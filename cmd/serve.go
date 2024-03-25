package cmd

import (
	"bookstack/repository"
	"bookstack/repository/gorm"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func serveCommand() *cobra.Command {
	cmd := cobra.Command{
		Use:   "serve",
		Short: "Start the server",
		Run: func(cmd *cobra.Command, args []string) {
			engine, err := c.getDatabase()
			if err != nil {
				panic(err)
			}
			db, err := engine.DB()
			if err != nil {
				panic(err)
			}
			defer db.Close()
			fmt.Println("database connection was established")

			repo := gorm.NewGormRepository(engine)

			server, err := newServer(engine, repo)
			if err != nil {
				panic(err)
			}

			fmt.Printf("server is running on port %d\n", c.Port)
			server.Start(fmt.Sprintf(":%d", c.Port))
		},
	}
	return &cmd
}

type Server struct {
	Router *echo.Echo
	Repo   repository.Repository
}

func (s *Server) Start(address string) error {
	return s.Router.Start(address)
}
