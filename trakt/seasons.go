//Package trakt ...
package trakt

import (
	//"fmt"
	"net/url"
	"time"
)

var (
	showSeasonsURL       = Hyperlink("shows/{showTraktID}/seasons")
	showSeasonsExtendedURL       = Hyperlink("shows/{showTraktID}/seasons?extended={extraInfo}")
	showSeasonsNumberURL = Hyperlink("shows/{showTraktID}/seasons/{seasonNumber}")
	showSeasonsCommentsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/comments/{sort}")
	showSeasonsListsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/lists/{tipo}/{sort}")
	showSeasonsPeopleURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/people")	
	showSeasonsPeopleExtendedURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/people?extended=guest_stars")
	showSeasonsRatingsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/ratings")	
	showSeasonsStatsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/stats")
	showSeasonsWatchingURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/watching")
)

//Seasons Create a ShowsService with the base url.URL
func (c *Client) Seasons() (seasons *SeasonsService) {
	seasons = &SeasonsService{client: c}
	return
}
//SeasonsService ...
type SeasonsService struct {
	client *Client
}

// All returns all the seasons of a particular Show. The seasons do not include
// the episodes.
func (r *SeasonsService) All(showTraktID string,extraInfo string) (seasons *Season, result *Result) {
	var url *url.URL
	if extraInfo == ""{
		url, _ = showSeasonsURL.Expand(M{"showTraktID": showTraktID})
	}else {
		url, _ = showSeasonsExtendedURL.Expand(M{"showTraktID": showTraktID,"extraInfo":extraInfo})
	}
	result = r.client.get(url, &seasons)
	return
}

// ByNumber returns a specific season of a particular Show.
func (r *SeasonsService) ByNumber(showTraktID string, seasonNumber string) (season []Episodes, result *Result) {
	url, _ := showSeasonsNumberURL.Expand(M{
		"showTraktID":  showTraktID,
		"seasonNumber": seasonNumber,
	})
	result = r.client.get(url, &season)
	return
}
//SeasonComments ...
func (r *SeasonsService) SeasonComments(traktID string, seasonN string,sort string) (season *SeasonComment, result *Result) {
	url, _ := showSeasonsCommentsURL.Expand(M{"traktID": traktID,"seasonN": seasonN,"sort": sort})
	result = r.client.get(url, &season)
	return
}
//SeasonList ...
func (r *SeasonsService) SeasonList(traktID string,seasonN string, tipo string,sort string) (season *SeasonList, result *Result) {
	url, _ := showSeasonsListsURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"tipo":tipo,"sort": sort})
	result = r.client.get(url, &season)
	return
}
//SeasonPeople ...
func (r *SeasonsService) SeasonPeople(traktID string,seasonN string,extended string) (season *SeasonCast, result *Result) {
	var url *url.URL
	if extended=="false"{
		url, _ = showSeasonsPeopleURL.Expand(M{"traktID": traktID,"seasonN":seasonN})
	}else{
		url, _ = showSeasonsPeopleExtendedURL.Expand(M{"traktID": traktID,"seasonN":seasonN})
	}

	result = r.client.get(url, &season)
	return
}
//SeasonRatings ...
func (r *SeasonsService) SeasonRatings(traktID string,seasonN string) (season *SeasonRating, result *Result) {
	url, _ := showSeasonsRatingsURL.Expand(M{"traktID": traktID,"seasonN":seasonN})
	result = r.client.get(url, &season)
	return
}
//SeasonStats ...
func (r *SeasonsService) SeasonStats(traktID string,seasonN string) (season *SeasonStats, result *Result) {
	url, _ := showSeasonsStatsURL.Expand(M{"traktID": traktID,"seasonN":seasonN})
	result = r.client.get(url, &season)
	return
}
//SeasonWatching ...
func (r *SeasonsService) SeasonWatching(traktID string,seasonN string) (season *SeasonUser, result *Result) {
	url, _ := showSeasonsWatchingURL.Expand(M{"traktID": traktID,"seasonN":seasonN})
	result = r.client.get(url, &season)
	return
}

