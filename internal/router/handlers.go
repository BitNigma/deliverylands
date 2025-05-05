package router

import (
	"delivery/internal/views/index"

	"github.com/gofiber/fiber/v2"
)

type Auth struct {
	Islogin    bool
	Email      string
	KYC        bool
	ErrorPass  error
	ErrorEmail error
}

// Handle main page
func HandleHome(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.MainPage(c).Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandleAbout(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.AboutPage(c).Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandleContacts(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.ContactsPage(c).Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandleTerms(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.TermsPage(c).Render(c.UserContext(), c.Context().Response.BodyWriter())
}

func HandlePolicy(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.PrivacyHanle(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleWallet(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.WalletPage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleDelivery(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.DeliveryPage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandlePlans(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.PlansPage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleSecurity(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.SecurityHanle(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleFraud(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.FraudHanle(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleClients(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.APIHandle(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleCompliance(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.HandleCompliance(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandlePayment(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.PaymentPage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleB2B(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.B2BPage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleExchange(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.ExchangePage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleMessage(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.WalletPage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleFAQ(c *fiber.Ctx) error {
	c.Set("Content-type", "text/html")
	return index.FaqPage(c).Render(c.UserContext(), c.Response().BodyWriter())
}

func HandleWallets(c *fiber.Ctx) error {
	return c.Redirect("/personal/kyc")
}
