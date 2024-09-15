package vrchat

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/samber/lo"
)

type Client struct {
	client *resty.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		client: resty.New().SetBaseURL(baseURL),
	}
}

// CheckUserExistsParams represents the parameters for the CheckUserExists request
type CheckUserExistsParams struct {
	Email         string `json:"email"`
	DisplayName   string `json:"displayName"`
	Username      string `json:"username"`
	ExcludeUserId string `json:"excludeUserId"`
}

func (c *Client) CheckUserExists(params CheckUserExistsParams) (*UserExistsResponse, error) {
	path := "/auth/exists"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Email) {
		queryParams["email"] = fmt.Sprintf("%v", params.Email)
	}
	if lo.IsNotEmpty(params.DisplayName) {
		queryParams["displayName"] = fmt.Sprintf("%v", params.DisplayName)
	}
	if lo.IsNotEmpty(params.Username) {
		queryParams["username"] = fmt.Sprintf("%v", params.Username)
	}
	if lo.IsNotEmpty(params.ExcludeUserId) {
		queryParams["excludeUserId"] = fmt.Sprintf("%v", params.ExcludeUserId)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result UserExistsResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetCurrentUser() (*CurrentUserLoginResponse, error) {
	path := "/auth/user"

	// Create request
	req := c.client.R()
	// Set response object
	var result CurrentUserLoginResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) Verify2Fa() (*Verify2FaResponse, error) {
	path := "/auth/twofactorauth/totp/verify"

	// Create request
	req := c.client.R()
	// Set response object
	var result Verify2FaResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) VerifyRecoveryCode() (*Verify2FaResponse, error) {
	path := "/auth/twofactorauth/otp/verify"

	// Create request
	req := c.client.R()
	// Set response object
	var result Verify2FaResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) Verify2FaEmailCode() (*Verify2FaEmailCodeResponse, error) {
	path := "/auth/twofactorauth/emailotp/verify"

	// Create request
	req := c.client.R()
	// Set response object
	var result Verify2FaEmailCodeResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) VerifyAuthToken() (*VerifyAuthTokenResponse, error) {
	path := "/auth"

	// Create request
	req := c.client.R()
	// Set response object
	var result VerifyAuthTokenResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) Logout() (*LogoutSuccess, error) {
	path := "/logout"

	// Create request
	req := c.client.R()
	// Set response object
	var result LogoutSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteUserParams represents the parameters for the DeleteUser request
type DeleteUserParams struct {
	UserId string `json:"userId"`
}

func (c *Client) DeleteUser(params DeleteUserParams) (*DeleteUserResponse, error) {
	path := "/users/{userId}/delete"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result DeleteUserResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetOwnAvatarParams represents the parameters for the GetOwnAvatar request
type GetOwnAvatarParams struct {
	UserId string `json:"userId"`
}

func (c *Client) GetOwnAvatar(params GetOwnAvatarParams) (*AvatarResponse, error) {
	path := "/users/{userId}/avatar"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result AvatarResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) CreateAvatar() (*AvatarResponse, error) {
	path := "/avatars"

	// Create request
	req := c.client.R()
	// Set response object
	var result AvatarResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// SearchAvatarsParams represents the parameters for the SearchAvatars request
type SearchAvatarsParams struct {
	Featured        bool          `json:"featured"`
	Sort            SortOption    `json:"sort"`
	UserId          string        `json:"userId"`
	N               int64         `json:"n"`
	Order           OrderOption   `json:"order"`
	Offset          int64         `json:"offset"`
	Tag             string        `json:"tag"`
	Notag           string        `json:"notag"`
	ReleaseStatus   ReleaseStatus `json:"releaseStatus"`
	MaxUnityVersion string        `json:"maxUnityVersion"`
	MinUnityVersion string        `json:"minUnityVersion"`
	Platform        string        `json:"platform"`
}

func (c *Client) SearchAvatars(params SearchAvatarsParams) (*AvatarListResponse, error) {
	path := "/avatars"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Featured) {
		queryParams["featured"] = fmt.Sprintf("%v", params.Featured)
	}
	if lo.IsNotEmpty(params.Sort) {
		queryParams["sort"] = fmt.Sprintf("%v", params.Sort)
	}
	if lo.IsNotEmpty(params.UserId) {
		queryParams["userId"] = fmt.Sprintf("%v", params.UserId)
	}
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Order) {
		queryParams["order"] = fmt.Sprintf("%v", params.Order)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}
	if lo.IsNotEmpty(params.Tag) {
		queryParams["tag"] = fmt.Sprintf("%v", params.Tag)
	}
	if lo.IsNotEmpty(params.Notag) {
		queryParams["notag"] = fmt.Sprintf("%v", params.Notag)
	}
	if lo.IsNotEmpty(params.ReleaseStatus) {
		queryParams["releaseStatus"] = fmt.Sprintf("%v", params.ReleaseStatus)
	}
	if lo.IsNotEmpty(params.MaxUnityVersion) {
		queryParams["maxUnityVersion"] = fmt.Sprintf("%v", params.MaxUnityVersion)
	}
	if lo.IsNotEmpty(params.MinUnityVersion) {
		queryParams["minUnityVersion"] = fmt.Sprintf("%v", params.MinUnityVersion)
	}
	if lo.IsNotEmpty(params.Platform) {
		queryParams["platform"] = fmt.Sprintf("%v", params.Platform)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result AvatarListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteAvatarParams represents the parameters for the DeleteAvatar request
type DeleteAvatarParams struct {
	AvatarId string `json:"avatarId"`
}

func (c *Client) DeleteAvatar(params DeleteAvatarParams) (*AvatarResponse, error) {
	path := "/avatars/{avatarId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{avatarId}", fmt.Sprintf("%v", params.AvatarId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result AvatarResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetAvatarParams represents the parameters for the GetAvatar request
type GetAvatarParams struct {
	AvatarId string `json:"avatarId"`
}

func (c *Client) GetAvatar(params GetAvatarParams) (*AvatarResponse, error) {
	path := "/avatars/{avatarId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{avatarId}", fmt.Sprintf("%v", params.AvatarId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result AvatarResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UpdateAvatarParams represents the parameters for the UpdateAvatar request
type UpdateAvatarParams struct {
	AvatarId string `json:"avatarId"`
}

func (c *Client) UpdateAvatar(params UpdateAvatarParams) (*AvatarResponse, error) {
	path := "/avatars/{avatarId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{avatarId}", fmt.Sprintf("%v", params.AvatarId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result AvatarResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// SelectAvatarParams represents the parameters for the SelectAvatar request
type SelectAvatarParams struct {
	AvatarId string `json:"avatarId"`
}

func (c *Client) SelectAvatar(params SelectAvatarParams) (*CurrentUserResponse, error) {
	path := "/avatars/{avatarId}/select"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{avatarId}", fmt.Sprintf("%v", params.AvatarId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result CurrentUserResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// SelectFallbackAvatarParams represents the parameters for the SelectFallbackAvatar request
type SelectFallbackAvatarParams struct {
	AvatarId string `json:"avatarId"`
}

func (c *Client) SelectFallbackAvatar(params SelectFallbackAvatarParams) (*CurrentUserResponse, error) {
	path := "/avatars/{avatarId}/selectFallback"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{avatarId}", fmt.Sprintf("%v", params.AvatarId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result CurrentUserResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetFavoritedAvatarsParams represents the parameters for the GetFavoritedAvatars request
type GetFavoritedAvatarsParams struct {
	Featured        bool          `json:"featured"`
	Sort            SortOption    `json:"sort"`
	N               int64         `json:"n"`
	Order           OrderOption   `json:"order"`
	Offset          int64         `json:"offset"`
	Search          string        `json:"search"`
	Tag             string        `json:"tag"`
	Notag           string        `json:"notag"`
	ReleaseStatus   ReleaseStatus `json:"releaseStatus"`
	MaxUnityVersion string        `json:"maxUnityVersion"`
	MinUnityVersion string        `json:"minUnityVersion"`
	Platform        string        `json:"platform"`
	UserId          string        `json:"userId"`
}

func (c *Client) GetFavoritedAvatars(params GetFavoritedAvatarsParams) (*AvatarListResponse, error) {
	path := "/avatars/favorites"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Featured) {
		queryParams["featured"] = fmt.Sprintf("%v", params.Featured)
	}
	if lo.IsNotEmpty(params.Sort) {
		queryParams["sort"] = fmt.Sprintf("%v", params.Sort)
	}
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Order) {
		queryParams["order"] = fmt.Sprintf("%v", params.Order)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}
	if lo.IsNotEmpty(params.Search) {
		queryParams["search"] = fmt.Sprintf("%v", params.Search)
	}
	if lo.IsNotEmpty(params.Tag) {
		queryParams["tag"] = fmt.Sprintf("%v", params.Tag)
	}
	if lo.IsNotEmpty(params.Notag) {
		queryParams["notag"] = fmt.Sprintf("%v", params.Notag)
	}
	if lo.IsNotEmpty(params.ReleaseStatus) {
		queryParams["releaseStatus"] = fmt.Sprintf("%v", params.ReleaseStatus)
	}
	if lo.IsNotEmpty(params.MaxUnityVersion) {
		queryParams["maxUnityVersion"] = fmt.Sprintf("%v", params.MaxUnityVersion)
	}
	if lo.IsNotEmpty(params.MinUnityVersion) {
		queryParams["minUnityVersion"] = fmt.Sprintf("%v", params.MinUnityVersion)
	}
	if lo.IsNotEmpty(params.Platform) {
		queryParams["platform"] = fmt.Sprintf("%v", params.Platform)
	}
	if lo.IsNotEmpty(params.UserId) {
		queryParams["userId"] = fmt.Sprintf("%v", params.UserId)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result AvatarListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetSteamTransactions() (*TransactionListResponse, error) {
	path := "/Steam/transactions"

	// Create request
	req := c.client.R()
	// Set response object
	var result TransactionListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetSteamTransactionParams represents the parameters for the GetSteamTransaction request
type GetSteamTransactionParams struct {
	TransactionId string `json:"transactionId"`
}

func (c *Client) GetSteamTransaction(params GetSteamTransactionParams) (*TransactionResponse, error) {
	path := "/Steam/transactions/{transactionId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{transactionId}", fmt.Sprintf("%v", params.TransactionId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result TransactionResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetCurrentSubscriptions() (*UserSubscriptionListResponse, error) {
	path := "/auth/user/subscription"

	// Create request
	req := c.client.R()
	// Set response object
	var result UserSubscriptionListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetSubscriptions() (*SubscriptionListResponse, error) {
	path := "/subscriptions"

	// Create request
	req := c.client.R()
	// Set response object
	var result SubscriptionListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetLicenseGroupParams represents the parameters for the GetLicenseGroup request
type GetLicenseGroupParams struct {
	LicenseGroupId string `json:"licenseGroupId"`
}

func (c *Client) GetLicenseGroup(params GetLicenseGroupParams) (*LicenseGroupResponse, error) {
	path := "/licenseGroups/{licenseGroupId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{licenseGroupId}", fmt.Sprintf("%v", params.LicenseGroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result LicenseGroupResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetFavoritesParams represents the parameters for the GetFavorites request
type GetFavoritesParams struct {
	N      int64  `json:"n"`
	Offset int64  `json:"offset"`
	Tag    string `json:"tag"`
}

func (c *Client) GetFavorites(params GetFavoritesParams) (*FavoriteListResponse, error) {
	path := "/favorites"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}
	if lo.IsNotEmpty(params.Tag) {
		queryParams["tag"] = fmt.Sprintf("%v", params.Tag)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FavoriteListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) AddFavorite() (*FavoriteResponse, error) {
	path := "/favorites"

	// Create request
	req := c.client.R()
	// Set response object
	var result FavoriteResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// RemoveFavoriteParams represents the parameters for the RemoveFavorite request
type RemoveFavoriteParams struct {
	FavoriteId string `json:"favoriteId"`
}

func (c *Client) RemoveFavorite(params RemoveFavoriteParams) (*FavoriteRemovedSuccess, error) {
	path := "/favorites/{favoriteId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{favoriteId}", fmt.Sprintf("%v", params.FavoriteId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FavoriteRemovedSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetFavoriteParams represents the parameters for the GetFavorite request
type GetFavoriteParams struct {
	FavoriteId string `json:"favoriteId"`
}

func (c *Client) GetFavorite(params GetFavoriteParams) (*FavoriteResponse, error) {
	path := "/favorites/{favoriteId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{favoriteId}", fmt.Sprintf("%v", params.FavoriteId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FavoriteResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetFavoriteGroupsParams represents the parameters for the GetFavoriteGroups request
type GetFavoriteGroupsParams struct {
	N      int64 `json:"n"`
	Offset int64 `json:"offset"`
}

func (c *Client) GetFavoriteGroups(params GetFavoriteGroupsParams) (*FavoriteGroupListResponse, error) {
	path := "/favorite/groups"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FavoriteGroupListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// ClearFavoriteGroupParams represents the parameters for the ClearFavoriteGroup request
type ClearFavoriteGroupParams struct {
	// FavoriteGroupType enum
	FavoriteGroupType string `json:"favoriteGroupType"`
	FavoriteGroupName string `json:"favoriteGroupName"`
	UserId            string `json:"userId"`
}

func (c *Client) ClearFavoriteGroup(params ClearFavoriteGroupParams) (*FavoriteGroupClearedSuccess, error) {
	path := "/favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{favoriteGroupType}", fmt.Sprintf("%v", params.FavoriteGroupType))
	path = strings.ReplaceAll(path, "{favoriteGroupName}", fmt.Sprintf("%v", params.FavoriteGroupName))
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FavoriteGroupClearedSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetFavoriteGroupParams represents the parameters for the GetFavoriteGroup request
type GetFavoriteGroupParams struct {
	// FavoriteGroupType enum
	FavoriteGroupType string `json:"favoriteGroupType"`
	FavoriteGroupName string `json:"favoriteGroupName"`
	UserId            string `json:"userId"`
}

func (c *Client) GetFavoriteGroup(params GetFavoriteGroupParams) (*FavoriteGroupResponse, error) {
	path := "/favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{favoriteGroupType}", fmt.Sprintf("%v", params.FavoriteGroupType))
	path = strings.ReplaceAll(path, "{favoriteGroupName}", fmt.Sprintf("%v", params.FavoriteGroupName))
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FavoriteGroupResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UpdateFavoriteGroupParams represents the parameters for the UpdateFavoriteGroup request
type UpdateFavoriteGroupParams struct {
	// FavoriteGroupType enum
	FavoriteGroupType string `json:"favoriteGroupType"`
	FavoriteGroupName string `json:"favoriteGroupName"`
	UserId            string `json:"userId"`
}

func (c *Client) UpdateFavoriteGroup(params UpdateFavoriteGroupParams) error {
	path := "/favorite/group/{favoriteGroupType}/{favoriteGroupName}/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{favoriteGroupType}", fmt.Sprintf("%v", params.FavoriteGroupType))
	path = strings.ReplaceAll(path, "{favoriteGroupName}", fmt.Sprintf("%v", params.FavoriteGroupName))
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// GetFilesParams represents the parameters for the GetFiles request
type GetFilesParams struct {
	N      int64 `json:"n"`
	Offset int64 `json:"offset"`
}

func (c *Client) GetFiles(params GetFilesParams) (*FileListResponse, error) {
	path := "/files"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FileListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) CreateFile() (*FileResponse, error) {
	path := "/file"

	// Create request
	req := c.client.R()
	// Set response object
	var result FileResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetFileParams represents the parameters for the GetFile request
type GetFileParams struct {
	FileId string `json:"fileId"`
}

func (c *Client) GetFile(params GetFileParams) (*FileResponse, error) {
	path := "/file/{fileId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{fileId}", fmt.Sprintf("%v", params.FileId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FileResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// CreateFileVersionParams represents the parameters for the CreateFileVersion request
type CreateFileVersionParams struct {
	FileId string `json:"fileId"`
}

func (c *Client) CreateFileVersion(params CreateFileVersionParams) (*FileResponse, error) {
	path := "/file/{fileId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{fileId}", fmt.Sprintf("%v", params.FileId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FileResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteFileParams represents the parameters for the DeleteFile request
type DeleteFileParams struct {
	FileId string `json:"fileId"`
}

func (c *Client) DeleteFile(params DeleteFileParams) (*FileResponse, error) {
	path := "/file/{fileId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{fileId}", fmt.Sprintf("%v", params.FileId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FileResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteFileVersionParams represents the parameters for the DeleteFileVersion request
type DeleteFileVersionParams struct {
	FileId    string `json:"fileId"`
	VersionId int64  `json:"versionId"`
}

func (c *Client) DeleteFileVersion(params DeleteFileVersionParams) (*FileResponse, error) {
	path := "/file/{fileId}/{versionId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{fileId}", fmt.Sprintf("%v", params.FileId))
	path = strings.ReplaceAll(path, "{versionId}", fmt.Sprintf("%v", params.VersionId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FileResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DownloadFileVersionParams represents the parameters for the DownloadFileVersion request
type DownloadFileVersionParams struct {
	FileId    string `json:"fileId"`
	VersionId int64  `json:"versionId"`
}

func (c *Client) DownloadFileVersion(params DownloadFileVersionParams) (*RawFileResponse, error) {
	path := "/file/{fileId}/{versionId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{fileId}", fmt.Sprintf("%v", params.FileId))
	path = strings.ReplaceAll(path, "{versionId}", fmt.Sprintf("%v", params.VersionId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result RawFileResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// FinishFileDataUploadParams represents the parameters for the FinishFileDataUpload request
type FinishFileDataUploadParams struct {
	FileId    string `json:"fileId"`
	VersionId int64  `json:"versionId"`
	// FileType enum
	FileType string `json:"fileType"`
}

func (c *Client) FinishFileDataUpload(params FinishFileDataUploadParams) (*FileResponse, error) {
	path := "/file/{fileId}/{versionId}/{fileType}/finish"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{fileId}", fmt.Sprintf("%v", params.FileId))
	path = strings.ReplaceAll(path, "{versionId}", fmt.Sprintf("%v", params.VersionId))
	path = strings.ReplaceAll(path, "{fileType}", fmt.Sprintf("%v", params.FileType))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FileResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// StartFileDataUploadParams represents the parameters for the StartFileDataUpload request
type StartFileDataUploadParams struct {
	FileId    string `json:"fileId"`
	VersionId int64  `json:"versionId"`
	// FileType enum
	FileType string `json:"fileType"`
}

func (c *Client) StartFileDataUpload(params StartFileDataUploadParams) (*FileUploadUrlResponse, error) {
	path := "/file/{fileId}/{versionId}/{fileType}/start"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{fileId}", fmt.Sprintf("%v", params.FileId))
	path = strings.ReplaceAll(path, "{versionId}", fmt.Sprintf("%v", params.VersionId))
	path = strings.ReplaceAll(path, "{fileType}", fmt.Sprintf("%v", params.FileType))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FileUploadUrlResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetFileDataUploadStatusParams represents the parameters for the GetFileDataUploadStatus request
type GetFileDataUploadStatusParams struct {
	FileId    string `json:"fileId"`
	VersionId int64  `json:"versionId"`
	// FileType enum
	FileType string `json:"fileType"`
}

func (c *Client) GetFileDataUploadStatus(params GetFileDataUploadStatusParams) (*FileVersionUploadStatusResponse, error) {
	path := "/file/{fileId}/{versionId}/{fileType}/status"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{fileId}", fmt.Sprintf("%v", params.FileId))
	path = strings.ReplaceAll(path, "{versionId}", fmt.Sprintf("%v", params.VersionId))
	path = strings.ReplaceAll(path, "{fileType}", fmt.Sprintf("%v", params.FileType))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FileVersionUploadStatusResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetFriendsParams represents the parameters for the GetFriends request
type GetFriendsParams struct {
	Offset  int64 `json:"offset"`
	N       int64 `json:"n"`
	Offline bool  `json:"offline"`
}

func (c *Client) GetFriends(params GetFriendsParams) (*LimitedUserListResponse, error) {
	path := "/auth/user/friends"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offline) {
		queryParams["offline"] = fmt.Sprintf("%v", params.Offline)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result LimitedUserListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteFriendRequestParams represents the parameters for the DeleteFriendRequest request
type DeleteFriendRequestParams struct {
	UserId string `json:"userId"`
}

func (c *Client) DeleteFriendRequest(params DeleteFriendRequestParams) (*DeleteFriendSuccess, error) {
	path := "/user/{userId}/friendRequest"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result DeleteFriendSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// FriendParams represents the parameters for the Friend request
type FriendParams struct {
	UserId string `json:"userId"`
}

func (c *Client) Friend(params FriendParams) (*NotificationResponse, error) {
	path := "/user/{userId}/friendRequest"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result NotificationResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetFriendStatusParams represents the parameters for the GetFriendStatus request
type GetFriendStatusParams struct {
	UserId string `json:"userId"`
}

func (c *Client) GetFriendStatus(params GetFriendStatusParams) (*FriendStatusResponse, error) {
	path := "/user/{userId}/friendStatus"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FriendStatusResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UnfriendParams represents the parameters for the Unfriend request
type UnfriendParams struct {
	UserId string `json:"userId"`
}

func (c *Client) Unfriend(params UnfriendParams) (*UnfriendSuccess, error) {
	path := "/auth/user/friends/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result UnfriendSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// SearchGroupsParams represents the parameters for the SearchGroups request
type SearchGroupsParams struct {
	Offset int64 `json:"offset"`
	N      int64 `json:"n"`
}

func (c *Client) SearchGroups(params SearchGroupsParams) (*LimitedGroupListResponse, error) {
	path := "/groups"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result LimitedGroupListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) CreateGroup() (*GroupResponse, error) {
	path := "/groups"

	// Create request
	req := c.client.R()
	// Set response object
	var result GroupResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UpdateGroupParams represents the parameters for the UpdateGroup request
type UpdateGroupParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) UpdateGroup(params UpdateGroupParams) (*GroupResponse, error) {
	path := "/groups/{groupId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteGroupParams represents the parameters for the DeleteGroup request
type DeleteGroupParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) DeleteGroup(params DeleteGroupParams) (*DeleteGroupSuccess, error) {
	path := "/groups/{groupId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result DeleteGroupSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetGroupParams represents the parameters for the GetGroup request
type GetGroupParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) GetGroup(params GetGroupParams) (*GroupResponse, error) {
	path := "/groups/{groupId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteGroupAnnouncementParams represents the parameters for the DeleteGroupAnnouncement request
type DeleteGroupAnnouncementParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) DeleteGroupAnnouncement(params DeleteGroupAnnouncementParams) (*DeleteGroupAnnouncementSuccess, error) {
	path := "/groups/{groupId}/announcement"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result DeleteGroupAnnouncementSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetGroupAnnouncementsParams represents the parameters for the GetGroupAnnouncements request
type GetGroupAnnouncementsParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) GetGroupAnnouncements(params GetGroupAnnouncementsParams) (*GroupAnnouncementResponse, error) {
	path := "/groups/{groupId}/announcement"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupAnnouncementResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// CreateGroupAnnouncementParams represents the parameters for the CreateGroupAnnouncement request
type CreateGroupAnnouncementParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) CreateGroupAnnouncement(params CreateGroupAnnouncementParams) (*GroupAnnouncementResponse, error) {
	path := "/groups/{groupId}/announcement"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupAnnouncementResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetGroupAuditLogsParams represents the parameters for the GetGroupAuditLogs request
type GetGroupAuditLogsParams struct {
	GroupId   string    `json:"groupId"`
	N         int64     `json:"n"`
	Offset    int64     `json:"offset"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
}

func (c *Client) GetGroupAuditLogs(params GetGroupAuditLogsParams) (*GroupAuditLogListResponse, error) {
	path := "/groups/{groupId}/auditLogs"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}
	if lo.IsNotEmpty(params.StartDate) {
		queryParams["startDate"] = fmt.Sprintf("%v", params.StartDate)
	}
	if lo.IsNotEmpty(params.EndDate) {
		queryParams["endDate"] = fmt.Sprintf("%v", params.EndDate)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupAuditLogListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetGroupBansParams represents the parameters for the GetGroupBans request
type GetGroupBansParams struct {
	GroupId string `json:"groupId"`
	N       int64  `json:"n"`
	Offset  int64  `json:"offset"`
}

func (c *Client) GetGroupBans(params GetGroupBansParams) (*GroupMemberListResponse, error) {
	path := "/groups/{groupId}/bans"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupMemberListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// BanGroupMemberParams represents the parameters for the BanGroupMember request
type BanGroupMemberParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) BanGroupMember(params BanGroupMemberParams) (*GroupMemberResponse, error) {
	path := "/groups/{groupId}/bans"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupMemberResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UnbanGroupMemberParams represents the parameters for the UnbanGroupMember request
type UnbanGroupMemberParams struct {
	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
}

func (c *Client) UnbanGroupMember(params UnbanGroupMemberParams) (*GroupMemberResponse, error) {
	path := "/groups/{groupId}/bans/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupMemberResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// CreateGroupGalleryParams represents the parameters for the CreateGroupGallery request
type CreateGroupGalleryParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) CreateGroupGallery(params CreateGroupGalleryParams) (*GroupGalleryResponse, error) {
	path := "/groups/{groupId}/galleries"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupGalleryResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteGroupGalleryParams represents the parameters for the DeleteGroupGallery request
type DeleteGroupGalleryParams struct {
	GroupId        string `json:"groupId"`
	GroupGalleryId string `json:"groupGalleryId"`
}

func (c *Client) DeleteGroupGallery(params DeleteGroupGalleryParams) (*DeleteGroupGallerySuccess, error) {
	path := "/groups/{groupId}/galleries/{groupGalleryId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{groupGalleryId}", fmt.Sprintf("%v", params.GroupGalleryId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result DeleteGroupGallerySuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetGroupGalleryImagesParams represents the parameters for the GetGroupGalleryImages request
type GetGroupGalleryImagesParams struct {
	GroupId        string `json:"groupId"`
	GroupGalleryId string `json:"groupGalleryId"`
	N              int64  `json:"n"`
	Offset         int64  `json:"offset"`
}

func (c *Client) GetGroupGalleryImages(params GetGroupGalleryImagesParams) (*GroupGalleryImageListResponse, error) {
	path := "/groups/{groupId}/galleries/{groupGalleryId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{groupGalleryId}", fmt.Sprintf("%v", params.GroupGalleryId))
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupGalleryImageListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UpdateGroupGalleryParams represents the parameters for the UpdateGroupGallery request
type UpdateGroupGalleryParams struct {
	GroupId        string `json:"groupId"`
	GroupGalleryId string `json:"groupGalleryId"`
}

func (c *Client) UpdateGroupGallery(params UpdateGroupGalleryParams) (*GroupGalleryResponse, error) {
	path := "/groups/{groupId}/galleries/{groupGalleryId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{groupGalleryId}", fmt.Sprintf("%v", params.GroupGalleryId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupGalleryResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// AddGroupGalleryImageParams represents the parameters for the AddGroupGalleryImage request
type AddGroupGalleryImageParams struct {
	GroupId        string `json:"groupId"`
	GroupGalleryId string `json:"groupGalleryId"`
}

func (c *Client) AddGroupGalleryImage(params AddGroupGalleryImageParams) (*GroupGalleryImageResponse, error) {
	path := "/groups/{groupId}/galleries/{groupGalleryId}/images"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{groupGalleryId}", fmt.Sprintf("%v", params.GroupGalleryId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupGalleryImageResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteGroupGalleryImageParams represents the parameters for the DeleteGroupGalleryImage request
type DeleteGroupGalleryImageParams struct {
	GroupId             string `json:"groupId"`
	GroupGalleryId      string `json:"groupGalleryId"`
	GroupGalleryImageId string `json:"groupGalleryImageId"`
}

func (c *Client) DeleteGroupGalleryImage(params DeleteGroupGalleryImageParams) (*DeleteGroupGalleryImageSuccess, error) {
	path := "/groups/{groupId}/galleries/{groupGalleryId}/images/{groupGalleryImageId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{groupGalleryId}", fmt.Sprintf("%v", params.GroupGalleryId))
	path = strings.ReplaceAll(path, "{groupGalleryImageId}", fmt.Sprintf("%v", params.GroupGalleryImageId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result DeleteGroupGalleryImageSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetGroupInstancesParams represents the parameters for the GetGroupInstances request
type GetGroupInstancesParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) GetGroupInstances(params GetGroupInstancesParams) (*GroupInstanceListResponse, error) {
	path := "/groups/{groupId}/instances"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupInstanceListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetGroupInvitesParams represents the parameters for the GetGroupInvites request
type GetGroupInvitesParams struct {
	GroupId string `json:"groupId"`
	N       int64  `json:"n"`
	Offset  int64  `json:"offset"`
}

func (c *Client) GetGroupInvites(params GetGroupInvitesParams) (*GroupMemberListResponse, error) {
	path := "/groups/{groupId}/invites"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupMemberListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// CreateGroupInviteParams represents the parameters for the CreateGroupInvite request
type CreateGroupInviteParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) CreateGroupInvite(params CreateGroupInviteParams) error {
	path := "/groups/{groupId}/invites"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// DeleteGroupInviteParams represents the parameters for the DeleteGroupInvite request
type DeleteGroupInviteParams struct {
	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
}

func (c *Client) DeleteGroupInvite(params DeleteGroupInviteParams) error {
	path := "/groups/{groupId}/invites/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// JoinGroupParams represents the parameters for the JoinGroup request
type JoinGroupParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) JoinGroup(params JoinGroupParams) (*GroupMemberResponse, error) {
	path := "/groups/{groupId}/join"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupMemberResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// LeaveGroupParams represents the parameters for the LeaveGroup request
type LeaveGroupParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) LeaveGroup(params LeaveGroupParams) error {
	path := "/groups/{groupId}/leave"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// GetGroupMembersParams represents the parameters for the GetGroupMembers request
type GetGroupMembersParams struct {
	GroupId string          `json:"groupId"`
	N       int64           `json:"n"`
	Offset  int64           `json:"offset"`
	Sort    GroupSearchSort `json:"sort"`
}

func (c *Client) GetGroupMembers(params GetGroupMembersParams) (*GroupMemberListResponse, error) {
	path := "/groups/{groupId}/members"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}
	if lo.IsNotEmpty(params.Sort) {
		queryParams["sort"] = fmt.Sprintf("%v", params.Sort)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupMemberListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// KickGroupMemberParams represents the parameters for the KickGroupMember request
type KickGroupMemberParams struct {
	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
}

func (c *Client) KickGroupMember(params KickGroupMemberParams) error {
	path := "/groups/{groupId}/members/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// GetGroupMemberParams represents the parameters for the GetGroupMember request
type GetGroupMemberParams struct {
	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
}

func (c *Client) GetGroupMember(params GetGroupMemberParams) (*GroupLimitedMemberResponse, error) {
	path := "/groups/{groupId}/members/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupLimitedMemberResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UpdateGroupMemberParams represents the parameters for the UpdateGroupMember request
type UpdateGroupMemberParams struct {
	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
}

func (c *Client) UpdateGroupMember(params UpdateGroupMemberParams) (*GroupLimitedMemberResponse, error) {
	path := "/groups/{groupId}/members/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupLimitedMemberResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// RemoveGroupMemberRoleParams represents the parameters for the RemoveGroupMemberRole request
type RemoveGroupMemberRoleParams struct {
	GroupId     string `json:"groupId"`
	UserId      string `json:"userId"`
	GroupRoleId string `json:"groupRoleId"`
}

func (c *Client) RemoveGroupMemberRole(params RemoveGroupMemberRoleParams) (*GroupRoleIdListResponse, error) {
	path := "/groups/{groupId}/members/{userId}/roles/{groupRoleId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))
	path = strings.ReplaceAll(path, "{groupRoleId}", fmt.Sprintf("%v", params.GroupRoleId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupRoleIdListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// AddGroupMemberRoleParams represents the parameters for the AddGroupMemberRole request
type AddGroupMemberRoleParams struct {
	GroupId     string `json:"groupId"`
	UserId      string `json:"userId"`
	GroupRoleId string `json:"groupRoleId"`
}

func (c *Client) AddGroupMemberRole(params AddGroupMemberRoleParams) (*GroupRoleIdListResponse, error) {
	path := "/groups/{groupId}/members/{userId}/roles/{groupRoleId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))
	path = strings.ReplaceAll(path, "{groupRoleId}", fmt.Sprintf("%v", params.GroupRoleId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupRoleIdListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetGroupPermissionsParams represents the parameters for the GetGroupPermissions request
type GetGroupPermissionsParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) GetGroupPermissions(params GetGroupPermissionsParams) (*GroupPermissionListResponse, error) {
	path := "/groups/{groupId}/permissions"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupPermissionListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetGroupPostParams represents the parameters for the GetGroupPost request
type GetGroupPostParams struct {
	GroupId string `json:"groupId"`
	N       int64  `json:"n"`
	Offset  int64  `json:"offset"`
}

func (c *Client) GetGroupPost(params GetGroupPostParams) (*GroupPostResponse, error) {
	path := "/groups/{groupId}/posts"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupPostResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// AddGroupPostParams represents the parameters for the AddGroupPost request
type AddGroupPostParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) AddGroupPost(params AddGroupPostParams) (*GroupPostResponse, error) {
	path := "/groups/{groupId}/posts"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupPostResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteGroupPostParams represents the parameters for the DeleteGroupPost request
type DeleteGroupPostParams struct {
	GroupId        string `json:"groupId"`
	NotificationId string `json:"notificationId"`
}

func (c *Client) DeleteGroupPost(params DeleteGroupPostParams) (*GroupPostResponseSuccess, error) {
	path := "/groups/{groupId}/posts/{notificationId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{notificationId}", fmt.Sprintf("%v", params.NotificationId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupPostResponseSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UpdateGroupPostParams represents the parameters for the UpdateGroupPost request
type UpdateGroupPostParams struct {
	GroupId        string `json:"groupId"`
	NotificationId string `json:"notificationId"`
}

func (c *Client) UpdateGroupPost(params UpdateGroupPostParams) (*GroupPostResponse, error) {
	path := "/groups/{groupId}/posts/{notificationId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{notificationId}", fmt.Sprintf("%v", params.NotificationId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupPostResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// CancelGroupRequestParams represents the parameters for the CancelGroupRequest request
type CancelGroupRequestParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) CancelGroupRequest(params CancelGroupRequestParams) error {
	path := "/groups/{groupId}/requests"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// GetGroupRequestsParams represents the parameters for the GetGroupRequests request
type GetGroupRequestsParams struct {
	GroupId string `json:"groupId"`
	N       int64  `json:"n"`
	Offset  int64  `json:"offset"`
}

func (c *Client) GetGroupRequests(params GetGroupRequestsParams) (*GroupMemberListResponse, error) {
	path := "/groups/{groupId}/requests"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupMemberListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// RespondGroupJoinRequestParams represents the parameters for the RespondGroupJoinRequest request
type RespondGroupJoinRequestParams struct {
	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
}

func (c *Client) RespondGroupJoinRequest(params RespondGroupJoinRequestParams) error {
	path := "/groups/{groupId}/requests/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// GetGroupRolesParams represents the parameters for the GetGroupRoles request
type GetGroupRolesParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) GetGroupRoles(params GetGroupRolesParams) (*GroupRoleListResponse, error) {
	path := "/groups/{groupId}/roles"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupRoleListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// CreateGroupRoleParams represents the parameters for the CreateGroupRole request
type CreateGroupRoleParams struct {
	GroupId string `json:"groupId"`
}

func (c *Client) CreateGroupRole(params CreateGroupRoleParams) (*GroupRoleResponse, error) {
	path := "/groups/{groupId}/roles"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupRoleResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteGroupRoleParams represents the parameters for the DeleteGroupRole request
type DeleteGroupRoleParams struct {
	GroupId     string `json:"groupId"`
	GroupRoleId string `json:"groupRoleId"`
}

func (c *Client) DeleteGroupRole(params DeleteGroupRoleParams) (*GroupRoleListResponse, error) {
	path := "/groups/{groupId}/roles/{groupRoleId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{groupRoleId}", fmt.Sprintf("%v", params.GroupRoleId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupRoleListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UpdateGroupRoleParams represents the parameters for the UpdateGroupRole request
type UpdateGroupRoleParams struct {
	GroupId     string `json:"groupId"`
	GroupRoleId string `json:"groupRoleId"`
}

func (c *Client) UpdateGroupRole(params UpdateGroupRoleParams) (*GroupRoleListResponse, error) {
	path := "/groups/{groupId}/roles/{groupRoleId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{groupId}", fmt.Sprintf("%v", params.GroupId))
	path = strings.ReplaceAll(path, "{groupRoleId}", fmt.Sprintf("%v", params.GroupRoleId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupRoleListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// InviteUserParams represents the parameters for the InviteUser request
type InviteUserParams struct {
	UserId string `json:"userId"`
}

func (c *Client) InviteUser(params InviteUserParams) (*SendNotificationResponse, error) {
	path := "/invite/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result SendNotificationResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// InviteMyselfToParams represents the parameters for the InviteMyselfTo request
type InviteMyselfToParams struct {
	WorldId    string `json:"worldId"`
	InstanceId string `json:"instanceId"`
}

func (c *Client) InviteMyselfTo(params InviteMyselfToParams) (*SendNotificationResponse, error) {
	path := "/invite/myself/to/{worldId}:{instanceId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))
	path = strings.ReplaceAll(path, "{instanceId}", fmt.Sprintf("%v", params.InstanceId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result SendNotificationResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// RequestInviteParams represents the parameters for the RequestInvite request
type RequestInviteParams struct {
	UserId string `json:"userId"`
}

func (c *Client) RequestInvite(params RequestInviteParams) (*NotificationResponse, error) {
	path := "/requestInvite/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result NotificationResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// RespondInviteParams represents the parameters for the RespondInvite request
type RespondInviteParams struct {
	NotificationId string `json:"notificationId"`
}

func (c *Client) RespondInvite(params RespondInviteParams) (*NotificationResponse, error) {
	path := "/invite/{notificationId}/response"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{notificationId}", fmt.Sprintf("%v", params.NotificationId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result NotificationResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetInviteMessagesParams represents the parameters for the GetInviteMessages request
type GetInviteMessagesParams struct {
	UserId      string            `json:"userId"`
	MessageType InviteMessageType `json:"messageType"`
}

func (c *Client) GetInviteMessages(params GetInviteMessagesParams) (*InviteMessageListResponse, error) {
	path := "/message/{userId}/{messageType}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))
	path = strings.ReplaceAll(path, "{messageType}", fmt.Sprintf("%v", params.MessageType))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result InviteMessageListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// ResetInviteMessageParams represents the parameters for the ResetInviteMessage request
type ResetInviteMessageParams struct {
	UserId      string            `json:"userId"`
	MessageType InviteMessageType `json:"messageType"`
	Slot        int64             `json:"slot"`
}

func (c *Client) ResetInviteMessage(params ResetInviteMessageParams) (*InviteMessageListResponse, error) {
	path := "/message/{userId}/{messageType}/{slot}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))
	path = strings.ReplaceAll(path, "{messageType}", fmt.Sprintf("%v", params.MessageType))
	path = strings.ReplaceAll(path, "{slot}", fmt.Sprintf("%v", params.Slot))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result InviteMessageListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetInviteMessageParams represents the parameters for the GetInviteMessage request
type GetInviteMessageParams struct {
	UserId      string            `json:"userId"`
	MessageType InviteMessageType `json:"messageType"`
	Slot        int64             `json:"slot"`
}

func (c *Client) GetInviteMessage(params GetInviteMessageParams) (*InviteMessageResponse, error) {
	path := "/message/{userId}/{messageType}/{slot}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))
	path = strings.ReplaceAll(path, "{messageType}", fmt.Sprintf("%v", params.MessageType))
	path = strings.ReplaceAll(path, "{slot}", fmt.Sprintf("%v", params.Slot))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result InviteMessageResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UpdateInviteMessageParams represents the parameters for the UpdateInviteMessage request
type UpdateInviteMessageParams struct {
	UserId      string            `json:"userId"`
	MessageType InviteMessageType `json:"messageType"`
	Slot        int64             `json:"slot"`
}

func (c *Client) UpdateInviteMessage(params UpdateInviteMessageParams) (*InviteMessageListResponse, error) {
	path := "/message/{userId}/{messageType}/{slot}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))
	path = strings.ReplaceAll(path, "{messageType}", fmt.Sprintf("%v", params.MessageType))
	path = strings.ReplaceAll(path, "{slot}", fmt.Sprintf("%v", params.Slot))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result InviteMessageListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) CreateInstance() (*InstanceResponse, error) {
	path := "/instances"

	// Create request
	req := c.client.R()
	// Set response object
	var result InstanceResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// CloseInstanceParams represents the parameters for the CloseInstance request
type CloseInstanceParams struct {
	WorldId    string `json:"worldId"`
	InstanceId string `json:"instanceId"`
}

func (c *Client) CloseInstance(params CloseInstanceParams) (*InstanceResponse, error) {
	path := "/instances/{worldId}:{instanceId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))
	path = strings.ReplaceAll(path, "{instanceId}", fmt.Sprintf("%v", params.InstanceId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result InstanceResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetInstanceParams represents the parameters for the GetInstance request
type GetInstanceParams struct {
	WorldId    string `json:"worldId"`
	InstanceId string `json:"instanceId"`
}

func (c *Client) GetInstance(params GetInstanceParams) (*InstanceResponse, error) {
	path := "/instances/{worldId}:{instanceId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))
	path = strings.ReplaceAll(path, "{instanceId}", fmt.Sprintf("%v", params.InstanceId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result InstanceResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetShortNameParams represents the parameters for the GetShortName request
type GetShortNameParams struct {
	WorldId    string `json:"worldId"`
	InstanceId string `json:"instanceId"`
}

func (c *Client) GetShortName(params GetShortNameParams) (*InstanceShortNameResponse, error) {
	path := "/instances/{worldId}:{instanceId}/shortName"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))
	path = strings.ReplaceAll(path, "{instanceId}", fmt.Sprintf("%v", params.InstanceId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result InstanceShortNameResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// SendSelfInviteParams represents the parameters for the SendSelfInvite request
type SendSelfInviteParams struct {
	WorldId    string `json:"worldId"`
	InstanceId string `json:"instanceId"`
}

func (c *Client) SendSelfInvite(params SendSelfInviteParams) (*InstanceSelfInviteSuccess, error) {
	path := "/instances/{worldId}:{instanceId}/invite"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))
	path = strings.ReplaceAll(path, "{instanceId}", fmt.Sprintf("%v", params.InstanceId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result InstanceSelfInviteSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetInstanceByShortName() (*InstanceResponse, error) {
	path := "/instances/s/{shortName}"

	// Create request
	req := c.client.R()
	// Set response object
	var result InstanceResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetNotificationsParams represents the parameters for the GetNotifications request
type GetNotificationsParams struct {
	N      int64 `json:"n"`
	Offset int64 `json:"offset"`
}

func (c *Client) GetNotifications(params GetNotificationsParams) (*NotificationListResponse, error) {
	path := "/auth/user/notifications"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result NotificationListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// AcceptFriendRequestParams represents the parameters for the AcceptFriendRequest request
type AcceptFriendRequestParams struct {
	NotificationId string `json:"notificationId"`
}

func (c *Client) AcceptFriendRequest(params AcceptFriendRequestParams) (*FriendSuccess, error) {
	path := "/auth/user/notifications/{notificationId}/accept"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{notificationId}", fmt.Sprintf("%v", params.NotificationId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result FriendSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// MarkNotificationAsReadParams represents the parameters for the MarkNotificationAsRead request
type MarkNotificationAsReadParams struct {
	NotificationId string `json:"notificationId"`
}

func (c *Client) MarkNotificationAsRead(params MarkNotificationAsReadParams) (*NotificationResponse, error) {
	path := "/auth/user/notifications/{notificationId}/see"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{notificationId}", fmt.Sprintf("%v", params.NotificationId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result NotificationResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteNotificationParams represents the parameters for the DeleteNotification request
type DeleteNotificationParams struct {
	NotificationId string `json:"notificationId"`
}

func (c *Client) DeleteNotification(params DeleteNotificationParams) (*NotificationResponse, error) {
	path := "/auth/user/notifications/{notificationId}/hide"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{notificationId}", fmt.Sprintf("%v", params.NotificationId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result NotificationResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) ClearNotifications() (*ClearNotificationsSuccess, error) {
	path := "/auth/user/notifications/clear"

	// Create request
	req := c.client.R()
	// Set response object
	var result ClearNotificationsSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetAssignedPermissions() (*PermissionListResponse, error) {
	path := "/auth/permissions"

	// Create request
	req := c.client.R()
	// Set response object
	var result PermissionListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetPermissionParams represents the parameters for the GetPermission request
type GetPermissionParams struct {
	PermissionId string `json:"permissionId"`
}

func (c *Client) GetPermission(params GetPermissionParams) (*PermissionResponse, error) {
	path := "/permissions/{permissionId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{permissionId}", fmt.Sprintf("%v", params.PermissionId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result PermissionResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) ClearAllPlayerModerations() (*PlayerModerationClearAllSuccess, error) {
	path := "/auth/user/playermoderations"

	// Create request
	req := c.client.R()
	// Set response object
	var result PlayerModerationClearAllSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetPlayerModerations() (*PlayerModerationListResponse, error) {
	path := "/auth/user/playermoderations"

	// Create request
	req := c.client.R()
	// Set response object
	var result PlayerModerationListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) ModerateUser() (*PlayerModerationResponse, error) {
	path := "/auth/user/playermoderations"

	// Create request
	req := c.client.R()
	// Set response object
	var result PlayerModerationResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) DeletePlayerModeration() (*PlayerModerationRemovedSuccess, error) {
	path := "/auth/user/playermoderations/{playerModerationId}"

	// Create request
	req := c.client.R()
	// Set response object
	var result PlayerModerationRemovedSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetPlayerModeration() (*PlayerModerationResponse, error) {
	path := "/auth/user/playermoderations/{playerModerationId}"

	// Create request
	req := c.client.R()
	// Set response object
	var result PlayerModerationResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) UnmoderateUser() (*PlayerModerationUnmoderatedSuccess, error) {
	path := "/auth/user/unplayermoderate"

	// Create request
	req := c.client.R()
	// Set response object
	var result PlayerModerationUnmoderatedSuccess
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetConfig() (*ApiConfigResponse, error) {
	path := "/config"

	// Create request
	req := c.client.R()
	// Set response object
	var result ApiConfigResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetInfoPushParams represents the parameters for the GetInfoPush request
type GetInfoPushParams struct {
	Require string `json:"require"`
	Include string `json:"include"`
}

func (c *Client) GetInfoPush(params GetInfoPushParams) (*InfoPushListResponse, error) {
	path := "/infoPush"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Require) {
		queryParams["require"] = fmt.Sprintf("%v", params.Require)
	}
	if lo.IsNotEmpty(params.Include) {
		queryParams["include"] = fmt.Sprintf("%v", params.Include)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result InfoPushListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetCssParams represents the parameters for the GetCss request
type GetCssParams struct {
	// Variant enum
	Variant string `json:"variant"`
	Branch  string `json:"branch"`
}

func (c *Client) GetCss(params GetCssParams) error {
	path := "/css/app.css"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Variant) {
		queryParams["variant"] = fmt.Sprintf("%v", params.Variant)
	}
	if lo.IsNotEmpty(params.Branch) {
		queryParams["branch"] = fmt.Sprintf("%v", params.Branch)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// GetJavaScriptParams represents the parameters for the GetJavaScript request
type GetJavaScriptParams struct {
	// Variant enum
	Variant string `json:"variant"`
	Branch  string `json:"branch"`
}

func (c *Client) GetJavaScript(params GetJavaScriptParams) error {
	path := "/js/app.js"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Variant) {
		queryParams["variant"] = fmt.Sprintf("%v", params.Variant)
	}
	if lo.IsNotEmpty(params.Branch) {
		queryParams["branch"] = fmt.Sprintf("%v", params.Branch)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

func (c *Client) GetHealth() (*ApiHealthResponse, error) {
	path := "/health"

	// Create request
	req := c.client.R()
	// Set response object
	var result ApiHealthResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetCurrentOnlineUsers() (*CurrentOnlineUsersResponse, error) {
	path := "/visits"

	// Create request
	req := c.client.R()
	// Set response object
	var result CurrentOnlineUsersResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetSystemTime() (*SystemTimeResponse, error) {
	path := "/time"

	// Create request
	req := c.client.R()
	// Set response object
	var result SystemTimeResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// SearchUsersParams represents the parameters for the SearchUsers request
type SearchUsersParams struct {
	N      int64 `json:"n"`
	Offset int64 `json:"offset"`
}

func (c *Client) SearchUsers(params SearchUsersParams) (*LimitedUserListResponse, error) {
	path := "/users"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result LimitedUserListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) GetUserByName() (*UserResponse, error) {
	path := "/users/{username}/name"

	// Create request
	req := c.client.R()
	// Set response object
	var result UserResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetUserParams represents the parameters for the GetUser request
type GetUserParams struct {
	UserId string `json:"userId"`
}

func (c *Client) GetUser(params GetUserParams) (*UserResponse, error) {
	path := "/users/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result UserResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UpdateUserParams represents the parameters for the UpdateUser request
type UpdateUserParams struct {
	UserId string `json:"userId"`
}

func (c *Client) UpdateUser(params UpdateUserParams) (*CurrentUserResponse, error) {
	path := "/users/{userId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result CurrentUserResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetUserGroupsParams represents the parameters for the GetUserGroups request
type GetUserGroupsParams struct {
	UserId string `json:"userId"`
}

func (c *Client) GetUserGroups(params GetUserGroupsParams) (*LimitedUserGroupListResponse, error) {
	path := "/users/{userId}/groups"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result LimitedUserGroupListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetUserGroupRequestsParams represents the parameters for the GetUserGroupRequests request
type GetUserGroupRequestsParams struct {
	UserId string `json:"userId"`
}

func (c *Client) GetUserGroupRequests(params GetUserGroupRequestsParams) (*GroupListResponse, error) {
	path := "/users/{userId}/groups/requested"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result GroupListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetUserRepresentedGroupParams represents the parameters for the GetUserRepresentedGroup request
type GetUserRepresentedGroupParams struct {
	UserId string `json:"userId"`
}

func (c *Client) GetUserRepresentedGroup(params GetUserRepresentedGroupParams) error {
	path := "/users/{userId}/groups/represented"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{userId}", fmt.Sprintf("%v", params.UserId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// SearchWorldsParams represents the parameters for the SearchWorlds request
type SearchWorldsParams struct {
	Featured        bool          `json:"featured"`
	Sort            SortOption    `json:"sort"`
	UserId          string        `json:"userId"`
	N               int64         `json:"n"`
	Order           OrderOption   `json:"order"`
	Offset          int64         `json:"offset"`
	Search          string        `json:"search"`
	Tag             string        `json:"tag"`
	Notag           string        `json:"notag"`
	ReleaseStatus   ReleaseStatus `json:"releaseStatus"`
	MaxUnityVersion string        `json:"maxUnityVersion"`
	MinUnityVersion string        `json:"minUnityVersion"`
	Platform        string        `json:"platform"`
	Fuzzy           bool          `json:"fuzzy"`
}

func (c *Client) SearchWorlds(params SearchWorldsParams) (*LimitedWorldListResponse, error) {
	path := "/worlds"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Featured) {
		queryParams["featured"] = fmt.Sprintf("%v", params.Featured)
	}
	if lo.IsNotEmpty(params.Sort) {
		queryParams["sort"] = fmt.Sprintf("%v", params.Sort)
	}
	if lo.IsNotEmpty(params.UserId) {
		queryParams["userId"] = fmt.Sprintf("%v", params.UserId)
	}
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Order) {
		queryParams["order"] = fmt.Sprintf("%v", params.Order)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}
	if lo.IsNotEmpty(params.Search) {
		queryParams["search"] = fmt.Sprintf("%v", params.Search)
	}
	if lo.IsNotEmpty(params.Tag) {
		queryParams["tag"] = fmt.Sprintf("%v", params.Tag)
	}
	if lo.IsNotEmpty(params.Notag) {
		queryParams["notag"] = fmt.Sprintf("%v", params.Notag)
	}
	if lo.IsNotEmpty(params.ReleaseStatus) {
		queryParams["releaseStatus"] = fmt.Sprintf("%v", params.ReleaseStatus)
	}
	if lo.IsNotEmpty(params.MaxUnityVersion) {
		queryParams["maxUnityVersion"] = fmt.Sprintf("%v", params.MaxUnityVersion)
	}
	if lo.IsNotEmpty(params.MinUnityVersion) {
		queryParams["minUnityVersion"] = fmt.Sprintf("%v", params.MinUnityVersion)
	}
	if lo.IsNotEmpty(params.Platform) {
		queryParams["platform"] = fmt.Sprintf("%v", params.Platform)
	}
	if lo.IsNotEmpty(params.Fuzzy) {
		queryParams["fuzzy"] = fmt.Sprintf("%v", params.Fuzzy)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result LimitedWorldListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

func (c *Client) CreateWorld() (*WorldResponse, error) {
	path := "/worlds"

	// Create request
	req := c.client.R()
	// Set response object
	var result WorldResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Post(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetActiveWorldsParams represents the parameters for the GetActiveWorlds request
type GetActiveWorldsParams struct {
	Featured        bool          `json:"featured"`
	Sort            SortOption    `json:"sort"`
	N               int64         `json:"n"`
	Order           OrderOption   `json:"order"`
	Offset          int64         `json:"offset"`
	Search          string        `json:"search"`
	Tag             string        `json:"tag"`
	Notag           string        `json:"notag"`
	ReleaseStatus   ReleaseStatus `json:"releaseStatus"`
	MaxUnityVersion string        `json:"maxUnityVersion"`
	MinUnityVersion string        `json:"minUnityVersion"`
	Platform        string        `json:"platform"`
}

func (c *Client) GetActiveWorlds(params GetActiveWorldsParams) (*LimitedWorldListResponse, error) {
	path := "/worlds/active"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Featured) {
		queryParams["featured"] = fmt.Sprintf("%v", params.Featured)
	}
	if lo.IsNotEmpty(params.Sort) {
		queryParams["sort"] = fmt.Sprintf("%v", params.Sort)
	}
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Order) {
		queryParams["order"] = fmt.Sprintf("%v", params.Order)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}
	if lo.IsNotEmpty(params.Search) {
		queryParams["search"] = fmt.Sprintf("%v", params.Search)
	}
	if lo.IsNotEmpty(params.Tag) {
		queryParams["tag"] = fmt.Sprintf("%v", params.Tag)
	}
	if lo.IsNotEmpty(params.Notag) {
		queryParams["notag"] = fmt.Sprintf("%v", params.Notag)
	}
	if lo.IsNotEmpty(params.ReleaseStatus) {
		queryParams["releaseStatus"] = fmt.Sprintf("%v", params.ReleaseStatus)
	}
	if lo.IsNotEmpty(params.MaxUnityVersion) {
		queryParams["maxUnityVersion"] = fmt.Sprintf("%v", params.MaxUnityVersion)
	}
	if lo.IsNotEmpty(params.MinUnityVersion) {
		queryParams["minUnityVersion"] = fmt.Sprintf("%v", params.MinUnityVersion)
	}
	if lo.IsNotEmpty(params.Platform) {
		queryParams["platform"] = fmt.Sprintf("%v", params.Platform)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result LimitedWorldListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetFavoritedWorldsParams represents the parameters for the GetFavoritedWorlds request
type GetFavoritedWorldsParams struct {
	Featured        bool          `json:"featured"`
	Sort            SortOption    `json:"sort"`
	N               int64         `json:"n"`
	Order           OrderOption   `json:"order"`
	Offset          int64         `json:"offset"`
	Search          string        `json:"search"`
	Tag             string        `json:"tag"`
	Notag           string        `json:"notag"`
	ReleaseStatus   ReleaseStatus `json:"releaseStatus"`
	MaxUnityVersion string        `json:"maxUnityVersion"`
	MinUnityVersion string        `json:"minUnityVersion"`
	Platform        string        `json:"platform"`
	UserId          string        `json:"userId"`
}

func (c *Client) GetFavoritedWorlds(params GetFavoritedWorldsParams) (*LimitedWorldListResponse, error) {
	path := "/worlds/favorites"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Featured) {
		queryParams["featured"] = fmt.Sprintf("%v", params.Featured)
	}
	if lo.IsNotEmpty(params.Sort) {
		queryParams["sort"] = fmt.Sprintf("%v", params.Sort)
	}
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Order) {
		queryParams["order"] = fmt.Sprintf("%v", params.Order)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}
	if lo.IsNotEmpty(params.Search) {
		queryParams["search"] = fmt.Sprintf("%v", params.Search)
	}
	if lo.IsNotEmpty(params.Tag) {
		queryParams["tag"] = fmt.Sprintf("%v", params.Tag)
	}
	if lo.IsNotEmpty(params.Notag) {
		queryParams["notag"] = fmt.Sprintf("%v", params.Notag)
	}
	if lo.IsNotEmpty(params.ReleaseStatus) {
		queryParams["releaseStatus"] = fmt.Sprintf("%v", params.ReleaseStatus)
	}
	if lo.IsNotEmpty(params.MaxUnityVersion) {
		queryParams["maxUnityVersion"] = fmt.Sprintf("%v", params.MaxUnityVersion)
	}
	if lo.IsNotEmpty(params.MinUnityVersion) {
		queryParams["minUnityVersion"] = fmt.Sprintf("%v", params.MinUnityVersion)
	}
	if lo.IsNotEmpty(params.Platform) {
		queryParams["platform"] = fmt.Sprintf("%v", params.Platform)
	}
	if lo.IsNotEmpty(params.UserId) {
		queryParams["userId"] = fmt.Sprintf("%v", params.UserId)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result LimitedWorldListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetRecentWorldsParams represents the parameters for the GetRecentWorlds request
type GetRecentWorldsParams struct {
	Featured        bool          `json:"featured"`
	Sort            SortOption    `json:"sort"`
	N               int64         `json:"n"`
	Order           OrderOption   `json:"order"`
	Offset          int64         `json:"offset"`
	Search          string        `json:"search"`
	Tag             string        `json:"tag"`
	Notag           string        `json:"notag"`
	ReleaseStatus   ReleaseStatus `json:"releaseStatus"`
	MaxUnityVersion string        `json:"maxUnityVersion"`
	MinUnityVersion string        `json:"minUnityVersion"`
	Platform        string        `json:"platform"`
	UserId          string        `json:"userId"`
}

func (c *Client) GetRecentWorlds(params GetRecentWorldsParams) (*LimitedWorldListResponse, error) {
	path := "/worlds/recent"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	if lo.IsNotEmpty(params.Featured) {
		queryParams["featured"] = fmt.Sprintf("%v", params.Featured)
	}
	if lo.IsNotEmpty(params.Sort) {
		queryParams["sort"] = fmt.Sprintf("%v", params.Sort)
	}
	if lo.IsNotEmpty(params.N) {
		queryParams["n"] = fmt.Sprintf("%v", params.N)
	}
	if lo.IsNotEmpty(params.Order) {
		queryParams["order"] = fmt.Sprintf("%v", params.Order)
	}
	if lo.IsNotEmpty(params.Offset) {
		queryParams["offset"] = fmt.Sprintf("%v", params.Offset)
	}
	if lo.IsNotEmpty(params.Search) {
		queryParams["search"] = fmt.Sprintf("%v", params.Search)
	}
	if lo.IsNotEmpty(params.Tag) {
		queryParams["tag"] = fmt.Sprintf("%v", params.Tag)
	}
	if lo.IsNotEmpty(params.Notag) {
		queryParams["notag"] = fmt.Sprintf("%v", params.Notag)
	}
	if lo.IsNotEmpty(params.ReleaseStatus) {
		queryParams["releaseStatus"] = fmt.Sprintf("%v", params.ReleaseStatus)
	}
	if lo.IsNotEmpty(params.MaxUnityVersion) {
		queryParams["maxUnityVersion"] = fmt.Sprintf("%v", params.MaxUnityVersion)
	}
	if lo.IsNotEmpty(params.MinUnityVersion) {
		queryParams["minUnityVersion"] = fmt.Sprintf("%v", params.MinUnityVersion)
	}
	if lo.IsNotEmpty(params.Platform) {
		queryParams["platform"] = fmt.Sprintf("%v", params.Platform)
	}
	if lo.IsNotEmpty(params.UserId) {
		queryParams["userId"] = fmt.Sprintf("%v", params.UserId)
	}

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result LimitedWorldListResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// DeleteWorldParams represents the parameters for the DeleteWorld request
type DeleteWorldParams struct {
	WorldId string `json:"worldId"`
}

func (c *Client) DeleteWorld(params DeleteWorldParams) error {
	path := "/worlds/{worldId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// GetWorldParams represents the parameters for the GetWorld request
type GetWorldParams struct {
	WorldId string `json:"worldId"`
}

func (c *Client) GetWorld(params GetWorldParams) (*WorldResponse, error) {
	path := "/worlds/{worldId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result WorldResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UpdateWorldParams represents the parameters for the UpdateWorld request
type UpdateWorldParams struct {
	WorldId string `json:"worldId"`
}

func (c *Client) UpdateWorld(params UpdateWorldParams) (*WorldResponse, error) {
	path := "/worlds/{worldId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result WorldResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// GetWorldMetadataParams represents the parameters for the GetWorldMetadata request
type GetWorldMetadataParams struct {
	WorldId string `json:"worldId"`
}

func (c *Client) GetWorldMetadata(params GetWorldMetadataParams) (*WorldMetadataResponse, error) {
	path := "/worlds/{worldId}/metadata"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result WorldMetadataResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// UnpublishWorldParams represents the parameters for the UnpublishWorld request
type UnpublishWorldParams struct {
	WorldId string `json:"worldId"`
}

func (c *Client) UnpublishWorld(params UnpublishWorldParams) error {
	path := "/worlds/{worldId}/publish"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Delete(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// GetWorldPublishStatusParams represents the parameters for the GetWorldPublishStatus request
type GetWorldPublishStatusParams struct {
	WorldId string `json:"worldId"`
}

func (c *Client) GetWorldPublishStatus(params GetWorldPublishStatusParams) (*WorldPublishStatusResponse, error) {
	path := "/worlds/{worldId}/publish"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result WorldPublishStatusResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}

// PublishWorldParams represents the parameters for the PublishWorld request
type PublishWorldParams struct {
	WorldId string `json:"worldId"`
}

func (c *Client) PublishWorld(params PublishWorldParams) error {
	path := "/worlds/{worldId}/publish"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)

	// Send request
	resp, err := req.Put(path)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return nil
}

// GetWorldInstanceParams represents the parameters for the GetWorldInstance request
type GetWorldInstanceParams struct {
	WorldId    string `json:"worldId"`
	InstanceId string `json:"instanceId"`
}

func (c *Client) GetWorldInstance(params GetWorldInstanceParams) (*InstanceResponse, error) {
	path := "/worlds/{worldId}/{instanceId}"
	// Replace path parameters and prepare query parameters
	queryParams := make(map[string]string)
	path = strings.ReplaceAll(path, "{worldId}", fmt.Sprintf("%v", params.WorldId))
	path = strings.ReplaceAll(path, "{instanceId}", fmt.Sprintf("%v", params.InstanceId))

	// Create request
	req := c.client.R()
	// Set query parameters
	req.SetQueryParams(queryParams)
	// Set response object
	var result InstanceResponse
	req.SetResult(&result)

	// Send request
	resp, err := req.Get(path)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}

	// Check for successful status code
	if resp.StatusCode() < 200 || resp.StatusCode() >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.String())
	}
	return &result, nil
}
