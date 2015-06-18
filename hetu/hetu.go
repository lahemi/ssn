package hetu

import (
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

var (
	checkKeys = []string{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A",
		"B", "C", "D", "E", "F", "H", "J", "K", "L", "M", "N",
		"P", "R", "S", "T", "U", "V", "W", "X", "Y",
	}
	centuries = map[string]string{
		"18": "+",
		"19": "-",
		"20": "A",
	}
)

func die(str ...interface{}) {
	fmt.Fprintln(os.Stderr, str...)
	os.Exit(1)
}

func randomRange(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

// This is equivalent to time.daysIn(m, year).
func daysIn(m, year int) int {
	return time.Date(year, time.Month(m)+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func mod31(num string) int {
	i, err := strconv.Atoi(num)
	if err != nil {
		die(err)
	}
	return i % 31
}

func Create(start, end int) string {
	year := randomRange(start, end)
	month := randomRange(1, 12)
	day := randomRange(1, daysIn(month, year))
	century := strconv.Itoa(year)[0:2]
    decade := strconv.Itoa(year)[2:4]

	centurySep := centuries[century]

	orderNum := randomRange(2, 889)

	checkNum := fmt.Sprintf("%02d%02d%s%03d", day, month, decade, orderNum)

	checkNumIndex := mod31(checkNum)

	key := checkKeys[checkNumIndex]

	return fmt.Sprintf("%02d%02d%s%s%03d%s", day, month, decade, centurySep, orderNum, key)
}

func Validate(hetu string) bool {
	if m, _ := regexp.MatchString(`^\d{6}[+-A]\d{3}[a-zA-Z0-9]$`, hetu); m {
		snum := hetu[0:6] + hetu[7:10]
		return checkKeys[mod31(snum)] == string(hetu[len(hetu)-1])
	}
	return false
}
