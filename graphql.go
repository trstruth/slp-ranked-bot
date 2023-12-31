package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const query string = `{"operationName":"AccountManagementPageQuery","variables":{"cc":"%s","uid":"%s"},"query":"fragment profileFields on NetplayProfile {\n  id\n  ratingOrdinal\n  ratingUpdateCount\n  wins\n  losses\n  dailyGlobalPlacement\n  dailyRegionalPlacement\n  continent\n  characters {\n    id\n    character\n    gameCount\n    __typename\n  }\n  __typename\n}\n\nfragment userProfilePage on User {\n  fbUid\n  displayName\n  connectCode {\n    code\n    __typename\n  }\n  status\n  activeSubscription {\n    level\n    hasGiftSub\n    __typename\n  }\n  rankedNetplayProfile {\n    ...profileFields\n    __typename\n  }\n  netplayProfiles {\n    ...profileFields\n    season {\n      id\n      startedAt\n      endedAt\n      name\n      status\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nquery AccountManagementPageQuery($cc: String!, $uid: String!) {\n  getUser(fbUid: $uid) {\n    ...userProfilePage\n    __typename\n  }\n  getConnectCode(code: $cc) {\n    user {\n      ...userProfilePage\n      __typename\n    }\n    __typename\n  }\n}\n"}`

type GraphQLResponse struct {
	Data struct {
		GetConnectCode struct {
			User struct {
				DisplayName string `json:"displayName"`
				ConnectCode struct {
					Code string `json:"code"`
				}
				RankedNetplayProfile struct {
					RatingOrdinal float32 `json:"ratingOrdinal"`
				} `json:"rankedNetplayProfile"`
			} `json:"user"`
		} `json:"getConnectCode"`
	} `json:"data"`
}

type UserData struct {
	DisplayName   string
	ConnectCode   string
	RatingOrdinal float32
}

// fetchUserData returns a UserData struct containing the fields queried from slippi
// note that the caller must provide the connect code in the following format: `BIG#9`
// the letters in the connectCode MUST be capitalized.
func fetchUserData(connectCode string) (*UserData, error) {
	// Schema
	payload := []byte(fmt.Sprintf(query, connectCode, connectCode))
	url := "https://gql-gateway-dot-slippi.uc.r.appspot.com/graphql"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %s", err)
	}

	// Set request headers

	// Add any other headers you may need from the example
	// ...
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Referer", "https://slippi.gg")
	req.Header.Set("Origin", "https://slippi.gg")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code:", resp.StatusCode)
		// Handle error as needed
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("Error reading response body: %s", err)
		}
		return nil, fmt.Errorf("Response Body: %s", string(responseBody))
	}

	// Assuming 'resp' is the HTTP response
	var graphQLResponse GraphQLResponse

	// Parse the JSON response
	err = json.NewDecoder(resp.Body).Decode(&graphQLResponse)
	if err != nil {
		return nil, fmt.Errorf("Error decoding JSON: %s", err)
	}

	displayName := graphQLResponse.Data.GetConnectCode.User.ConnectCode.Code
	ratingOrdinal := graphQLResponse.Data.GetConnectCode.User.RankedNetplayProfile.RatingOrdinal

	return &UserData{
		ConnectCode:   connectCode,
		DisplayName:   displayName,
		RatingOrdinal: ratingOrdinal,
	}, nil
}
