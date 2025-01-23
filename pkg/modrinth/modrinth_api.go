package modrinth

import (
	"time"
)

type ModrinthServer struct {
	ApiKey string `json:"apiKey"`
	Name   string `json:"name"`
	Url    string `json:"url"`
}

type License struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type Gallery struct {
	Url         string    `json:"url"`
	Featured    bool      `json:"featured"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Created     time.Time `json:"created"`
	Ordering    int       `json:"ordering"`
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

type Project struct {
	Id               string    `json:"id"`
	Team             string    `json:"team"`
	BodyUrl          string    `json:"body_url,omitempty"`
	ModeratorMessage string    `json:"moderator_message,omitempty"`
	Published        time.Time `json:"published"`
	Updated          time.Time `json:"updated"`
	Approved         time.Time `json:"approved,omitempty"`
	Queued           time.Time `json:"queued,omitempty"`
	Followers        int       `json:"followers"`
	License          License   `json:"license"`
	Versions         []string  `json:"versions,omitempty"`
	GameVersions     []string  `json:"game_versions,omitempty"`
	Loaders          []string  `json:"loaders,omitempty"`
	Gallery          []Gallery `json:"gallery"`
}

type Hit struct {
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

type ProjectVersion struct {
	Name            string                     `json:"name,omitempty"`
	VersionNumber   string                     `json:"version_number,omitempty"`
	Changelog       string                     `json:"changelog,omitempty"`
	Dependencies    []ProjectVersionDependency `json:"dependencies,omitempty"`
	GameVersions    []string                   `json:"game_versions,omitempty"`
	VersionType     string                     `json:"version_type,omitempty"`
	Loaders         []string                   `json:"loaders,omitempty"`
	Featured        bool                       `json:"featured,omitempty"`
	Status          string                     `json:"status,omitempty"`
	RequestedStatus string                     `json:"requested_status,omitempty"`
	Id              string                     `json:"id"`
	ProjectId       string                     `json:"project_id"`
	AuthorId        string                     `json:"author_id"`
	DatePublished   time.Time                  `json:"date_published"`
	Downloads       int                        `json:"downloads"`
	ChangelogUrl    string                     `json:"changelog_url,omitempty"`
	Files           []ProjectVersionFile       `json:"files"`
}

type ProjectVersionDependency struct {
	VersionId      string `json:"version_id,omitempty"`
	ProjectId      string `json:"project_id,omitempty"`
	FileName       string `json:"file_name,omitempty"`
	DependencyType string `json:"dependency_type"`
}

type ProjectVersionFile struct {
	Hashes   []ProjectVersionFileHashes `json:"hashes"`
	Url      string                     `json:"url"`
	Filename string                     `json:"filename"`
	Primary  bool                       `json:"primary"`
	Size     int                        `json:"size"`
	FileType string                     `json:"file_type,omitempty"`
}

type ProjectVersionFileHashes struct {
	Sha512 string `json:"sha512"`
	Sha1   string `json:"sha1"`
}
