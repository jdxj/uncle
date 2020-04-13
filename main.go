package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Data 结构是数据的包装
type Data struct {
	Data map[string]*Province `json:"data"`
}

// Province 省
type Province struct {
	Name   string           `json:"name"`
	Cities map[string]*City `json:"cities"`
}

// City 市
type City struct {
	Name     string     `json:"name"`
	Airports []*Airport `json:"airports"`
}

// Airport 机场
type Airport struct {
	Name  string   `json:"name"`
	Areas []string `json:"areas"` // 区
}

func main() {
	path := "jc.txt"
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("can't found '%s' file", path))
	}
	defer file.Close()

	buf := bufio.NewReader(file)

	result := &Data{
		Data: make(map[string]*Province),
	}
	dataMap := result.Data

	var data string
	for data, err = buf.ReadString('\n'); err == nil; data, err = buf.ReadString('\n') {
		parts := strings.Split(data, ";")
		if len(parts) != 5 {
			panic(fmt.Sprintf("data format error"))
		}

		provinceName := parts[0]
		if province, ok := dataMap[provinceName]; !ok {
			province = &Province{
				Name:   provinceName,
				Cities: make(map[string]*City),
			}

			city := &City{
				Name:     parts[1],
				Airports: make([]*Airport, 0),
			}

			airport := &Airport{
				Name:  parts[2],
				Areas: strings.Split(parts[3], ","),
			}

			city.Airports = append(city.Airports, airport)
			province.Cities[city.Name] = city
			dataMap[provinceName] = province

		} else {
			cityName := parts[1]
			if city, ok := province.Cities[cityName]; !ok {
				city = &City{
					Name:     cityName,
					Airports: make([]*Airport, 0),
				}

				airport := &Airport{
					Name:  parts[2],
					Areas: strings.Split(parts[3], ","),
				}

				city.Airports = append(city.Airports, airport)
				province.Cities[cityName] = city

			} else {
				airport := &Airport{
					Name:  parts[2],
					Areas: strings.Split(parts[3], ","),
				}

				city.Airports = append(city.Airports, airport)
			}
		}
	}

	print, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("%s\n", print)
}
