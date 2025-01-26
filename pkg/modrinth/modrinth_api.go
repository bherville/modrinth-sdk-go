package modrinth

import (
	"time"
)

type ModrinthServer struct {
	ApiKey string `json:"apiKey"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}

type ApiServerInformation struct {
	About         string `json:"about"`
	Documentation string `json:"documentation"`
	Name          string `json:"name"`
	Version       string `json:"version"`
}

type ApiError struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}

type Hit struct {
	Title             string    `json:"title"`
	Slug              string    `json:"slug"`
	ProjectId         string    `json:"project_id"`
	Author            string    `json:"author"`
	DisplayCategories []string  `json:"display_categories,omitempty"`
	Versions          []string  `json:"versions"`
	Followers         int       `json:"followers"`
	DateCreated       time.Time `json:"data_created"`
	DateModified      time.Time `json:"date_modified"`
	LatestVersion     string    `json:"latest_version,omitempty"`
	License           string    `json:"license"`
	Gallery           []string  `json:"gallery,omitempty"`
	FeaturedGallery   string    `json:"featured_gallery,omitempty"`
	ClientSide        string    `json:"client_side,omitempty"`
	ServerSide        string    `json:"server_side,omitempty"`
	Color             int       `json:"color,omitempty"`
}

type SearchQuery struct {
	Query  string     `json:"query"`
	Facets [][]string `json:"facets,omitempty"`
}

type SearchResult struct {
	Hits      []Hit `json:"hits"`
	Offset    int   `json:"offset"`
	Limit     int   `json:"limit"`
	TotalHits int   `json:"total_hits"`
}

type ProjectVersionDependency struct {
	VersionId      string `json:"version_id,omitempty"`
	ProjectId      string `json:"project_id,omitempty"`
	FileName       string `json:"file_name,omitempty"`
	DependencyType string `json:"dependency_type"`
}

type Project struct {
	ClientSide           string    `json:"client_side"`
	ServerSide           string    `json:"server_side"`
	GameVersions         []string  `json:"game_versions"`
	ID                   string    `json:"id"`
	Slug                 string    `json:"slug"`
	ProjectType          string    `json:"project_type"`
	Team                 string    `json:"team"`
	Organization         any       `json:"organization"`
	Title                string    `json:"title"`
	Description          string    `json:"description"`
	Body                 string    `json:"body"`
	BodyURL              any       `json:"body_url"`
	Published            time.Time `json:"published"`
	Updated              time.Time `json:"updated"`
	Approved             time.Time `json:"approved"`
	Queued               any       `json:"queued"`
	Status               string    `json:"status"`
	RequestedStatus      any       `json:"requested_status"`
	ModeratorMessage     any       `json:"moderator_message"`
	License              License   `json:"license"`
	Downloads            int       `json:"downloads"`
	Followers            int       `json:"followers"`
	Categories           []string  `json:"categories"`
	AdditionalCategories []any     `json:"additional_categories"`
	Loaders              []string  `json:"loaders"`
	Versions             []string  `json:"versions"`
	IconURL              string    `json:"icon_url"`
	IssuesURL            any       `json:"issues_url"`
	SourceURL            any       `json:"source_url"`
	WikiURL              any       `json:"wiki_url"`
	DiscordURL           string    `json:"discord_url"`
	DonationUrls         []any     `json:"donation_urls"`
	Gallery              []Gallery `json:"gallery"`
	Color                int       `json:"color"`
	ThreadID             string    `json:"thread_id"`
	MonetizationStatus   string    `json:"monetization_status"`
}
type License struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  any    `json:"url"`
}
type Gallery struct {
	URL         string    `json:"url"`
	RawURL      string    `json:"raw_url"`
	Featured    bool      `json:"featured"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	Ordering    int       `json:"ordering"`
}

type ProjectVersion struct {
	GameVersions    []string             `json:"game_versions"`
	Loaders         []string             `json:"loaders"`
	ID              string               `json:"id"`
	ProjectID       string               `json:"project_id"`
	AuthorID        string               `json:"author_id"`
	Featured        bool                 `json:"featured"`
	Name            string               `json:"name"`
	VersionNumber   string               `json:"version_number"`
	Changelog       string               `json:"changelog"`
	ChangelogURL    any                  `json:"changelog_url"`
	DatePublished   time.Time            `json:"date_published"`
	Downloads       int                  `json:"downloads"`
	VersionType     string               `json:"version_type"`
	Status          string               `json:"status"`
	RequestedStatus any                  `json:"requested_status"`
	Files           []ProjectVersionFile `json:"files"`
	Dependencies    []any                `json:"dependencies"`
}
type ProjectVersionFileHash struct {
	Sha512 string `json:"sha512"`
	Sha1   string `json:"sha1"`
}
type ProjectVersionFile struct {
	Hashes   ProjectVersionFileHash `json:"hashes"`
	URL      string                 `json:"url"`
	Filename string                 `json:"filename"`
	Primary  bool                   `json:"primary"`
	Size     int                    `json:"size"`
	FileType any                    `json:"file_type"`
}
