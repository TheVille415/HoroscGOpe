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
	case "Aries\n":
		return 1
	case "Taurus\n":
		return 2
	case "Gemini\n":
		return 3
	case "Cancer\n":
		return 4
	case "Leo\n":
		return 5
	case "Virgo\n":
		return 6
	case "Libra\n":
		return 7
	case "Scorpio\n":
		return 8
	case "Sagittarius\n":
		return 9
	case "Capricorn\n":
		return 10
	case "Aquarius\n":
		return 11
	case "Pisces\n":
		return 12
	}
	return 0
}

type Store struct {
	Horoscope string
}

// get the horoscope for the day
func scrapeHoro(sign int) string {
	// Get the data from https://www.horoscope.com/us/index.aspx
	// Instantiate default collector
	c := colly.NewCollector()
	character := make([]Store, 0)

	// On every a element which has href attribute call callback
	c.OnHTML("div.main-horoscope", func(e *colly.HTMLElement) {
		newInfo := Store{}
		newInfo.Horoscope = e.ChildText("p:nth-child(2)")
		character = append(character, newInfo)
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
	c.OnScraped(func(r *colly.Response) {
		fmt.Print(character[0].Horoscope)
	})
	return character[0].Horoscope
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your Zodiac sign?: ")
	sign, _ := reader.ReadString('\n')
	zodiacSign := getMonthOfBirth(sign)
	if zodiacSign == 0 {
		fmt.Println("Invalid sign")
		return
	}
	todayHoroscope := scrapeHoro(zodiacSign)
	fmt.Print("\n")
	fmt.Println("Your horoscope for today is:", todayHoroscope)
}
