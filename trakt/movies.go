package trakt

import (
	"net/url"
	"time"
)

var (
	MovieURL         = Hyperlink("movies/{traktID}?extended=full")
	MoviesPopularURL = Hyperlink("movies/popular")
    MoviesTrendingURL = Hyperlink("movies/trending")
    MoviesRecommendedURL = Hyperlink("movies/recommended/{period}")
    MoviesPlayedURL = Hyperlink("movies/played/{period}")
    MoviesWatchedURL = Hyperlink("movies/watched/{period}")
    MoviesCollectedURL = Hyperlink("movies/collected/{period}")
	MoviesAnticipatedURL = Hyperlink("movies/anticipated")
	MoviesUpdatesURL = Hyperlink("movies/updates/{start_date}")
	MoviesUpdatesIdURL = Hyperlink("movies/updates/id/{start_date}")	
	MoviesSearchURL  = Hyperlink("search?query={query}&type=movie")
	MovieAliasURL    = Hyperlink("movies/{traktID}/aliases")
	MovieBoxOfficeURL    = Hyperlink("/movies/boxoffice")	
	MovieReleasesURL = Hyperlink("/movies/{traktID}/releases/{country}")
	MovieTranslationsURL    = Hyperlink("movies/{traktID}/translations/{lang}")	
	MovieCommentsURL    = Hyperlink("movies/{traktID}/comments/{sort}")
	MovieListsURL    = Hyperlink("movies/{traktID}/lists/{tipo}/{sort}")	
	MoviePeopleURL    = Hyperlink("movies/{traktID}/people")	
	MoviePeopleExtendedURL    = Hyperlink("movies/{traktID}/people?extended=guest_stars")	
	MovieRatingsURL    = Hyperlink("movies/{traktID}/ratings")	
	MovieRelatedURL    = Hyperlink("movies/{traktID}/related")	
	MovieStatsURL    = Hyperlink("movies/{traktID}/stats")
	MovieWatchingURL    = Hyperlink("movies/{traktID}/watching")
	MoviesByIDURL    = Hyperlink("search?id_type={id_type}&id={id}&type=movie")
)

// Create a MoviesService with the base url.URL
func (c *Client) Movies() (Movies *MoviesService) {
	Movies = &MoviesService{client: c}
	return
}

type MoviesService struct {
	client *Client
}

// One returns a single Movie identified by a Trakt ID. It also returns a Result
// object to inspect the returned response of the server.
func (r *MoviesService) One(traktID string) (Movie *MovieInfo, result *Result) {
	url, _ := MovieURL.Expand(M{"traktID":traktID})
	result = r.client.get(url, &Movie)
	return
}

func (r *MoviesService) Alias(traktID string) (Movie *MovieAlias, result *Result) {
	url, _ := MovieAliasURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &Movie)
	return
}

func (r *MoviesService) BoxOffice() (Movie []MovieBoxOffice, result *Result) {
	url, _ := MovieBoxOfficeURL.Expand(M{})
	result = r.client.get(url, &Movie)
	return
}

func (r *MoviesService) Releases(traktID string, countryC string) (Movie *MovieRelease, result *Result) {
	url, _ := MovieReleasesURL.Expand(M{"traktID": traktID,"country": countryC})
	result = r.client.get(url, &Movie)
	return
}

func (r *MoviesService) Translations(traktID string, lang string) (Movie *MovieTranslations, result *Result) {
	url, _ := MovieTranslationsURL.Expand(M{"traktID": traktID,"lang": lang})
	result = r.client.get(url, &Movie)
	return
}

func (r *MoviesService) Comments(traktID string, sort string) (Movie *MovieComment, result *Result) {
	url, _ := MovieCommentsURL.Expand(M{"traktID": traktID,"sort": sort})
	result = r.client.get(url, &Movie)
	return
}

func (r *MoviesService) List(traktID string, tipo string,sort string) (Movie *MovieList, result *Result) {
	url, _ := MovieListsURL.Expand(M{"traktID": traktID,"tipo":tipo,"sort": sort})
	result = r.client.get(url, &Movie)
	return
}



func (r *MoviesService) People(traktID string) (Movie *MovieCast, result *Result) {
	var url *url.URL
	url, _ = MoviePeopleURL.Expand(M{"traktID": traktID})

	result = r.client.get(url, &Movie)
	return
}

func (r *MoviesService) Ratings(traktID string) (Movie *MovieRatings, result *Result) {
	url, _ := MovieRatingsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &Movie)
	return
}

func (r *MoviesService) Related(traktID string) (Movie *MovieRelated, result *Result) {
	url, _ := MovieRelatedURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &Movie)
	return
}

func (r *MoviesService) Stats(traktID string) (Movie *MovieStats, result *Result) {
	url, _ := MovieStatsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &Movie)
	return
}

func (r *MoviesService) Watching(traktID string) (Movie []User, result *Result) {
	url, _ := MovieWatchingURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &Movie)
	return
}


