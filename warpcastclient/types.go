package warpcastclient

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

// SuggestedUsersResponse is the response from the Farcaster API v2 for the suggested users endpoint.
type SuggestedUsersResponse struct {
	Result struct {
		Users []struct {
			FID         int64  `json:"fid"`
			Username    string `json:"username"`
			DisplayName string `json:"displayName"`
			PFP         struct {
				URL      string `json:"url"`
				Verified bool   `json:"verified"`
			} `json:"pfp"`
			Profile struct {
				Bio struct {
					Text            string   `json:"text"`
					Mentions        []string `json:"mentions"`
					ChannelMentions []string `json:"channelMentions"`
				} `json:"bio"`
				Location struct {
					PlaceID     string `json:"placeId"`
					Description string `json:"description"`
				} `json:"location"`
			} `json:"profile"`
			FollowerCount     int64  `json:"followerCount"`
			FollowingCount    int64  `json:"followingCount"`
			ActiveOnFcNetwork bool   `json:"activeOnFcNetwork"`
			ReferrerUsername  string `json:"referrerUsername,omitempty"`
			ViewerContext     struct {
				Following           bool `json:"following"`
				FollowedBy          bool `json:"followedBy"`
				EnableNotifications bool `json:"enableNotifications"`
			} `json:"viewerContext"`
		} `json:"users"`
	} `json:"result"`
	Next struct {
		Cursor string `json:"cursor"`
	} `json:"next"`
}
