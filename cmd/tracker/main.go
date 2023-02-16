package main

import (
	"database/sql"
	"math"
	server "net/http"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/shonjord/tracker/cmd/config"
	"github.com/shonjord/tracker/internal/pkg/application/handler"
	"github.com/shonjord/tracker/internal/pkg/infrastructure/http"
	"github.com/shonjord/tracker/internal/pkg/infrastructure/mysql"
	"github.com/shonjord/tracker/internal/pkg/presentation/rest"
	"github.com/shonjord/tracker/internal/pkg/presentation/rest/action"
	log "github.com/sirupsen/logrus"
)

var (
	spec config.Specification
)

const (
	retries = 4
)

func init() {
	if err := config.LoadEnvironmentVariables(&spec); err != nil {
		log.WithError(err).Fatal("error while loading environment variables.")
	}

	level, err := log.ParseLevel(spec.Log.Level)
	if err != nil {
		log.WithError(err).Fatal("log level could not be parsed.")
	}

	log.SetLevel(level)
}

func main() {
	var (
		mysqlConnection *sql.DB
		err             error
	)

	// infrastructure layer
	retry(func() error {
		mysqlConnection, err = sql.Open(spec.Database.Driver, spec.Database.DSN)
		if err != nil {
			return err
		}

		if err = mysqlConnection.Ping(); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.WithError(err).Fatal("connection to DB was not possible.")
	}

	dbConnection := mysql.NewConnection(mysqlConnection)
	employeeRepository := mysql.NewEmployeeRepository(dbConnection)
	computerRepository := mysql.NewComputerRepository(dbConnection)
	employeeRepository.WithComputerRepository(computerRepository)
	computerRepository.WithEmployeeRepository(employeeRepository)

	retryableClient := http.NewRetryable()
	adminNotificationClient := http.NewAdminNotificationClient(spec.AdminNotification.Url, retryableClient)

	// application layer
	assignNewComputerToEmployeeHandler := handler.NewAssignNewComputerToEmployee(
		employeeRepository,
		computerRepository,
		adminNotificationClient,
	)
	assignComputerToEmployeeHandler := handler.NewAssignComputerToEmployee(
		employeeRepository,
		computerRepository,
		adminNotificationClient,
	)
	unassignComputerFromEmployeeHandler := handler.NewUnassignComputerFromEmployee(
		employeeRepository,
		computerRepository,
	)
	getAllComputersHandler := handler.NewGetAllComputers(computerRepository)
	getComputerByUUIDHandler := handler.NewGetComputerByUUID(computerRepository)
	getComputersAssignedToEmployeeHandler := handler.NewGetComputersAssignedToEmployee(employeeRepository)

	// presentation layer
	assignNewComputerToEmployeeAction := action.NewAssignNewComputerToEmployee(
		assignNewComputerToEmployeeHandler,
	)
	assignComputerToEmployeeAction := action.NewAssignComputerToEmployee(
		assignComputerToEmployeeHandler,
	)
	unassignComputerFromEmployeeAction := action.NewUnassignNewComputerToEmployee(
		unassignComputerFromEmployeeHandler,
	)
	getAllComputersAction := action.NewGetAllComputers(getAllComputersHandler)
	getComputerByUUIDAction := action.NewGetComputerByUUID(getComputerByUUIDHandler)
	getComputersAssignedToEmployeeAction := action.NewGetComputersAssignedToEmployee(
		getComputersAssignedToEmployeeHandler,
	)

	// rest endpoints
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Route("/employees", func(r chi.Router) {
			r.Route("/{employeeUUID}", func(r chi.Router) {
				r.Post("/computers", rest.Handler(assignNewComputerToEmployeeAction))
				r.Get("/computers", rest.Handler(getComputersAssignedToEmployeeAction))
			})
		})

		r.Route("/computers", func(r chi.Router) {
			r.Get("/", rest.Handler(getAllComputersAction))
			r.Route("/{computerUUID}", func(r chi.Router) {
				r.Get("/", rest.Handler(getComputerByUUIDAction))
				r.Delete("/unassign-employee", rest.Handler(unassignComputerFromEmployeeAction))
				r.Post("/assign-employee", rest.Handler(assignComputerToEmployeeAction))
			})
		})
	})

	if err = server.ListenAndServe(spec.Server.Port, router); err != nil {
		log.WithError(err).Fatal("error initializing server.")
	}
}

// retry receives a callback and executes it, if an errors is encountered
// it retries after the $retries attempts.
func retry(f func() error) {
	count := 0

	for {
		if err := f(); err == nil {
			break
		}

		count++

		if retries == count {
			break
		}

		value := time.Duration(math.Pow(3, float64(count)))

		time.Sleep(value * time.Second)
	}
}
