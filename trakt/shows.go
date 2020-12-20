//Package trakt ...
package trakt

import (
	"net/url"
	"time"
)

var (
	showURL         = Hyperlink("shows/{traktID}")
	showsPopularURL = Hyperlink("shows/popular")
    showsTrendingURL = Hyperlink("shows/trending")
    showsRecommendedURL = Hyperlink("shows/recommended/{period}")
    showsPlayedURL = Hyperlink("shows/played/{period}")
    showsWatchedURL = Hyperlink("shows/watched/{period}")
    showsCollectedURL = Hyperlink("shows/collected/{period}")
	showsAnticipatedURL = Hyperlink("shows/anticipated")
	showsUpdatesURL = Hyperlink("shows/updates/{start_date}")
	showsUpdatesIDURL = Hyperlink("shows/updates/id/{start_date}")	
	showsSearchURL  = Hyperlink("search?query={query}&type=show")
	showAliasURL    = Hyperlink("shows/{traktID}/aliases")
	showCertificationsURL    = Hyperlink("shows/{traktID}/certifications")	
	showTranslationsURL    = Hyperlink("shows/{traktID}/translations/{lang}")	
	showCommentsURL    = Hyperlink("shows/{traktID}/comments/{sort}")
	showListsURL    = Hyperlink("shows/{traktID}/lists/{tipo}/{sort}")	
	showPeopleURL    = Hyperlink("shows/{traktID}/people")	
	showPeopleExtendedURL    = Hyperlink("shows/{traktID}/people?extended=guest_stars")	
	showRatingsURL    = Hyperlink("shows/{traktID}/ratings")	
	showRelatedURL    = Hyperlink("shows/{traktID}/related")	
	showStatsURL    = Hyperlink("shows/{traktID}/stats")
	showWatchingURL    = Hyperlink("shows/{traktID}/watching")
	showNextEpURL    = Hyperlink("shows/{traktID}/next_episode")
	showLastEpURL    = Hyperlink("shows/{traktID}/last_episode")
	showProgressURL = Hyperlink("shows/{traktID}/progress/collection?hidden={hiddenB}&specials={specialsB}&count_specials={countspecialsB}")
	showWatchedProgressURL = Hyperlink("shows/{traktID}/progress/watched?hidden={hiddenB}&specials={specialsB}&count_specials={countspecialsB}")
	showsByIDURL    = Hyperlink("search?id_type={id_type}&id={id}&type=show")
)

//Shows Create a ShowsService with the base url.URL
func (c *Client) Shows() (shows *ShowsService) {
	shows = &ShowsService{client: c}
	return
}
//ShowsService ...
type ShowsService struct {
	client *Client
}

// One returns a single show identified by a Trakt ID. It also returns a Result
// object to inspect the returned response of the server.
func (r *ShowsService) One(traktID string) (show *Show, result *Result) {
	url, _ := showURL.Expand(M{"traktID":traktID})
	result = r.client.get(url, &show)
	return
}
//Alias ...
func (r *ShowsService) Alias(traktID string) (show *ShowAlias, result *Result) {
	url, _ := showAliasURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}
//Certifications ...
func (r *ShowsService) Certifications(traktID string) (show *ShowCert, result *Result) {
	url, _ := showCertificationsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}
//Translations ...
func (r *ShowsService) Translations(traktID string, lang string) (show *ShowTranslations, result *Result) {
	url, _ := showTranslationsURL.Expand(M{"traktID": traktID,"lang": lang})
	result = r.client.get(url, &show)
	return
}
//Comments ...
func (r *ShowsService) Comments(traktID string, sort string) (show *ShowComment, result *Result) {
	url, _ := showCommentsURL.Expand(M{"traktID": traktID,"sort": sort})
	result = r.client.get(url, &show)
	return
}
//List ...
func (r *ShowsService) List(traktID string, tipo string,sort string) (show *ShowList, result *Result) {
	url, _ := showListsURL.Expand(M{"traktID": traktID,"tipo":tipo,"sort": sort})
	result = r.client.get(url, &show)
	return
}
//CollectionProgress ...
func (r *ShowsService) CollectionProgress(traktID string, hiddenB string, specialB string, countspecialsB string) (show *ShowProgress, result *Result) {
	url, _ := showProgressURL.Expand(M{"traktID": traktID,"hiddenB":hiddenB,"specialsB": specialB,"countspecialsB":countspecialsB})
	result = r.client.get(url, &show)
	return
}
//WatchedProgress ...
func (r *ShowsService) WatchedProgress(traktID string, hiddenB string, specialB string, countspecialsB string) (show *ShowProgress, result *Result) {
	url, _ := showWatchedProgressURL.Expand(M{"traktID": traktID,"hiddenB":hiddenB,"specialsB": specialB,"countspecialsB":countspecialsB})
	result = r.client.get(url, &show)
	return
}


