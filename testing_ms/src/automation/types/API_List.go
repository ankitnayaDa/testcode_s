package types

var LogInAPI = "http://localhost:3000/api/register"

var ListKittyAPI = "http://localhost:3000/api/kitties"

type ListKittyResponse []struct {
	Name       string `json:"name"`
	PictureURL string `json:"pictureUrl"`
}

