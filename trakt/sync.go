//Package trakt ...
package trakt

import (
	//"fmt"
	"net/url"
	"time"
)

var (

	syncLastActivitiesURL = Hyperlink("sync/last_activities")
	syncPlayBackURL = Hyperlink("sync/playback/{type}?start_at={startDate}&end_at={endDate}")
	syncPlayBackIDURL = Hyperlink("sync/playback/{id}")
	syncCollectionTypeURL = Hyperlink("sync/collection/{type}")
	syncCollectionTypeURLExtended = Hyperlink("sync/collection/{type}?extended=metadata")
	syncCollectionURL = Hyperlink("sync/collection")
	syncCollectionRemoveURL = Hyperlink("sync/collection/remove")
	syncWatchedTypeURL = Hyperlink("sync/watched/{type}")
	syncWatchedTypeURLShowsExtended = Hyperlink("sync/watched/{type}?extended=noseasons")
	syncHistoryDateURL = Hyperlink("sync/history/type/id?start_at={startDate}&end_at={endDate}")
	syncHistoryURL = Hyperlink("sync/history")
	syncHistoriRemoveURL = Hyperlink("sync/history/remove")
	syncRatingsTypeURL = Hyperlink("sync/ratings/type/rating")
	syncRatingsURL = Hyperlink("sync/ratings")
	syncRatingsRemoveURL = Hyperlink("sync/ratings/remove")
	syncWatclistTypeURL = Hyperlink("sync/watchlist/type/sort")
	syncWatchListURL = Hyperlink("sync/watchlist")
	syncWatchListRemoveURL = Hyperlink("sync/watchlist/remove")
	syncRecommendationsTypeURL = Hyperlink("sync/recommendations/type/sort")
	syncRecommendationsURL = Hyperlink("sync/recommendations")
	syncRecommendationsRemoveURL = Hyperlink("sync/recommendations/remove")

)

//Sync Create a SyncService with the base url.URL
func (c *Client) Sync() (seasons *SyncService) {
	seasons = &SyncService{client: c}
	return
}
//SyncService ...
type SyncService struct {
	client *Client
}

// LastActivities ...
func (r *SyncService) LastActivities() (sync *DataLastActivities, result *Result) {
	url, _ := syncLastActivitiesURL.Expand(M{})
	result = r.client.get(url, &sync)
	return
}
// PlayBack ...
func (r *SyncService) PlayBack(types string, startDate string,endDate string) (sync *PlayBackProgress, result *Result) {
	url, _ := syncPlayBackURL.Expand(M{"type":types,"startDate": startDate,"endDate": endDate})
	result = r.client.get(url, &sync)
	return
}
// PlayBackRemove ...
func (r *SyncService) PlayBackRemove(id string) (sync *PlayBackProgress, result *Result) {
	url, _ := syncPlayBackIDURL.Expand(M{"id":id})
	result = r.client.delete(url, &sync)
	return
}
// GetCollectionMovies ...
func (r *SyncService) GetCollectionMovies(types string, extended string) (syncMovie *MovieCollection, result *Result) {
	var url *url.URL
	if extended=="false"{
		url, _ = syncCollectionTypeURL.Expand(M{"type": types})
	}else{
		url, _ = syncCollectionTypeURLExtended.Expand(M{"type": types})
	}
	result = r.client.get(url, &syncMovie)	

	return

}
// GetCollectionShows ...
func (r *SyncService) GetCollectionShows(types string, extended string) (syncShow *ShowCollection, result *Result) {
	var url *url.URL
	if extended=="false"{
		url, _ = syncCollectionTypeURL.Expand(M{"type": types})
	}else{
		url, _ = syncCollectionTypeURLExtended.Expand(M{"type": types})
	}
	result = r.client.get(url, &syncShow)	

	return

}

// GetWatchedShows ...
func (r *SyncService) GetWatchedShows(types string, extended string) (syncShow *WatchedShows, result *Result) {
	var url *url.URL
	if extended=="false"{
		url, _ = syncCollectionTypeURL.Expand(M{"type": types})
	}else{
		url, _ = syncWatchedTypeURLShowsExtended.Expand(M{"type": types})
	}
	result = r.client.get(url, &syncShow)	

	return

}


//PlayBackProgress ...
type PlayBackProgress []struct {
	Progress float64       `json:"progress"`
	PausedAt time.Time `json:"paused_at"`
	ID       int       `json:"id"`
	Type     string    `json:"type"`
	Movie    struct {
		Title string `json:"title"`
		Year  int    `json:"year"`
		Ids   struct {
			Trakt int    `json:"trakt"`
			Slug  string `json:"slug"`
			Imdb  string `json:"imdb"`
			Tmdb  int    `json:"tmdb"`
		} `json:"ids"`
	} `json:"movie,omitempty"`
	Episode struct {
		Season int    `json:"season"`
		Number int    `json:"number"`
		Title  string `json:"title"`
		Ids    struct {
			Trakt int    `json:"trakt"`
			Tvdb  int    `json:"tvdb"`
			Imdb  string `json:"imdb"`
			Tmdb  int    `json:"tmdb"`
		} `json:"ids"`
	} `json:"episode,omitempty"`
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
	} `json:"show,omitempty"`
}

