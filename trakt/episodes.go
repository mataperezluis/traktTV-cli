package trakt

import (
	//"fmt"
	"net/url"
	"time"
)

var (

	ShowEpisodesTranslationsURL    = Hyperlink("shows/{traktID}/seasons/{seasonNumber}/episodes/{episodes}/translations/{lang}")	
	ShowEpisodesNumberURL = Hyperlink("shows/{showTraktID}/seasons/{seasonNumber}/episodes/{episodes}")
	ShowEpisodesCommentsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/comments/{sort}")
	ShowEpisodesListsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/lists/{tipo}/{sort}")
	ShowEpisodesPeopleURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/people")	
	ShowEpisodesPeopleExtendedURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/people?extended=guest_stars")
	ShowEpisodesRatingsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/ratings")	
	ShowEpisodesStatsURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/stats")
	ShowEpisodesWatchingURL    = Hyperlink("shows/{traktID}/seasons/{seasonN}/episodes/{episodes}/watching")
)

// Create a ShowsService with the base url.URL
func (c *Client) Episodes() (episodes *EpisodesService) {
	episodes = &EpisodesService{client: c}
	return
}

type EpisodesService struct {
	client *Client
}


// ByNumber returns a specific season of a particular Show.
func (r *EpisodesService) ByNumber(showTraktID string, seasonNumber string,episode string) (episodes *Episode, result *Result) {
	url, _ := ShowEpisodesNumberURL.Expand(M{
		"showTraktID":  showTraktID,
		"seasonNumber": seasonNumber,
		"episodes": episode,
	})
	result = r.client.get(url, &episodes)
	return
}

func (r *EpisodesService) Translations(traktID string, seasonNumber string,episode string, lang string) (episodes *ShowTranslations, result *Result) {
	url, _ := ShowEpisodesTranslationsURL.Expand(M{"traktID": traktID,"seasonNumber": seasonNumber,"episodes": episode,"lang": lang})
	result = r.client.get(url, &episodes)
	return
}

func (r *EpisodesService) Comments(traktID string, seasonN string,episode string,sort string) (episodes *EpisodeComments, result *Result) {
	url, _ := ShowEpisodesCommentsURL.Expand(M{"traktID": traktID,"seasonN": seasonN,"episodes": episode,"sort": sort})
	result = r.client.get(url, &episodes)
	return
}

func (r *EpisodesService) List(traktID string,seasonN string,episode string, tipo string,sort string) (episodes *EpisodesList, result *Result) {
	url, _ := ShowEpisodesListsURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode,"tipo":tipo,"sort": sort})
	result = r.client.get(url, &episodes)
	return
}

func (r *EpisodesService) People(traktID string,seasonN string,episode string,extended string) (episodes *EpisodeCast, result *Result) {
	var url *url.URL
	if extended=="false"{
		url, _ = ShowEpisodesPeopleURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode})
	}else{
		url, _ = ShowEpisodesPeopleExtendedURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode})
	}

	result = r.client.get(url, &episodes)
	return
}

func (r *EpisodesService) Ratings(traktID string,seasonN string,episode string) (episodes *SeasonRating, result *Result) {
	url, _ := ShowEpisodesRatingsURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode})
	result = r.client.get(url, &episodes)
	return
}

func (r *EpisodesService) Stats(traktID string,seasonN string,episode string) (episodes *SeasonStats, result *Result) {
	url, _ := ShowEpisodesStatsURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode})
	result = r.client.get(url, &episodes)
	return
}

func (r *EpisodesService) Watching(traktID string,seasonN string,episode string) (episodes *SeasonUser, result *Result) {
	url, _ := ShowEpisodesWatchingURL.Expand(M{"traktID": traktID,"seasonN":seasonN,"episodes": episode})
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
type IdsEpisode struct {
	Trakt int    `json:"trakt"`
	Tvdb  int    `json:"tvdb"`
	Imdb  string `json:"imdb"`
	Tmdb  int    `json:"tmdb"`
}

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