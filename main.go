package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/tidwall/gjson"
)

func main() {
	url := "https://bracketchallenge.nhl.com/brackets/"
	client := &http.Client{}
	counter := 1
	max := 458223

	// max := 458223

	for ; counter <= max; counter++ {
		bracketNum := strconv.Itoa(counter)
		req, err := http.NewRequest("GET", url+bracketNum, nil)
		if err != nil {
			fmt.Println(bracketNum)
			return
		}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("error getting response from client when streaming from location:\n%v\n%s", response, err.Error())
			fmt.Println(bracketNum)
			return
		}
		if response.StatusCode != 200 {
			fmt.Println("failed response (stream from external source): status code %d", response.StatusCode)
			fmt.Println(bracketNum)
			return
		}

		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("error reading data from response body:\n%s", err.Error())
			fmt.Println(bracketNum)
			return
		}
		strContent := string(content)
		re := regexp.MustCompile(`"bracket":{"info":{.*},`)
		bracket := re.FindString(strContent)
		picks := gjson.Get(bracket, "picks").Array()
		name := gjson.Get(bracket, "info.name").String()
		possiblePoints := gjson.Get(bracket, "info.possible_points").String()
		userPicks := map[string]map[string]string{}
		userPicks["meta"] = map[string]string{"name": name, "possible_points": possiblePoints}
		for _, ele := range picks {
			seriesId := ele.Get("series_id").String()
			pickId := ele.Get("pick_id").String()
			tb1 := ele.Get("tie_breaker").String()
			tb2 := ele.Get("tie_breaker_2").String()
			userPicks["series_id"+seriesId] = map[string]string{"pickId": pickId, "tb1": tb1, "tb2": tb2}
		}

		contentJSON, err := json.Marshal(userPicks)
		// fmt.Println(string(contentJSON))
		// fmt.Println(bracket + "\n")
		// d1 := []byte(bracket)
		err = ioutil.WriteFile("./brackets/"+bracketNum+".txt", contentJSON, 0644)
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
