package warpcastclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

const (
	userEndpoint           = "https://client.warpcast.com/v2/user?fid=%d"
	verificationsEndpoint  = "https://client.warpcast.com/v2/verifications?fid=%d&limit=100"
	recentUsersEndpoint    = "https://api.warpcast.com/v2/recent-users?filter=off&limit=%d"
	suggestedUsersEndpoint = "https://client.warpcast.com/v2/suggested-users?limit=10&randomized=true"

	updatedUsersByIteration = 200
	protocolEthereum        = "ethereum"
	userAgent               = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
)

// https://client.warpcast.com/v2/discover-channels?limit=10"
// https://client.warpcast.com/v2/user-by-username?username=p4u

// UserProfileByFID returns the user profile from the Farcaster API v2.
func UserProfileByFID(fid uint64) (*UserProfile, error) {
	var profile *UserProfile
	// Create a new HTTP request
	req, err := http.NewRequest("GET", fmt.Sprintf(userEndpoint, fid), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	// Set a custom user-agent
	req.Header.Set("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user profile: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read user profile: %w", err)
	}
	if err := json.Unmarshal(data, &profile); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user profile: %w", err)
	}
	return profile, nil
}

// AddressesByFID returns the verified Ethereum addresses from the Warpcast API.
func AddressesByFID(fid uint64) ([]string, error) {
	var verifications *VerificationResponse
	// Create a new HTTP request
	req, err := http.NewRequest("GET", fmt.Sprintf(verificationsEndpoint, fid), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	// Set a custom user-agent
	req.Header.Set("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get user verifications: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read user verifications: %w", err)
	}
	if err := json.Unmarshal(data, &verifications); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user verifications: %w", err)
	}
	var addresses []string
	for _, v := range verifications.Result.Verifications {
		if v.Protocol == protocolEthereum {
			addresses = append(addresses, common.HexToAddress(v.Address).Hex())
		}
	}
	return addresses, nil
}

// LastRegisteredFID returns the last registered FID from the Warpcast API.
func LastRegisteredFID() (uint64, error) {
	var recentUsers struct {
		Result struct {
			Users []struct {
				Fid uint64 `json:"fid"`
			} `json:"users"`
		} `json:"result"`
	}
	// Create a new HTTP request
	req, err := http.NewRequest("GET", fmt.Sprintf(recentUsersEndpoint, 1), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}
	// Set a custom user-agent
	req.Header.Set("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to get recent users: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to read recent users: %w", err)
	}
	if err := json.Unmarshal(data, &recentUsers); err != nil {
		return 0, fmt.Errorf("failed to unmarshal recent users: %w", err)
	}
	if len(recentUsers.Result.Users) == 0 {
		return 0, errors.New("no recent users")
	}
	return recentUsers.Result.Users[0].Fid, nil
}

func SuggestedUsers() (*SuggestedUsersResponse, error) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", suggestedUsersEndpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	// Set a custom user-agent
	req.Header.Set("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get suggested users: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read suggested users: %w", err)
	}
	var response SuggestedUsersResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal suggested users: %w", err)
	}
	return &response, nil
}
