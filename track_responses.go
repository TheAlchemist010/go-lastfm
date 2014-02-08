package lastfm

type TrackStreamableResponse struct {
	FullTrack  int `xml:"fulltrack,attr"`
	Streamable int `xml:",chardata"`
}

type TrackResponse struct {
	Name       string                  `xml:"name"`
	ID         int64                   `xml:"id"`
	MBID       string                  `xml:"mbid"`
	URL        string                  `xml:"url"`
	Duration   int64                   `xml:"duration"`
	Streamable TrackStreamableResponse `xml:"streamable"`
	Listeners  int64                   `xml:"listeners"`
	Playcount  int64                   `xml:"playcount"`
	Artist     ArtistResponse          `xml:"artist"`
	Image      []LastfmImageResponse   `xml:"image"`
	UserLoved  int                     `xml:"userloved"`
	Album      AlbumResponse           `xml:"album"`
	TopTags    TopTagsResponse         `xml:"toptags"`
	Wiki       LastfmWikiResponse      `xml:"wiki"`
}

type TrackResponseNoArtistStruct struct {
	Name       string                  `xml:"name"`
	ID         int64                   `xml:"id"`
	MBID       string                  `xml:"mbid"`
	URL        string                  `xml:"url"`
	Duration   int64                   `xml:"duration"`
	Streamable TrackStreamableResponse `xml:"streamable"`
	Listeners  int64                   `xml:"listeners"`
	Playcount  int64                   `xml:"playcount"`
	Artist     string                  `xml:"artist"`
	Image      []LastfmImageResponse   `xml:"image"`
	UserLoved  int                     `xml:"userloved"`
	Album      AlbumResponse           `xml:"album"`
	TopTags    TopTagsResponse         `xml:"toptags"`
	Wiki       LastfmWikiResponse      `xml:"wiki"`
}
type TrackInfoResponse struct {
	LastfmStatusResponse
	Track TrackResponse `xml:"track"`
}

type TopTracksResponse struct {
	LastfmStatusResponse
	TopTracks []TrackResponse `xml:"toptracks>track"`
}

type TrackSimilarResponse struct {
	LastfmStatusResponse
	SimilarTracks []TrackResponse `xml:"similartracks>track"`
}

type TrackSingleCorrectionResponse struct {
	Index           int           `xml:"index,attr"`
	ArtistCorrected int           `xml:"artistcorrected,attr"`
	TrackCorrected  int           `xml:"trackcorrected,attr"`
	Track           TrackResponse `xml:"track"`
}

type TrackCorrectionResponse struct {
	LastfmStatusResponse
	TrackCorrections []TrackSingleCorrectionResponse `xml:"corrections>correction"`
}

type TrackSearchResponse struct {
	LastfmStatusResponse
	LastfmOpenSearchResponse
	TrackMatches []TrackResponseNoArtistStruct `xml:"results>trackmatches>track"`
}

type TracksLovedResponse struct {
	LastfmStatusResponse
	LovedTracks []TrackResponse `xml:"lovedtracks>track"`
}

type RecentTrackResponse struct {
	TrackResponse
	NowPlaying string             `xml:"nowplaying,attr"`
	Date       LastfmDateResponse `xml:"date"`
}

type TracksRecentResponse struct {
	LastfmStatusResponse
	RecentTracks []RecentTrackResponse `xml:"recenttracks>track"`
}
