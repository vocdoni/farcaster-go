package warpcastapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

const (
	farcasterV2APIuser          = "https://client.warpcast.com/v2/user?fid=%d"
	farcasterV2APIverifications = "https://client.warpcast.com/v2/verifications?fid=%d&limit=100"
	farcasterV2APIrecentUsers   = "https://api.warpcast.com/v2/recent-users?filter=off&limit=%d"
	updatedUsersByIteration     = 200
	protocolEthereum            = "ethereum"
	userAgent                   = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
)

// UserProfile represents the user profile from the Farcaster API v2.
type UserProfile struct {
	Result struct {
		User struct {
			Fid         uint64 `json:"fid"`
			Username    string `json:"username"`
			DisplayName string `json:"displayName"`
			Pfp         struct {
				Url      string `json:"url"`
				Verified bool   `json:"verified"`
			} `json:"pfp"`
			Profile struct {
				Bio struct {
					Text            string   `json:"text"`
					Mentions        []string `json:"mentions"`
					ChannelMentions []string `json:"channelMentions"`
				} `json:"bio"`
				Location struct {
					PlaceId     string `json:"placeId"`
					Description string `json:"description"`
				} `json:"location"`
			} `json:"profile"`
			FollowerCount     int  `json:"followerCount"`
			FollowingCount    int  `json:"followingCount"`
			ActiveOnFcNetwork bool `json:"activeOnFcNetwork"`
			ViewerContext     struct {
				Following            bool `json:"following"`
				FollowedBy           bool `json:"followedBy"`
				CanSendDirectCasts   bool `json:"canSendDirectCasts"`
				HasUploadedInboxKeys bool `json:"hasUploadedInboxKeys"`
			} `json:"viewerContext"`
		} `json:"user"`
		InviterIsReferrer bool          `json:"inviterIsReferrer"`
		CollectionsOwned  []interface{} `json:"collectionsOwned"`
		Extras            struct {
			Fid            uint64 `json:"fid"`
			CustodyAddress string `json:"custodyAddress"`
		} `json:"extras"`
	} `json:"result"`
}

// VerificationResponse is the response from the Farcaster API v2 for the verifications endpoint.
type VerificationResponse struct {
	Result struct {
		Verifications []struct {
			FID       int    `json:"fid"`
			Address   string `json:"address"`
			Timestamp int64  `json:"timestamp"`
			Version   string `json:"version"`
			Protocol  string `json:"protocol"`
		} `json:"verifications"`
	} `json:"result"`
}

// UserProfileByFID returns the user profile from the Farcaster API v2.
func UserProfileByFID(fid uint64) (*UserProfile, error) {
	var profile *UserProfile
	// Create a new HTTP request
	req, err := http.NewRequest("GET", fmt.Sprintf(farcasterV2APIuser, fid), nil)
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
	req, err := http.NewRequest("GET", fmt.Sprintf(farcasterV2APIverifications, fid), nil)
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
	req, err := http.NewRequest("GET", fmt.Sprintf(farcasterV2APIrecentUsers, 1), nil)
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
