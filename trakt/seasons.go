package trakt

import (
	//"fmt"
	"net/url"
	"time"
)

var (
	ShowSeasonsURL       = Hyperlink("shows/{showTraktID}/seasons")
	ShowSeasonsExtendedURL       = Hyperlink("shows/{showTraktID}/seasons?extended={extraInfo}")
	ShowSeasonsNumberURL = Hyperlink("shows/{showTraktID}/seasons/{seasonNumber}")
	ShowSeasonsCommentsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/comments/{sort}")
	ShowSeasonsListsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/lists/{tipo}/{sort}")
	ShowSeasonsPeopleURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/people")	
	ShowSeasonsRatingsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/ratings")	
	ShowSeasonsStatsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/stats")
	ShowSeasonsWatchingURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/watching")
)

// Create a ShowsService with the base url.URL
func (c *Client) Seasons() (seasons *SeasonsService) {
	seasons = &SeasonsService{client: c}
	return
}

type SeasonsService struct {
	client *Client
}

// All returns all the seasons of a particular Show. The seasons do not include
// the episodes.
func (r *SeasonsService) All(showTraktID string,extraInfo string) (seasons *Season, result *Result) {
	var url *url.URL
	if extraInfo == ""{
		url, _ = ShowSeasonsURL.Expand(M{"showTraktID": showTraktID})
	}else {
		url, _ = ShowSeasonsExtendedURL.Expand(M{"showTraktID": showTraktID,"extraInfo":extraInfo})
	}
	result = r.client.get(url, &seasons)
	return
}

// ByNumber returns a specific season of a particular Show.
func (r *SeasonsService) ByNumber(showTraktID string, seasonNumber string) (season []Episodes, result *Result) {
	url, _ := ShowSeasonsNumberURL.Expand(M{
		"showTraktID":  showTraktID,
		"seasonNumber": seasonNumber,
	})
	result = r.client.get(url, &season)
	return
}

func (r *SeasonsService) SeasonComments(traktID string, seasonN string,sort string) (season *SeasonComment, result *Result) {
	url, _ := ShowSeasonsCommentsURL.Expand(M{"traktID": traktID,"seasonN": seasonN,"sort": sort})
	result = r.client.get(url, &season)
	return
}

func (r *SeasonsService) SeasonList(traktID string,seasonN string, tipo string,sort string) (season *SeasonList, result *Result) {
	url, _ := ShowListsURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"tipo":tipo,"sort": sort})
	result = r.client.get(url, &season)
	return
}

func (r *SeasonsService) SeasonPeople(traktID string,seasonN string) (show *ShowCast, result *Result) {
	url, _ := ShowSeasonsPeopleURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *SeasonsService) SeasonRatings(traktID string,seasonN string) (show *ShowRatings, result *Result) {
	url, _ := ShowSeasonsRatingsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *SeasonsService) SeasonStats(traktID string,seasonN string) (show *ShowStats, result *Result) {
	url, _ := ShowSeasonsStatsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *SeasonsService) SeasonWatching(traktID string,seasonN string) (show []User, result *Result) {
	url, _ := ShowSeasonsWatchingURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

// Season struct for the Trakt v2 API
type Season []struct {
	Number   int        `json:"number"`
	Ids      Ids        `json:"ids"`
	Episodes []Episodes `json:"episodes"`
}
type IdsSeason struct {
	Trakt int `json:"trakt"`
	Tvdb  int `json:"tvdb"`
	Tmdb  int `json:"tmdb"`
}
type IdsEpisodes struct {
	Trakt int    `json:"trakt"`
	Tvdb  int    `json:"tvdb"`
	Imdb  string `json:"imdb"`
	Tmdb  int    `json:"tmdb"`
}
type Episodes struct {
	Season int    `json:"season"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Ids    IdsEpisodes    `json:"ids"`
}

type SeasonComment []struct {
	ID        int       `json:"id"`
	ParentID  int       `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comment   string    `json:"comment"`
	Spoiler   bool      `json:"spoiler"`
	Review    bool      `json:"review"`
	Replies   int       `json:"replies"`
	Likes     int       `json:"likes"`
	UserStats struct {
		Rating         int `json:"rating"`
		PlayCount      int `json:"play_count"`
		CompletedCount int `json:"completed_count"`
	} `json:"user_stats"`
	User struct {
		Username string `json:"username"`
		Private  bool   `json:"private"`
		Name     string `json:"name"`
		Vip      bool   `json:"vip"`
		VipEp    bool   `json:"vip_ep"`
		Ids      struct {
			Slug string `json:"slug"`
		} `json:"ids"`
	} `json:"user"`
}



type SeasonList []struct {
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Privacy        string    `json:"privacy"`
	DisplayNumbers bool      `json:"display_numbers"`
	AllowComments  bool      `json:"allow_comments"`
	SortBy         string    `json:"sort_by"`
	SortHow        string    `json:"sort_how"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	ItemCount      int       `json:"item_count"`
	CommentCount   int       `json:"comment_count"`
	Likes          int       `json:"likes"`
	Ids            struct {
		Trakt int    `json:"trakt"`
		Slug  string `json:"slug"`
	} `json:"ids"`
	User struct {
		Username string `json:"username"`
		Private  bool   `json:"private"`
		Name     string `json:"name"`
		Vip      bool   `json:"vip"`
		VipEp    bool   `json:"vip_ep"`
		Ids      struct {
			Slug string `json:"slug"`
		} `json:"ids"`
	} `json:"user"`
}