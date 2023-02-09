package entities

type Fintech struct {
	Name         string `json:"name"`
	Document     string `json:"document"`
	IconUrl      string `json:"iconUrl"`
	SocialReason string `json:"socialReason"`
}

var Fintechs []Fintech
