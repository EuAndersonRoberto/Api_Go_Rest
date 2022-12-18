package models

type Personalidade struct {
	Id       int    `json:"id"`       //Aqui definmos com o json como será exibido dentro da nossa API.
	Nome     string `json:"nome"`     //Aqui definmos com o json como será exibido dentro da nossa API.
	Historia string `json:"historia"` //Aqui definmos com o json como será exibido dentro da nossa API.
}

var Personalidades []Personalidade