// Season struct for the Trakt v2 API
type Season []struct {
	Number   int        `json:"number"`
	Ids      Ids        `json:"ids"`
	Episodes []Episodes `json:"episodes"`
}
//IdsSeason ...
type IdsSeason struct {
	Trakt int `json:"trakt"`
	Tvdb  int `json:"tvdb"`
	Tmdb  int `json:"tmdb"`
}
//IdsEpisodes ...
type IdsEpisodes struct {
	Trakt int    `json:"trakt"`
	Tvdb  int    `json:"tvdb"`
	Imdb  string `json:"imdb"`
	Tmdb  int    `json:"tmdb"`
}
//Episodes ...
type Episodes struct {
	Season int    `json:"season"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Ids    IdsEpisodes    `json:"ids"`
}
//SeasonComment ...
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
//SeasonUser ...
type SeasonUser []struct {
	Username string `json:"username"`
	Private  bool   `json:"private"`
	Name     string `json:"name"`
	Vip      bool   `json:"vip"`
	VipEp    bool   `json:"vip_ep"`
	Ids      struct {
		Slug string `json:"slug"`
	} `json:"ids"`
}


//SeasonList ...
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
//SeasonStats ...
type SeasonStats struct {
	Watchers          int `json:"watchers"`
	Plays             int `json:"plays"`
	Collectors        int `json:"collectors"`
	CollectedEpisodes int `json:"collected_episodes"`
	Comments          int `json:"comments"`
	Lists             int `json:"lists"`
	Votes             int `json:"votes"`
}
//SeasonRating ...	
type SeasonRating struct {
	Rating       float64 `json:"rating"`
	Votes        int `json:"votes"`
	Distribution struct {
		Num1  int `json:"1"`
		Num2  int `json:"2"`
		Num3  int `json:"3"`
		Num4  int `json:"4"`
		Num5  int `json:"5"`
		Num6  int `json:"6"`
		Num7  int `json:"7"`
		Num8  int `json:"8"`
		Num9  int `json:"9"`
		Num10 int `json:"10"`
	} `json:"distribution"`
}

//SeasonCast ...
type SeasonCast struct {
	Cast []struct {
		Characters   []string `json:"characters"`
		EpisodeCount int      `json:"episode_count"`
		Person       struct {
			Name string `json:"name"`
			Ids  struct {
				Trakt  int         `json:"trakt"`
				Slug   string      `json:"slug"`
				Imdb   string      `json:"imdb"`
				Tmdb   int         `json:"tmdb"`
				Tvrage interface{} `json:"tvrage"`
			} `json:"ids"`
		} `json:"person"`
	} `json:"cast"`
	GuestStars []struct {
		Characters   []string `json:"characters"`
		EpisodeCount int      `json:"episode_count"`
		Person       struct {
			Name string `json:"name"`
			Ids  struct {
				Trakt  int    `json:"trakt"`
				Slug   string `json:"slug"`
				Imdb   string `json:"imdb"`
				Tmdb   int    `json:"tmdb"`
				Tvrage int    `json:"tvrage"`
			} `json:"ids"`
		} `json:"person"`
	} `json:"guest_stars"`
	Crew struct {
		Production []struct {
			Jobs         []string `json:"jobs"`
			EpisodeCount int      `json:"episode_count"`
			Person       struct {
				Name string `json:"name"`
				Ids  struct {
					Trakt  int         `json:"trakt"`
					Slug   string      `json:"slug"`
					Imdb   string      `json:"imdb"`
					Tmdb   int         `json:"tmdb"`
					Tvrage interface{} `json:"tvrage"`
				} `json:"ids"`
			} `json:"person"`
		} `json:"production"`
		Sound []struct {
			Jobs         []string `json:"jobs"`
			EpisodeCount int      `json:"episode_count"`
			Person       struct {
				Name string `json:"name"`
				Ids  struct {
					Trakt  int         `json:"trakt"`
					Slug   string      `json:"slug"`
					Imdb   string      `json:"imdb"`
					Tmdb   int         `json:"tmdb"`
					Tvrage interface{} `json:"tvrage"`
				} `json:"ids"`
			} `json:"person"`
		} `json:"sound"`
		CostumeMakeUp []struct {
			Jobs         []string `json:"jobs"`
			EpisodeCount int      `json:"episode_count"`
			Person       struct {
				Name string `json:"name"`
				Ids  struct {
					Trakt  int         `json:"trakt"`
					Slug   string      `json:"slug"`
					Imdb   string      `json:"imdb"`
					Tmdb   int         `json:"tmdb"`
					Tvrage interface{} `json:"tvrage"`
				} `json:"ids"`
			} `json:"person"`
		} `json:"costume & make-up"`
		Writing []struct {
			Jobs         []string `json:"jobs"`
			EpisodeCount int      `json:"episode_count"`
			Person       struct {
				Name string `json:"name"`
				Ids  struct {
					Trakt  int    `json:"trakt"`
					Slug   string `json:"slug"`
					Imdb   string `json:"imdb"`
					Tmdb   int    `json:"tmdb"`
					Tvrage int    `json:"tvrage"`
				} `json:"ids"`
			} `json:"person"`
		} `json:"writing"`
		Art []struct {
			Jobs         []string `json:"jobs"`
			EpisodeCount int      `json:"episode_count"`
			Person       struct {
				Name string `json:"name"`
				Ids  struct {
					Trakt  int         `json:"trakt"`
					Slug   string      `json:"slug"`
					Imdb   string      `json:"imdb"`
					Tmdb   int         `json:"tmdb"`
					Tvrage interface{} `json:"tvrage"`
				} `json:"ids"`
			} `json:"person"`
		} `json:"art"`
		Directing []struct {
			Jobs         []string `json:"jobs"`
			EpisodeCount int      `json:"episode_count"`
			Person       struct {
				Name string `json:"name"`
				Ids  struct {
					Trakt  int    `json:"trakt"`
					Slug   string `json:"slug"`
					Imdb   string `json:"imdb"`
					Tmdb   int    `json:"tmdb"`
					Tvrage int    `json:"tvrage"`
				} `json:"ids"`
			} `json:"person"`
		} `json:"directing"`
	} `json:"crew"`
}