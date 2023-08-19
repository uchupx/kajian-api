package kajian

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/kajian-api/internal"
)

func InitServer() {
	e := echo.New()

	//  intialize routing
	internal.InitRoute(e)
	e.Logger.Fatal(e.Start(":1323"))
}