//DataLastActivities ...
type DataLastActivities struct {
	All    time.Time `json:"all"`
	Movies struct {
		WatchedAt         time.Time `json:"watched_at"`
		CollectedAt       time.Time `json:"collected_at"`
		RatedAt           time.Time `json:"rated_at"`
		WatchlistedAt     time.Time `json:"watchlisted_at"`
		RecommendationsAt time.Time `json:"recommendations_at"`
		CommentedAt       time.Time `json:"commented_at"`
		PausedAt          time.Time `json:"paused_at"`
		HiddenAt          time.Time `json:"hidden_at"`
	} `json:"movies"`
	Episodes struct {
		WatchedAt     time.Time `json:"watched_at"`
		CollectedAt   time.Time `json:"collected_at"`
		RatedAt       time.Time `json:"rated_at"`
		WatchlistedAt time.Time `json:"watchlisted_at"`
		CommentedAt   time.Time `json:"commented_at"`
		PausedAt      time.Time `json:"paused_at"`
	} `json:"episodes"`
	Shows struct {
		RatedAt           time.Time `json:"rated_at"`
		WatchlistedAt     time.Time `json:"watchlisted_at"`
		RecommendationsAt time.Time `json:"recommendations_at"`
		CommentedAt       time.Time `json:"commented_at"`
		HiddenAt          time.Time `json:"hidden_at"`
	} `json:"shows"`
	Seasons struct {
		RatedAt       time.Time `json:"rated_at"`
		WatchlistedAt time.Time `json:"watchlisted_at"`
		CommentedAt   time.Time `json:"commented_at"`
		HiddenAt      time.Time `json:"hidden_at"`
	} `json:"seasons"`
	Comments struct {
		LikedAt time.Time `json:"liked_at"`
	} `json:"comments"`
	Lists struct {
		LikedAt     time.Time `json:"liked_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		CommentedAt time.Time `json:"commented_at"`
	} `json:"lists"`
	Watchlist struct {
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"watchlist"`
	Recommendations struct {
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"recommendations"`
	Account struct {
		SettingsAt  time.Time `json:"settings_at"`
		FollowedAt  time.Time `json:"followed_at"`
		FollowingAt time.Time `json:"following_at"`
		PendingAt   time.Time `json:"pending_at"`
	} `json:"account"`
}

//MovieCollection ...
type MovieCollection []struct {
	CollectedAt time.Time `json:"collected_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Movie       struct {
		Title string `json:"title"`
		Year  int    `json:"year"`
		Ids   struct {
			Trakt int    `json:"trakt"`
			Slug  string `json:"slug"`
			Imdb  string `json:"imdb"`
			Tmdb  int    `json:"tmdb"`
		} `json:"ids"`
	} `json:"movie"`
	Metadata struct {
		MediaType     string `json:"media_type"`
		Resolution    string `json:"resolution"`
		Hdr           string `json:"hdr"`
		Audio         string `json:"audio"`
		AudioChannels string `json:"audio_channels"`
		ThreeD        bool   `json:"3d"`
	} `json:"metadata"`
}

//ShowCollection ...
type ShowCollection []struct {
	LastCollectedAt time.Time `json:"last_collected_at"`
	LastUpdatedAt   time.Time `json:"last_updated_at"`
	Show            struct {
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
	Seasons []struct {
		Number   int `json:"number"`
		Episodes []struct {
			Number      int       `json:"number"`
			CollectedAt time.Time `json:"collected_at"`
			Metadata    struct {
				MediaType     string `json:"media_type"`
				Resolution    string `json:"resolution"`
				Hdr           string `json:"hdr"`
				Audio         string `json:"audio"`
				AudioChannels string `json:"audio_channels"`
				ThreeD        bool   `json:"3d"`
			} `json:"metadata"`
		} `json:"episodes"`
	} `json:"seasons"`
}

type WatchedShows []struct {
	Plays         int         `json:"plays"`
	LastWatchedAt time.Time   `json:"last_watched_at"`
	LastUpdatedAt time.Time   `json:"last_updated_at"`
	ResetAt       interface{} `json:"reset_at,omitempty"`
	Show          struct {
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
	Seasons []struct {
		Number   int `json:"number"`
		Episodes []struct {
			Number        int       `json:"number"`
			Plays         int       `json:"plays"`
			LastWatchedAt time.Time `json:"last_watched_at"`
		} `json:"episodes"`
	} `json:"seasons"`
}

type WatchedMovies []struct {
	Plays         int       `json:"plays"`
	LastWatchedAt time.Time `json:"last_watched_at"`
	LastUpdatedAt time.Time `json:"last_updated_at"`
	Movie         struct {
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