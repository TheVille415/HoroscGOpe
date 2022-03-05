package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

// get the month of birth
func getMonthOfBirth(month string) int {
	switch month {
	case "Aries":
		return 1
	case "Taurus":
		return 2
	case "Gemini":
		return 3
	case "Cancer":
		return 4
	case "Leo":
		return 5
	case "Virgo":
		return 6
	case "Libra":
		return 7
	case "Scorpio":
		return 8
	case "Sagittarius":
		return 9
	case "Capricorn":
		return 10
	case "Aquarius":
		return 11
	case "Pisces":
		return 12
	}
	return 0
}

type Store struct {
	Horoscope string
}

func scrapeHoro(sign int) string {
	// Get the data from https://www.horoscope.com/us/index.aspx
	// Instantiate default collector
	c := colly.NewCollector()
	character := make([]Store, 1)

	// On every a element which has href attribute call callback
	c.OnHTML("main-horoscope", func(e *colly.HTMLElement) {
		e.ForEach("p", func(_ int, e *colly.HTMLElement) {
			newInfo := Store{}
			newInfo.Horoscope = e.ChildText("p")
			fmt.Println(e.ChildText("p"))
			character = append(character, newInfo)
		})
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Received error:", e)
	})

	// Start scraping on https://lucifer.fandom.com/wiki/Episodes
	c.Visit("https://www.horoscope.com/us/horoscopes/general/horoscope-general-daily-today.aspx?sign=" + fmt.Sprintf("%d", sign))

	return character[0].Horoscope
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your Zodiac sign?: ")
	sign, _ := reader.ReadString('\n')
	zodiacSign := getMonthOfBirth(sign)
	todayHoroscope := scrapeHoro(zodiacSign)
	fmt.Println("Your horoscope for today is: ", todayHoroscope)
}
