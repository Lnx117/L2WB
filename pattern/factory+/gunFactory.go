package main

import "fmt"

//Фабрика это функция которая вызывет конструктор конкреного оружия который возвращает объект любого оружия,
//но с необходимыми параметрами
func getGun(gunType string) (iGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
