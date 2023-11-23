package main

import (
	"mobile-legend-api/api/routes"
	"mobile-legend-api/config"
	detailresult "mobile-legend-api/pkg/detail_result"
	"mobile-legend-api/pkg/hero"
	"mobile-legend-api/pkg/herorole"
	"mobile-legend-api/pkg/role"
	"mobile-legend-api/pkg/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){

	db:= config.GetConnection()
	userRepository := user.NewUserRepository()
	userService := user.NewUserService(userRepository,db)
	roleRepository:= role.NewRoleRepository()
	roleService := role.NewRoleService(roleRepository,db)
	heroroleRepository:= herorole.NewHeroRoleRepository()
	heroroleService := herorole.NewHeroRoleService(heroroleRepository,db)
	heroRepository:= hero.NewHeroRepository()
	heroService := hero.NewHeroService(heroRepository,db)

	detailResultRepository:= detailresult.NewDetailResultRepository()
	detailResultService := detailresult.NewDetailResultService(detailResultRepository,db)
	app:= fiber.New(
		fiber.Config{
			Prefork: true,
		})
	userApi:= app.Group("/api/v1/users")
	routes.UserRouter(userApi, userService)
	roleApi := app.Group("/api/v1/roles")
	routes.RoleRouter(roleApi,roleService)
	heroroleApi := app.Group("/api/v1/heroroles")
	routes.HeroRoleRouter(heroroleApi,heroroleService)
	heroApi := app.Group("/api/v1/heroes")
	routes.HeroRouter(heroApi,heroService)
	detailresultApi:= app.Group("/api/v1/detail-result")
	routes.DetailResultRouter(detailresultApi,detailResultService)
	app.Use(cors.New());
	app.Listen(":3000")

}