func (r *MoviesService) OneOfType(id string, idType string) (Movie *Movie, result *Result) {
	Movies := []MovieResult{}
	url, _ := MoviesByIDURL.Expand(M{"id_type": idType, "id": id})
	result = r.client.get(url, &Movies)
	if len(Movies) > 0 {
		return Movies[0].Movie, result
	}
	return nil, result
}

func (r *MoviesService) AllPopular() (Movies []MovieData, result *Result) {
	url, _ := MoviesPopularURL.Expand(M{})
	result = r.client.get(url, &Movies)
	return
}

func (r *MoviesService) Trending() (Movies []MovieTrending, result *Result) {
	url, _ := MoviesTrendingURL.Expand(M{})
	result = r.client.get(url, &Movies)
	return
}

func (r *MoviesService) Recommended(period string) (Movies []MovieRecommended, result *Result) {
	url, _ := MoviesRecommendedURL.Expand(M{"period": period})
	result = r.client.get(url, &Movies)
	return
}

func (r *MoviesService) Played(period string) (Movies []MoviePlayed, result *Result) {
	url, _ := MoviesPlayedURL.Expand(M{"period": period})
	result = r.client.get(url, &Movies)
	return
}

func (r *MoviesService) Watched(period string) (Movies []MoviePlayed, result *Result) {
	url, _ := MoviesWatchedURL.Expand(M{"period": period})
	result = r.client.get(url, &Movies)
	return
}

func (r *MoviesService) Collected(period string) (Movies []MoviePlayed, result *Result) {
	url, _ := MoviesCollectedURL.Expand(M{"period": period})
	result = r.client.get(url, &Movies)
	return
}

func (r *MoviesService) Anticipated() (Movies []MovieAnticipated, result *Result) {
	url, _ := MoviesAnticipatedURL.Expand(M{})
	result = r.client.get(url, &Movies)
	return
}

func (r *MoviesService) Updates(startDate string) (Movies []MovieUpdate, result *Result) {
	url, _ := MoviesUpdatesURL.Expand(M{"start_date": startDate})
	result = r.client.get(url, &Movies)
	return
}

func (r *MoviesService) UpdatesId(startDate string) (MoviesId UpdatesIdval, result *Result) {
	url, _ := MoviesUpdatesIdURL.Expand(M{"start_date": startDate})
	result = r.client.get(url, &MoviesId)
	return
}


func (r *MoviesService) Search(query string) (Movies []MovieResult, result *Result) {
	url, _ := MoviesSearchURL.Expand(M{"query": query})
	result = r.client.get(url, &Movies)
	return
}

// Movie struct for the Trakt v2 API
type MovieInfo struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   struct {
		Trakt int    `json:"trakt"`
		Slug  string `json:"slug"`
		Imdb  string `json:"imdb"`
		Tmdb  int    `json:"tmdb"`
	} `json:"ids"`
	Tagline               string      `json:"tagline"`
	Overview              string      `json:"overview"`
	Released              string      `json:"released"`
	Runtime               int         `json:"runtime"`
	Country               string      `json:"country"`
	UpdatedAt             time.Time   `json:"updated_at"`
	Trailer               interface{} `json:"trailer"`
	Homepage              string      `json:"homepage"`
	Status                string      `json:"status"`
	Rating                int         `json:"rating"`
	Votes                 int         `json:"votes"`
	CommentCount          int         `json:"comment_count"`
	Language              string      `json:"language"`
	AvailableTranslations []string    `json:"available_translations"`
	Genres                []string    `json:"genres"`
	Certification         string      `json:"certification"`
}


type MovieRecommended struct {
	UserCount int  `json:"user_count"`
	Movie     MovieData `json:"Movie"`
}

type MovieAnticipated struct {
	ListCount int  `json:"list_count"`
	Movie     MovieData `json:"Movie"`
}

type MovieTranslations []struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Language string `json:"language"`
}

type MovieTrending struct {
	Watchers int  `json:"watchers"`
	Movie     MovieData `json:"Movie"`
}
type MovieUpdate struct {
	UpdatedAt time.Time `json:"updated_at"`
	Movie     MovieData `json:"Movie"`
}

type MoviePlayed struct {
	WatcherCount   int  `json:"watcher_count"`
	PlayCount      int  `json:"play_count"`
	CollectedCount int  `json:"collected_count"`
	Movie           MovieData `json:"Movie"`
}

type MovieBoxOffice struct {
	Revenue   int  `json:"revenue"`
	Movie           MovieData `json:"Movie"`
}

type MovieAlias []struct {
	Title   string `json:"title"`
	Country string `json:"country"`
}

	
type MovieCert []struct {
	Certification string `json:"certification"`
	Country       string `json:"country"`
}

type IdsMovie struct {
	Trakt int    `json:"trakt"`
	Slug  string `json:"slug"`
	Imdb  string `json:"imdb"`
	Tmdb  int    `json:"tmdb"`
}
type MovieData struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   IdsMovie    `json:"ids"`
}

