package main

import (
	"fmt"
	"strconv"
	_ "strconv"
)

var numSystems = [4]string{"Binary", "Decimal", "Octal", "Hexadecimal"}

//func printMenu(){
//	fmt.Println("1.Decimal to Binary 2.Decimal to Hexadecimal 3.Decimal to octal")
//	fmt.Println("4.Binary to Decimal 5.Binary to Hexadecimal 6.Binary to octal")
//	fmt.Println("7.Hexadecimal to Decimal 8.Hexadecimal to Binary 9.Hexadecimal to Octal")
//	fmt.Println("10.Octal to Decimal 11.Octal to Binary 12.Octal to Hexadecimal")
//	fmt.Println("13.Exit\n")
//
//}

//func getParams(choice string) Params{
//
//}
//
//func promptInput(prompt string) string{
//	fmt.Println(prompt)
//	return prompt
//}
//

type Params struct {
	frm, to string
}

//Takes a number as a string to be converted
//Uses params to decide how to perform the conversion
//func convert(number string, params Params) string{
//
//	return " "
//}
//
func createMenuDict() map[string]Params {
	menuDict := make(map[string]Params)

	index := 1
	for _, xVal := range numSystems {
		for _, yVal := range numSystems {
			//Ignore redu
			if xVal != yVal {
				//Todo: Use to populate a map of params for convert function
				choice := strconv.Itoa(index)
				//fmt.Printf("%s %s %s\n",choice,xVal,yVal)
				menuDict[choice] = Params{xVal, yVal}
				index++
			}
		}
	}
	return menuDict
}

func makeMenu(menuDict map[string]Params) string {
	var menu string
	for index := 1; index < 13; index++ {

		choice := strconv.Itoa(index)
		params := menuDict[choice]
		menu += fmt.Sprintf("%s. %s to %s ", choice, params.frm, params.to)
		if index%3 == 0 {
			menu += "\n"
		}

	}
	menu += "\n13.Exit"

	return menu
}

func main() {

	//Initializatio
	menuDict := createMenuDict()
	menu := makeMenu(menuDict)

	fmt.Println(menu)

}
