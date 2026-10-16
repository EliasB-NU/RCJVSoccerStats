package data

import (
	"crypto/tls"

	"github.com/gofiber/fiber/v3/client"
)

// Preserve the TLS behavior of the previous Fiber v2 requests.
var catigoalClient = client.New().SetTLSConfig(&tls.Config{InsecureSkipVerify: true})
