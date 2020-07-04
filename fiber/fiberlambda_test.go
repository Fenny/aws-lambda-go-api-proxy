package fiberadapter_test

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	fiberadaptor "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FiberLambda tests", func() {
	Context("Simple ping request", func() {
		It("Proxies the event correctly", func() {
			log.Println("Starting test")
			app := fiber.New()
			app.Get("/ping", func(c *fiber.Ctx) {
				log.Println("Handler!!")
				c.SendString("pong")
			})

			adapter := fiberadaptor.New(app)

			req := events.APIGatewayProxyRequest{
				Path:       "/ping",
				HTTPMethod: "GET",
			}

			resp, err := adapter.Proxy(req)

			Expect(err).To(BeNil())
			Expect(resp.StatusCode).To(Equal(200))
		})
	})
})