//People ...
func (r *ShowsService) People(traktID string,extended string) (show *ShowCast, result *Result) {
	var url *url.URL
	if extended=="false"{
		url, _ = showPeopleURL.Expand(M{"traktID": traktID})
	}else{
		url, _ = showPeopleExtendedURL.Expand(M{"traktID": traktID})
	}
	result = r.client.get(url, &show)
	return
}
//Ratings ...
func (r *ShowsService) Ratings(traktID string) (show *ShowRatings, result *Result) {
	url, _ := showRatingsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}
//Related ...
func (r *ShowsService) Related(traktID string) (show *ShowRelated, result *Result) {
	url, _ := showRelatedURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}
//Stats ...
func (r *ShowsService) Stats(traktID string) (show *ShowStats, result *Result) {
	url, _ := showStatsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}
//Watching ...
func (r *ShowsService) Watching(traktID string) (show []User, result *Result) {
	url, _ := showWatchingURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}
//NextEpisode ...
func (r *ShowsService) NextEpisode(traktID string) (show *ShowNext, result *Result) {
	url, _ := showNextEpURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}
//LastEpisode ...
func (r *ShowsService) LastEpisode(traktID string) (show *ShowNext, result *Result) {
	url, _ := showLastEpURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &show)
	return
}
//OneOfType ...
func (r *ShowsService) OneOfType(id string, idType string) (show *Show, result *Result) {
	shows := []ShowResult{}
	url, _ := showsByIDURL.Expand(M{"id_type": idType, "id": id})
	result = r.client.get(url, &shows)
	if len(shows) > 0 {
		return shows[0].Show, result
	}
	return nil, result
}
//AllPopular ...
func (r *ShowsService) AllPopular() (shows []Show, result *Result) {
	url, _ := showsPopularURL.Expand(M{})
	result = r.client.get(url, &shows)
	return
}
//Trending ...
func (r *ShowsService) Trending() (shows []ShowTrending, result *Result) {
	url, _ := showsTrendingURL.Expand(M{})
	result = r.client.get(url, &shows)
	return
}
//Recommended ...
func (r *ShowsService) Recommended(period string) (shows []ShowRecommended, result *Result) {
	url, _ := showsRecommendedURL.Expand(M{"period": period})
	result = r.client.get(url, &shows)
	return
}
//Played ...
func (r *ShowsService) Played(period string) (shows []ShowPlayed, result *Result) {
	url, _ := showsPlayedURL.Expand(M{"period": period})
	result = r.client.get(url, &shows)
	return
}
//Watched ...
func (r *ShowsService) Watched(period string) (shows []ShowPlayed, result *Result) {
	url, _ := showsWatchedURL.Expand(M{"period": period})
	result = r.client.get(url, &shows)
	return
}
//Collected ...
func (r *ShowsService) Collected(period string) (shows []ShowPlayed, result *Result) {
	url, _ := showsCollectedURL.Expand(M{"period": period})
	result = r.client.get(url, &shows)
	return
}
//Anticipated ...
func (r *ShowsService) Anticipated() (shows []ShowAnticipated, result *Result) {
	url, _ := showsAnticipatedURL.Expand(M{})
	result = r.client.get(url, &shows)
	return
}
//Updates ...
func (r *ShowsService) Updates(startDate string) (shows []ShowUpdate, result *Result) {
	url, _ := showsUpdatesURL.Expand(M{"start_date": startDate})
	result = r.client.get(url, &shows)
	return
}
//UpdatesID ...
func (r *ShowsService) UpdatesID(startDate string) (showsID UpdatesIDval, result *Result) {
	url, _ := showsUpdatesIDURL.Expand(M{"start_date": startDate})
	result = r.client.get(url, &showsID)
	return
}

