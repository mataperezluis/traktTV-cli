//Package trakt ...
package trakt

import (
	"net/url"
	"time"
)

var (
	movieURL         = Hyperlink("movies/{traktID}?extended=full")
	moviesPopularURL = Hyperlink("movies/popular")
    moviesTrendingURL = Hyperlink("movies/trending")
    moviesRecommendedURL = Hyperlink("movies/recommended/{period}")
    moviesPlayedURL = Hyperlink("movies/played/{period}")
    moviesWatchedURL = Hyperlink("movies/watched/{period}")
    moviesCollectedURL = Hyperlink("movies/collected/{period}")
	moviesAnticipatedURL = Hyperlink("movies/anticipated")
	moviesUpdatesURL = Hyperlink("movies/updates/{start_date}")
	moviesUpdatesIDURL = Hyperlink("movies/updates/id/{start_date}")	
	moviesSearchURL  = Hyperlink("search?query={query}&type=movie")
	movieAliasURL    = Hyperlink("movies/{traktID}/aliases")
	movieBoxOfficeURL    = Hyperlink("/movies/boxoffice")	
	movieReleasesURL = Hyperlink("/movies/{traktID}/releases/{country}")
	movieTranslationsURL    = Hyperlink("movies/{traktID}/translations/{lang}")	
	movieCommentsURL    = Hyperlink("movies/{traktID}/comments/{sort}")
	movieListsURL    = Hyperlink("movies/{traktID}/lists/{tipo}/{sort}")	
	moviePeopleURL    = Hyperlink("movies/{traktID}/people")	
	moviePeopleExtendedURL    = Hyperlink("movies/{traktID}/people?extended=guest_stars")	
	movieRatingsURL    = Hyperlink("movies/{traktID}/ratings")	
	movieRelatedURL    = Hyperlink("movies/{traktID}/related")	
	movieStatsURL    = Hyperlink("movies/{traktID}/stats")
	movieWatchingURL    = Hyperlink("movies/{traktID}/watching")
	moviesByIDURL    = Hyperlink("search?id_type={id_type}&id={id}&type=movie")
)

//Movies Create a MoviesService with the base url.URL
func (c *Client) Movies() (Movies *MoviesService) {
	Movies = &MoviesService{client: c}
	return
}
//MoviesService ...
type MoviesService struct {
	client *Client
}

// One returns a single Movie identified by a Trakt ID. It also returns a Result
// object to inspect the returned response of the server.
func (r *MoviesService) One(traktID string) (Movie *MovieInfo, result *Result) {
	url, _ := movieURL.Expand(M{"traktID":traktID})
	result = r.client.get(url, &Movie)
	return
}
//Alias ...
func (r *MoviesService) Alias(traktID string) (Movie *MovieAlias, result *Result) {
	url, _ := movieAliasURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &Movie)
	return
}
//BoxOffice ...
func (r *MoviesService) BoxOffice() (Movie []MovieBoxOffice, result *Result) {
	url, _ := movieBoxOfficeURL.Expand(M{})
	result = r.client.get(url, &Movie)
	return
}
//Releases ...
func (r *MoviesService) Releases(traktID string, countryC string) (Movie *MovieRelease, result *Result) {
	url, _ := movieReleasesURL.Expand(M{"traktID": traktID,"country": countryC})
	result = r.client.get(url, &Movie)
	return
}
//Translations ...
func (r *MoviesService) Translations(traktID string, lang string) (Movie *MovieTranslations, result *Result) {
	url, _ := movieTranslationsURL.Expand(M{"traktID": traktID,"lang": lang})
	result = r.client.get(url, &Movie)
	return
}
//Comments ...
func (r *MoviesService) Comments(traktID string, sort string) (Movie *MovieComment, result *Result) {
	url, _ := movieCommentsURL.Expand(M{"traktID": traktID,"sort": sort})
	result = r.client.get(url, &Movie)
	return
}
//List ...
func (r *MoviesService) List(traktID string, tipo string,sort string) (Movie *MovieList, result *Result) {
	url, _ := movieListsURL.Expand(M{"traktID": traktID,"tipo":tipo,"sort": sort})
	result = r.client.get(url, &Movie)
	return
}


//People ...
func (r *MoviesService) People(traktID string) (Movie *MovieCast, result *Result) {
	var url *url.URL
	url, _ = moviePeopleURL.Expand(M{"traktID": traktID})

	result = r.client.get(url, &Movie)
	return
}
//Ratings ...
func (r *MoviesService) Ratings(traktID string) (Movie *MovieRatings, result *Result) {
	url, _ := movieRatingsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &Movie)
	return
}
//Related ...
func (r *MoviesService) Related(traktID string) (Movie *MovieRelated, result *Result) {
	url, _ := movieRelatedURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &Movie)
	return
}
//Stats ...
func (r *MoviesService) Stats(traktID string) (Movie *MovieStats, result *Result) {
	url, _ := movieStatsURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &Movie)
	return
}
//Watching ...
func (r *MoviesService) Watching(traktID string) (Movie []User, result *Result) {
	url, _ := movieWatchingURL.Expand(M{"traktID": traktID})
	result = r.client.get(url, &Movie)
	return
}