type MovieResultTrending struct {
	Score float64 `json:"score"`
	Movie  *MovieTrending   `json:"Movie"`
	Type  string  `json:"type"`
}

type MovieResult struct {
	Score float64 `json:"score"`
	Movie  *Movie   `json:"Movie"`
	Type  string  `json:"type"`
}

type MovieComment []struct {
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

type MovieList []struct {
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



type MovieRatings struct {
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

type MovieRelated []struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   IdsMovie `json:"ids"`
}



type MovieStats struct {
	Watchers          int `json:"watchers"`
	Plays             int `json:"plays"`
	Collectors        int `json:"collectors"`
	Comments          int `json:"comments"`
	Lists             int `json:"lists"`
	Votes             int `json:"votes"`
	Recommended       int `json:"recommended"`
}

type MovieNext struct {
	Season int    `json:"season"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Ids    IdsMovie `json:"ids"`
}

type MovieProgress struct {
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

type MovieRelease []struct {
	Country       string      `json:"country"`
	Certification string      `json:"certification"`
	ReleaseDate   string      `json:"release_date"`
	ReleaseType   string      `json:"release_type"`
	Note          interface{} `json:"note"`
}

//------------------------cast struct -------------------------------------

type MovieCast struct {
		Cast []struct {
			Characters []string `json:"characters"`
			Person     struct {
				Name string `json:"name"`
				Ids  struct {
					Trakt int    `json:"trakt"`
					Slug  string `json:"slug"`
					Imdb  string `json:"imdb"`
					Tmdb  int    `json:"tmdb"`
				} `json:"ids"`
			} `json:"person"`
		} `json:"cast"`
		Crew struct {
			Production []struct {
				Jobs   []string `json:"jobs,omitempty"`
				Person struct {
					Name string `json:"name"`
					Ids  struct {
						Trakt int    `json:"trakt"`
						Slug  string `json:"slug"`
						Imdb  string `json:"imdb"`
						Tmdb  int    `json:"tmdb"`
					} `json:"ids"`
				} `json:"person"`
				Job []string `json:"job,omitempty"`
			} `json:"production"`
			Art []struct {
				Jobs   []string `json:"jobs"`
				Person struct {
					Name string `json:"name"`
					Ids  struct {
						Trakt int    `json:"trakt"`
						Slug  string `json:"slug"`
						Imdb  string `json:"imdb"`
						Tmdb  int    `json:"tmdb"`
					} `json:"ids"`
				} `json:"person"`
			} `json:"art"`
			Crew []struct {
				Jobs   []string `json:"jobs"`
				Person struct {
					Name string `json:"name"`
					Ids  struct {
						Trakt int    `json:"trakt"`
						Slug  string `json:"slug"`
						Imdb  string `json:"imdb"`
						Tmdb  int    `json:"tmdb"`
					} `json:"ids"`
				} `json:"person"`
			} `json:"crew"`
			CostumeMakeUp []struct {
				Jobs   []string `json:"jobs"`
				Person struct {
					Name string `json:"name"`
					Ids  struct {
						Trakt int    `json:"trakt"`
						Slug  string `json:"slug"`
						Imdb  string `json:"imdb"`
						Tmdb  int    `json:"tmdb"`
					} `json:"ids"`
				} `json:"person"`
			} `json:"costume & make-up"`
			Directing []struct {
				Jobs   []string `json:"jobs"`
				Person struct {
					Name string `json:"name"`
					Ids  struct {
						Trakt int    `json:"trakt"`
						Slug  string `json:"slug"`
						Imdb  string `json:"imdb"`
						Tmdb  int    `json:"tmdb"`
					} `json:"ids"`
				} `json:"person"`
			} `json:"directing"`
			Writing []struct {
				Jobs   []string `json:"jobs,omitempty"`
				Person struct {
					Name string `json:"name"`
					Ids  struct {
						Trakt int    `json:"trakt"`
						Slug  string `json:"slug"`
						Imdb  string `json:"imdb"`
						Tmdb  int    `json:"tmdb"`
					} `json:"ids"`
				} `json:"person"`
				Job []string `json:"job,omitempty"`
			} `json:"writing"`
			Sound []struct {
				Jobs   []string `json:"jobs"`
				Person struct {
					Name string `json:"name"`
					Ids  struct {
						Trakt int    `json:"trakt"`
						Slug  string `json:"slug"`
						Imdb  string `json:"imdb"`
						Tmdb  int    `json:"tmdb"`
					} `json:"ids"`
				} `json:"person"`
			} `json:"sound"`
			Camera []struct {
				Jobs   []string `json:"jobs"`
				Person struct {
					Name string `json:"name"`
					Ids  struct {
						Trakt int    `json:"trakt"`
						Slug  string `json:"slug"`
						Imdb  string `json:"imdb"`
						Tmdb  int    `json:"tmdb"`
					} `json:"ids"`
				} `json:"person"`
			} `json:"camera"`
		} `json:"crew"`
	}