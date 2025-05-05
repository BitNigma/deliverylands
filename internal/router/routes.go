package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store *session.Store
)

func Setup(app *fiber.App) {

	/* Sessions Config */
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		// CookieSecure: true, for https
		Expiration:   time.Hour * 2,
		CookieSecure: true,
	})

	// Static files
	app.Static("/static", "./static")

	/* Views */

	app.Get("/", HandleHome)
	app.Get("/about", HandleAbout)
	app.Get("/contacts", HandleContacts)
	app.Get("/terms", HandleTerms)
	app.Get("/privacy", HandlePolicy)
	app.Get("/fraud", HandleFraud)
	app.Get("/security", HandleSecurity)
	app.Get("/b2b", HandleB2B)
	app.Get("/payment", HandlePayment)
	app.Get("/delivery", HandleDelivery)
	app.Get("/plans", HandlePlans)
	app.Get("/payments", HandlePayment)
	app.Get("/clients", HandleClients)
	app.Get("/compliance", HandleCompliance)
	app.Get("/rules", HandleCompliance)
	app.Get("/exchange", HandleExchange)
	app.Get("/faq", HandleFAQ)
	/* Page Not und Management */

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendFile("./static/page-error-404.html")
	})
}
