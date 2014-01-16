package wonging

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

// -1 means cannot find
func checkDealerContain(d *Dealer, dealers []*Dealer) int {
	index := -1
	for i := 0; i < len(dealers); i++ {
		if d.id == dealers[i].id {
			index = i
		}
	}
	return index
}

func checkPlayerContain(p *Player, players []*Player) int {
	index := -1
	for i := 0; i < len(players); i++ {
		if p.id == players[i].id {
			index = i
		}
	}
	return index
}

func checkShoeContain(c *Card, shoe *Deck) int {
	index := -1
	for i := 0; i < len(shoe.cards); i++ {
		if c.value == shoe.cards[i].value && c.suit == shoe.cards[i].suit {
			index = i
		}
	}
	return index
}

func combineCounters(counters []*Counter) *Counter {
	combinedCounter := new(Counter).initialize()
	for _, counter := range counters {
		combinedCounter.HiLo += counter.HiLo
		combinedCounter.HiOpt1 += counter.HiOpt1
		combinedCounter.HiOpt2 += counter.HiOpt2
		combinedCounter.KO += counter.KO
		combinedCounter.Omega2 += counter.Omega2
		combinedCounter.Red7 += counter.Red7
		combinedCounter.ZenCount += counter.ZenCount
	}
	return combinedCounter
}

func checkCardsValueEqual(c1 *Card, c2 *Card) bool {
	return c1.value == c2.value
}

func MinFloatS(v []float64) int {
	var index int
	var m float64
	if len(v) > 0 {
		m = v[0]

	}
	for i := 1; i < len(v); i++ {
		if v[i] < m {
			m = v[i]
			index = i
		}
	}
	return index
}

func MaxFloatS(v []float64) int {
	var index int
	var m float64
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
			index = i
		}
	}
	return index
}

func randInt(min, max int) int {
	if min == max {
		return min
	} else {
		return min + rand.Intn(max-min)
	}
}

// Colorify things

func BlackText(s string) string {
	return "\x1b[30;1m" + s + "\x1b[0m"
}

func RedText(s string) string {
	return "\x1b[31;1m" + s + "\x1b[0m"
}

func GreenText(s string) string {
	return "\x1b[32;1m" + s + "\x1b[0m"
}

func YellowText(s string) string {
	return "\x1b[33;1m" + s + "\x1b[0m"
}

func BlueText(s string) string {
	return "\x1b[34;1m" + s + "\x1b[0m"
}

func MagentaText(s string) string {
	return "\x1b[35;1m" + s + "\x1b[0m"
}

func CyanText(s string) string {
	return "\x1b[36;1m" + s + "\x1b[0m"
}

func WhiteText(s string) string {
	return "\x1b[37;1m" + s + "\x1b[0m"
}

// Logging overriding

func logAbsolutePath() string {
	dir := "log/"
	os.Remove(dir)

	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			//directory doesn't exist, create it
			fmt.Printf("Directory %s doesn't exist, creating it\n", dir)
			os.Mkdir(dir, 0777)
		}
	}

	absPath, _ := filepath.Abs(dir)
	return absPath + "/"
}

func logToWithLevel(fileName, level, logStr string) {
	fileName = logAbsolutePath() + fileName + ".log"
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf("[%s] %s", level, logStr)
}

func logTo(fileName, logStr string) {

	fileName = logAbsolutePath() + fileName + ".log"
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(logStr)
}

// HTML generation
func htmlAbsolutePath() string {
	dir := "html/"
	os.Remove(dir)

	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			//directory doesn't exist, create it
			fmt.Printf("Directory %s doesn't exist, creating it\n", dir)
			os.Mkdir(dir, 0777)
		}
	}
	absPath, _ := filepath.Abs(dir)
	return absPath + "/"
}
func generateHTMLMap(fileName, htmlString string) {
	fileName = htmlAbsolutePath() + fileName + ".html"
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}
	defer f.Close()
	io.WriteString(f, htmlString)
}
