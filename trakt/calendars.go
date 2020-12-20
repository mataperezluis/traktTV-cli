// Package trakt ...
package trakt

import (
	//"net/url"
	"time"
)

var (

	calendarsGETShowsURL = Hyperlink("calendars/my/shows/{start_date}/{days}")
	calendarsGETNewShows= Hyperlink("calendars/my/shows/new/{start_date}/{days}")
	calendarsGETSeasonsPremiere = Hyperlink("calendars/my/shows/premieres/{start_date}/{days}")
	calendarsGETMovies = Hyperlink("calendars/my/movies/{start_date}/{days}")
	calendarsGETDVD = Hyperlink("calendars/my/dvd/{start_date}/{days}")
	calendarsGETAllShows = Hyperlink("calendars/all/shows/start_date/days")
	calendarsGETAllNewShows = Hyperlink("calendars/all/shows/new/start_date/days")
	calendarsGETSeasonPremire = Hyperlink("calendars/all/shows/premieres/start_date/days")
	calendarsGETAllMovies = Hyperlink("calendars/all/movies/start_date/days")
	calendarsGETAllDVD = Hyperlink("calendars/all/dvd/start_date/days")
)

//Calendars with the base url.URL
func (c *Client) Calendars() (Calendars *CalendarsService) {
	Calendars = &CalendarsService{client: c}
	return
}

//CalendarsService ...
type CalendarsService struct {
	client *Client
}


//MyShows ...
func (r *CalendarsService) MyShows(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := calendarsGETShowsURL.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}

//MyNewShows ...
func (r *CalendarsService) MyNewShows(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := calendarsGETNewShows.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}

//MySeasonPremiere ...
func (r *CalendarsService) MySeasonPremiere(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := calendarsGETSeasonsPremiere.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}

//MyMovies ...
func (r *CalendarsService) MyMovies(startdate string, days string) (calendars *CalendarMovies, result *Result) {
	url, _ := calendarsGETMovies.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendars)
	return
}

//MyDVD ...
func (r *CalendarsService) MyDVD(startdate string, days string) (calendars *CalendarMovies, result *Result) {
	url, _ := calendarsGETDVD.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendars)
	return
}

//AllShows ...
func (r *CalendarsService) AllShows(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := calendarsGETAllShows.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}

//AllNewShows ...
func (r *CalendarsService) AllNewShows(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := calendarsGETAllNewShows.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}

//AllPremieres ...
func (r *CalendarsService) AllPremieres(startdate string, days string) (calendar *CalendarShow, result *Result) {
	url, _ := calendarsGETSeasonPremire.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendar)
	return
}


//AllMovies ...
func (r *CalendarsService) AllMovies(startdate string, days string) (calendars *CalendarMovies, result *Result) {
	url, _ := calendarsGETAllMovies.Expand(M{"start_date":startdate,"days":days})

	result = r.client.get(url, &calendars)
	return
}

//AllDVD ...
func (r *CalendarsService) AllDVD(startdate string, days string) (calendars *CalendarMovies, result *Result) {
	url, _ := calendarsGETAllDVD.Expand(M{"start_date":startdate,"days":days})
	result = r.client.get(url, &calendars)
	return
}

//CalendarShow ...
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
//CalendarMovies ...
type CalendarMovies []struct {
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
