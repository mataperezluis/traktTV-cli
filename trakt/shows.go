package trakt

import (
	"time"
)

var (
	ShowURL         = Hyperlink("shows/{traktID}")
	ShowsPopularURL = Hyperlink("shows/popular")
    ShowsTrendingURL = Hyperlink("shows/trending")
    ShowsRecommendedURL = Hyperlink("shows/recommended/{period}")
    ShowsPlayedURL = Hyperlink("shows/played/{period}")
    ShowsWatchedURL = Hyperlink("shows/watched/{period}")
    ShowsCollectedURL = Hyperlink("shows/collected/{period}")
	ShowsAnticipatedURL = Hyperlink("shows/anticipated")
	ShowsUpdatesURL = Hyperlink("shows/updates/{start_date}")
	ShowsUpdatesIdURL = Hyperlink("shows/updates/id/{start_date}")	
	ShowsSearchURL  = Hyperlink("search?query={query}&type=show")
	ShowAliasURL    = Hyperlink("shows/{traktID}/aliases")
	ShowCertificationsURL    = Hyperlink("shows/{traktID}/certifications")	
	ShowTranslationsURL    = Hyperlink("shows/{traktID}/translations/{lang}")	
	ShowCommentsURL    = Hyperlink("shows/{traktID}/comments/{sort}")
	ShowListsURL    = Hyperlink("shows/{traktID}/lists/{tipo}/{sort}")	
	ShowPeopleURL    = Hyperlink("shows/{traktID}/people")	
	ShowRatingsURL    = Hyperlink("shows/{traktID}/ratings")	
	ShowRelatedURL    = Hyperlink("shows/{traktID}/related")	
	ShowStatsURL    = Hyperlink("shows/{traktID}/stats")
	ShowWatchingURL    = Hyperlink("shows/{traktID}/watching")
	ShowNextEpURL    = Hyperlink("shows/{traktID}/next_episode")
	ShowLastEpURL    = Hyperlink("shows/{traktID}/last_episode")
	ShowsByIDURL    = Hyperlink("search?id_type={id_type}&id={id}&type=show")
)

// Create a ShowsService with the base url.URL
func (c *Client) Shows() (shows *ShowsService) {
	shows = &ShowsService{client: c}
	return
}

type ShowsService struct {
	client *Client
}

// One returns a single show identified by a Trakt ID. It also returns a Result
// object to inspect the returned response of the server.
func (r *ShowsService) One(traktID string) (show *Show, result *Result) {
	url, _ := ShowURL.Expand(M{"traktID":traktID})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) Alias(traktID string) (show *ShowAlias, result *Result) {
	url, _ := ShowAliasURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) Certifications(traktID string) (show *ShowCert, result *Result) {
	url, _ := ShowCertificationsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) Translations(traktID string, lang string) (show *ShowTranslations, result *Result) {
	url, _ := ShowTranslationsURL.Expand(M{"traktID": traktID,"lang": lang})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) Comments(traktID string, sort string) (show *ShowComment, result *Result) {
	url, _ := ShowCommentsURL.Expand(M{"traktID": traktID,"sort": sort})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) List(traktID string, tipo string,sort string) (show *ShowList, result *Result) {
	url, _ := ShowListsURL.Expand(M{"traktID": traktID,"tipo":tipo,"sort": sort})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) People(traktID string) (show *ShowCast, result *Result) {
	url, _ := ShowPeopleURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) Ratings(traktID string) (show *ShowRatings, result *Result) {
	url, _ := ShowRatingsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) Related(traktID string) (show *ShowRelated, result *Result) {
	url, _ := ShowRelatedURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) Stats(traktID string) (show *ShowStats, result *Result) {
	url, _ := ShowStatsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) Watching(traktID string) (show []User, result *Result) {
	url, _ := ShowWatchingURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) NextEpisode(traktID string) (show *ShowNext, result *Result) {
	url, _ := ShowNextEpURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) LastEpisode(traktID string) (show *ShowNext, result *Result) {
	url, _ := ShowLastEpURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}

