package entities

type Fintech struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Document     string `json:"document"`
	IconUrl      string `json:"iconUrl"`
	SocialReason string `json:"socialReason"`
}
