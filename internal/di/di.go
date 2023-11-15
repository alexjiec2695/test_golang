package di

import (
	"test/internal/handler"
	"test/internal/repository/drug"
	"test/internal/repository/postgres"
	user2 "test/internal/repository/user"
	"test/internal/repository/vaccination"
	"test/internal/server"
	drugService "test/internal/service/drug"
	"test/internal/service/user"
	vaccinationService "test/internal/service/vaccination"
)

func Start() {
	db, err := postgres.NewConnection()
	if err != nil {
		panic(err)
	}
	userRespository := user2.NewUserRepository(db)
	drugRepository := drug.NewDrugRepository(db)
	vaccinationRepository := vaccination.NewVaccinationRepository(db)

	serviceUser := user.NewUser(userRespository)
	serviceDrug := drugService.NewDrug(drugRepository)
	serviceVaccination := vaccinationService.NewVaccination(vaccinationRepository, drugRepository)

	newServer := server.NewServer()
	h := handler.NewSecurityHandler(newServer, serviceUser, serviceDrug, serviceVaccination)

	h.Start()

}