func (r *ShowsService) OneOfType(id string, idType string) (show *Show, result *Result) {
	shows := []ShowResult{}
	url, _ := ShowsByIDURL.Expand(M{"id_type": idType, "id": id})
	result = r.client.get(url, &shows)
	if len(shows) > 0 {
		return shows[0].Show, result
	}
	return nil, result
}

func (r *ShowsService) AllPopular() (shows []Show, result *Result) {
	url, _ := ShowsPopularURL.Expand(M{})
	result = r.client.get(url, &shows)
	return
}

func (r *ShowsService) Trending() (shows []ShowTrending, result *Result) {
	url, _ := ShowsTrendingURL.Expand(M{})
	result = r.client.get(url, &shows)
	return
}

func (r *ShowsService) Recommended(period string) (shows []ShowRecommended, result *Result) {
	url, _ := ShowsRecommendedURL.Expand(M{"period": period})
	result = r.client.get(url, &shows)
	return
}

func (r *ShowsService) Played(period string) (shows []ShowPlayed, result *Result) {
	url, _ := ShowsPlayedURL.Expand(M{"period": period})
	result = r.client.get(url, &shows)
	return
}

func (r *ShowsService) Watched(period string) (shows []ShowPlayed, result *Result) {
	url, _ := ShowsWatchedURL.Expand(M{"period": period})
	result = r.client.get(url, &shows)
	return
}

func (r *ShowsService) Collected(period string) (shows []ShowPlayed, result *Result) {
	url, _ := ShowsCollectedURL.Expand(M{"period": period})
	result = r.client.get(url, &shows)
	return
}

func (r *ShowsService) Anticipated() (shows []ShowAnticipated, result *Result) {
	url, _ := ShowsAnticipatedURL.Expand(M{})
	result = r.client.get(url, &shows)
	return
}

func (r *ShowsService) Updates(startDate string) (shows []ShowUpdate, result *Result) {
	url, _ := ShowsUpdatesURL.Expand(M{"start_date": startDate})
	result = r.client.get(url, &shows)
	return
}

func (r *ShowsService) UpdatesId(startDate string) (showsId UpdatesIdval, result *Result) {
	url, _ := ShowsUpdatesIdURL.Expand(M{"start_date": startDate})
	result = r.client.get(url, &showsId)
	return
}


func (r *ShowsService) Search(query string) (shows []ShowResult, result *Result) {
	url, _ := ShowsSearchURL.Expand(M{"query": query})
	result = r.client.get(url, &shows)
	return
}

// Show struct for the Trakt v2 API
type Show struct {
	AiredEpisodes int `json:"aired_episodes"`
	Airs          struct {
		Day      string `json:"day"`
		Time     string `json:"time"`
		Timezone string `json:"timezone"`
	} `json:"airs"`
	AvailableTranslations []string `json:"available_translations"`
	Certification         string   `json:"certification"`
	Country               string   `json:"country"`
	FirstAired            string   `json:"first_aired"`
	Genres                []string `json:"genres"`
	Homepage              string   `json:"homepage"`
	IDs                   struct {
		Imdb   string `json:"imdb"`
		Slug   string `json:"slug"`
		Tmdb   int    `json:"tmdb"`
		Trakt  int    `json:"trakt"`
		Tvdb   int    `json:"tvdb"`
		Tvrage int    `json:"tvrage"`
	} `json:"ids"`
	Images struct {
		Banner struct {
			Full string `json:"full"`
		} `json:"banner"`
		Clearart struct {
			Full string `json:"full"`
		} `json:"clearart"`
		Fanart struct {
			Full   string `json:"full"`
			Medium string `json:"medium"`
			Thumb  string `json:"thumb"`
		} `json:"fanart"`
		Logo struct {
			Full string `json:"full"`
		} `json:"logo"`
		Poster struct {
			Full   string `json:"full"`
			Medium string `json:"medium"`
			Thumb  string `json:"thumb"`
		} `json:"poster"`
		Thumb struct {
			Full string `json:"full"`
		} `json:"thumb"`
	} `json:"images"`
	Language  string  `json:"language"`
	Network   string  `json:"network"`
	Overview  string  `json:"overview"`
	Rating    float64 `json:"rating"`
	Runtime   float64 `json:"runtime"`
	Status    string  `json:"status"`
	Title     string  `json:"title"`
	Trailer   string  `json:"trailer"`
	UpdatedAt string  `json:"updated_at"`
	Votes     int     `json:"votes"`
	Year      int     `json:"year"`
}

