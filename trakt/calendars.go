package trakt

import (
	//"net/url"
	"time"
)

var (

	CalendarsGETShowsUrl = Hyperlink("calendars/my/shows/{start_date}/{days}")
	CalendarsGETNewShows= Hyperlink("calendars/my/shows/new/{start_date}/{days}")
	CalendarsGETSeasonsPremiere = Hyperlink("calendars/my/shows/premieres/{start_date}/{days}")
	CalendarsGETMovies = Hyperlink("calendars/my/movies/{start_date}/{days}")
	CalendarsGETDVD = Hyperlink("calendars/my/dvd/{start_date}/{days}")
	CalendarsGETAllShows = Hyperlink("calendars/all/shows/start_date/days")
	CalendarsGETAllNewShows = Hyperlink("calendars/all/shows/new/start_date/days")
	CalendarsGETSeasonPremire = Hyperlink("calendars/all/shows/premieres/start_date/days")
	CalendarsGETAllMovies = Hyperlink("calendars/all/movies/start_date/days")
	CalendarsGETAllDVD = Hyperlink("calendars/all/dvd/start_date/days")
)

// Create a CalendarsService with the base url.URL
func (c *Client) Calendars() (Calendars *CalendarsService) {
	Calendars = &CalendarsService{client: c}
	return
}

type CalendarsService struct {
	client *Client
}

// One returns a single Movie identified by a Trakt ID. It also returns a Result
// object to inspect the returned response of the server.
func (r *CalendarsService) MyShows(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := CalendarsGETShowsUrl.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}

func (r *CalendarsService) MyNewShows(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := CalendarsGETNewShows.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}

func (r *CalendarsService) MySeasonPremiere(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := CalendarsGETSeasonsPremiere.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}

func (r *CalendarsService) MyMovies(startdate string, days string) (calendars *calendarMovies, result *Result) {
	url, _ := CalendarsGETMovies.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendars)
	return
}

func (r *CalendarsService) MyDVD(startdate string, days string) (calendars *calendarMovies, result *Result) {
	url, _ := CalendarsGETDVD.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendars)
	return
}

func (r *CalendarsService) AllShows(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := CalendarsGETAllShows.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}

func (r *CalendarsService) AllNewShows(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := CalendarsGETAllNewShows.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}

func (r *CalendarsService) AllPremieres(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := CalendarsGETSeasonPremire.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}



func (r *CalendarsService) AllMovies(startdate string, days string) (calendars *calendarMovies, result *Result) {
	url, _ := CalendarsGETAllMovies.Expand(M{"start_date":startdate,"days":days})

	result = r.client.get(url, &calendars)
	return
}

func (r *CalendarsService) AllDVD(startdate string, days string) (calendars *calendarMovies, result *Result) {
	url, _ := CalendarsGETAllDVD.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendars)
	return
}


type CalendarShow []struct {
	FirstAired time.Time `json:"first_aired"`
	Episode    struct {
		Season int    `json:"season"`
		Number int    `json:"number"`
		Title  string `json:"title"`
		Ids    struct {
			Trakt int    `json:"trakt"`
			Tvdb  int    `json:"tvdb"`
			Imdb  string `json:"imdb"`
			Tmdb  int    `json:"tmdb"`
		} `json:"ids"`
	} `json:"episode"`
	Show struct {
		Title string `json:"title"`
		Year  int    `json:"year"`
		Ids   struct {
			Trakt int    `json:"trakt"`
			Slug  string `json:"slug"`
			Tvdb  int    `json:"tvdb"`
			Imdb  string `json:"imdb"`
			Tmdb  int    `json:"tmdb"`
		} `json:"ids"`
	} `json:"show"`
}

type calendarMovies []struct {
	Released string `json:"released"`
	Movie    struct {
		Title string `json:"title"`
		Year  int    `json:"year"`
		Ids   struct {
			Trakt int    `json:"trakt"`
			Slug  string `json:"slug"`
			Imdb  string `json:"imdb"`
			Tmdb  int    `json:"tmdb"`
		} `json:"ids"`
	} `json:"movie"`
}
