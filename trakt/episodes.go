//Package trakt ...
package trakt

import (
	//"fmt"
	"net/url"
	"time"
)

var (

	showEpisodesTranslationsURL    = Hyperlink("shows/{traktID}/seasons/{seasonNumber}/episodes/{episodes}/translations/{lang}")	
	showEpisodesNumberURL = Hyperlink("shows/{showTraktID}/seasons/{seasonNumber}/episodes/{episodes}")
	showEpisodesCommentsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/comments/{sort}")
	showEpisodesListsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/lists/{tipo}/{sort}")
	showEpisodesPeopleURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/people")	
	showEpisodesPeopleExtendedURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/people?extended=guest_stars")
	showEpisodesRatingsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/ratings")	
	showEpisodesStatsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/stats")
	showEpisodesWatchingURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/watching")
)

//Episodes Create a ShowsService with the base url.URL
func (c *Client) Episodes() (episodes *EpisodesService) {
	episodes = &EpisodesService{client: c}
	return
}

//EpisodesService ...
type EpisodesService struct {
	client *Client
}


// ByNumber returns a specific season of a particular Show.
func (r *EpisodesService) ByNumber(showTraktID string, seasonNumber string,episode string) (episodes *Episode, result *Result) {
	url, _ := showEpisodesNumberURL.Expand(M{
		"showTraktID":  showTraktID,
		"seasonNumber": seasonNumber,
		"episodes": episode,
	})
	result = r.client.get(url, &episodes)
	return
}

//Translations ...
func (r *EpisodesService) Translations(traktID string, seasonNumber string,episode string, lang string) (episodes *ShowTranslations, result *Result) {
	url, _ := showEpisodesTranslationsURL.Expand(M{"traktID": traktID,"seasonNumber": seasonNumber,"episodes": episode,"lang": lang})
	result = r.client.get(url, &episodes)
	return
}

//Comments ...
func (r *EpisodesService) Comments(traktID string, seasonN string,episode string,sort string) (episodes *EpisodeComments, result *Result) {
	url, _ := showEpisodesCommentsURL.Expand(M{"traktID": traktID,"seasonN": seasonN,"episodes": episode,"sort": sort})
	result = r.client.get(url, &episodes)
	return
}
//List ...
func (r *EpisodesService) List(traktID string,seasonN string,episode string, tipo string,sort string) (episodes *EpisodesList, result *Result) {
	url, _ := showEpisodesListsURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode,"tipo":tipo,"sort": sort})
	result = r.client.get(url, &episodes)
	return
}
// People ....
func (r *EpisodesService) People(traktID string,seasonN string,episode string,extended string) (episodes *EpisodeCast, result *Result) {
	var url *url.URL
	if extended=="false"{
		url, _ = showEpisodesPeopleURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode})
	}else{
		url, _ = showEpisodesPeopleExtendedURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode})
	}

	result = r.client.get(url, &episodes)
	return
}
//Ratings ...
func (r *EpisodesService) Ratings(traktID string,seasonN string,episode string) (episodes *SeasonRating, result *Result) {
	url, _ := showEpisodesRatingsURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode})
	result = r.client.get(url, &episodes)
	return
}
//Stats ...
func (r *EpisodesService) Stats(traktID string,seasonN string,episode string) (episodes *SeasonStats, result *Result) {
	url, _ := showEpisodesStatsURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode})
	result = r.client.get(url, &episodes)
	return
}
//Watching ...
func (r *EpisodesService) Watching(traktID string,seasonN string,episode string) (episodes *SeasonUser, result *Result) {
	url, _ := showEpisodesWatchingURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode})
	result = r.client.get(url, &episodes)
	return
}

// Episode struct for the Trakt v2 API
type Episode struct {
	Season int    `json:"season"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Ids    IdsEpisode    `json:"ids"`
}
//IdsEpisode ...
type IdsEpisode struct {
	Trakt int    `json:"trakt"`
	Tvdb  int    `json:"tvdb"`
	Imdb  string `json:"imdb"`
	Tmdb  int    `json:"tmdb"`
}
//EpisodeComments ...
type EpisodeComments []struct {
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

//EpisodesList ...
type EpisodesList []struct {
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

//EpisodeCast ...
type EpisodeCast struct {
	Cast []struct {
		Characters []string `json:"characters"`
		Person     struct {
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
		Characters []string `json:"characters"`
		Person     struct {
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
		Writing []struct {
			Jobs   []string `json:"jobs"`
			Person struct {
				Name string `json:"name"`
				Ids  struct {
					Trakt  int         `json:"trakt"`
					Slug   string      `json:"slug"`
					Imdb   string      `json:"imdb"`
					Tmdb   int         `json:"tmdb"`
					Tvrage interface{} `json:"tvrage"`
				} `json:"ids"`
			} `json:"person"`
		} `json:"writing"`
		Directing []struct {
			Jobs   []string `json:"jobs"`
			Person struct {
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
		Camera []struct {
			Jobs   []string `json:"jobs"`
			Person struct {
				Name string `json:"name"`
				Ids  struct {
					Trakt  int         `json:"trakt"`
					Slug   string      `json:"slug"`
					Imdb   string      `json:"imdb"`
					Tmdb   int         `json:"tmdb"`
					Tvrage interface{} `json:"tvrage"`
				} `json:"ids"`
			} `json:"person"`
		} `json:"camera"`
		Editing []struct {
			Jobs   []string `json:"jobs"`
			Person struct {
				Name string `json:"name"`
				Ids  struct {
					Trakt  int         `json:"trakt"`
					Slug   string      `json:"slug"`
					Imdb   string      `json:"imdb"`
					Tmdb   int         `json:"tmdb"`
					Tvrage interface{} `json:"tvrage"`
				} `json:"ids"`
			} `json:"person"`
		} `json:"editing"`
	} `json:"crew"`
}