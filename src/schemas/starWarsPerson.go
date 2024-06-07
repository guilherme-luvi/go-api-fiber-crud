package schemas

type StarWarsPerson struct {
	Name      string `json:"name"`
	Height    string `json:"height"`
	Mass      string `json:"mass"`
	HairColor string `json:"hair_color"`
	SkinColor string `json:"skin_color"`
	EyeColor  string `json:"eye_color"`
}
