package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	counter := 229112
	max := 458223
	// 229111
	// max := 458223

	for ; counter <= max; counter++ {
		bracketNum := strconv.Itoa(counter)
		byteData, err := ioutil.ReadFile("./brackets/" + bracketNum + ".txt")
		if err != nil {
			fmt.Println("Error reading file: ", "./brackets/"+bracketNum+".txt => ", err)
		}

		err = ioutil.WriteFile("./brackets-2/"+bracketNum+".txt", byteData, 0644)
		if err != nil {
			fmt.Println(bracketNum)
			return
		}
		// data, err := ioutil.ReadFile("./brackets/" + bracketNum + ".txt")

		// name = gjson.Get(string(data), "meta.name").String()
		// fmt.Println(name)

	}

	// 458223

}
