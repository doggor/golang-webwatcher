package jobs

import (
	"fmt"
	"regexp"
	"time"
	"webwatcher/internal/utils"
)

// for extract the bookable date of the HTML content
var htmlBookableDateMatcher = regexp.MustCompile(`Bookable Days of Operation</h2>\s*<div class="text_strong">\s*(\d+\/\d+)\s*～\s*(\d+\/\d+)\s`)
var dateMatcher = regexp.MustCompile(`(\d+)\/(\d+)`)

// CheckNaebaBusBookingDates define the job of the NaebaBusBookingScanner
func CheckNaebaBusBookingDates() {
	webPageURL := "https://japanbusonline.com/en/CourseSearch/11100290001"

	//get the web page html
	content, err := utils.FetchPage(webPageURL)

	if err != nil {
		fmt.Println(err)
		return
	}

	//find the available booking date
	matched := htmlBookableDateMatcher.FindStringSubmatch(content)

	if matched == nil || len(matched) != 3 {
		fmt.Println("Booking days not found.")
		return
	}

	//extract the end of the available booking date
	endDate, err := toWinterDate(matched[2])

	if err != nil {
		fmt.Println("Cannot parse booking date string.")
		return
	}

	//the date of the booking I want to make
	targetDate := time.Date(2020, time.January, 7, 0, 0, 0, 0, time.UTC)

	//check if the date is open for booking, and notify me
	if endDate.After(targetDate) {
		utils.SendToTelegram("It's time to book the bus: " + webPageURL)
	} else {
		utils.SendToTelegram("Be patient. The bookable date is only from " + matched[1] + " to " + matched[2])
	}
}

// toDate parse date string from format `12/31` to time.Time
// in which it maintances the winter session where Sep < Aug
func toWinterDate(dateString string) (time.Time, error) {
	date, err := time.Parse(`01/02`, dateString)

	if err != nil {
		return time.Now(), err
	}

	if date.Month() >= 9 {
		return date.AddDate(2019, 0, 0), nil
	}

	return date.AddDate(2020, 0, 0), nil
}
