package models

type Asset struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Chart       *Chart    `json:"chart,omitempty"`
	Insight     *Insight  `json:"insight,omitempty"`
	Audience    *Audience `json:"audience,omitempty"`
}

type Chart struct {
	Title      string `json:"title"`
	AxesTitles string `json:"axes_titles"`
	Data       string `json:"data"`
}

type Insight struct {
	Text string `json:"text"`
}

type Audience struct {
	Gender             string `json:"gender"`
	BirthCountry       string `json:"birth_country"`
	AgeGroup           string `json:"age_group"`
	HoursOnSocialMedia int    `json:"hours_on_social_media"`
	PurchasesLastMonth int    `json:"purchases_last_month"`
}
