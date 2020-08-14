package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"

	"github.com/tidwall/gjson"
)

// series [1-8]
type FirstRoundSeries struct {
	toWin    [8]string
	numGames map[string]int
}

// series [9-14]
type MiddleSeries struct {
	toWin map[int]string
}

// series 15
type FinalSeries struct {
	toWin      [1]string
	totalGoals map[string]int
}

func main() {
	var userBracketNum string
	flag.StringVar(&userBracketNum, "userBracketNum", "1", "Bracket num to check against")
	var withTiebreaks bool
	flag.BoolVar(&withTiebreaks, "withTiebreaks", false, "Compare with tiebreakers")
	flag.Parse()

	fmt.Printf("userBracketNum: %v\n", userBracketNum)
	fmt.Printf("withTiebreaks: %v\n", withTiebreaks)

	userQuarterFinals := FirstRoundSeries{
		toWin:    [8]string{},
		numGames: map[string]int{},
	}

	userMiddleSeries := MiddleSeries{
		toWin: map[int]string{},
	}

	userFinals := FinalSeries{
		toWin:      [1]string{},
		totalGoals: map[string]int{},
	}

	byteData, err := ioutil.ReadFile("./brackets/" + userBracketNum + ".txt")
	if err != nil {
		fmt.Println("1Error reading file: ", "./brackets/"+userBracketNum+".txt => ", err)
	}
	userBracket := string(byteData)

	for seriesNum := 1; seriesNum <= 15; seriesNum++ {
		seriesID := "series_id" + strconv.Itoa(seriesNum)
		pickID := gjson.Get(userBracket, seriesID+".pickId").String()
		teamAbbrv := getTeamAbbrv(pickID)
		if seriesNum >= 1 && seriesNum <= 8 {
			userQuarterFinals.toWin[seriesNum-1] = teamAbbrv
			numGames := int(gjson.Get(userBracket, seriesID+".tb1").Int())
			userQuarterFinals.numGames[teamAbbrv] = numGames
			continue
		} else if seriesNum >= 9 && seriesNum < 15 {
			userMiddleSeries.toWin[seriesNum-1] = teamAbbrv
			continue
		} else {
			userFinals.toWin[0] = teamAbbrv
			totalGoals := int(gjson.Get(userBracket, seriesID+".tb2").Int())
			userFinals.totalGoals[teamAbbrv] = totalGoals
		}
	}

	totalMatches := 0
	counter := 1
	max := 458223
	// max := 458223
	/////////////////////////////////////////////////////////////////////
	for ; counter <= max; counter++ {
		stringUserBracketNum, err := strconv.Atoi(userBracketNum)
		if err != nil {
			fmt.Printf("Error onverting braketNum from string to int: %v", err.Error())
		}
		if counter == stringUserBracketNum {
			continue
		}
		quarterFinals := FirstRoundSeries{
			toWin:    [8]string{},
			numGames: map[string]int{},
		}

		middleSeries := MiddleSeries{
			toWin: map[int]string{},
		}

		finals := FinalSeries{
			toWin:      [1]string{},
			totalGoals: map[string]int{},
		}

		bracketNum := strconv.Itoa(counter)
		byteData, err = ioutil.ReadFile("./brackets/" + bracketNum + ".txt")
		if err != nil {
			fmt.Println("2Error reading file: ", "./brackets/"+bracketNum+".txt => ", err)
		}
		data := string(byteData)
		for seriesNum := 1; seriesNum <= 15; seriesNum++ {
			seriesID := "series_id" + strconv.Itoa(seriesNum)
			pickID := gjson.Get(data, seriesID+".pickId").String()
			teamAbbrv := getTeamAbbrv(pickID)
			if seriesNum >= 1 && seriesNum <= 8 {
				quarterFinals.toWin[seriesNum-1] = teamAbbrv
				numGames := int(gjson.Get(data, seriesID+".tb1").Int())
				quarterFinals.numGames[teamAbbrv] = numGames
				continue
			} else if seriesNum >= 9 && seriesNum < 15 {
				middleSeries.toWin[seriesNum-1] = teamAbbrv
				continue
			} else {
				finals.toWin[0] = teamAbbrv
				totalGoals := int(gjson.Get(data, seriesID+".tb2").Int())
				finals.totalGoals[teamAbbrv] = totalGoals
			}
		}
		totalMatches += compareBracket(quarterFinals, middleSeries, finals, userQuarterFinals, userMiddleSeries, userFinals, withTiebreaks)
	}
	fmt.Printf("Matches: %v\n", totalMatches)
}

func compareBracket(qf FirstRoundSeries, ms MiddleSeries, f FinalSeries, uqf FirstRoundSeries, ums MiddleSeries, uf FinalSeries, useTiebreaks bool) int {
	fmt.Println(uf.toWin[0])
	if uf.toWin[0] != f.toWin[0] {
		return 0
	}

	fmt.Println(ums.toWin)
	if !reflect.DeepEqual(ums.toWin, ms.toWin) {
		return 0
	}

	for seriesNum := 1; seriesNum <= 8; seriesNum++ {
		fmt.Println(uqf.toWin[seriesNum-1])
		if uqf.toWin[seriesNum-1] != qf.toWin[seriesNum-1] {
			return 0
		}
	}

	return 1
}

func getTeamAbbrv(teamNum string) string {
	teams := map[string]string{"2": "ARI", "3": "BOS", "5": "CGY", "6": "CAR", "7": "CHI", "8": "COL", "9": "CBJ", "10": "DAL", "16": "MTL", "19": "NYI", "22": "PHI", "25": "STL", "26": "TBL", "28": "VAN", "29": "VGK", "30": "WSH"}
	return teams[teamNum]
}
