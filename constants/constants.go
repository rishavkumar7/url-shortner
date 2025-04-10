package constants

const (
	BaseUrl                  = "https://url-shortner-bjjsj2rjo-rishav-kumars-projects-4b2abfc6.vercel.app/"
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
