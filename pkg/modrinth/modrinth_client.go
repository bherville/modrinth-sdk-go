package modrinth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var (
	WaitForBackupSeconds int64 = 5
)

type ApiEnvironment string

const (
	ApiEndpointApiServerInfo string = ""
	ApiEndpointProject       string = "v2/project"
	ApiEndpointSearch        string = "v2/search"
	ApiEndpointVersion       string = "version"

	EnvironmentProduction ApiEnvironment = "Production"
	EnvironmentStaging    ApiEnvironment = "Staging"
)

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
	println(fmt.Sprintf("Attempting to call URL: %s", apiUrl))

	dataToSend := url.Values{}
	// var proccessedDataToSend strings.Reader

	for k, v := range data {
		dataToSend.Set(k, v)
	}

	if data != nil {
		switch method {
		case http.MethodGet:
			apiUrl = fmt.Sprintf("%s?%s", apiUrl, dataToSend.Encode())
			// case http.MethodPost:
			// 	proccessedDataToSend = *strings.NewReader(dataToSend.Encode())
		}
	}

	println(fmt.Sprintf("apiUrl: %s", apiUrl))
	// os.Exit(0)

	req, _ := http.NewRequest(method, apiUrl, nil)
	req.Header.Add("Accept", "application/json")
	// req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", modrinthServer.ApiKey))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	println(string(body))

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

	encjson, _ := json.Marshal(searchQuery.Facets)
	fmt.Println(string(encjson))

	queryMap := map[string]string{
		"query": searchQuery.Query,
	}

	if searchQuery.Facets != nil {
		queryMap["facets"] = string(encjson)
	}

	err := callApi(&searchResult, modrinthServer, http.MethodGet, ApiEndpointSearch, nil, queryMap)
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

func GetProjectVersions(modrinthServer ModrinthServer, projectName string) (*[]ProjectVersion, error) {
	var projectVersions []ProjectVersion
	err := callApi(&projectVersions, modrinthServer, http.MethodGet, ApiEndpointProject, []string{projectName, ApiEndpointVersion}, nil)
	if err != nil {
		return nil, err
	}

	return &projectVersions, nil
}
