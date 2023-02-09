package main

import (
	"fmt"
	"go-fintechs-api/entities"
	"go-fintechs-api/routes"
)

func main() {
	entities.Fintechs = []entities.Fintech{{Name: "Nubank", Document: "18.236.120/0001-58", IconUrl: "https://commons.wikimedia.org/wiki/File:Nubank_logo_2021.svg#/media/File:Nubank_logo_2021.svg", SocialReason: "NU FINANCEIRA S.A. – SOCIEDADE DE CRÉDITO, FINANCIAMENTO E INVESTIMENTO"}}
	fmt.Println("init go server")
	routes.HandleRequest()
}
