package types

var LogInAPI = "http://localhost:3000/api/register"

var ListKittyAPI = "http://localhost:3000/api/kitties"

type ListKittyResponse []struct {
	Name       string `json:"name"`
	PictureURL string `json:"pictureUrl"`
}

type CatRank struct {
	Name       string `json:"name"`
	Rank         int   `json:"Rank"`
	Awesomeness  int   `json:"Awesomeness"`
}

type CatNewName struct {
	NewName string `json:"newName"`
}