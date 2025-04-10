package constants

const (
	BaseUrl                  = "https://url-shortner-mini.vercel.app/"
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
