package handler

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func (s *security) Start() {
	middleware := newAuthMiddleware()
	s.server.Post("signup", s.signup)
	s.server.Post("login", s.login)

	s.server.Post("drugs", middleware, s.registryDrug)
	s.server.Put("drugs/:id", middleware, s.updateDrug)
	s.server.Get("drugs", middleware, s.getAllDrug)
	s.server.Delete("drugs/:id", middleware, s.deleteDrug)

	s.server.Post("vaccination", middleware, s.registryVaccination)
	s.server.Put("vaccination/:id", middleware, s.updateVaccination)
	s.server.Get("vaccination", middleware, s.getAllVaccination)
	s.server.Delete("vaccination/:id", middleware, s.deleteVaccination)

	s.server.Listen(":3000")
}

func newAuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("secret"))},
	})
}
