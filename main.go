package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/justtaldevelops/hcaptcha-solver-go"
	"strconv"
	"time"
)

// SolveRequest is a request sent by a user to the solver.
type SolveRequest struct {
	SiteURL string                 `json:"site_url"`
	Options hcaptcha.SolverOptions `json:"options"`
	Proxies []string               `json:"proxies"`
}

// SolveResponse is a response to the SolveRequest.
type SolveResponse struct {
	CaptchaCode string `json:"captcha_code"`
}

func main() {
	con := loadConfig()

	app := fiber.New()

	app.Post("/solve", func(c *fiber.Ctx) error {
		if string(c.Request().Header.Peek("authorization")) != con.AuthorizationHeader {
			return c.SendStatus(fiber.StatusForbidden)
		}

		var request SolveRequest
		err := json.Unmarshal(c.Body(), &request)
		if err != nil {
			return c.SendString(err.Error())
		}
		var s *hcaptcha.Solver
		if len(request.Proxies) != 0 {
			s, err = hcaptcha.NewSolverWithProxies(request.SiteURL, request.Proxies, request.Options)
			if err != nil {
				return c.SendString(err.Error())
			}
		} else {
			s, err = hcaptcha.NewSolver(request.SiteURL, request.Options)
			if err != nil {
				return c.SendString(err.Error())
			}
		}
		solution, err := s.Solve(time.Now().Add(time.Duration(con.SolveTimeout) * time.Second))
		if err != nil {
			return c.SendString(err.Error())
		}
		b, err := json.Marshal(SolveResponse{CaptchaCode: solution})
		if err != nil {
			return c.SendString(err.Error())
		}
		return c.Send(b)
	})

	fmt.Println("hCaptcha solver API is now running on port " + strconv.Itoa(con.Port) + "!")

	app.Listen(":" + strconv.Itoa(con.Port))
}
