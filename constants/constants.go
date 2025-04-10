package constants

const (
	BaseUrl                  = "http://localhost:9000/"
	UrlShortnerPath          = "/short"
	RedirectToLongUrlPath    = "/:code"
	ShortUrlCodeCharacterSet = "abcdefghijklmnopqrstuvwxyz"
	ShortUrlCodeLength       = 6
)

const (
	Database      = "urlshortner"
	UrlCollection = "urls"
)

const (
	BindError     = "Error while binding the request body"
	ExistError    = "Url code is already in use"
	NotFoundError = "Url not found"
)