type UpdatesIdval []int

type ShowRecommended struct {
	UserCount int  `json:"user_count"`
	Show     ShowData `json:"show"`
}

type ShowAnticipated struct {
	ListCount int  `json:"list_count"`
	Show     ShowData `json:"show"`
}

type ShowTranslations []struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Language string `json:"language"`
}

type ShowTrending struct {
	Watchers int  `json:"watchers"`
	Show     ShowData `json:"show"`
}
type ShowUpdate struct {
	UpdatedAt time.Time `json:"updated_at"`
	Show     ShowData `json:"show"`
}

type ShowPlayed struct {
	WatcherCount   int  `json:"watcher_count"`
	PlayCount      int  `json:"play_count"`
	CollectedCount int  `json:"collected_count"`
	CollectorCount int  `json:"collector_count"`
	Show           ShowData `json:"show"`
}

type ShowAlias []struct {
	Title   string `json:"title"`
	Country string `json:"country"`
}

	
type ShowCert []struct {
	Certification string `json:"certification"`
	Country       string `json:"country"`
}

type Ids struct {
	Trakt int    `json:"trakt"`
	Slug  string `json:"slug"`
	Tvdb  int    `json:"tvdb"`
	Imdb  string `json:"imdb"`
	Tmdb  int    `json:"tmdb"`
}
type ShowData struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   Ids    `json:"ids"`
}

type ShowResultTrending struct {
	Score float64 `json:"score"`
	Show  *ShowTrending   `json:"show"`
	Type  string  `json:"type"`
}

type ShowResult struct {
	Score float64 `json:"score"`
	Show  *Show   `json:"show"`
	Type  string  `json:"type"`
}

type ShowComment []struct {
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
	Ids            IdsM       `json:"ids"`
	User           User      `json:"user"`
}

type ShowList []struct {
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
	Ids            IdsM       `json:"ids"`
	User           User      `json:"user"`
}
type IdsM struct {
	Trakt int    `json:"trakt"`
	Slug  string `json:"slug"`
}
type IdsUsr struct {
	Slug string `json:"slug"`
}
type User struct {
	Username string `json:"username"`
	Private  bool   `json:"private"`
	Name     string `json:"name"`
	Vip      bool   `json:"vip"`
	VipEp    bool   `json:"vip_ep"`
	Ids      IdsUsr    `json:"ids"`
}


type ShowRatings struct {
	Rating       float64 `json:"rating"`
	Votes        int     `json:"votes"`
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

type ShowRelated []struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   Ids `json:"ids"`
}

type ShowStats struct {
	Watchers          int `json:"watchers"`
	Plays             int `json:"plays"`
	Collectors        int `json:"collectors"`
	CollectedEpisodes int `json:"collected_episodes"`
	Comments          int `json:"comments"`
	Lists             int `json:"lists"`
	Votes             int `json:"votes"`
	Recommended       int `json:"recommended"`
}

type ShowNext struct {
	Season int    `json:"season"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Ids    Ids `json:"ids"`
}

//------------------------cast struct -------------------------------------

type ShowCast struct {
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
	Crew struct {
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
		Production []struct {
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
		VisualEffects []struct {
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
		} `json:"visual effects"`
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