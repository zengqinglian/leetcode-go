package alertUsingSameKeyCardThreeorMoreTimesinaOneHourPeriod

import (
	"sort"
	"strconv"
	"strings"
)

/*
Runtime: 163 ms, faster than 90.00% of Go online submissions for Alert Using Same Key-Card Three or More Times in a One Hour Period.
Memory Usage: 16 MB, less than 95.00% of Go online submissions for Alert Using Same Key-Card Three or More Times in a One Hour Period.
*/
func alertNames(keyName []string, keyTime []string) []string {
	keyCardMap := make(map[string][]int)
	for i, v := range keyName {
		_, ok := keyCardMap[v]
		if !ok {
			timeSlice := make([]int, 0)
			keyCardMap[v] = timeSlice
		}
		keyCardMap[v] = append(keyCardMap[v], getTime(keyTime[i]))
	}
	res := make([]string, 0, len(keyName))
	for k, v := range keyCardMap {
		sort.Ints(v)
		if len(v) >= 3 {
			for i := 2; i < len(v); i++ {
				if v[i]-v[i-2] <= 60 {
					res = append(res, k)
					break
				}
			}

		}
	}
	sort.Strings(res)
	return res
}

func getTime(s string) int {
	timeArray := strings.Split(s, ":")
	h, _ := strconv.Atoi(timeArray[0])
	m, _ := strconv.Atoi(timeArray[1])
	//mts := split[0]+"h"+split[1]+"m"
	//minutes, _ := time.ParseDuration(mts)
	return h*60 + m
}
