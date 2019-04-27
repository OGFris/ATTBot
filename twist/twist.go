package twist

type AnimesFeed struct {
	Type        string `json:"-"`
	Title       string `json:"-"`
	Description string `json:"-"`
	Link        string `json:"-"`
	Image       struct {
		Url   string `json:"url"`
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"-"`
	Generator      string           `json:"-"`
	Lastbuilddate  string           `json:"-"`
	AtomLink       struct{}         `json:"-"`
	Language       string           `json:"-"`
	Managingeditor string           `json:"-"`
	Webmaster      string           `json:"-"`
	Docs           string           `json:"-"`
	Ttl            int              `json:"-"`
	Items          []AnimesFeedItem `json:"items"`
}

type AnimesFeedItem struct {
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Guid           struct{} `json:"-"`
	Pubdate        string   `json:"pubdate"`
	AnimeTitle     string   `json:"anime:title"`
	AnimeOngoing   int      `json:"anime:ongoing"`
	AnimetwistSlug string   `json:"animetwist:slug"`
	AnimetwistID   int      `json:"animetwist:id"`
	KitsuID        int      `json:"kitsu:id"`
	MalID          int      `json:"mal:id"`
}

type EpisodesFeed struct {
	Type        string `json:"-"`
	Title       string `json:"-"`
	Description string `json:"-"`
	Link        string `json:"-"`
	Image       struct {
		Url   string `json:"url"`
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"-"`
	Generator      string             `json:"-"`
	Lastbuilddate  string             `json:"lastbuilddate"`
	AtomLink       struct{}           `json:"-"`
	Language       string             `json:"-"`
	Managingeditor string             `json:"-"`
	Webmaster      string             `json:"-"`
	Docs           string             `json:"-"`
	Ttl            int                `json:"-"`
	Items          []EpisodesFeedItem `json:"items"`
}

type EpisodesFeedItem struct {
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Guid          struct{} `json:"-"`
	Pubdate       string   `json:"pubdate"`
	AnimeTitle    string   `json:"anime:title"`
	EpisodeNumber int      `json:"episode:number"`
	AnimetwistID  int      `json:"animetwist:id"`
	KitsuID       int      `json:"kitsu:id"`
	MalID         int      `json:"mal:id"`
}
