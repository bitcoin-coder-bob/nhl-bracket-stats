package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"

	"github.com/tidwall/gjson"
)

func main() {
	counter := 1
	max := 458223
	// max := 458223

	// series [1-8]
	type FirstRoundSeries struct {
		toWin    map[string]int
		numGames map[string]map[int]int
	}

	// series [9-14]
	type MiddleSeries struct {
		toWin map[int]map[string]int
	}

	// series 15
	type FinalSeries struct {
		toWin      map[string]int
		totalGoals map[string]map[int]int
	}

	quarterFinals := FirstRoundSeries{
		toWin:    map[string]int{},
		numGames: map[string]map[int]int{},
	}

	middleSeries := MiddleSeries{
		toWin: map[int]map[string]int{},
	}

	finals := FinalSeries{
		toWin:      map[string]int{},
		totalGoals: map[string]map[int]int{},
	}

	for ; counter <= max; counter++ {
		bracketNum := strconv.Itoa(counter)
		byteData, err := ioutil.ReadFile("./brackets/" + bracketNum + ".txt")
		if err != nil {
			fmt.Println("Error reading file: ", "./brackets/"+bracketNum+".txt => ", err)
		}
		data := string(byteData)

		for seriesNum := 1; seriesNum <= 15; seriesNum++ {
			seriesID := "series_id" + strconv.Itoa(seriesNum)
			pickID := gjson.Get(data, seriesID+".pickId").String()
			teamAbbrv := getTeamAbbrv(pickID)
			if seriesNum >= 1 && seriesNum <= 8 {
				quarterFinals.toWin[teamAbbrv]++
				numGames := int(gjson.Get(data, seriesID+".tb1").Int())
				if quarterFinals.numGames[teamAbbrv] == nil {
					quarterFinals.numGames[teamAbbrv] = map[int]int{numGames: 1}
				} else {
					quarterFinals.numGames[teamAbbrv][numGames]++
				}
				continue
			} else if seriesNum >= 9 && seriesNum < 15 {
				if middleSeries.toWin[seriesNum] == nil {
					middleSeries.toWin[seriesNum] = map[string]int{teamAbbrv: 1}
				} else {
					middleSeries.toWin[seriesNum][teamAbbrv]++
				}
				continue
			} else {
				finals.toWin[teamAbbrv]++
				totalGoals := int(gjson.Get(data, seriesID+".tb2").Int())
				if finals.totalGoals[teamAbbrv] == nil {
					finals.totalGoals[teamAbbrv] = map[int]int{totalGoals: 1}
				} else {
					finals.totalGoals[teamAbbrv][totalGoals]++
				}
			}
		}
	}
	fmt.Println("--- ROUND OF 16 ---\n\n")
	fmt.Println("Vegas Golden Knights ------ Chicago Blackhawks")
	fmt.Println("-------------------------------------------")
	// numGames := quarterFinals.numGames["VGK"][4] + quarterFinals.numGames["VGK"][5] + quarterFinals.numGames["VGK"][6] + quarterFinals.numGames["VGK"][7] + quarterFinals.numGames["CHI"][4] + quarterFinals.numGames["CHI"][5] + quarterFinals.numGames["CHI"][6] + quarterFinals.numGames["CHI"][7]
	fmt.Printf("Win in 4 games: %v (%.2f%%)-------- Win in 4 games: %v (%.2f%%)\n", quarterFinals.numGames["VGK"][4], float64(quarterFinals.numGames["VGK"][4])/float64(max)*100, quarterFinals.numGames["CHI"][4], float64(quarterFinals.numGames["CHI"][4])/float64(max)*100)
	fmt.Printf("Win in 5 games: %v (%.2f%%)-------- Win in 5 games: %v (%.2f%%)\n", quarterFinals.numGames["VGK"][5], float64(quarterFinals.numGames["VGK"][5])/float64(max)*100, quarterFinals.numGames["CHI"][5], float64(quarterFinals.numGames["CHI"][5])/float64(max)*100)
	fmt.Printf("Win in 6 games: %v (%.2f%%)-------- Win in 6 games: %v (%.2f%%)\n", quarterFinals.numGames["VGK"][6], float64(quarterFinals.numGames["VGK"][6])/float64(max)*100, quarterFinals.numGames["CHI"][6], float64(quarterFinals.numGames["CHI"][6])/float64(max)*100)
	fmt.Printf("Win in 7 games: %v (%.2f%%)-------- Win in 7 games: %v (%.2f%%)\n", quarterFinals.numGames["VGK"][7], float64(quarterFinals.numGames["VGK"][7])/float64(max)*100, quarterFinals.numGames["CHI"][7], float64(quarterFinals.numGames["CHI"][7])/float64(max)*100)
	// fmt.Printf("No selection made: %v (%.2f%%)\n", max-numGames, float64(numGames)/float64(max)*100)
	fmt.Println("=============================================")
	fmt.Println("Colorado Avalanche --------- Arizona Coyotes")
	fmt.Println("-------------------------------------------")
	// numGames = quarterFinals.numGames["COL"][4] + quarterFinals.numGames["COL"][5] + quarterFinals.numGames["COL"][6] + quarterFinals.numGames["COL"][7] + quarterFinals.numGames["ARI"][4] + quarterFinals.numGames["ARI"][5] + quarterFinals.numGames["ARI"][6] + quarterFinals.numGames["ARI"][7]
	fmt.Printf("Win in 4 games: %v (%.2f%%)-------- Win in 4 games: %v (%.2f%%)\n", quarterFinals.numGames["COL"][4], float64(quarterFinals.numGames["COL"][4])/float64(max)*100, quarterFinals.numGames["ARI"][4], float64(quarterFinals.numGames["ARI"][4])/float64(max)*100)
	fmt.Printf("Win in 5 games: %v (%.2f%%)-------- Win in 5 games: %v (%.2f%%)\n", quarterFinals.numGames["COL"][5], float64(quarterFinals.numGames["COL"][5])/float64(max)*100, quarterFinals.numGames["ARI"][5], float64(quarterFinals.numGames["ARI"][5])/float64(max)*100)
	fmt.Printf("Win in 6 games: %v (%.2f%%)-------- Win in 6 games: %v (%.2f%%)\n", quarterFinals.numGames["COL"][6], float64(quarterFinals.numGames["COL"][6])/float64(max)*100, quarterFinals.numGames["ARI"][6], float64(quarterFinals.numGames["ARI"][6])/float64(max)*100)
	fmt.Printf("Win in 7 games: %v (%.2f%%)-------- Win in 7 games: %v (%.2f%%)\n", quarterFinals.numGames["COL"][7], float64(quarterFinals.numGames["COL"][7])/float64(max)*100, quarterFinals.numGames["ARI"][7], float64(quarterFinals.numGames["ARI"][7])/float64(max)*100)
	// fmt.Printf("No selection made: %v (%.2f%%)\n", max-numGames, float64(numGames)/float64(max)*100)
	fmt.Println("=============================================")
	fmt.Println("Dallas Stars --------------- Calgary Flames")
	fmt.Println("-------------------------------------------")
	// numGames = quarterFinals.numGames["DAL"][4] + quarterFinals.numGames["DAL"][5] + quarterFinals.numGames["DAL"][6] + quarterFinals.numGames["DAL"][7] + quarterFinals.numGames["CGY"][4] + quarterFinals.numGames["CGY"][5] + quarterFinals.numGames["CGY"][6] + quarterFinals.numGames["CGY"][7]
	fmt.Printf("Win in 4 games: %v (%.2f%%)-------- Win in 4 games: %v (%.2f%%)\n", quarterFinals.numGames["DAL"][4], float64(quarterFinals.numGames["DAL"][4])/float64(max)*100, quarterFinals.numGames["CGY"][4], float64(quarterFinals.numGames["CGY"][4])/float64(max)*100)
	fmt.Printf("Win in 5 games: %v (%.2f%%)-------- Win in 5 games: %v (%.2f%%)\n", quarterFinals.numGames["DAL"][5], float64(quarterFinals.numGames["DAL"][5])/float64(max)*100, quarterFinals.numGames["CGY"][5], float64(quarterFinals.numGames["CGY"][5])/float64(max)*100)
	fmt.Printf("Win in 6 games: %v (%.2f%%)-------- Win in 6 games: %v (%.2f%%)\n", quarterFinals.numGames["DAL"][6], float64(quarterFinals.numGames["DAL"][6])/float64(max)*100, quarterFinals.numGames["CGY"][6], float64(quarterFinals.numGames["CGY"][6])/float64(max)*100)
	fmt.Printf("Win in 7 games: %v (%.2f%%)-------- Win in 7 games: %v (%.2f%%)\n", quarterFinals.numGames["DAL"][7], float64(quarterFinals.numGames["DAL"][7])/float64(max)*100, quarterFinals.numGames["CGY"][7], float64(quarterFinals.numGames["CGY"][7])/float64(max)*100)
	// fmt.Printf("No selection made: %v (%.2f%%)\n", max-numGames, float64(numGames)/float64(max)*100)
	fmt.Println("=============================================")
	fmt.Println("St. Louis Blues --------- Vancouver Canucks")
	fmt.Println("-------------------------------------------")
	// numGames = quarterFinals.numGames["STL"][4] + quarterFinals.numGames["STL"][5] + quarterFinals.numGames["STL"][6] + quarterFinals.numGames["STL"][7] + quarterFinals.numGames["VAN"][4] + quarterFinals.numGames["VAN"][5] + quarterFinals.numGames["VAN"][6] + quarterFinals.numGames["VAN"][7]
	fmt.Printf("Win in 4 games: %v (%.2f%%)-------- Win in 4 games: %v (%.2f%%)\n", quarterFinals.numGames["STL"][4], float64(quarterFinals.numGames["STL"][4])/float64(max)*100, quarterFinals.numGames["VAN"][4], float64(quarterFinals.numGames["VAN"][4])/float64(max)*100)
	fmt.Printf("Win in 5 games: %v (%.2f%%)-------- Win in 5 games: %v (%.2f%%)\n", quarterFinals.numGames["STL"][5], float64(quarterFinals.numGames["STL"][5])/float64(max)*100, quarterFinals.numGames["VAN"][5], float64(quarterFinals.numGames["VAN"][5])/float64(max)*100)
	fmt.Printf("Win in 6 games: %v (%.2f%%)-------- Win in 6 games: %v (%.2f%%)\n", quarterFinals.numGames["STL"][6], float64(quarterFinals.numGames["STL"][6])/float64(max)*100, quarterFinals.numGames["VAN"][6], float64(quarterFinals.numGames["VAN"][6])/float64(max)*100)
	fmt.Printf("Win in 7 games: %v (%.2f%%)-------- Win in 7 games: %v (%.2f%%)\n", quarterFinals.numGames["STL"][7], float64(quarterFinals.numGames["STL"][7])/float64(max)*100, quarterFinals.numGames["VAN"][7], float64(quarterFinals.numGames["VAN"][7])/float64(max)*100)
	// fmt.Printf("No selection made: %v (%.2f%%)\n", max-numGames, float64(numGames)/float64(max)*100)
	fmt.Println("=============================================")
	fmt.Println("Philadelphia Flyers ------ Montreal Canadiens")
	fmt.Println("-------------------------------------------")
	// numGames = quarterFinals.numGames["PHI"][4] + quarterFinals.numGames["PHI"][5] + quarterFinals.numGames["PHI"][6] + quarterFinals.numGames["PHI"][7] + quarterFinals.numGames["MTL"][4] + quarterFinals.numGames["MTL"][5] + quarterFinals.numGames["MTL"][6] + quarterFinals.numGames["MTL"][7]
	fmt.Printf("Win in 4 games: %v (%.2f%%)-------- Win in 4 games: %v (%.2f%%)\n", quarterFinals.numGames["PHI"][4], float64(quarterFinals.numGames["PHI"][4])/float64(max)*100, quarterFinals.numGames["MTL"][4], float64(quarterFinals.numGames["MTL"][4])/float64(max)*100)
	fmt.Printf("Win in 5 games: %v (%.2f%%)-------- Win in 5 games: %v (%.2f%%)\n", quarterFinals.numGames["PHI"][5], float64(quarterFinals.numGames["PHI"][5])/float64(max)*100, quarterFinals.numGames["MTL"][5], float64(quarterFinals.numGames["MTL"][5])/float64(max)*100)
	fmt.Printf("Win in 6 games: %v (%.2f%%)-------- Win in 6 games: %v (%.2f%%)\n", quarterFinals.numGames["PHI"][6], float64(quarterFinals.numGames["PHI"][6])/float64(max)*100, quarterFinals.numGames["MTL"][6], float64(quarterFinals.numGames["MTL"][6])/float64(max)*100)
	fmt.Printf("Win in 7 games: %v (%.2f%%)-------- Win in 7 games: %v (%.2f%%)\n", quarterFinals.numGames["PHI"][7], float64(quarterFinals.numGames["PHI"][7])/float64(max)*100, quarterFinals.numGames["MTL"][7], float64(quarterFinals.numGames["MTL"][7])/float64(max)*100)
	// fmt.Printf("No selection made: %v (%.2f%%)\n", max-numGames, float64(numGames)/float64(max)*100)
	fmt.Println("=============================================")
	fmt.Println("Tampa Bay Lightning - Columbus Blue Jackets")
	fmt.Println("-------------------------------------------")
	// numGames = quarterFinals.numGames["TBL"][4] + quarterFinals.numGames["TBL"][5] + quarterFinals.numGames["TBL"][6] + quarterFinals.numGames["TBL"][7] + quarterFinals.numGames["CBJ"][4] + quarterFinals.numGames["CBJ"][5] + quarterFinals.numGames["CBJ"][6] + quarterFinals.numGames["CBJ"][7]
	fmt.Printf("Win in 4 games: %v (%.2f%%)-------- Win in 4 games: %v (%.2f%%)\n", quarterFinals.numGames["TBL"][4], float64(quarterFinals.numGames["TBL"][4])/float64(max)*100, quarterFinals.numGames["CBJ"][4], float64(quarterFinals.numGames["CBJ"][4])/float64(max)*100)
	fmt.Printf("Win in 5 games: %v (%.2f%%)-------- Win in 5 games: %v (%.2f%%)\n", quarterFinals.numGames["TBL"][5], float64(quarterFinals.numGames["TBL"][5])/float64(max)*100, quarterFinals.numGames["CBJ"][5], float64(quarterFinals.numGames["CBJ"][5])/float64(max)*100)
	fmt.Printf("Win in 6 games: %v (%.2f%%)-------- Win in 6 games: %v (%.2f%%)\n", quarterFinals.numGames["TBL"][6], float64(quarterFinals.numGames["TBL"][6])/float64(max)*100, quarterFinals.numGames["CBJ"][6], float64(quarterFinals.numGames["CBJ"][6])/float64(max)*100)
	fmt.Printf("Win in 7 games: %v (%.2f%%)-------- Win in 7 games: %v (%.2f%%)\n", quarterFinals.numGames["TBL"][7], float64(quarterFinals.numGames["TBL"][7])/float64(max)*100, quarterFinals.numGames["CBJ"][7], float64(quarterFinals.numGames["CBJ"][7])/float64(max)*100)
	// fmt.Printf("No selection made: %v (%.2f%%)\n", max-numGames, float64(numGames)/float64(max)*100)
	fmt.Println("=============================================")
	fmt.Println("Washington Capitals ---- New York Islanders")
	fmt.Println("-------------------------------------------")
	// numGames = quarterFinals.numGames["WSH"][4] + quarterFinals.numGames["WSH"][5] + quarterFinals.numGames["WSH"][6] + quarterFinals.numGames["WSH"][7] + quarterFinals.numGames["NYI"][4] + quarterFinals.numGames["NYI"][5] + quarterFinals.numGames["NYI"][6] + quarterFinals.numGames["NYI"][7]
	fmt.Printf("Win in 4 games: %v (%.2f%%)-------- Win in 4 games: %v (%.2f%%)\n", quarterFinals.numGames["WSH"][4], float64(quarterFinals.numGames["WSH"][4])/float64(max)*100, quarterFinals.numGames["NYI"][4], float64(quarterFinals.numGames["NYI"][4])/float64(max)*100)
	fmt.Printf("Win in 5 games: %v (%.2f%%)-------- Win in 5 games: %v (%.2f%%)\n", quarterFinals.numGames["WSH"][5], float64(quarterFinals.numGames["WSH"][5])/float64(max)*100, quarterFinals.numGames["NYI"][5], float64(quarterFinals.numGames["NYI"][5])/float64(max)*100)
	fmt.Printf("Win in 6 games: %v (%.2f%%)-------- Win in 6 games: %v (%.2f%%)\n", quarterFinals.numGames["WSH"][6], float64(quarterFinals.numGames["WSH"][6])/float64(max)*100, quarterFinals.numGames["NYI"][6], float64(quarterFinals.numGames["NYI"][6])/float64(max)*100)
	fmt.Printf("Win in 7 games: %v (%.2f%%)-------- Win in 7 games: %v (%.2f%%)\n", quarterFinals.numGames["WSH"][7], float64(quarterFinals.numGames["WSH"][7])/float64(max)*100, quarterFinals.numGames["NYI"][7], float64(quarterFinals.numGames["NYI"][7])/float64(max)*100)
	// fmt.Printf("No selection made: %v (%.2f%%)\n", max-numGames, float64(numGames)/float64(max)*100)
	fmt.Println("=============================================")
	fmt.Println("Boston Bruins ---------- Carolina Hurricanes")
	fmt.Println("-------------------------------------------")
	// numGames = quarterFinals.numGames["BOS"][4] + quarterFinals.numGames["BOS"][5] + quarterFinals.numGames["BOS"][6] + quarterFinals.numGames["BOS"][7] + quarterFinals.numGames["CAR"][4] + quarterFinals.numGames["CAR"][5] + quarterFinals.numGames["CAR"][6] + quarterFinals.numGames["CAR"][7]
	fmt.Printf("Win in 4 games: %v (%.2f%%)-------- Win in 4 games: %v (%.2f%%)\n", quarterFinals.numGames["BOS"][4], float64(quarterFinals.numGames["BOS"][4])/float64(max)*100, quarterFinals.numGames["CAR"][4], float64(quarterFinals.numGames["CAR"][4])/float64(max)*100)
	fmt.Printf("Win in 5 games: %v (%.2f%%)-------- Win in 5 games: %v (%.2f%%)\n", quarterFinals.numGames["BOS"][5], float64(quarterFinals.numGames["BOS"][5])/float64(max)*100, quarterFinals.numGames["CAR"][5], float64(quarterFinals.numGames["CAR"][5])/float64(max)*100)
	fmt.Printf("Win in 6 games: %v (%.2f%%)-------- Win in 6 games: %v (%.2f%%)\n", quarterFinals.numGames["BOS"][6], float64(quarterFinals.numGames["BOS"][6])/float64(max)*100, quarterFinals.numGames["CAR"][6], float64(quarterFinals.numGames["CAR"][6])/float64(max)*100)
	fmt.Printf("Win in 7 games: %v (%.2f%%)-------- Win in 7 games: %v (%.2f%%)\n", quarterFinals.numGames["BOS"][7], float64(quarterFinals.numGames["BOS"][7])/float64(max)*100, quarterFinals.numGames["CAR"][7], float64(quarterFinals.numGames["CAR"][7])/float64(max)*100)
	// fmt.Printf("No selection made: %v (%.2f%%)\n", max-numGames, float64(numGames)/float64(max)*100)
	fmt.Println("=============================================")

	fmt.Println("--- STANLEY CUP FINALS ---\n\n")
	sortedFinalsTeams := SortMapOfInts(finals.toWin)
	for _, team := range sortedFinalsTeams {
		if team == "" {
			fmt.Printf("DID NOT CHOOSE A WINNER LMAO"+": %v  (%.2f%%)\n", finals.toWin[team], float64(finals.toWin[team])/float64(max)*100)
		} else {
			fmt.Printf(team+": %v  (%.2f%%)\n", finals.toWin[team], float64(finals.toWin[team])/float64(max)*100)
		}
	}
}

func getTeamAbbrv(teamNum string) string {
	teams := map[string]string{"2": "ARI", "3": "BOS", "5": "CGY", "6": "CAR", "7": "CHI", "8": "COL", "9": "CBJ", "10": "DAL", "16": "MTL", "19": "NYI", "22": "PHI", "25": "STL", "26": "TBL", "28": "VAN", "29": "VGK", "30": "WSH"}
	return teams[teamNum]
}

func SortMapOfInts(i map[string]int) (ph []string) {
	ph = []string{}
	type kv struct {
		Key   string
		Value int
	}
	var entitiesAndFreqs []kv
	for k, v := range i {
		entitiesAndFreqs = append(entitiesAndFreqs, kv{k, v})
	}

	sort.Slice(entitiesAndFreqs, func(i, j int) bool {
		return entitiesAndFreqs[i].Value > entitiesAndFreqs[j].Value
	})

	for _, kv := range entitiesAndFreqs {
		ph = append(ph, kv.Key)
	}
	return ph
}
