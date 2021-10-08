package metadata

type DetailPageImage struct {
	Uri string `json:"uri"`
}

type DetailPageArtistInfo struct {
	Uri       string            `json:"uri"`
	Name      string            `json:"name"`
	Portraits []DetailPageImage `json:"portraits"`
	Verified  bool              `json:"verified"`
}
type DetailPageArtistHeaderImage struct {
	Image  string `json:"image"`
	Offset int    `json:"offset"`
}

type DetailPageArtistTopTrackRelease struct {
	Uri   string          `json:"uri"`
	Name  string          `json:"name"`
	Cover DetailPageImage `json:"cover"`
}
type DetailPageArtistTopTrack struct {
	Uri       string                          `json:"uri"`
	Playcount int                             `json:"playcount"`
	Name      string                          `json:"name"`
	Release   DetailPageArtistTopTrackRelease `json:"release"`
	Explicit  bool                            `json:"explicit"`
}
type DetailPageArtistUpcomingConcerts struct {
	InactiveArtist bool `json:"inactive_artist"`
}

type DetailPageArtistTopTracks struct {
	Tracks []DetailPageArtistTopTrack `json:"tracks"`
}

type DetailPageArtistRelatedArtist struct {
	Uri       string            `json:"uri"`
	Name      string            `json:"name"`
	Portraits []DetailPageImage `json:"portraits"`
}

type DetailPageArtistRelatedArtists struct {
	Artists []DetailPageArtistRelatedArtist `json:"artists"`
}
type DetailPageArtistBiography struct {
	// there is problems with the escaping of this for some reason
	// Text string `json:"text"`
}

type DetailPageGenericReleaseArtist struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

type DetailPageGenericReleaseTrack struct {
	Uri        string                           `json:"uri"`
	Playcount  int                              `json:"playcount"`
	Name       string                           `json:"name"`
	Popularity int                              `json:"popularity"`
	Number     int                              `json:"number"`
	Duration   int                              `json:"duration"`
	Explicit   bool                             `json:"explicit"`
	Playable   bool                             `json:"playable"`
	Artists    []DetailPageGenericReleaseArtist `json:"artists"`
}

type DetailPageGenericReleaseDisc struct {
	Number int                             `json:"number"`
	Name   string                          `json:"name"`
	Tracks []DetailPageGenericReleaseTrack `json:"tracks"`
}
type DetailPageGenericRelease struct {
	Uri        string                         `json:"uri"`
	Name       string                         `json:"name"`
	Cover      DetailPageImage                `json:"cover"`
	Year       int                            `json:"year"`
	Month      int                            `json:"month"`
	Day        int                            `json:"day"`
	TrackCount int                            `json:"track_count"`
	Discs      []DetailPageGenericReleaseDisc `json:"discs,omitempty"`
}

type DetailPageArtistReleasesTypeContainer struct {
	Releases   []DetailPageGenericRelease `json:"releases"`
	TotalCount int                        `json:"total_count"`
}
type DetailPageArtistReleases struct {
	Albums       DetailPageArtistReleasesTypeContainer `json:"albums"`
	Singles      DetailPageArtistReleasesTypeContainer `json:"singles"`
	AppearsOn    DetailPageArtistReleasesTypeContainer `json:"appears_on"`
	Compilations DetailPageArtistReleasesTypeContainer `json:"compilations"`
}
type DetailPageArtistMerchItems struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Link        string `json:"link"`
	ImageUri    string `json:"image_uri"`
	Price       string `json:"price"`
	UUID        string `json:"uuid"`
}
type DetailPageArtistMerch struct {
	Items []DetailPageArtistMerchItems `json:"items"`
}

type DetailPageArtistGallery struct {
	Images []DetailPageImage `json:"images"`
}

type DetailPageArtistPublishedPlaylist struct {
	Uri           string          `json:"uri"`
	Name          string          `json:"name"`
	Cover         DetailPageImage `json:"cover"`
	FollowerCount int             `json:"follower_count"`
}
type DetailPageArtistPublishedPlaylists struct {
	Playlists []DetailPageArtistPublishedPlaylist `json:"playlists"`
}

type DetailPageArtistMonthlyListeners struct {
	ListenerCount int `json:"listener_count"`
}
type DetailPageArtistCreatorAbout struct {
	MonthlyListeners    int `json:"monthlyListeners"`
	GlobalChartPosition int `json:"globalChartPosition"`
}
type DetailPageArtist struct {
	Uri                string                             `json:"uri"`
	Info               DetailPageArtistInfo               `json:"info"`
	HeaderImage        DetailPageArtistHeaderImage        `json:"header_image"`
	TopTracks          DetailPageArtistTopTracks          `json:"top_tracks"`
	UpcomingConcerts   DetailPageArtistUpcomingConcerts   `json:"upcoming_concerts"`
	RelatedArtists     DetailPageArtistRelatedArtists     `json:"related_artists"`
	Biography          DetailPageArtistBiography          `json:"biography"`
	Releases           DetailPageArtistReleases           `json:"releases"`
	Merch              DetailPageArtistMerch              `json:"merch"`
	Gallery            DetailPageArtistGallery            `json:"gallery"`
	LatestRelease      DetailPageGenericRelease           `json:"latest_release"`
	PublishedPlaylists DetailPageArtistPublishedPlaylists `json:"published_playlists"`
	MonthlyListeners   DetailPageArtistMonthlyListeners   `json:"monthly_listeners"`
	CreatorAbout       DetailPageArtistCreatorAbout       `json:"creator_about"`
}
