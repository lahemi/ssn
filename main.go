package main

import (
	"flag"
	"fmt"
    "os"
	"time"

	"github.com/lahemi/ssn/hetu"
)

var (
	repeatGen          int
	validateInput      string
	yearStart, yearEnd int
)

func die(str ...interface{}) {
    fmt.Fprintln(os.Stderr, str...)
    os.Exit(1)
}

func argRange(arg int) bool {
	if arg >= 1800 && arg <= 2099 {
		return true
	}
	return false
}

func init() {
	flag.StringVar(&validateInput, "validate", "", "SSN to validate")
	flag.IntVar(&yearStart, "start", 1800, "Start year for SSN creation")
	flag.IntVar(&yearEnd, "end", int(time.Now().Year()), "End year for SSN creation")
	flag.IntVar(&repeatGen, "repeat", 1, "How many SSNs to generate")
	flag.Parse()
}

func main() {
	switch validateInput {
	case "":
		if repeatGen > 0 {
			if !argRange(yearStart) || !argRange(yearEnd) || yearEnd == 1800 {
				die("Year not in range!")
			}
			for i := 0; i < repeatGen; i++ {
				fmt.Println(hetu.Create(yearStart, yearEnd))
			}
		}
	default:
		switch hetu.Validate(validateInput) {
		case true:
			fmt.Printf("%s is valid!\n", validateInput)
		default:
			fmt.Printf("%s is not valid!\n", validateInput)
		}
	}
}

// table tests!
// validate multiple
