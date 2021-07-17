package main

import "fmt"

func country_freq(list []string) map[string]int {

	freq := make(map[string]int)

	for _, item := range list {

		_, exist := freq[item]

		if exist {
			freq[item] += 1
		} else {
			freq[item] = 1
		}
	}
	return freq
}

func main() {
	mapa_cidades := make(map[string]string)
	mapa_cidades["Rio de Janeiro"] = "Brasil"
	mapa_cidades["Brasilia"] = "Brasil"
	mapa_cidades["Palmas"] = "Brasil"
	mapa_cidades["Manaus"] = "Brasil"
	mapa_cidades["Los Angeles"] = "EUA"
	mapa_cidades["Nova Iorque"] = "EUA"
	mapa_cidades["Kyoto"] = "Japão"

	pais_mapa := []string{}

	for _, v := range mapa_cidades {
		pais_mapa = append(pais_mapa, v)
	}

	frequencia := country_freq(pais_mapa)

	fmt.Println(frequencia)
}
