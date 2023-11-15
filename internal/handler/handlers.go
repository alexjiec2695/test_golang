package handler

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"test/internal/entities"
	"test/internal/service/drug"
	"test/internal/service/jwt"
	"test/internal/service/user"
	"test/internal/service/vaccination"
	"time"
)

type security struct {
	server              *fiber.App
	userExecutor        user.Executor
	drugExecutor        drug.Executor
	vaccinationExecutor vaccination.Executor
}

func NewSecurityHandler(server *fiber.App, signupExecutor user.Executor, drugExecutor drug.Executor, vaccinationExecutor vaccination.Executor) security {
	return security{
		server:              server,
		userExecutor:        signupExecutor,
		drugExecutor:        drugExecutor,
		vaccinationExecutor: vaccinationExecutor,
	}
}

/*user*/

func (s *security) signup(ctx *fiber.Ctx) error {
	u := entities.User{}

	err := ctx.BodyParser(&u)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	return s.userExecutor.Registry(u)
}

func (s *security) login(ctx *fiber.Ctx) error {
	u := entities.User{}

	err := ctx.BodyParser(&u)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	err = s.userExecutor.Login(u.Email, u.Password)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{"Data": err.Error()})
	}

	t, err := jwt.Generate()

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"Token": t})
}

/*Drug*/

func (s *security) registryDrug(ctx *fiber.Ctx) error {
	d := entities.Drug{}
	err := ctx.BodyParser(&d)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	t, err := time.Parse("2006-01-02", d.AvailableAt)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	d.AvailableAt = t.Format("2006-01-02")

	return s.drugExecutor.Registry(d)
}

func (s *security) updateDrug(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	d := entities.Drug{}
	err := ctx.BodyParser(&d)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	t, err := time.Parse("2006-01-02", d.AvailableAt)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	d.AvailableAt = t.Format("2006-01-02")

	err = s.drugExecutor.UpdateItem(id, d)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	return ctx.SendStatus(http.StatusOK)
}

func (s *security) getAllDrug(ctx *fiber.Ctx) error {

	registries, err := s.drugExecutor.GetAllItems()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(registries)
}

func (s *security) deleteDrug(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := s.drugExecutor.DeleteItem(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	return ctx.SendStatus(http.StatusOK)
}

/*Vaccination*/

func (s *security) registryVaccination(ctx *fiber.Ctx) error {
	v := entities.Vaccination{}
	err := ctx.BodyParser(&v)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	t, err := time.Parse("2006-01-02", v.Date)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	v.Date = t.Format("2006-01-02")

	err = s.vaccinationExecutor.Registry(v)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	return ctx.SendStatus(http.StatusOK)
}

func (s *security) updateVaccination(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	v := entities.Vaccination{}
	err := ctx.BodyParser(&v)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	t, err := time.Parse("2006-01-02", v.Date)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	v.Date = t.Format("2006-01-02")

	err = s.vaccinationExecutor.UpdateItem(id, v)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	return ctx.SendStatus(http.StatusOK)
}

func (s *security) getAllVaccination(ctx *fiber.Ctx) error {
	registries, err := s.vaccinationExecutor.GetAllItems()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(registries)
}

func (s *security) deleteVaccination(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := s.vaccinationExecutor.DeleteItem(id)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Data": err.Error()})
	}

	return ctx.SendStatus(http.StatusOK)
}
