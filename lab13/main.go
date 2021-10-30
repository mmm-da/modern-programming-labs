package main

import (
	"encoding/json"
	"io/ioutil"
	"lab10/pkg/errorhandlers"
	"lab10/pkg/gem"
)

// Алмазный фонд.
// Корневой элемент назвать Gem. Драгоценные и полудрагоценные
// камни, содержащиеся в павильоне, имеют следующие характеристики:
// − Name – название камня;
// − Preciousness – может быть драгоценным либо полудрагоценным;
// − Origin – место добывания;
// − Visual parameters (должно быть несколько) – могут быть: цвет
// (зеленый, красный, желтый и т. д.), прозрачность (измеряется в процентах 0–
// 100%), способы огранки (количество граней 4–15);
// − Value – вес камня (измеряется в каратах).
// С помощью XSL преобразовать XML-файл в формат XML, где корневым элементом будет место происхождения.

func main() {
	jsonBytes := make([]byte, 0)
	jsonBytes, err := ioutil.ReadFile("test.json")
	errorhandlers.PanicIfErr(err)

	gemArray := make([]gem.Gem, 0)
	err = json.Unmarshal(jsonBytes, &gemArray)
	errorhandlers.PanicIfErr(err)

	gemMap := make(map[string][]gem.Gem)

	for _, item := range gemArray {
		if val, ok := gemMap[item.Origin]; ok {
			gemMap[item.Origin] = append(val, item)
		} else {
			gemMap[item.Origin] = make([]gem.Gem, 0)
			gemMap[item.Origin] = append(val, item)
		}
	}

	jsonBytes, err = json.MarshalIndent(gemMap, "", "  ")
	errorhandlers.PanicIfErr(err)

	ioutil.WriteFile("result.json", jsonBytes, 0655)
}
