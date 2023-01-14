package api

import "fmt"

func (api *Api) startRouting() {
	v1Group := api.router.Group("/api/v1")
	usersGroup := v1Group.Group("/holiday")

	fmt.Println()

	usersGroup.GET("", api.handler.GetHoliday)

	api.logger.Info(
		fmt.Sprintf("Loading routes, base path: %s", usersGroup.BasePath()),
	)
}
