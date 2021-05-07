package repositories

import (
	"github.com/douglasdennys/clean-architecture-go/src/application/ports/repositories/log"
	application "github.com/douglasdennys/clean-architecture-go/src/application/usecases/log"
	"github.com/douglasdennys/clean-architecture-go/src/drivers/database/mongodb/repositories"
	"github.com/douglasdennys/clean-architecture-go/src/drivers/database/mongodb/shared"
)

func MakeDbAddLogRepositoryFactory() log.AddLogRepository {
	collection := shared.GetCollection("logs")
	repo := repositories.NewLogMongoRepository(collection)
	return application.NewAddLogBusiness(repo)
}
