package modrinth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	WaitForBackupSeconds int64 = 5
	contextLogger        *log.Entry
)

type ApiEnvironment string

const (
	ApiEndpointApiServerInfo string = ""
	ApiEndpointProject       string = "v2/project"
	ApiEndpointSearch        string = "v2/search"
	ApiEndpointVersion       string = "version"
	ApiEndpointVersionFile   string = "v2/version_file"

	EnvironmentProduction ApiEnvironment = "Production"
	EnvironmentStaging    ApiEnvironment = "Staging"
)

func init() {
	contextLogger = log.WithFields(log.Fields{
		"library": "modrinth-sdk-go",
	})
}

func NewServer(apiEnvironment ApiEnvironment) ModrinthServer {
	switch apiEnvironment {
	case EnvironmentStaging:
		return ModrinthServer{
			Name: string(EnvironmentStaging),
			Url:  "https://staging-api.modrinth.com",
		}
	default:
		return ModrinthServer{
			Name: string(EnvironmentProduction),
			Url:  "https://api.modrinth.com",
		}
	}
}

func buildApiUrl(modrinthServer ModrinthServer, endpoint string, subPaths []string) string {
	url := fmt.Sprintf("%s/%s", modrinthServer.Url, endpoint)

	for _, path := range subPaths {
		url = fmt.Sprintf("%s/%s", url, path)
	}
	return url
}

func callApi[T any](apiObject *T, modrinthServer ModrinthServer, method string, endpoint string, subPaths []string, data map[string]string) error {
	apiUrl := buildApiUrl(modrinthServer, endpoint, subPaths)

	dataToSend := url.Values{}

	for k, v := range data {
		dataToSend.Set(k, v)
	}

	if data != nil {
		switch method {
		case http.MethodGet:
			apiUrl = fmt.Sprintf("%s?%s", apiUrl, dataToSend.Encode())
		}
	}

	contextLogger.Trace(fmt.Sprintf("apiUrl: %s", apiUrl))

	req, _ := http.NewRequest(method, apiUrl, nil)
	req.Header.Add("Accept", "application/json")
	// req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", modrinthServer.ApiKey))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	log.Trace(fmt.Sprintf("Response Body: %s", string(body)))

	if res.StatusCode != http.StatusOK {
		var apiError ApiError

		err = json.Unmarshal(body, &apiError)
		if err != nil {
			return err
		}

		return fmt.Errorf("api call failed with errors: %s", apiError)
	}

	err = json.Unmarshal(body, &apiObject)
	return err
}

func GetApiServerInfo(modrinthServer ModrinthServer) (*ApiServerInformation, error) {
	var serverInfo ApiServerInformation

	err := callApi(&serverInfo, modrinthServer, http.MethodGet, ApiEndpointApiServerInfo, nil, nil)
	if err != nil {
		return nil, err
	}

	return &serverInfo, nil
}

func Search(modrinthServer ModrinthServer, searchQuery SearchQuery) (*SearchResult, error) {
	var searchResult SearchResult

	encjson, err := json.Marshal(searchQuery.Facets)
	if err != nil {
		return nil, err
	}

	queryMap := map[string]string{
		"query": searchQuery.Query,
	}

	if searchQuery.Facets != nil {
		queryMap["facets"] = string(encjson)
	}

	err = callApi(&searchResult, modrinthServer, http.MethodGet, ApiEndpointSearch, nil, queryMap)
	if err != nil {
		return nil, err
	}

	return &searchResult, nil
}

func GetProject(modrinthServer ModrinthServer, projectName string) (*Project, error) {
	var project Project
	err := callApi(&project, modrinthServer, http.MethodGet, ApiEndpointProject, []string{projectName}, nil)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func GetProjectVersions(modrinthServer ModrinthServer, projectName string, loaders []string, gameVersions []string) ([]ProjectVersion, error) {
	var projectVersions []ProjectVersion
	params := make(map[string]string)

	if loaders != nil {
		loadersJson, err := json.Marshal(loaders)
		if err != nil {
			return nil, err
		}
		params["loaders"] = string(loadersJson)
	}
	if gameVersions != nil {
		gameVersionsJson, err := json.Marshal(gameVersions)
		if err != nil {
			return nil, err
		}
		params["game_versions"] = string(gameVersionsJson)
	}

	err := callApi(&projectVersions, modrinthServer, http.MethodGet, ApiEndpointProject, []string{projectName, ApiEndpointVersion}, params)
	if err != nil {
		return nil, err
	}

	return projectVersions, nil
}

func GetProjectVersion(modrinthServer ModrinthServer, project Project, versionIdOrNumber string) (ProjectVersion, error) {
	var projectVersion ProjectVersion
	err := callApi(&projectVersion, modrinthServer, http.MethodGet, ApiEndpointProject, []string{project.ID, "version"}, nil)
	if err != nil {
		return projectVersion, err
	}

	return projectVersion, nil
}

func GetProjectDependencies(modrinthServer ModrinthServer, projectId string) (ProjectDependencies, error) {
	var projectVersionDependencies ProjectDependencies
	err := callApi(&projectVersionDependencies, modrinthServer, http.MethodGet, ApiEndpointProject, []string{projectId, "dependencies"}, nil)
	if err != nil {
		return projectVersionDependencies, err
	}

	return projectVersionDependencies, nil
}

func GetProjectVersionFromHash(modrinthServer ModrinthServer, versionHash string, versionHashAlgo string) (ProjectVersion, error) {
	var projectVersion ProjectVersion
	err := callApi(&projectVersion, modrinthServer, http.MethodGet, ApiEndpointVersionFile, []string{versionHash}, map[string]string{
		"algorithm": versionHashAlgo,
	})
	if err != nil {
		return projectVersion, err
	}

	return projectVersion, nil
}

// GetVersionByID retrieves a specific version by its ID
func GetVersionByID(modrinthServer ModrinthServer, versionID string) (ProjectVersion, error) {
	var projectVersion ProjectVersion
	err := callApi(&projectVersion, modrinthServer, http.MethodGet, "v2/version", []string{versionID}, nil)
	if err != nil {
		return projectVersion, err
	}

	return projectVersion, nil
}

// GetMultipleVersions retrieves multiple versions by their IDs
func GetMultipleVersions(modrinthServer ModrinthServer, versionIDs []string) ([]ProjectVersion, error) {
	var projectVersions []ProjectVersion

	idsJson, err := json.Marshal(versionIDs)
	if err != nil {
		return nil, err
	}

	err = callApi(&projectVersions, modrinthServer, http.MethodGet, "v2/versions", nil, map[string]string{
		"ids": string(idsJson),
	})
	if err != nil {
		return nil, err
	}

	return projectVersions, nil
}

func DownloadProjectVersion(modrinthServer ModrinthServer, projectVersionFile ProjectVersionFile, destination string) error {
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(projectVersionFile.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
