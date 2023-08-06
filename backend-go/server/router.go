package server

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/Hotpot-protocol1/hotpot-global/config"
	"github.com/Hotpot-protocol1/hotpot-global/db"
	user "github.com/Hotpot-protocol1/hotpot-global/server/handlers/tickets"
	eventservice "github.com/Hotpot-protocol1/hotpot-global/services/event"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) (err error) {
	err = cv.validator.Struct(i)
	return
}

func Router(cfg config.Conf, db db.DBHandler) (*echo.Echo, error) {
	router := echo.New()
	router.Validator = &CustomValidator{validator: validator.New()}
	cors := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*", "GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"*", "Accept", "Authorization", "Content-Type", "X-CSRF-Token", "x-auth", "Access-Control-Allow-Origin", "Access-Control-Allow-Methods", "Access-Control-Allow-Credentials"},
		ExposeHeaders:    []string{"*", "Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(
		cors,
		middleware.Recover(),
		middleware.Logger(),
	)

	log := cfg.Log.New()
	userHandler := user.New(db, log)
	infura := eventservice.InitializeInfura(cfg.ProxyContract, cfg.Infura.BaseURLWS, cfg.Infura.APIKey)
	infura.Start(db.UserTickets(), log)

	// buyer := "0xB838b0b5Ff5f856b6defb75e843fd7D8d606f856"
	// seller := "0xB203a89D86B6B0F8fa65b278A97D835DF1C58c96"
	// operator := contractservice.InitializeOperator(cfg.Contract.OperatorPrivKey, cfg.Infura.BaseURL, cfg.Infura.APIKey)
	// err := operator.Execute(buyer, seller, 250000000000000000)
	// if err != nil {
	// 	fmt.Println("Error is ", err)
	// }

	// USER endpoints
	router.GET("/user/:wallet_address/pot/:pot_id", userHandler.GetUserTicketsForPot)
	router.GET("/user/:wallet_address/pot", userHandler.GetPotsWithRaffleTimestamp)

	// DEBUG endpoints
	router.GET("/status", func(c echo.Context) error {
		return c.JSON(http.StatusOK, `{"status": "ok"}`)
	})

	router.GET("/try", func(c echo.Context) error {
		// err := operator.Execute(buyer, seller, 250000000000000000)
		// if err != nil {
		// 	fmt.Println("Error is ", err)
		// }

		return c.NoContent(http.StatusOK)
	})

	router.GET("/debug/*", echo.WrapHandler(http.DefaultServeMux))

	return router, nil
}
