package models

type Portfolio struct {
	Img			string `json:"img"`
	Title			string `json:"title"`
	Description 	string `json:"description"`
	Category 	string `json:"categorys"`
	Link		string `json:"link"`
}

type PortfolioList struct {
	Portfolio []Portfolio `json:"portfolio"`
}