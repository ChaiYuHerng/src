/*
 * Nudm_UECM
 *
 * Nudm Context Management Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package uecontextmanagement

import (
	"free5gc/lib/logger_util"
	"free5gc/src/udm/logger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// NewRouter returns a new router.
func NewRouter() *gin.Engine {
	router := logger_util.NewGinWithLogrus(logger.GinLog)
	AddService(router)
	return router
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nudm-uecm/v1")

	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Pattern, route.HandlerFunc)
		case "POST":
			group.POST(route.Pattern, route.HandlerFunc)
		case "PUT":
			group.PUT(route.Pattern, route.HandlerFunc)
		case "DELETE":
			group.DELETE(route.Pattern, route.HandlerFunc)
		case "PATCH":
			group.PATCH(route.Pattern, route.HandlerFunc)
		}
	}

	return group
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		"GET",
		"/",
		Index,
	},

	{
		"GetAmf3gppAccess",
		strings.ToUpper("Get"),
		"/:ueId/registrations/amf-3gpp-access",
		HTTPGetAmf3gppAccess,
	},

	{
		"GetAmfNon3gppAccess",
		strings.ToUpper("Get"),
		"/:ueId/registrations/amf-non-3gpp-access",
		HTTPGetAmfNon3gppAccess,
	},

	{
		"RegistrationAmf3gppAccess",
		strings.ToUpper("Put"),
		"/:ueId/registrations/amf-3gpp-access",
		HTTPRegistrationAmf3gppAccess,
	},

	{
		"Register",
		strings.ToUpper("Put"),
		"/:ueId/registrations/amf-non-3gpp-access",
		HTTPRegistrationAmfNon3gppAccess,
	},

	{
		"UpdateAmf3gppAccess",
		strings.ToUpper("Patch"),
		"/:ueId/registrations/amf-3gpp-access",
		HTTPUpdateAmf3gppAccess,
	},

	{
		"UpdateAmfNon3gppAccess",
		strings.ToUpper("Patch"),
		"/:ueId/registrations/amf-non-3gpp-access",
		HTTPUpdateAmfNon3gppAccess,
	},

	{
		"DeregistrationSmfRegistrations",
		strings.ToUpper("Delete"),
		"/:ueId/registrations/smf-registrations/:pduSessionId",
		HTTPDeregistrationSmfRegistrations,
	},

	{
		"RegistrationSmfRegistrations",
		strings.ToUpper("Put"),
		"/:ueId/registrations/smf-registrations/:pduSessionId",
		HTTPRegistrationSmfRegistrations,
	},

	{
		"GetSmsf3gppAccess",
		strings.ToUpper("Get"),
		"/:ueId/registrations/smsf-3gpp-access",
		HTTPGetSmsf3gppAccess,
	},

	{
		"DeregistrationSmsf3gppAccess",
		strings.ToUpper("Delete"),
		"/:ueId/registrations/smsf-3gpp-access",
		HTTPDeregistrationSmsf3gppAccess,
	},

	{
		"DeregistrationSmsfNon3gppAccess",
		strings.ToUpper("Delete"),
		"/:ueId/registrations/smsf-non-3gpp-access",
		HTTPDeregistrationSmsfNon3gppAccess,
	},

	{
		"GetSmsfNon3gppAccess",
		strings.ToUpper("Get"),
		"/:ueId/registrations/smsf-non-3gpp-access",
		HTTPGetSmsfNon3gppAccess,
	},

	{
		"UpdateSMSFReg3GPP",
		strings.ToUpper("Put"),
		"/:ueId/registrations/smsf-3gpp-access",
		HTTPUpdateSMSFReg3GPP,
	},

	{
		"RegistrationSmsfNon3gppAccess",
		strings.ToUpper("Put"),
		"/:ueId/registrations/smsf-non-3gpp-access",
		HTTPRegistrationSmsfNon3gppAccess,
	},
}