//OneOfType ...
func (r *MoviesService) OneOfType(id string, idType string) (Movie *Movie, result *Result) {
	Movies := []MovieResult{}
	url, _ := moviesByIDURL.Expand(M{"id_type": idType, "id": id})
	result = r.client.get(url, &Movies)
	if len(Movies) > 0 {
		return Movies[0].Movie, result
	}
	return nil, result
}
//AllPopular ...
func (r *MoviesService) AllPopular() (Movies []MovieData, result *Result) {
	url, _ := moviesPopularURL.Expand(M{})
	result = r.client.get(url, &Movies)
	return
}
//Trending ...
func (r *MoviesService) Trending() (Movies []MovieTrending, result *Result) {
	url, _ := moviesTrendingURL.Expand(M{})
	result = r.client.get(url, &Movies)
	return
}
//Recommended ...
func (r *MoviesService) Recommended(period string) (Movies []MovieRecommended, result *Result) {
	url, _ := moviesRecommendedURL.Expand(M{"period": period})
	result = r.client.get(url, &Movies)
	return
}
//Played ...
func (r *MoviesService) Played(period string) (Movies []MoviePlayed, result *Result) {
	url, _ := moviesPlayedURL.Expand(M{"period": period})
	result = r.client.get(url, &Movies)
	return
}
//Watched ...
func (r *MoviesService) Watched(period string) (Movies []MoviePlayed, result *Result) {
	url, _ := moviesWatchedURL.Expand(M{"period": period})
	result = r.client.get(url, &Movies)
	return
}
//Collected ...
func (r *MoviesService) Collected(period string) (Movies []MoviePlayed, result *Result) {
	url, _ := moviesCollectedURL.Expand(M{"period": period})
	result = r.client.get(url, &Movies)
	return
}
//Anticipated ...
func (r *MoviesService) Anticipated() (Movies []MovieAnticipated, result *Result) {
	url, _ := moviesAnticipatedURL.Expand(M{})
	result = r.client.get(url, &Movies)
	return
}
//Updates ...
func (r *MoviesService) Updates(startDate string) (Movies []MovieUpdate, result *Result) {
	url, _ := moviesUpdatesURL.Expand(M{"start_date": startDate})
	result = r.client.get(url, &Movies)
	return
}
//UpdatesID ...
func (r *MoviesService) UpdatesID(startDate string) (MoviesID UpdatesIDval, result *Result) {
	url, _ := moviesUpdatesIDURL.Expand(M{"start_date": startDate})
	result = r.client.get(url, &MoviesID)
	return
}

//Search ...
func (r *MoviesService) Search(query string) (Movies []MovieResult, result *Result) {
	url, _ := moviesSearchURL.Expand(M{"query": query})
	result = r.client.get(url, &Movies)
	return
}

//MovieInfo struct for the Trakt v2 API
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

//MovieRecommended ...
type MovieRecommended struct {
	UserCount int  `json:"user_count"`
	Movie     MovieData `json:"Movie"`
}
//MovieAnticipated ...
type MovieAnticipated struct {
	ListCount int  `json:"list_count"`
	Movie     MovieData `json:"Movie"`
}
//MovieTranslations ...
type MovieTranslations []struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Language string `json:"language"`
}
//MovieTrending ...
type MovieTrending struct {
	Watchers int  `json:"watchers"`
	Movie     MovieData `json:"Movie"`
}
//MovieUpdate ...
type MovieUpdate struct {
	UpdatedAt time.Time `json:"updated_at"`
	Movie     MovieData `json:"Movie"`
}
//MoviePlayed ...
type MoviePlayed struct {
	WatcherCount   int  `json:"watcher_count"`
	PlayCount      int  `json:"play_count"`
	CollectedCount int  `json:"collected_count"`
	Movie           MovieData `json:"Movie"`
}
//MovieBoxOffice ...
type MovieBoxOffice struct {
	Revenue   int  `json:"revenue"`
	Movie           MovieData `json:"Movie"`
}
//MovieAlias ...
type MovieAlias []struct {
	Title   string `json:"title"`
	Country string `json:"country"`
}

//MovieCert ...	
type MovieCert []struct {
	Certification string `json:"certification"`
	Country       string `json:"country"`
}
//IdsMovie ...
type IdsMovie struct {
	Trakt int    `json:"trakt"`
	Slug  string `json:"slug"`
	Imdb  string `json:"imdb"`
	Tmdb  int    `json:"tmdb"`
}
//MovieData ...
type MovieData struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   IdsMovie    `json:"ids"`
}
//MovieResultTrending ...
type MovieResultTrending struct {
	Score float64 `json:"score"`
	Movie  *MovieTrending   `json:"Movie"`
	Type  string  `json:"type"`
}
//MovieResult ...
type MovieResult struct {
	Score float64 `json:"score"`
	Movie  *Movie   `json:"Movie"`
	Type  string  `json:"type"`
}
//MovieComment ...
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
//MovieList ...
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


//MovieRatings ...
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
//MovieRelated ...
type MovieRelated []struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Ids   IdsMovie `json:"ids"`
}


//MovieStats ...
type MovieStats struct {
	Watchers          int `json:"watchers"`
	Plays             int `json:"plays"`
	Collectors        int `json:"collectors"`
	Comments          int `json:"comments"`
	Lists             int `json:"lists"`
	Votes             int `json:"votes"`
	Recommended       int `json:"recommended"`
}
//MovieNext ...
type MovieNext struct {
	Season int    `json:"season"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	Ids    IdsMovie `json:"ids"`
}
//MovieProgress ...
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
//MovieRelease ...
type MovieRelease []struct {
	Country       string      `json:"country"`
	Certification string      `json:"certification"`
	ReleaseDate   string      `json:"release_date"`
	ReleaseType   string      `json:"release_type"`
	Note          interface{} `json:"note"`
}


//MovieCast ...
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