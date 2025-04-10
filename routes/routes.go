package routes

import (
	"net/http"

	"github.com/rishavkumar7/url-shortner/constants"
	"github.com/rishavkumar7/url-shortner/controllers"
)

var urlShortner = Routes{
	Route{
		Name:        "Url Shortner Service",
		Method:      http.MethodPost,
		Pattern:     constants.UrlShortnerPath,
		HandlerFunc: controllers.ShortTheUrl,
	},
	Route{
		Name:        "Redirect to Long url",
		Method:      http.MethodGet,
		Pattern:     constants.RedirectToLongUrlPath,
		HandlerFunc: controllers.RedirectToLongUrl,
	},
}