//Search ...
func (r *ShowsService) Search(query string) (shows []ShowResult, result *Result) {
	url, _ := showsSearchURL.Expand(M{"query": query})
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
//UpdatesIDval ...
type UpdatesIDval []int
//ShowRecommended ...
type ShowRecommended struct {
	UserCount int  `json:"user_count"`
	Show     ShowData `json:"show"`
}
//ShowAnticipated ...
type ShowAnticipated struct {
	ListCount int  `json:"list_count"`
	Show     ShowData `json:"show"`
}
//ShowTranslations ...
type ShowTranslations []struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Language string `json:"language"`
}
//ShowTrending ...
type ShowTrending struct {
	Watchers int  `json:"watchers"`
	Show     ShowData `json:"show"`
}
//ShowUpdate ...
type ShowUpdate struct {
	UpdatedAt time.Time `json:"updated_at"`
	Show     ShowData `json:"show"`
}
//ShowPlayed ...
type ShowPlayed struct {
	WatcherCount   int  `json:"watcher_count"`
	PlayCount      int  `json:"play_count"`
	CollectedCount int  `json:"collected_count"`
	CollectorCount int  `json:"collector_count"`
	Show           ShowData `json:"show"`
}
//ShowAlias ...
type ShowAlias []struct {
	Title   string `json:"title"`
	Country string `json:"country"`
}

//ShowCert ...	
type ShowCert []struct {
	Certification string `json:"certification"`
	Country       string `json:"country"`
}
//Ids ...
type Ids struct {
	Trakt int    `json:"trakt"`
	Slug  string `json:"slug"`
	Tvdb  int    `json:"tvdb"`
	Imdb  string `json:"imdb"`
	Tmdb  int    `json:"tmdb"`
}
//ShowData ...
type ShowData struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   Ids    `json:"ids"`
}
//ShowResultTrending ...
type ShowResultTrending struct {
	Score float64 `json:"score"`
	Show  *ShowTrending   `json:"show"`
	Type  string  `json:"type"`
}
//ShowResult ...
type ShowResult struct {
	Score float64 `json:"score"`
	Show  *Show   `json:"show"`
	Type  string  `json:"type"`
}
//ShowComment ...
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
//ShowList ...
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
//IdsM ...
type IdsM struct {
	Trakt int    `json:"trakt"`
	Slug  string `json:"slug"`
}
//IdsUsr ...
type IdsUsr struct {
	Slug string `json:"slug"`
}
// User ...
type User struct {
	Username string `json:"username"`
	Private  bool   `json:"private"`
	Name     string `json:"name"`
	Vip      bool   `json:"vip"`
	VipEp    bool   `json:"vip_ep"`
	Ids      IdsUsr    `json:"ids"`
}

//ShowRatings ...
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
//ShowRelated ...
type ShowRelated []struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   Ids `json:"ids"`
}
//ShowStats ...
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
//ShowNext ...
type ShowNext struct {
	Season int    `json:"season"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Ids    Ids `json:"ids"`
}
//ShowProgress ...
type ShowProgress struct {
	Aired           int       `json:"aired"`
	Completed       int       `json:"completed"`
	LastCollectedAt time.Time `json:"last_collected_at"`
	Seasons         []struct {
		Number    int    `json:"number"`
		Title     string `json:"title"`
		Aired     int    `json:"aired"`
		Completed int    `json:"completed"`
		Episodes  []struct {
			Number      int       `json:"number"`
			Completed   bool      `json:"completed"`
			CollectedAt time.Time `json:"collected_at"`
		} `json:"episodes"`
	} `json:"seasons"`
	HiddenSeasons []struct {
		Number int `json:"number"`
		Ids    struct {
			Trakt int `json:"trakt"`
			Tvdb  int `json:"tvdb"`
			Tmdb  int `json:"tmdb"`
		} `json:"ids"`
	} `json:"hidden_seasons"`
	NextEpisode struct {
		Season int    `json:"season"`
		Number int    `json:"number"`
		Title  string `json:"title"`
		Ids    struct {
			Trakt int         `json:"trakt"`
			Tvdb  int         `json:"tvdb"`
			Imdb  interface{} `json:"imdb"`
			Tmdb  interface{} `json:"tmdb"`
		} `json:"ids"`
	} `json:"next_episode"`
	LastEpisode struct {
		Season int    `json:"season"`
		Number int    `json:"number"`
		Title  string `json:"title"`
		Ids    struct {
			Trakt int         `json:"trakt"`
			Tvdb  int         `json:"tvdb"`
			Imdb  interface{} `json:"imdb"`
			Tmdb  interface{} `json:"tmdb"`
		} `json:"ids"`
	} `json:"last_episode"`
}

//ShowCast ...
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