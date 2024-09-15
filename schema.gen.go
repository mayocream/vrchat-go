package vrchat

import (
	"time"
)

// UserExists Status object representing if a queried user by username or userId exists or not. This model is primarily used by the `/auth/exists` endpoint, which in turn is used during registration. Please see the documentation on that endpoint for more information on usage.
type UserExists struct {
	// NameOk Is the username valid?
	NameOk bool `json:"nameOk"`

	// UserExists Status if a user exist with that username or userId.
	UserExists bool `json:"userExists"`
}

type Response struct {
	Message    string `json:"message,omitempty"`
	StatusCode int64  `json:"status_code"`
}

type Error struct {
	Error Response `json:"error,omitempty"`
}

type AccountDeletionLog struct {
	// DateTime Date and time of the deletion request.
	DateTime time.Time `json:"dateTime,omitempty"`

	// DeletionScheduled When the deletion is scheduled to happen, standard is 14 days after the request.
	DeletionScheduled time.Time `json:"deletionScheduled,omitempty"`

	// Message Typically "Deletion requested" or "Deletion canceled". Other messages like "Deletion completed" may exist, but are these are not possible to see as a regular user.
	Message string `json:"message,omitempty"`
}

// UserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
type UserId string

type BadgeId string

type Badge struct {
	// AssignedAt only present in CurrentUser badges
	AssignedAt       time.Time `json:"assignedAt,omitempty"`
	BadgeDescription string    `json:"badgeDescription"`
	BadgeId          BadgeId   `json:"badgeId"`

	// BadgeImageUrl direct url to image
	BadgeImageUrl string `json:"badgeImageUrl"`
	BadgeName     string `json:"badgeName"`

	// Hidden only present in CurrentUser badges
	Hidden    bool `json:"hidden,omitempty"`
	Showcased bool `json:"showcased"`

	// UpdatedAt only present in CurrentUser badges
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type AvatarId string

// CurrentAvatarImageUrl When profilePicOverride is not empty, use it instead.
type CurrentAvatarImageUrl string

// CurrentAvatarThumbnailImageUrl When profilePicOverride is not empty, use it instead.
type CurrentAvatarThumbnailImageUrl string

// Tag Tags are a way to grant various access, assign restrictions or other kinds of metadata to various to objects such as worlds, users and avatars.
//
// System tags starting with `system_` are granted automatically by the system, while admin tags with `admin_` are granted manually. More prefixes such as `language_ ` (to indicate that a player can speak the tagged language), and `author_tag_` (provided by a world author for search and sorting) exist as well.
type Tag string

// DeveloperType "none" User is a normal user
// "trusted" Unknown
// "internal" Is a VRChat Developer
// "moderator" Is a VRChat Moderator
//
// Staff can hide their developerType at will.
type DeveloperType string

const (
	DeveloperTypeNone      DeveloperType = "none"
	DeveloperTypeTrusted   DeveloperType = "trusted"
	DeveloperTypeInternal  DeveloperType = "internal"
	DeveloperTypeModerator DeveloperType = "moderator"
)

// WorldId WorldID be "offline" on User profiles if you are not friends with that user.
type WorldId string

// Platform This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
type Platform string

type PastDisplayName struct {
	DisplayName string    `json:"displayName"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GroupId string

type CurrentUserPresence struct {
	AvatarThumbnail string    `json:"avatarThumbnail,omitempty"`
	DisplayName     string    `json:"displayName,omitempty"`
	Groups          []GroupId `json:"groups,omitempty"`

	// Id A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id       UserId `json:"id,omitempty"`
	Instance string `json:"instance,omitempty"`

	// InstanceType either an InstanceType or an empty string
	InstanceType string `json:"instanceType,omitempty"`
	IsRejoining  string `json:"isRejoining,omitempty"`

	// Platform either a Platform or an empty string
	Platform           string `json:"platform,omitempty"`
	ProfilePicOverride string `json:"profilePicOverride,omitempty"`

	// Status either a UserStatus or empty string
	Status              string `json:"status,omitempty"`
	TravelingToInstance string `json:"travelingToInstance,omitempty"`

	// TravelingToWorld WorldID be "offline" on User profiles if you are not friends with that user.
	TravelingToWorld WorldId `json:"travelingToWorld,omitempty"`

	// World WorldID be "offline" on User profiles if you are not friends with that user.
	World WorldId `json:"world,omitempty"`
}

// UserState * "online" User is online in VRChat
// * "active" User is online, but not in VRChat
// * "offline" User is offline
//
// Always offline when returned through `getCurrentUser` (/auth/user).
type UserState string

const (
	UserStateOffline UserState = "offline"
	UserStateActive  UserState = "active"
	UserStateOnline  UserState = "online"
)

// UserStatus Defines the User's current status, for example "ask me", "join me" or "offline. This status is a combined indicator of their online activity and privacy preference.
type UserStatus string

const (
	UserStatusActive  UserStatus = "active"
	UserStatusJoinMe  UserStatus = "join me"
	UserStatusAskMe   UserStatus = "ask me"
	UserStatusBusy    UserStatus = "busy"
	UserStatusOffline UserStatus = "offline"
)

type CurrentUser struct {
	AcceptedPrivacyVersion int64  `json:"acceptedPrivacyVersion,omitempty"`
	AcceptedTosVersion     int64  `json:"acceptedTOSVersion"`
	AccountDeletionDate    string `json:"accountDeletionDate,omitempty"`

	// AccountDeletionLog
	AccountDeletionLog []AccountDeletionLog `json:"accountDeletionLog,omitempty"`

	// ActiveFriends
	ActiveFriends      []UserId `json:"activeFriends,omitempty"`
	AllowAvatarCopying bool     `json:"allowAvatarCopying"`

	// Badges
	Badges []Badge `json:"badges,omitempty"`
	Bio    string  `json:"bio"`

	// BioLinks
	BioLinks              []string `json:"bioLinks"`
	CurrentAvatar         AvatarId `json:"currentAvatar"`
	CurrentAvatarAssetUrl string   `json:"currentAvatarAssetUrl"`

	// CurrentAvatarImageUrl When profilePicOverride is not empty, use it instead.
	CurrentAvatarImageUrl CurrentAvatarImageUrl `json:"currentAvatarImageUrl"`
	CurrentAvatarTags     []Tag                 `json:"currentAvatarTags"`

	// CurrentAvatarThumbnailImageUrl When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl CurrentAvatarThumbnailImageUrl `json:"currentAvatarThumbnailImageUrl"`
	DateJoined                     string                         `json:"date_joined"`

	// DeveloperType "none" User is a normal user
	// "trusted" Unknown
	// "internal" Is a VRChat Developer
	// "moderator" Is a VRChat Moderator
	//
	// Staff can hide their developerType at will.
	DeveloperType  DeveloperType `json:"developerType"`
	DisplayName    string        `json:"displayName"`
	EmailVerified  bool          `json:"emailVerified"`
	FallbackAvatar AvatarId      `json:"fallbackAvatar,omitempty"`

	// FriendGroupNames Always empty array.
	FriendGroupNames          []string `json:"friendGroupNames"`
	FriendKey                 string   `json:"friendKey"`
	Friends                   []UserId `json:"friends"`
	GoogleDetails             any      `json:"googleDetails,omitempty"`
	GoogleId                  string   `json:"googleId,omitempty"`
	HasBirthday               bool     `json:"hasBirthday"`
	HasEmail                  bool     `json:"hasEmail"`
	HasLoggedInFromClient     bool     `json:"hasLoggedInFromClient"`
	HasPendingEmail           bool     `json:"hasPendingEmail"`
	HideContentFilterSettings bool     `json:"hideContentFilterSettings,omitempty"`

	// HomeLocation WorldID be "offline" on User profiles if you are not friends with that user.
	HomeLocation WorldId `json:"homeLocation"`

	// Id A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id               UserId    `json:"id"`
	IsBoopingEnabled bool      `json:"isBoopingEnabled,omitempty"`
	IsFriend         bool      `json:"isFriend"`
	LastActivity     time.Time `json:"last_activity,omitempty"`
	LastLogin        time.Time `json:"last_login"`
	LastMobile       time.Time `json:"last_mobile"`

	// LastPlatform This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	LastPlatform           Platform `json:"last_platform"`
	ObfuscatedEmail        string   `json:"obfuscatedEmail"`
	ObfuscatedPendingEmail string   `json:"obfuscatedPendingEmail"`
	OculusId               string   `json:"oculusId"`
	OfflineFriends         []UserId `json:"offlineFriends,omitempty"`
	OnlineFriends          []UserId `json:"onlineFriends,omitempty"`

	// PastDisplayNames
	PastDisplayNames            []PastDisplayName   `json:"pastDisplayNames"`
	PicoId                      string              `json:"picoId,omitempty"`
	Presence                    CurrentUserPresence `json:"presence,omitempty"`
	ProfilePicOverride          string              `json:"profilePicOverride"`
	ProfilePicOverrideThumbnail string              `json:"profilePicOverrideThumbnail"`
	Pronouns                    string              `json:"pronouns"`

	// State * "online" User is online in VRChat
	// * "active" User is online, but not in VRChat
	// * "offline" User is offline
	//
	// Always offline when returned through `getCurrentUser` (/auth/user).
	State UserState `json:"state"`

	// Status Defines the User's current status, for example "ask me", "join me" or "offline. This status is a combined indicator of their online activity and privacy preference.
	Status                   UserStatus `json:"status"`
	StatusDescription        string     `json:"statusDescription"`
	StatusFirstTime          bool       `json:"statusFirstTime"`
	StatusHistory            []string   `json:"statusHistory"`
	SteamDetails             any        `json:"steamDetails"`
	SteamId                  string     `json:"steamId"`
	Tags                     []Tag      `json:"tags"`
	TwoFactorAuthEnabled     bool       `json:"twoFactorAuthEnabled"`
	TwoFactorAuthEnabledDate time.Time  `json:"twoFactorAuthEnabledDate,omitempty"`
	Unsubscribe              bool       `json:"unsubscribe"`
	UpdatedAt                time.Time  `json:"updated_at,omitempty"`
	UserIcon                 string     `json:"userIcon"`
	UserLanguage             string     `json:"userLanguage,omitempty"`
	UserLanguageCode         string     `json:"userLanguageCode,omitempty"`

	// Username -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429).
	Username string `json:"username,omitempty"`
	ViveId   string `json:"viveId,omitempty"`
}

type TwoFactorAuthCode struct {
	Code string `json:"code"`
}

type Verify2FaResult struct {
	Verified bool `json:"verified"`
}

type TwoFactorEmailCode struct {
	Code string `json:"code"`
}

type Verify2FaEmailCodeResult struct {
	Verified bool `json:"verified"`
}

type VerifyAuthTokenResult struct {
	Ok    bool   `json:"ok"`
	Token string `json:"token"`
}

type Success struct {
	Success Response `json:"success,omitempty"`
}

type ReleaseStatus string

const (
	ReleaseStatusPublic  ReleaseStatus = "public"
	ReleaseStatusPrivate ReleaseStatus = "private"
	ReleaseStatusHidden  ReleaseStatus = "hidden"
	ReleaseStatusAll     ReleaseStatus = "all"
)

type UnityPackageId string

type UnityPackage struct {
	AssetUrl       string         `json:"assetUrl,omitempty"`
	AssetUrlObject any            `json:"assetUrlObject,omitempty"`
	AssetVersion   int64          `json:"assetVersion"`
	CreatedAt      time.Time      `json:"created_at,omitempty"`
	Id             UnityPackageId `json:"id"`
	ImpostorUrl    string         `json:"impostorUrl,omitempty"`

	// Platform This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	Platform        Platform `json:"platform"`
	PluginUrl       string   `json:"pluginUrl,omitempty"`
	PluginUrlObject any      `json:"pluginUrlObject,omitempty"`
	ScanStatus      string   `json:"scanStatus,omitempty"`
	UnitySortNumber int64    `json:"unitySortNumber,omitempty"`
	UnityVersion    string   `json:"unityVersion"`
	Variant         string   `json:"variant,omitempty"`
}

type Avatar struct {
	// AssetUrl Not present from general serach `/avatars`, only on specific requests `/avatars/{avatarId}`.
	AssetUrl string `json:"assetUrl,omitempty"`

	// AssetUrlObject Not present from general serach `/avatars`, only on specific requests `/avatars/{avatarId}`.
	// **Deprecation:** `Object` has unknown usage/fields, and is always empty. Use normal `Url` field instead.
	AssetUrlObject any `json:"assetUrlObject,omitempty"`

	// AuthorId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId      UserId        `json:"authorId"`
	AuthorName    string        `json:"authorName"`
	CreatedAt     time.Time     `json:"created_at"`
	Description   string        `json:"description"`
	Featured      bool          `json:"featured"`
	Id            AvatarId      `json:"id"`
	ImageUrl      string        `json:"imageUrl"`
	Name          string        `json:"name"`
	ReleaseStatus ReleaseStatus `json:"releaseStatus"`

	// Tags
	Tags              []Tag  `json:"tags"`
	ThumbnailImageUrl string `json:"thumbnailImageUrl"`
	UnityPackageUrl   string `json:"unityPackageUrl"`

	// UnityPackageUrlObject **Deprecation:** `Object` has unknown usage/fields, and is always empty. Use normal `Url` field instead.

	UnityPackageUrlObject struct {
		UnityPackageUrl string `json:"unityPackageUrl,omitempty"`
	} `json:"unityPackageUrlObject"`
	UnityPackages []UnityPackage `json:"unityPackages"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Version       int64          `json:"version"`
}

type SortOption string

const (
	SortOptionPopularity          SortOption = "popularity"
	SortOptionHeat                SortOption = "heat"
	SortOptionTrust               SortOption = "trust"
	SortOptionShuffle             SortOption = "shuffle"
	SortOptionRandom              SortOption = "random"
	SortOptionFavorites           SortOption = "favorites"
	SortOptionReportScore         SortOption = "reportScore"
	SortOptionReportCount         SortOption = "reportCount"
	SortOptionPublicationDate     SortOption = "publicationDate"
	SortOptionLabsPublicationDate SortOption = "labsPublicationDate"
	SortOptionCreated             SortOption = "created"
	SortOptionCreatedAt           SortOption = "_created_at"
	SortOptionUpdated             SortOption = "updated"
	SortOptionUpdatedAt           SortOption = "_updated_at"
	SortOptionOrder               SortOption = "order"
	SortOptionRelevance           SortOption = "relevance"
	SortOptionMagic               SortOption = "magic"
	SortOptionName                SortOption = "name"
)

type OrderOption string

const (
	OrderOptionAscending  OrderOption = "ascending"
	OrderOptionDescending OrderOption = "descending"
)

type CreateAvatarRequest struct {
	AssetUrl      string        `json:"assetUrl,omitempty"`
	Description   string        `json:"description,omitempty"`
	Id            AvatarId      `json:"id,omitempty"`
	ImageUrl      string        `json:"imageUrl"`
	Name          string        `json:"name"`
	ReleaseStatus ReleaseStatus `json:"releaseStatus,omitempty"`

	// Tags
	Tags            []Tag   `json:"tags,omitempty"`
	UnityPackageUrl string  `json:"unityPackageUrl,omitempty"`
	UnityVersion    string  `json:"unityVersion,omitempty"`
	Version         float64 `json:"version,omitempty"`
}

type UpdateAvatarRequest struct {
	AssetUrl      string        `json:"assetUrl,omitempty"`
	Description   string        `json:"description,omitempty"`
	Id            AvatarId      `json:"id,omitempty"`
	ImageUrl      string        `json:"imageUrl,omitempty"`
	Name          string        `json:"name,omitempty"`
	ReleaseStatus ReleaseStatus `json:"releaseStatus,omitempty"`

	// Tags
	Tags            []Tag   `json:"tags,omitempty"`
	UnityPackageUrl string  `json:"unityPackageUrl,omitempty"`
	UnityVersion    string  `json:"unityVersion,omitempty"`
	Version         float64 `json:"version,omitempty"`
}

type TransactionId string

type TransactionStatus string

const (
	TransactionStatusActive     TransactionStatus = "active"
	TransactionStatusFailed     TransactionStatus = "failed"
	TransactionStatusExpired    TransactionStatus = "expired"
	TransactionStatusChargeback TransactionStatus = "chargeback"
)

type SubscriptionPeriod string

const (
	SubscriptionPeriodHour  SubscriptionPeriod = "hour"
	SubscriptionPeriodDay   SubscriptionPeriod = "day"
	SubscriptionPeriodWeek  SubscriptionPeriod = "week"
	SubscriptionPeriodMonth SubscriptionPeriod = "month"
	SubscriptionPeriodYear  SubscriptionPeriod = "year"
)

type Subscription struct {
	Amount          float64            `json:"amount"`
	Description     string             `json:"description"`
	GooglePlanId    string             `json:"googlePlanId,omitempty"`
	GoogleProductId string             `json:"googleProductId,omitempty"`
	Id              string             `json:"id"`
	OculusSku       string             `json:"oculusSku,omitempty"`
	Period          SubscriptionPeriod `json:"period"`
	PicoSku         string             `json:"picoSku,omitempty"`
	SteamItemId     string             `json:"steamItemId"`
	Tier            float64            `json:"tier"`
}

type TransactionSteamWalletInfo struct {
	Country  string `json:"country"`
	Currency string `json:"currency"`
	State    string `json:"state"`
	Status   string `json:"status"`
}

type TransactionSteamInfo struct {
	// OrderId Steam Order ID
	OrderId string `json:"orderId"`

	// SteamId Steam User ID
	SteamId string `json:"steamId"`

	// SteamUrl Empty
	SteamUrl string `json:"steamUrl"`

	// TransId Steam Transaction ID, NOT the same as VRChat TransactionID
	TransId    string                     `json:"transId"`
	WalletInfo TransactionSteamWalletInfo `json:"walletInfo"`
}

// TransactionAgreement Represents a single Transaction, which is likely between VRChat and Steam.
type TransactionAgreement struct {
	Agreement      string  `json:"agreement"`
	AgreementId    string  `json:"agreementId"`
	BillingType    string  `json:"billingType"`
	Currency       string  `json:"currency"`
	EndDate        string  `json:"endDate"`
	FailedAttempts float64 `json:"failedAttempts"`
	Frequency      float64 `json:"frequency"`
	ItemId         float64 `json:"itemId"`
	LastAmount     float64 `json:"lastAmount"`
	LastAmountVat  float64 `json:"lastAmountVat"`
	LastPayment    string  `json:"lastPayment"`
	NextPayment    string  `json:"nextPayment"`
	Outstanding    float64 `json:"outstanding"`
	Period         string  `json:"period"`
	RecurringAmt   float64 `json:"recurringAmt"`
	StartDate      string  `json:"startDate"`

	// Status This is NOT TransactionStatus, but whatever Steam return.
	Status      string `json:"status"`
	TimeCreated string `json:"timeCreated"`
}

type Transaction struct {
	// Agreement Represents a single Transaction, which is likely between VRChat and Steam.
	Agreement       TransactionAgreement `json:"agreement,omitempty"`
	CreatedAt       time.Time            `json:"created_at"`
	Error           string               `json:"error"`
	Id              TransactionId        `json:"id"`
	IsGift          bool                 `json:"isGift,omitempty"`
	IsTokens        bool                 `json:"isTokens,omitempty"`
	Sandbox         bool                 `json:"sandbox"`
	Status          TransactionStatus    `json:"status"`
	Steam           TransactionSteamInfo `json:"steam,omitempty"`
	Subscription    Subscription         `json:"subscription"`
	UpdatedAt       time.Time            `json:"updated_at"`
	UserDisplayName string               `json:"userDisplayName,omitempty"`

	// UserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	UserId UserId `json:"userId,omitempty"`
}

type LicenseGroupId string

type UserSubscription struct {
	Active        bool               `json:"active"`
	Amount        float64            `json:"amount"`
	CreatedAt     time.Time          `json:"created_at"`
	Description   string             `json:"description"`
	Expires       time.Time          `json:"expires"`
	Id            string             `json:"id"`
	IsGift        bool               `json:"isGift"`
	LicenseGroups []LicenseGroupId   `json:"licenseGroups"`
	Period        SubscriptionPeriod `json:"period"`
	Starts        string             `json:"starts,omitempty"`
	Status        TransactionStatus  `json:"status"`
	SteamItemId   string             `json:"steamItemId,omitempty"`

	// Store Which "Store" it came from. Right now only Stores are "Steam" and "Admin".
	Store         string        `json:"store"`
	Tier          float64       `json:"tier"`
	TransactionId TransactionId `json:"transactionId"`
	UpdatedAt     time.Time     `json:"updated_at"`
}

type LicenseType string

const (
	LicenseTypeAvatar       LicenseType = "avatar"
	LicenseTypeLicenseGroup LicenseType = "licenseGroup"
	LicenseTypePermission   LicenseType = "permission"
	LicenseTypeProduct      LicenseType = "product"
)

type LicenseAction string

const (
	LicenseActionWear LicenseAction = "wear"
	LicenseActionHave LicenseAction = "have"
)

type License struct {
	ForAction LicenseAction `json:"forAction"`

	// ForId Either a AvatarID, LicenseGroupID, PermissionID or ProductID. This depends on the `forType` field.
	ForId   string      `json:"forId"`
	ForName string      `json:"forName"`
	ForType LicenseType `json:"forType"`
}

type LicenseGroup struct {
	Description string         `json:"description"`
	Id          LicenseGroupId `json:"id"`
	Licenses    []License      `json:"licenses"`
	Name        string         `json:"name"`
}

type FavoriteId string

type FavoriteType string

const (
	FavoriteTypeWorld  FavoriteType = "world"
	FavoriteTypeFriend FavoriteType = "friend"
	FavoriteTypeAvatar FavoriteType = "avatar"
)

type Favorite struct {
	// FavoriteId MUST be either AvatarID, UserID or WorldID.
	FavoriteId string     `json:"favoriteId"`
	Id         FavoriteId `json:"id"`

	// Tags
	Tags []Tag        `json:"tags"`
	Type FavoriteType `json:"type"`
}

type AddFavoriteRequest struct {
	// FavoriteId Must be either AvatarID, WorldID or UserID.
	FavoriteId string `json:"favoriteId"`

	// Tags Tags indicate which group this favorite belongs to. Adding multiple groups makes it show up in all. Removing it from one in that case removes it from all.
	Tags []Tag        `json:"tags"`
	Type FavoriteType `json:"type"`
}

type FavoriteGroupId string

type FavoriteGroupVisibility string

const (
	FavoriteGroupVisibilityPrivate FavoriteGroupVisibility = "private"
	FavoriteGroupVisibilityFriends FavoriteGroupVisibility = "friends"
	FavoriteGroupVisibilityPublic  FavoriteGroupVisibility = "public"
)

type FavoriteGroup struct {
	DisplayName      string          `json:"displayName"`
	Id               FavoriteGroupId `json:"id"`
	Name             string          `json:"name"`
	OwnerDisplayName string          `json:"ownerDisplayName"`

	// OwnerId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId UserId `json:"ownerId"`

	// Tags
	Tags       []Tag                   `json:"tags"`
	Type       FavoriteType            `json:"type"`
	Visibility FavoriteGroupVisibility `json:"visibility"`
}

type UpdateFavoriteGroupRequest struct {
	DisplayName string `json:"displayName,omitempty"`

	// Tags Tags on FavoriteGroups are believed to do nothing.
	Tags       []Tag                   `json:"tags,omitempty"`
	Visibility FavoriteGroupVisibility `json:"visibility,omitempty"`
}

type FileId string

type MimeType string

const (
	MimeTypeImageJpeg                  MimeType = "image/jpeg"
	MimeTypeImageJpg                   MimeType = "image/jpg"
	MimeTypeImagePng                   MimeType = "image/png"
	MimeTypeImageWebp                  MimeType = "image/webp"
	MimeTypeImageGif                   MimeType = "image/gif"
	MimeTypeImageBmp                   MimeType = "image/bmp"
	MimeTypeImageSvgXml                MimeType = "image/svg＋xml"
	MimeTypeImageTiff                  MimeType = "image/tiff"
	MimeTypeApplicationXAvatar         MimeType = "application/x-avatar"
	MimeTypeApplicationXWorld          MimeType = "application/x-world"
	MimeTypeApplicationGzip            MimeType = "application/gzip"
	MimeTypeApplicationXRsyncSignature MimeType = "application/x-rsync-signature"
	MimeTypeApplicationXRsyncDelta     MimeType = "application/x-rsync-delta"
	MimeTypeApplicationOctetStream     MimeType = "application/octet-stream"
)

type FileStatus string

const (
	FileStatusWaiting  FileStatus = "waiting"
	FileStatusComplete FileStatus = "complete"
	FileStatusNone     FileStatus = "none"
	FileStatusQueued   FileStatus = "queued"
)

type FileData struct { // Category enum
	Category    string     `json:"category"`
	FileName    string     `json:"fileName"`
	Md5         string     `json:"md5,omitempty"`
	SizeInBytes int64      `json:"sizeInBytes"`
	Status      FileStatus `json:"status"`
	UploadId    string     `json:"uploadId"`
	Url         string     `json:"url"`
}

type FileVersion struct {
	CreatedAt time.Time `json:"created_at"`

	// Deleted Usually only present if `true`
	Deleted   bool       `json:"deleted,omitempty"`
	Delta     FileData   `json:"delta,omitempty"`
	File      FileData   `json:"file,omitempty"`
	Signature FileData   `json:"signature,omitempty"`
	Status    FileStatus `json:"status"`

	// Version Incremental version counter, can only be increased.
	Version int64 `json:"version"`
}

type File struct {
	Extension string   `json:"extension"`
	Id        FileId   `json:"id"`
	MimeType  MimeType `json:"mimeType"`
	Name      string   `json:"name"`

	// OwnerId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId UserId `json:"ownerId"`

	// Tags
	Tags []Tag `json:"tags"`

	// Versions
	Versions []FileVersion `json:"versions"`
}

type CreateFileRequest struct {
	Extension string   `json:"extension"`
	MimeType  MimeType `json:"mimeType"`
	Name      string   `json:"name"`

	// Tags
	Tags []Tag `json:"tags,omitempty"`
}

type CreateFileVersionRequest struct {
	FileMd5              string  `json:"fileMd5,omitempty"`
	FileSizeInBytes      float64 `json:"fileSizeInBytes,omitempty"`
	SignatureMd5         string  `json:"signatureMd5"`
	SignatureSizeInBytes float64 `json:"signatureSizeInBytes"`
}

type FinishFileDataUploadRequest struct {
	// Etags Array of ETags uploaded.
	Etags []string `json:"etags,omitempty"`

	// MaxParts Always a zero in string form, despite how many parts uploaded.
	MaxParts string `json:"maxParts"`

	// NextPartNumber Always a zero in string form, despite how many parts uploaded.
	NextPartNumber string `json:"nextPartNumber"`
}

type FileUploadUrl struct {
	Url string `json:"url"`
}

type FileVersionUploadStatus struct {
	// Etags Unknown
	Etags          []any   `json:"etags"`
	FileName       string  `json:"fileName"`
	MaxParts       float64 `json:"maxParts"`
	NextPartNumber float64 `json:"nextPartNumber"`
	Parts          []any   `json:"parts"`
	UploadId       string  `json:"uploadId"`
}

type LimitedUser struct {
	Bio string `json:"bio,omitempty"`

	// BioLinks
	BioLinks []string `json:"bioLinks,omitempty"`

	// CurrentAvatarImageUrl When profilePicOverride is not empty, use it instead.
	CurrentAvatarImageUrl CurrentAvatarImageUrl `json:"currentAvatarImageUrl,omitempty"`
	CurrentAvatarTags     []Tag                 `json:"currentAvatarTags,omitempty"`

	// CurrentAvatarThumbnailImageUrl When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl CurrentAvatarThumbnailImageUrl `json:"currentAvatarThumbnailImageUrl,omitempty"`

	// DeveloperType "none" User is a normal user
	// "trusted" Unknown
	// "internal" Is a VRChat Developer
	// "moderator" Is a VRChat Moderator
	//
	// Staff can hide their developerType at will.
	DeveloperType  DeveloperType `json:"developerType"`
	DisplayName    string        `json:"displayName"`
	FallbackAvatar AvatarId      `json:"fallbackAvatar,omitempty"`
	FriendKey      string        `json:"friendKey,omitempty"`

	// Id A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id       UserId `json:"id"`
	IsFriend bool   `json:"isFriend"`

	// LastPlatform This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	LastPlatform       Platform `json:"last_platform"`
	Location           string   `json:"location,omitempty"`
	ProfilePicOverride string   `json:"profilePicOverride,omitempty"`
	Pronouns           string   `json:"pronouns,omitempty"`

	// Status Defines the User's current status, for example "ask me", "join me" or "offline. This status is a combined indicator of their online activity and privacy preference.
	Status            UserStatus `json:"status"`
	StatusDescription string     `json:"statusDescription"`

	// Tags <- Always empty.
	Tags     []Tag  `json:"tags"`
	UserIcon string `json:"userIcon,omitempty"`

	// Username -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429).
	Username string `json:"username,omitempty"`
}

type NotificationType string

const (
	NotificationTypeFriendRequest         NotificationType = "friendRequest"
	NotificationTypeInvite                NotificationType = "invite"
	NotificationTypeInviteResponse        NotificationType = "inviteResponse"
	NotificationTypeMessage               NotificationType = "message"
	NotificationTypeRequestInvite         NotificationType = "requestInvite"
	NotificationTypeRequestInviteResponse NotificationType = "requestInviteResponse"
	NotificationTypeVotetokick            NotificationType = "votetokick"
)

type Notification struct {
	CreatedAt time.Time `json:"created_at"`

	// Details **NOTICE:** This is not a JSON object when received from the REST API, but it is when received from the Websocket API. When received from the REST API, this is a json **encoded** object, meaning you have to json-de-encode to get the NotificationDetail object depending on the NotificationType.
	Details string `json:"details"`
	Id      string `json:"id"`
	Message string `json:"message"`

	// ReceiverUserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	ReceiverUserId UserId `json:"receiverUserId,omitempty"`

	// Seen Not included in notification objects received from the Websocket API
	Seen bool `json:"seen,omitempty"`

	// SenderUserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	SenderUserId UserId `json:"senderUserId"`

	// SenderUsername -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429).
	SenderUsername string           `json:"senderUsername,omitempty"`
	Type           NotificationType `json:"type"`
}

type FriendStatus struct {
	IncomingRequest bool `json:"incomingRequest"`
	IsFriend        bool `json:"isFriend"`
	OutgoingRequest bool `json:"outgoingRequest"`
}

type GroupShortCode string

type GroupDiscriminator string

type GroupMemberStatus string

const (
	GroupMemberStatusInactive    GroupMemberStatus = "inactive"
	GroupMemberStatusMember      GroupMemberStatus = "member"
	GroupMemberStatusRequested   GroupMemberStatus = "requested"
	GroupMemberStatusInvited     GroupMemberStatus = "invited"
	GroupMemberStatusBanned      GroupMemberStatus = "banned"
	GroupMemberStatusUserblocked GroupMemberStatus = "userblocked"
)

type GroupGalleryId string

type GroupRoleId string

type GroupGallery struct {
	CreatedAt time.Time `json:"createdAt,omitempty"`

	// Description Description of the gallery.
	Description string         `json:"description,omitempty"`
	Id          GroupGalleryId `json:"id,omitempty"`

	// MembersOnly Whether the gallery is members only.
	MembersOnly bool `json:"membersOnly,omitempty"`

	// Name Name of the gallery.
	Name string `json:"name,omitempty"`

	// RoleIdsToAutoApprove
	RoleIdsToAutoApprove []GroupRoleId `json:"roleIdsToAutoApprove,omitempty"`

	// RoleIdsToManage
	RoleIdsToManage []GroupRoleId `json:"roleIdsToManage,omitempty"`

	// RoleIdsToSubmit
	RoleIdsToSubmit []GroupRoleId `json:"roleIdsToSubmit,omitempty"`

	// RoleIdsToView
	RoleIdsToView []GroupRoleId `json:"roleIdsToView,omitempty"`
	UpdatedAt     time.Time     `json:"updatedAt,omitempty"`
}

type LimitedGroup struct {
	BannerId      string             `json:"bannerId,omitempty"`
	BannerUrl     string             `json:"bannerUrl,omitempty"`
	CreatedAt     time.Time          `json:"createdAt,omitempty"`
	Description   string             `json:"description,omitempty"`
	Discriminator GroupDiscriminator `json:"discriminator,omitempty"`

	// Galleries
	Galleries        []GroupGallery    `json:"galleries,omitempty"`
	IconId           string            `json:"iconId,omitempty"`
	IconUrl          string            `json:"iconUrl,omitempty"`
	Id               GroupId           `json:"id,omitempty"`
	IsSearchable     bool              `json:"isSearchable,omitempty"`
	MemberCount      int64             `json:"memberCount,omitempty"`
	MembershipStatus GroupMemberStatus `json:"membershipStatus,omitempty"`
	Name             string            `json:"name,omitempty"`

	// OwnerId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId   UserId         `json:"ownerId,omitempty"`
	Rules     string         `json:"rules,omitempty"`
	ShortCode GroupShortCode `json:"shortCode,omitempty"`

	// Tags
	Tags []Tag `json:"tags,omitempty"`
}

type GroupJoinState string

const (
	GroupJoinStateClosed  GroupJoinState = "closed"
	GroupJoinStateInvite  GroupJoinState = "invite"
	GroupJoinStateRequest GroupJoinState = "request"
	GroupJoinStateOpen    GroupJoinState = "open"
)

type GroupPrivacy string

const (
	GroupPrivacyDefault GroupPrivacy = "default"
	GroupPrivacyPrivate GroupPrivacy = "private"
)

type GroupRoleTemplate string

const (
	GroupRoleTemplateDefault        GroupRoleTemplate = "default"
	GroupRoleTemplateManagedFree    GroupRoleTemplate = "managedFree"
	GroupRoleTemplateManagedInvite  GroupRoleTemplate = "managedInvite"
	GroupRoleTemplateManagedRequest GroupRoleTemplate = "managedRequest"
)

type CreateGroupRequest struct {
	BannerId     string            `json:"bannerId,omitempty"`
	Description  string            `json:"description,omitempty"`
	IconId       string            `json:"iconId,omitempty"`
	JoinState    GroupJoinState    `json:"joinState,omitempty"`
	Name         string            `json:"name"`
	Privacy      GroupPrivacy      `json:"privacy,omitempty"`
	RoleTemplate GroupRoleTemplate `json:"roleTemplate"`
	ShortCode    string            `json:"shortCode"`
}

type GroupMemberId string

type GroupMyMember struct {
	AcceptedByDisplayName string `json:"acceptedByDisplayName,omitempty"`

	// AcceptedById A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AcceptedById                UserId        `json:"acceptedById,omitempty"`
	BannedAt                    string        `json:"bannedAt,omitempty"`
	CreatedAt                   time.Time     `json:"createdAt,omitempty"`
	GroupId                     GroupId       `json:"groupId,omitempty"`
	Has2Fa                      bool          `json:"has2FA,omitempty"`
	HasJoinedFromPurchase       bool          `json:"hasJoinedFromPurchase,omitempty"`
	Id                          GroupMemberId `json:"id,omitempty"`
	IsRepresenting              bool          `json:"isRepresenting,omitempty"`
	IsSubscribedToAnnouncements bool          `json:"isSubscribedToAnnouncements,omitempty"`
	JoinedAt                    time.Time     `json:"joinedAt,omitempty"`
	LastPostReadAt              time.Time     `json:"lastPostReadAt,omitempty"`
	MRoleIds                    []string      `json:"mRoleIds,omitempty"`
	ManagerNotes                string        `json:"managerNotes,omitempty"`
	MembershipStatus            string        `json:"membershipStatus,omitempty"`
	Permissions                 []string      `json:"permissions,omitempty"`
	RoleIds                     []GroupRoleId `json:"roleIds,omitempty"`

	// UserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	UserId     UserId `json:"userId,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

type GroupRole struct {
	CreatedAt         time.Time   `json:"createdAt,omitempty"`
	Description       string      `json:"description,omitempty"`
	GroupId           GroupId     `json:"groupId,omitempty"`
	Id                GroupRoleId `json:"id,omitempty"`
	IsManagementRole  bool        `json:"isManagementRole,omitempty"`
	IsSelfAssignable  bool        `json:"isSelfAssignable,omitempty"`
	Name              string      `json:"name,omitempty"`
	Order             int64       `json:"order,omitempty"`
	Permissions       []string    `json:"permissions,omitempty"`
	RequiresPurchase  bool        `json:"requiresPurchase,omitempty"`
	RequiresTwoFactor bool        `json:"requiresTwoFactor,omitempty"`
	UpdatedAt         time.Time   `json:"updatedAt,omitempty"`
}

type Group struct {
	BannerId            string             `json:"bannerId,omitempty"`
	BannerUrl           string             `json:"bannerUrl,omitempty"`
	CreatedAt           time.Time          `json:"createdAt,omitempty"`
	Description         string             `json:"description,omitempty"`
	Discriminator       GroupDiscriminator `json:"discriminator,omitempty"`
	Galleries           []GroupGallery     `json:"galleries,omitempty"`
	IconId              string             `json:"iconId,omitempty"`
	IconUrl             string             `json:"iconUrl,omitempty"`
	Id                  GroupId            `json:"id,omitempty"`
	IsVerified          bool               `json:"isVerified,omitempty"`
	JoinState           GroupJoinState     `json:"joinState,omitempty"`
	Languages           []string           `json:"languages,omitempty"`
	LastPostCreatedAt   time.Time          `json:"lastPostCreatedAt,omitempty"`
	Links               []string           `json:"links,omitempty"`
	MemberCount         int64              `json:"memberCount,omitempty"`
	MemberCountSyncedAt time.Time          `json:"memberCountSyncedAt,omitempty"`
	MembershipStatus    GroupMemberStatus  `json:"membershipStatus,omitempty"`
	MyMember            GroupMyMember      `json:"myMember,omitempty"`
	Name                string             `json:"name,omitempty"`
	OnlineMemberCount   int64              `json:"onlineMemberCount,omitempty"`

	// OwnerId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId UserId       `json:"ownerId,omitempty"`
	Privacy GroupPrivacy `json:"privacy,omitempty"`

	// Roles Only returned if ?includeRoles=true is specified.
	Roles     []GroupRole    `json:"roles,omitempty"`
	Rules     string         `json:"rules,omitempty"`
	ShortCode GroupShortCode `json:"shortCode,omitempty"`
	Tags      []Tag          `json:"tags,omitempty"`

	// TransferTargetId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	TransferTargetId UserId    `json:"transferTargetId,omitempty"`
	UpdatedAt        time.Time `json:"updatedAt,omitempty"`
}

type UpdateGroupRequest struct {
	BannerId    string         `json:"bannerId,omitempty"`
	Description string         `json:"description,omitempty"`
	IconId      string         `json:"iconId,omitempty"`
	JoinState   GroupJoinState `json:"joinState,omitempty"`

	// Languages 3 letter language code
	Languages []string `json:"languages,omitempty"`
	Links     []string `json:"links,omitempty"`
	Name      string   `json:"name,omitempty"`
	Rules     string   `json:"rules,omitempty"`
	ShortCode string   `json:"shortCode,omitempty"`

	// Tags
	Tags []Tag `json:"tags,omitempty"`
}

type GroupAnnouncementId string

type GroupAnnouncement struct {
	// AuthorId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId  UserId              `json:"authorId,omitempty"`
	CreatedAt time.Time           `json:"createdAt,omitempty"`
	GroupId   GroupId             `json:"groupId,omitempty"`
	Id        GroupAnnouncementId `json:"id,omitempty"`
	ImageId   FileId              `json:"imageId,omitempty"`
	ImageUrl  string              `json:"imageUrl,omitempty"`
	Text      string              `json:"text,omitempty"`
	Title     string              `json:"title,omitempty"`
	UpdatedAt time.Time           `json:"updatedAt,omitempty"`
}

type CreateGroupAnnouncementRequest struct {
	ImageId FileId `json:"imageId,omitempty"`

	// SendNotification Send notification to group members.
	SendNotification bool `json:"sendNotification,omitempty"`

	// Text Announcement text
	Text string `json:"text,omitempty"`

	// Title Announcement title
	Title string `json:"title"`
}

type GroupAuditLogId string

type GroupAuditLogEntry struct {
	ActorDisplayName string `json:"actorDisplayName,omitempty"`

	// ActorId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	ActorId   UserId    `json:"actorId,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Data The data associated with the event. The format of this data is dependent on the event type.
	Data any `json:"data,omitempty"`

	// Description A human-readable description of the event.
	Description string `json:"description,omitempty"`

	// EventType The type of event that occurred. This is a string that is prefixed with the type of object that the event occurred on. For example, a group role update event would be prefixed with `group.role`.
	EventType string          `json:"eventType,omitempty"`
	GroupId   GroupId         `json:"groupId,omitempty"`
	Id        GroupAuditLogId `json:"id,omitempty"`

	// TargetId Typically GroupID or GroupRoleID, but could be other types of IDs.
	TargetId string `json:"targetId,omitempty"`
}

type PaginatedGroupAuditLogEntryList struct {
	// HasNext Whether there are more results after this page.
	HasNext bool `json:"hasNext,omitempty"`

	// Results
	Results []GroupAuditLogEntry `json:"results,omitempty"`

	// TotalCount The total number of results that the query would return if there were no pagination.
	TotalCount int64 `json:"totalCount,omitempty"`
}

// GroupMemberLimitedUser Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
type GroupMemberLimitedUser struct {
	CurrentAvatarTags              []Tag  `json:"currentAvatarTags,omitempty"`
	CurrentAvatarThumbnailImageUrl string `json:"currentAvatarThumbnailImageUrl,omitempty"`
	DisplayName                    string `json:"displayName,omitempty"`
	IconUrl                        string `json:"iconUrl,omitempty"`

	// Id A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id                 UserId `json:"id,omitempty"`
	ProfilePicOverride string `json:"profilePicOverride,omitempty"`
	ThumbnailUrl       string `json:"thumbnailUrl,omitempty"`
}

type GroupMember struct {
	// BannedAt Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
	BannedAt time.Time `json:"bannedAt,omitempty"`

	// CreatedAt Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
	CreatedAt             time.Time     `json:"createdAt,omitempty"`
	GroupId               GroupId       `json:"groupId,omitempty"`
	HasJoinedFromPurchase bool          `json:"hasJoinedFromPurchase,omitempty"`
	Id                    GroupMemberId `json:"id,omitempty"`

	// IsRepresenting Whether the user is representing the group. This makes the group show up above the name tag in-game.
	IsRepresenting              bool          `json:"isRepresenting,omitempty"`
	IsSubscribedToAnnouncements bool          `json:"isSubscribedToAnnouncements,omitempty"`
	JoinedAt                    time.Time     `json:"joinedAt,omitempty"`
	LastPostReadAt              time.Time     `json:"lastPostReadAt,omitempty"`
	MRoleIds                    []GroupRoleId `json:"mRoleIds,omitempty"`

	// ManagerNotes Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
	ManagerNotes     string            `json:"managerNotes,omitempty"`
	MembershipStatus GroupMemberStatus `json:"membershipStatus,omitempty"`
	RoleIds          []GroupRoleId     `json:"roleIds,omitempty"`

	// User Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
	User GroupMemberLimitedUser `json:"user,omitempty"`

	// UserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	UserId     UserId `json:"userId,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

type BanGroupMemberRequest struct {
	// UserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	UserId UserId `json:"userId"`
}

type CreateGroupGalleryRequest struct {
	// Description Description of the gallery.
	Description string `json:"description,omitempty"`

	// MembersOnly Whether the gallery is members only.
	MembersOnly bool `json:"membersOnly,omitempty"`

	// Name Name of the gallery.
	Name string `json:"name"`

	// RoleIdsToAutoApprove
	RoleIdsToAutoApprove []GroupRoleId `json:"roleIdsToAutoApprove,omitempty"`

	// RoleIdsToManage
	RoleIdsToManage []GroupRoleId `json:"roleIdsToManage,omitempty"`

	// RoleIdsToSubmit
	RoleIdsToSubmit []GroupRoleId `json:"roleIdsToSubmit,omitempty"`

	// RoleIdsToView
	RoleIdsToView []GroupRoleId `json:"roleIdsToView,omitempty"`
}

type GroupGalleryImageId string

type GroupGalleryImage struct {
	Approved   bool      `json:"approved,omitempty"`
	ApprovedAt time.Time `json:"approvedAt,omitempty"`

	// ApprovedByUserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	ApprovedByUserId UserId              `json:"approvedByUserId,omitempty"`
	CreatedAt        time.Time           `json:"createdAt,omitempty"`
	FileId           FileId              `json:"fileId,omitempty"`
	GalleryId        GroupGalleryId      `json:"galleryId,omitempty"`
	GroupId          GroupId             `json:"groupId,omitempty"`
	Id               GroupGalleryImageId `json:"id,omitempty"`
	ImageUrl         string              `json:"imageUrl,omitempty"`

	// SubmittedByUserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	SubmittedByUserId UserId `json:"submittedByUserId,omitempty"`
}

type UpdateGroupGalleryRequest struct {
	// Description Description of the gallery.
	Description string `json:"description,omitempty"`

	// MembersOnly Whether the gallery is members only.
	MembersOnly bool `json:"membersOnly,omitempty"`

	// Name Name of the gallery.
	Name string `json:"name,omitempty"`

	// RoleIdsToAutoApprove
	RoleIdsToAutoApprove []GroupRoleId `json:"roleIdsToAutoApprove,omitempty"`

	// RoleIdsToManage
	RoleIdsToManage []GroupRoleId `json:"roleIdsToManage,omitempty"`

	// RoleIdsToSubmit
	RoleIdsToSubmit []GroupRoleId `json:"roleIdsToSubmit,omitempty"`

	// RoleIdsToView
	RoleIdsToView []GroupRoleId `json:"roleIdsToView,omitempty"`
}

type AddGroupGalleryImageRequest struct {
	FileId FileId `json:"fileId"`
}

// InstanceId InstanceID can be "offline" on User profiles if you are not friends with that user and "private" if you are friends and user is in private instance.
type InstanceId string

// UdonProductId A unique ID of a Udon Product
type UdonProductId string

type World struct {
	// AuthorId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId    UserId    `json:"authorId"`
	AuthorName  string    `json:"authorName"`
	Capacity    int64     `json:"capacity"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	Favorites   int64     `json:"favorites,omitempty"`
	Featured    bool      `json:"featured"`
	Heat        int64     `json:"heat"`

	// Id WorldID be "offline" on User profiles if you are not friends with that user.
	Id       WorldId `json:"id"`
	ImageUrl string  `json:"imageUrl"`

	// Instances Will always be an empty list when unauthenticated.
	Instances           [][]any `json:"instances,omitempty"`
	LabsPublicationDate string  `json:"labsPublicationDate"`
	Name                string  `json:"name"`
	Namespace           string  `json:"namespace,omitempty"`

	// Occupants Will always be `0` when unauthenticated.
	Occupants        int64  `json:"occupants,omitempty"`
	Organization     string `json:"organization"`
	Popularity       int64  `json:"popularity"`
	PreviewYoutubeId string `json:"previewYoutubeId,omitempty"`

	// PrivateOccupants Will always be `0` when unauthenticated.
	PrivateOccupants int64 `json:"privateOccupants,omitempty"`

	// PublicOccupants Will always be `0` when unauthenticated.
	PublicOccupants     int64         `json:"publicOccupants,omitempty"`
	PublicationDate     string        `json:"publicationDate"`
	RecommendedCapacity int64         `json:"recommendedCapacity"`
	ReleaseStatus       ReleaseStatus `json:"releaseStatus"`

	// Tags
	Tags              []Tag           `json:"tags"`
	ThumbnailImageUrl string          `json:"thumbnailImageUrl"`
	UdonProducts      []UdonProductId `json:"udonProducts,omitempty"`

	// UnityPackages Empty if unauthenticated.
	UnityPackages []UnityPackage `json:"unityPackages,omitempty"`
	UpdatedAt     time.Time      `json:"updated_at"`
	Version       int64          `json:"version"`
	Visits        int64          `json:"visits"`
}

type GroupInstance struct {
	InstanceId string `json:"instanceId"`

	// Location InstanceID can be "offline" on User profiles if you are not friends with that user and "private" if you are friends and user is in private instance.
	Location    InstanceId `json:"location"`
	MemberCount int64      `json:"memberCount"`
	World       World      `json:"world"`
}

type CreateGroupInviteRequest struct {
	ConfirmOverrideBlock bool `json:"confirmOverrideBlock,omitempty"`

	// UserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	UserId UserId `json:"userId"`
}

type GroupSearchSort string

const (
	GroupSearchSortJoinedAtAsc  GroupSearchSort = "joinedAt:asc"
	GroupSearchSortJoinedAtDesc GroupSearchSort = "joinedAt:desc"
)

type GroupLimitedMember struct {
	// BannedAt Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
	BannedAt time.Time `json:"bannedAt,omitempty"`

	// CreatedAt Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
	CreatedAt             time.Time     `json:"createdAt,omitempty"`
	GroupId               GroupId       `json:"groupId,omitempty"`
	HasJoinedFromPurchase bool          `json:"hasJoinedFromPurchase,omitempty"`
	Id                    GroupMemberId `json:"id,omitempty"`

	// IsRepresenting Whether the user is representing the group. This makes the group show up above the name tag in-game.
	IsRepresenting              bool          `json:"isRepresenting,omitempty"`
	IsSubscribedToAnnouncements bool          `json:"isSubscribedToAnnouncements,omitempty"`
	JoinedAt                    time.Time     `json:"joinedAt,omitempty"`
	LastPostReadAt              time.Time     `json:"lastPostReadAt,omitempty"`
	MRoleIds                    []GroupRoleId `json:"mRoleIds,omitempty"`

	// ManagerNotes Only visible via the /groups/:groupId/members endpoint, **not** when fetching a specific user.
	ManagerNotes     string            `json:"managerNotes,omitempty"`
	MembershipStatus GroupMemberStatus `json:"membershipStatus,omitempty"`
	RoleIds          []GroupRoleId     `json:"roleIds,omitempty"`

	// UserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	UserId     UserId `json:"userId,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

type GroupUserVisibility string

const (
	GroupUserVisibilityVisible GroupUserVisibility = "visible"
	GroupUserVisibilityHidden  GroupUserVisibility = "hidden"
	GroupUserVisibilityFriends GroupUserVisibility = "friends"
)

type UpdateGroupMemberRequest struct {
	IsSubscribedToAnnouncements bool                `json:"isSubscribedToAnnouncements,omitempty"`
	ManagerNotes                string              `json:"managerNotes,omitempty"`
	Visibility                  GroupUserVisibility `json:"visibility,omitempty"`
}

// GroupRoleIdList
type GroupRoleIdList []GroupRoleId

// GroupPermission A permission that can be granted to a role in a group.
type GroupPermission struct {
	// AllowedToAdd Whether the user is allowed to add this permission to a role.
	AllowedToAdd bool `json:"allowedToAdd,omitempty"`

	// DisplayName The display name of the permission.
	DisplayName string `json:"displayName,omitempty"`

	// Help Human-readable description of the permission.
	Help string `json:"help,omitempty"`

	// IsManagementPermission Whether this permission is a "management" permission.
	IsManagementPermission bool `json:"isManagementPermission,omitempty"`

	// Name The name of the permission.
	Name string `json:"name,omitempty"`
}

type NotificationId string

type GroupPostVisibility string

const (
	GroupPostVisibilityGroup  GroupPostVisibility = "group"
	GroupPostVisibilityPublic GroupPostVisibility = "public"
)

type GroupPost struct {
	// AuthorId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId  UserId    `json:"authorId,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`

	// EditorId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	EditorId UserId         `json:"editorId,omitempty"`
	GroupId  GroupId        `json:"groupId,omitempty"`
	Id       NotificationId `json:"id,omitempty"`
	ImageId  FileId         `json:"imageId,omitempty"`
	ImageUrl string         `json:"imageUrl,omitempty"`

	// RoleId
	RoleId     GroupRoleIdList     `json:"roleId,omitempty"`
	Text       string              `json:"text,omitempty"`
	Title      string              `json:"title,omitempty"`
	UpdatedAt  time.Time           `json:"updatedAt,omitempty"`
	Visibility GroupPostVisibility `json:"visibility,omitempty"`
}

type CreateGroupPostRequest struct {
	ImageId FileId `json:"imageId,omitempty"`

	// RoleIds
	RoleIds GroupRoleIdList `json:"roleIds,omitempty"`

	// SendNotification Send notification to group members.
	SendNotification bool `json:"sendNotification"`

	// Text Post text
	Text string `json:"text"`

	// Title Post title
	Title      string              `json:"title"`
	Visibility GroupPostVisibility `json:"visibility"`
}

type GroupJoinRequestAction string

const (
	GroupJoinRequestActionAccept GroupJoinRequestAction = "accept"
	GroupJoinRequestActionReject GroupJoinRequestAction = "reject"
)

type RespondGroupJoinRequest struct {
	Action GroupJoinRequestAction `json:"action"`

	// Block Whether to block the user from requesting again
	Block bool `json:"block,omitempty"`
}

type CreateGroupRoleRequest struct {
	Description      string   `json:"description,omitempty"`
	Id               string   `json:"id,omitempty"`
	IsSelfAssignable bool     `json:"isSelfAssignable,omitempty"`
	Name             string   `json:"name,omitempty"`
	Permissions      []string `json:"permissions,omitempty"`
}

type UpdateGroupRoleRequest struct {
	Description      string   `json:"description,omitempty"`
	IsSelfAssignable bool     `json:"isSelfAssignable,omitempty"`
	Name             string   `json:"name,omitempty"`
	Order            int64    `json:"order,omitempty"`
	Permissions      []string `json:"permissions,omitempty"`
}

type InviteRequest struct {
	// InstanceId InstanceID can be "offline" on User profiles if you are not friends with that user and "private" if you are friends and user is in private instance.
	InstanceId  InstanceId `json:"instanceId"`
	MessageSlot int64      `json:"messageSlot,omitempty"`
}

type SentNotification struct {
	CreatedAt time.Time `json:"created_at"`
	Details   any       `json:"details"`
	Id        string    `json:"id"`
	Message   string    `json:"message"`

	// ReceiverUserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	ReceiverUserId UserId `json:"receiverUserId"`

	// SenderUserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	SenderUserId UserId `json:"senderUserId"`

	// SenderUsername -| **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429).
	SenderUsername string           `json:"senderUsername,omitempty"`
	Type           NotificationType `json:"type"`
}

type RequestInviteRequest struct {
	MessageSlot int64 `json:"messageSlot,omitempty"`
}

type InviteResponse struct {
	ResponseSlot int64 `json:"responseSlot"`
}

type InviteMessageType string

const (
	InviteMessageTypeMessage         InviteMessageType = "message"
	InviteMessageTypeResponse        InviteMessageType = "response"
	InviteMessageTypeRequest         InviteMessageType = "request"
	InviteMessageTypeRequestResponse InviteMessageType = "requestResponse"
)

type InviteMessageId string

type InviteMessage struct {
	CanBeUpdated bool              `json:"canBeUpdated"`
	Id           InviteMessageId   `json:"id"`
	Message      string            `json:"message"`
	MessageType  InviteMessageType `json:"messageType"`

	// RemainingCooldownMinutes Changes to 60 when updated, although probably server-side configurable.
	RemainingCooldownMinutes int64     `json:"remainingCooldownMinutes"`
	Slot                     int64     `json:"slot"`
	UpdatedAt                time.Time `json:"updatedAt"`
}

type UpdateInviteMessageRequest struct {
	Message string `json:"message"`
}

type InstanceType string

const (
	InstanceTypePublic  InstanceType = "public"
	InstanceTypeHidden  InstanceType = "hidden"
	InstanceTypeFriends InstanceType = "friends"
	InstanceTypePrivate InstanceType = "private"
	InstanceTypeGroup   InstanceType = "group"
)

// InstanceRegion Instance region
type InstanceRegion string

const (
	InstanceRegionUs      InstanceRegion = "us"
	InstanceRegionUse     InstanceRegion = "use"
	InstanceRegionEu      InstanceRegion = "eu"
	InstanceRegionJp      InstanceRegion = "jp"
	InstanceRegionUnknown InstanceRegion = "unknown"
)

// InstanceOwnerId A groupId if the instance type is "group", null if instance type is public, or a userId otherwise
type InstanceOwnerId string

// GroupAccessType Group access type when the instance type is "group"
type GroupAccessType string

const (
	GroupAccessTypePublic  GroupAccessType = "public"
	GroupAccessTypePlus    GroupAccessType = "plus"
	GroupAccessTypeMembers GroupAccessType = "members"
)

type CreateInstanceRequest struct {
	// CanRequestInvite Only applies to invite type instances to make them invite+
	CanRequestInvite bool `json:"canRequestInvite,omitempty"`

	// ClosedAt The time after which users won't be allowed to join the instance. This doesn't work for public instances.
	ClosedAt time.Time `json:"closedAt,omitempty"`

	// GroupAccessType Group access type when the instance type is "group"
	GroupAccessType GroupAccessType `json:"groupAccessType,omitempty"`

	// HardClose Currently unused, but will eventually be a flag to set if the closing of the instance should kick people.
	HardClose  bool `json:"hardClose,omitempty"`
	InviteOnly bool `json:"inviteOnly,omitempty"`

	// OwnerId A groupId if the instance type is "group", null if instance type is public, or a userId otherwise
	OwnerId      InstanceOwnerId `json:"ownerId,omitempty"`
	QueueEnabled bool            `json:"queueEnabled,omitempty"`

	// Region Instance region
	Region InstanceRegion `json:"region"`

	// RoleIds Group roleIds that are allowed to join if the type is "group" and groupAccessType is "member"
	RoleIds []GroupRoleId `json:"roleIds,omitempty"`
	Type    InstanceType  `json:"type"`

	// WorldId WorldID be "offline" on User profiles if you are not friends with that user.
	WorldId WorldId `json:"worldId"`
}

// Region API/Photon region.
type Region string

const (
	RegionUs      Region = "us"
	RegionUse     Region = "use"
	RegionUsw     Region = "usw"
	RegionEu      Region = "eu"
	RegionJp      Region = "jp"
	RegionUnknown Region = "unknown"
)

type InstancePlatforms struct {
	Android           int64 `json:"android"`
	Standalonewindows int64 `json:"standalonewindows"`
}

// Instance * `hidden` field is only present if InstanceType is `hidden` aka "Friends+", and is instance creator.
// * `friends` field is only present if InstanceType is `friends` aka "Friends", and is instance creator.
// * `private` field is only present if InstanceType is `private` aka "Invite" or "Invite+", and is instance creator.
type Instance struct {
	Active           bool  `json:"active"`
	CanRequestInvite bool  `json:"canRequestInvite"`
	Capacity         int64 `json:"capacity"`

	// ClientNumber Always returns "unknown".
	ClientNumber string    `json:"clientNumber"`
	ClosedAt     time.Time `json:"closedAt,omitempty"`

	// Friends A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Friends UserId `json:"friends,omitempty"`
	Full    bool   `json:"full"`

	// GroupAccessType Group access type when the instance type is "group"
	GroupAccessType   GroupAccessType `json:"groupAccessType,omitempty"`
	HardClose         bool            `json:"hardClose,omitempty"`
	HasCapacityForYou bool            `json:"hasCapacityForYou,omitempty"`

	// Hidden A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Hidden UserId `json:"hidden,omitempty"`

	// Id InstanceID can be "offline" on User profiles if you are not friends with that user and "private" if you are friends and user is in private instance.
	Id         InstanceId `json:"id"`
	InstanceId string     `json:"instanceId"`

	// Location InstanceID can be "offline" on User profiles if you are not friends with that user and "private" if you are friends and user is in private instance.
	Location InstanceId `json:"location"`
	NUsers   int64      `json:"n_users"`
	Name     string     `json:"name"`
	Nonce    string     `json:"nonce,omitempty"`

	// OwnerId A groupId if the instance type is "group", null if instance type is public, or a userId otherwise
	OwnerId   InstanceOwnerId `json:"ownerId,omitempty"`
	Permanent bool            `json:"permanent"`

	// PhotonRegion API/Photon region.
	PhotonRegion Region            `json:"photonRegion"`
	Platforms    InstancePlatforms `json:"platforms"`

	// Private A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Private             UserId `json:"private,omitempty"`
	QueueEnabled        bool   `json:"queueEnabled"`
	QueueSize           int64  `json:"queueSize"`
	RecommendedCapacity int64  `json:"recommendedCapacity"`

	// Region Instance region
	Region         InstanceRegion `json:"region"`
	RoleRestricted bool           `json:"roleRestricted,omitempty"`
	SecureName     string         `json:"secureName"`
	ShortName      string         `json:"shortName,omitempty"`
	Strict         bool           `json:"strict"`

	// Tags The tags array on Instances usually contain the language tags of the people in the instance.
	Tags      []Tag        `json:"tags"`
	Type      InstanceType `json:"type"`
	UserCount int64        `json:"userCount"`

	// Users The users field is present on instances created by the requesting user.
	Users []LimitedUser `json:"users,omitempty"`
	World World         `json:"world"`

	// WorldId WorldID be "offline" on User profiles if you are not friends with that user.
	WorldId WorldId `json:"worldId"`
}

type InstanceShortNameResponse struct {
	SecureName string `json:"secureName"`
	ShortName  string `json:"shortName,omitempty"`
}

type PermissionId string

type Permission struct {
	Data             any          `json:"data,omitempty"`
	Id               PermissionId `json:"id"`
	Name             string       `json:"name"`
	OwnerDisplayName string       `json:"ownerDisplayName"`

	// OwnerId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId UserId `json:"ownerId"`
}

type PlayerModerationId string

type PlayerModerationType string

const (
	PlayerModerationTypeMute        PlayerModerationType = "mute"
	PlayerModerationTypeUnmute      PlayerModerationType = "unmute"
	PlayerModerationTypeBlock       PlayerModerationType = "block"
	PlayerModerationTypeUnblock     PlayerModerationType = "unblock"
	PlayerModerationTypeInteractOn  PlayerModerationType = "interactOn"
	PlayerModerationTypeInteractOff PlayerModerationType = "interactOff"
)

type PlayerModeration struct {
	Created           time.Time          `json:"created"`
	Id                PlayerModerationId `json:"id"`
	SourceDisplayName string             `json:"sourceDisplayName"`

	// SourceUserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	SourceUserId      UserId `json:"sourceUserId"`
	TargetDisplayName string `json:"targetDisplayName"`

	// TargetUserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	TargetUserId UserId               `json:"targetUserId"`
	Type         PlayerModerationType `json:"type"`
}

type ModerateUserRequest struct {
	// Moderated A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Moderated UserId               `json:"moderated"`
	Type      PlayerModerationType `json:"type"`
}

// ApiConfigAnnouncement Public Announcement
type ApiConfigAnnouncement struct {
	// Name Announcement name
	Name string `json:"name"`

	// Text Announcement text
	Text string `json:"text"`
}

// DeploymentGroup Used to identify which API deployment cluster is currently responding.
//
// `blue` and `green` are used by Production.
// `grape`and `cherry` are used during Development.
//
// [Blue Green Deployment by Martin Fowler](https://martinfowler.com/bliki/BlueGreenDeployment.html)
type DeploymentGroup string

const (
	DeploymentGroupBlue   DeploymentGroup = "blue"
	DeploymentGroupGreen  DeploymentGroup = "green"
	DeploymentGroupGrape  DeploymentGroup = "grape"
	DeploymentGroupCherry DeploymentGroup = "cherry"
)

// ApiConfigDownloadUrlList Download links for various development assets.
type ApiConfigDownloadUrlList struct {
	// Bootstrap Download link for ???
	Bootstrap string `json:"bootstrap"`

	// Sdk2 Download link for legacy SDK2
	Sdk2 string `json:"sdk2"`

	// Sdk3Avatars Download link for SDK3 for Avatars
	Sdk3Avatars string `json:"sdk3-avatars"`

	// Sdk3Worlds Download link for SDK3 for Worlds
	Sdk3Worlds string `json:"sdk3-worlds"`

	// Vcc Download link for the Creator Companion
	Vcc string `json:"vcc"`
}

type DynamicContentRow struct {
	Index int64  `json:"index,omitempty"`
	Name  string `json:"name"`

	// Platform Usually "ThisPlatformSupported", but can also be other values such as "all" or platform specific identifiers.
	Platform      string `json:"platform"`
	SortHeading   string `json:"sortHeading"`
	SortOrder     string `json:"sortOrder"`
	SortOwnership string `json:"sortOwnership"`

	// Tag Tag to filter content for this row.
	Tag string `json:"tag,omitempty"`

	// Type Type is not present if it is a world.
	Type string `json:"type,omitempty"`
}

type ApiConfigEvents struct {
	// DistanceClose Unknown
	DistanceClose int64 `json:"distanceClose"`

	// DistanceFactor Unknown
	DistanceFactor int64 `json:"distanceFactor"`

	// DistanceFar Unknown
	DistanceFar int64 `json:"distanceFar"`

	// GroupDistance Unknown
	GroupDistance int64 `json:"groupDistance"`

	// MaximumBunchSize Unknown
	MaximumBunchSize int64 `json:"maximumBunchSize"`

	// NotVisibleFactor Unknown
	NotVisibleFactor int64 `json:"notVisibleFactor"`

	// PlayerOrderBucketSize Unknown
	PlayerOrderBucketSize int64 `json:"playerOrderBucketSize"`

	// PlayerOrderFactor Unknown
	PlayerOrderFactor int64 `json:"playerOrderFactor"`

	// SlowUpdateFactorThreshold Unknown
	SlowUpdateFactorThreshold int64 `json:"slowUpdateFactorThreshold"`

	// ViewSegmentLength Unknown
	ViewSegmentLength int64 `json:"viewSegmentLength"`
}

type ApiConfig struct {
	// Address VRChat's office address
	Address string `json:"address"`

	// Announcements Public Announcements
	Announcements []ApiConfigAnnouncement `json:"announcements"`

	// AppName Game name
	AppName string `json:"appName"`

	// AvailableLanguageCodes List of supported Languages
	AvailableLanguageCodes []string `json:"availableLanguageCodes"`

	// AvailableLanguages List of supported Languages
	AvailableLanguages []string `json:"availableLanguages"`

	// BuildVersionTag Build tag of the API server
	BuildVersionTag string `json:"buildVersionTag"`

	// ClientApiKey apiKey to be used for all other requests
	ClientApiKey string `json:"clientApiKey"`

	// ClientBpsCeiling Unknown
	ClientBpsCeiling int64 `json:"clientBPSCeiling"`

	// ClientDisconnectTimeout Unknown
	ClientDisconnectTimeout int64 `json:"clientDisconnectTimeout"`

	// ClientNetDispatchThread Unknown
	ClientNetDispatchThread bool `json:"clientNetDispatchThread,omitempty"`

	// ClientNetInThread Unknown
	ClientNetInThread bool `json:"clientNetInThread,omitempty"`

	// ClientNetInThread2 Unknown
	ClientNetInThread2 bool `json:"clientNetInThread2,omitempty"`

	// ClientNetInThreadMobile Unknown
	ClientNetInThreadMobile bool `json:"clientNetInThreadMobile,omitempty"`

	// ClientNetInThreadMobile2 Unknown
	ClientNetInThreadMobile2 bool `json:"clientNetInThreadMobile2,omitempty"`

	// ClientNetOutThread Unknown
	ClientNetOutThread bool `json:"clientNetOutThread,omitempty"`

	// ClientNetOutThread2 Unknown
	ClientNetOutThread2 bool `json:"clientNetOutThread2,omitempty"`

	// ClientNetOutThreadMobile Unknown
	ClientNetOutThreadMobile bool `json:"clientNetOutThreadMobile,omitempty"`

	// ClientNetOutThreadMobile2 Unknown
	ClientNetOutThreadMobile2 bool `json:"clientNetOutThreadMobile2,omitempty"`

	// ClientQr Unknown
	ClientQr int64 `json:"clientQR,omitempty"`

	// ClientReservedPlayerBps Unknown
	ClientReservedPlayerBps int64 `json:"clientReservedPlayerBPS"`

	// ClientSentCountAllowance Unknown
	ClientSentCountAllowance int64 `json:"clientSentCountAllowance"`

	// ContactEmail VRChat's contact email
	ContactEmail string `json:"contactEmail"`

	// CopyrightEmail VRChat's copyright-issues-related email
	CopyrightEmail string `json:"copyrightEmail"`

	// CurrentPrivacyVersion Current version number of the Privacy Agreement
	CurrentPrivacyVersion int64 `json:"currentPrivacyVersion,omitempty"`

	// CurrentTosVersion Current version number of the Terms of Service
	CurrentTosVersion int64    `json:"currentTOSVersion"`
	DefaultAvatar     AvatarId `json:"defaultAvatar"`

	// DeploymentGroup Used to identify which API deployment cluster is currently responding.
	//
	// `blue` and `green` are used by Production.
	// `grape`and `cherry` are used during Development.
	//
	// [Blue Green Deployment by Martin Fowler](https://martinfowler.com/bliki/BlueGreenDeployment.html)
	DeploymentGroup DeploymentGroup `json:"deploymentGroup"`

	// DevLanguageCodes Unknown
	DevLanguageCodes []string `json:"devLanguageCodes,omitempty"`

	// DevSdkUrl Link to download the development SDK, use downloadUrls instead
	DevSdkUrl string `json:"devSdkUrl"`

	// DevSdkVersion Version of the development SDK
	DevSdkVersion string `json:"devSdkVersion"`

	// DisCountdown Unknown, "dis" maybe for disconnect?
	DisCountdown time.Time `json:"dis-countdown"`

	// DisableAvProInProton Unknown
	DisableAvProInProton bool `json:"disableAVProInProton,omitempty"`

	// DisableAvatarCopying Toggles if copying avatars should be disabled
	DisableAvatarCopying bool `json:"disableAvatarCopying"`

	// DisableAvatarGating Toggles if avatar gating should be disabled. Avatar gating restricts uploading of avatars to people with the `system_avatar_access` Tag or `admin_avatar_access` Tag
	DisableAvatarGating bool `json:"disableAvatarGating"`

	// DisableCaptcha Unknown
	DisableCaptcha bool `json:"disableCaptcha,omitempty"`

	// DisableCommunityLabs Toggles if the Community Labs should be disabled
	DisableCommunityLabs bool `json:"disableCommunityLabs"`

	// DisableCommunityLabsPromotion Toggles if promotion out of Community Labs should be disabled
	DisableCommunityLabsPromotion bool `json:"disableCommunityLabsPromotion"`

	// DisableEmail Unknown
	DisableEmail bool `json:"disableEmail"`

	// DisableEventStream Toggles if Analytics should be disabled.
	DisableEventStream bool `json:"disableEventStream"`

	// DisableFeedbackGating Toggles if feedback gating should be disabled. Feedback gating restricts submission of feedback (reporting a World or User) to people with the `system_feedback_access` Tag.
	DisableFeedbackGating bool `json:"disableFeedbackGating"`

	// DisableFrontendBuilds Unknown, probably toggles compilation of frontend web builds? So internal flag?
	DisableFrontendBuilds bool `json:"disableFrontendBuilds"`

	// DisableHello Unknown
	DisableHello bool `json:"disableHello"`

	// DisableOculusSubs Toggles if signing up for Subscriptions in Oculus is disabled or not.
	DisableOculusSubs bool `json:"disableOculusSubs"`

	// DisableRegistration Toggles if new user account registration should be disabled.
	DisableRegistration bool `json:"disableRegistration"`

	// DisableSteamNetworking Toggles if Steam Networking should be disabled. VRChat these days uses Photon Unity Networking (PUN) instead.
	DisableSteamNetworking bool `json:"disableSteamNetworking"`

	// DisableTwoFactorAuth Toggles if 2FA should be disabled.
	DisableTwoFactorAuth bool `json:"disableTwoFactorAuth"`

	// DisableUdon Toggles if Udon should be universally disabled in-game.
	DisableUdon bool `json:"disableUdon"`

	// DisableUpgradeAccount Toggles if account upgrading "linking with Steam/Oculus" should be disabled.
	DisableUpgradeAccount bool `json:"disableUpgradeAccount"`

	// DownloadLinkWindows Download link for game on the Oculus Rift website.
	DownloadLinkWindows string `json:"downloadLinkWindows"`

	// DownloadUrls Download links for various development assets.
	DownloadUrls ApiConfigDownloadUrlList `json:"downloadUrls"`

	// DynamicWorldRows Array of DynamicWorldRow objects, used by the game to display the list of world rows
	DynamicWorldRows []DynamicContentRow `json:"dynamicWorldRows"`

	// EconomyPauseEnd Unknown
	EconomyPauseEnd string `json:"economyPauseEnd,omitempty"`

	// EconomyPauseStart Unknown
	EconomyPauseStart string `json:"economyPauseStart,omitempty"`

	// EconomyState Unknown
	EconomyState int64           `json:"economyState,omitempty"`
	Events       ApiConfigEvents `json:"events"`

	// HomeWorldId WorldID be "offline" on User profiles if you are not friends with that user.
	HomeWorldId WorldId `json:"homeWorldId"`

	// HomepageRedirectTarget Redirect target if you try to open the base API domain in your browser
	HomepageRedirectTarget string `json:"homepageRedirectTarget"`

	// HubWorldId WorldID be "offline" on User profiles if you are not friends with that user.
	HubWorldId WorldId `json:"hubWorldId"`

	// ImageHostUrlList A list of explicitly allowed origins that worlds can request images from via the Udon's [VRCImageDownloader#DownloadImage](https://creators.vrchat.com/worlds/udon/image-loading/#downloadimage).
	ImageHostUrlList []string `json:"imageHostUrlList"`

	// JobsEmail VRChat's job application email
	JobsEmail string `json:"jobsEmail"`

	// ModerationEmail VRChat's moderation related email
	ModerationEmail string `json:"moderationEmail"`

	// NotAllowedToSelectAvatarInPrivateWorldMessage Used in-game to notify a user they aren't allowed to select avatars in private worlds
	NotAllowedToSelectAvatarInPrivateWorldMessage string `json:"notAllowedToSelectAvatarInPrivateWorldMessage"`

	// PlayerUrlResolverSha1 Currently used youtube-dl.exe hash in SHA1-delimited format
	PlayerUrlResolverSha1 string `json:"player-url-resolver-sha1"`

	// PlayerUrlResolverVersion Currently used youtube-dl.exe version
	PlayerUrlResolverVersion string `json:"player-url-resolver-version"`

	// SdkDeveloperFaqUrl Link to the developer FAQ
	SdkDeveloperFaqUrl string `json:"sdkDeveloperFaqUrl"`

	// SdkDiscordUrl Link to the official VRChat Discord
	SdkDiscordUrl string `json:"sdkDiscordUrl"`

	// SdkNotAllowedToPublishMessage Used in the SDK to notify a user they aren't allowed to upload avatars/worlds yet
	SdkNotAllowedToPublishMessage string `json:"sdkNotAllowedToPublishMessage"`

	// SdkUnityVersion Unity version supported by the SDK
	SdkUnityVersion string `json:"sdkUnityVersion"`

	// ServerName Server name of the API server currently responding
	ServerName string `json:"serverName"`

	// StringHostUrlList A list of explicitly allowed origins that worlds can request strings from via the Udon's [VRCStringDownloader.LoadUrl](https://creators.vrchat.com/worlds/udon/string-loading/#ivrcstringdownload).
	StringHostUrlList []string `json:"stringHostUrlList"`

	// SupportEmail VRChat's support email
	SupportEmail string `json:"supportEmail"`

	// TimeOutWorldId WorldID be "offline" on User profiles if you are not friends with that user.
	TimeOutWorldId WorldId `json:"timeOutWorldId"`

	// TutorialWorldId WorldID be "offline" on User profiles if you are not friends with that user.
	TutorialWorldId WorldId `json:"tutorialWorldId"`

	// UpdateRateMsMaximum Unknown
	UpdateRateMsMaximum int64 `json:"updateRateMsMaximum"`

	// UpdateRateMsMinimum Unknown
	UpdateRateMsMinimum int64 `json:"updateRateMsMinimum"`

	// UpdateRateMsNormal Unknown
	UpdateRateMsNormal int64 `json:"updateRateMsNormal"`

	// UpdateRateMsUdonManual Unknown
	UpdateRateMsUdonManual int64 `json:"updateRateMsUdonManual"`

	// UploadAnalysisPercent Unknown
	UploadAnalysisPercent int64 `json:"uploadAnalysisPercent"`

	// UrlList List of allowed URLs that bypass the "Allow untrusted URL's" setting in-game
	UrlList []string `json:"urlList"`

	// UseReliableUdpForVoice Unknown
	UseReliableUdpForVoice bool `json:"useReliableUdpForVoice"`

	// ViveWindowsUrl Download link for game on the Steam website.
	ViveWindowsUrl string `json:"viveWindowsUrl"`

	// VoiceEnableDegradation Unknown, probably voice optimization testing
	VoiceEnableDegradation bool `json:"VoiceEnableDegradation"`

	// VoiceEnableReceiverLimiting Unknown, probably voice optimization testing
	VoiceEnableReceiverLimiting bool `json:"VoiceEnableReceiverLimiting"`

	// WhiteListedAssetUrls List of allowed URLs that are allowed to host avatar assets
	WhiteListedAssetUrls []string `json:"whiteListedAssetUrls"`
}

type InfoPushDataClickable struct { // Command enum
	Command string `json:"command"`

	// Parameters In case of OpenURL, this would contain the link.
	Parameters []string `json:"parameters,omitempty"`
}

type InfoPushDataArticleContent struct {
	ImageUrl  string                `json:"imageUrl,omitempty"`
	OnPressed InfoPushDataClickable `json:"onPressed,omitempty"`
	Text      string                `json:"text,omitempty"`
}

type InfoPushDataArticle struct {
	Content InfoPushDataArticleContent `json:"content,omitempty"`
}

type InfoPushData struct {
	Article     InfoPushDataArticle   `json:"article,omitempty"`
	ContentList DynamicContentRow     `json:"contentList,omitempty"`
	Description string                `json:"description,omitempty"`
	ImageUrl    string                `json:"imageUrl,omitempty"`
	Name        string                `json:"name,omitempty"`
	OnPressed   InfoPushDataClickable `json:"onPressed,omitempty"`
	Template    string                `json:"template,omitempty"`
	Version     string                `json:"version,omitempty"`
}

type InfoPush struct {
	CreatedAt time.Time    `json:"createdAt"`
	Data      InfoPushData `json:"data"`
	EndDate   time.Time    `json:"endDate,omitempty"`

	// Hash Unknown usage, MD5
	Hash          string        `json:"hash"`
	Id            string        `json:"id"`
	IsEnabled     bool          `json:"isEnabled"`
	Priority      int64         `json:"priority"`
	ReleaseStatus ReleaseStatus `json:"releaseStatus"`
	StartDate     time.Time     `json:"startDate,omitempty"`

	// Tags
	Tags      []Tag     `json:"tags"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ApiHealth struct {
	BuildVersionTag string `json:"buildVersionTag"`
	Ok              bool   `json:"ok"`
	ServerName      string `json:"serverName"`
}

type User struct {
	AllowAvatarCopying bool `json:"allowAvatarCopying"`

	// Badges
	Badges   []Badge  `json:"badges,omitempty"`
	Bio      string   `json:"bio"`
	BioLinks []string `json:"bioLinks"`

	// CurrentAvatarImageUrl When profilePicOverride is not empty, use it instead.
	CurrentAvatarImageUrl CurrentAvatarImageUrl `json:"currentAvatarImageUrl"`
	CurrentAvatarTags     []Tag                 `json:"currentAvatarTags"`

	// CurrentAvatarThumbnailImageUrl When profilePicOverride is not empty, use it instead.
	CurrentAvatarThumbnailImageUrl CurrentAvatarThumbnailImageUrl `json:"currentAvatarThumbnailImageUrl"`
	DateJoined                     string                         `json:"date_joined"`

	// DeveloperType "none" User is a normal user
	// "trusted" Unknown
	// "internal" Is a VRChat Developer
	// "moderator" Is a VRChat Moderator
	//
	// Staff can hide their developerType at will.
	DeveloperType DeveloperType `json:"developerType"`

	// DisplayName A users visual display name. This is what shows up in-game, and can different from their `username`. Changing display name is restricted to a cooldown period.
	DisplayName         string `json:"displayName"`
	FriendKey           string `json:"friendKey"`
	FriendRequestStatus string `json:"friendRequestStatus,omitempty"`

	// Id A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	Id UserId `json:"id"`

	// InstanceId InstanceID can be "offline" on User profiles if you are not friends with that user and "private" if you are friends and user is in private instance.
	InstanceId InstanceId `json:"instanceId,omitempty"`

	// IsFriend Either their `friendKey`, or empty string if you are not friends. Unknown usage.
	IsFriend bool `json:"isFriend"`

	// LastActivity Either a date-time or empty string.
	LastActivity string `json:"last_activity"`

	// LastLogin Either a date-time or empty string.
	LastLogin string `json:"last_login"`

	// LastPlatform This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	LastPlatform Platform `json:"last_platform"`

	// Location WorldID be "offline" on User profiles if you are not friends with that user.
	Location                    WorldId `json:"location,omitempty"`
	Note                        string  `json:"note,omitempty"`
	Platform                    string  `json:"platform"`
	ProfilePicOverride          string  `json:"profilePicOverride"`
	ProfilePicOverrideThumbnail string  `json:"profilePicOverrideThumbnail"`
	Pronouns                    string  `json:"pronouns"`

	// State * "online" User is online in VRChat
	// * "active" User is online, but not in VRChat
	// * "offline" User is offline
	//
	// Always offline when returned through `getCurrentUser` (/auth/user).
	State UserState `json:"state"`

	// Status Defines the User's current status, for example "ask me", "join me" or "offline. This status is a combined indicator of their online activity and privacy preference.
	Status            UserStatus `json:"status"`
	StatusDescription string     `json:"statusDescription"`

	// Tags
	Tags                []Tag  `json:"tags"`
	TravelingToInstance string `json:"travelingToInstance,omitempty"`
	TravelingToLocation string `json:"travelingToLocation,omitempty"`
	TravelingToWorld    string `json:"travelingToWorld,omitempty"`
	UserIcon            string `json:"userIcon"`

	// Username -| A users unique name, used during login. This is different from `displayName` which is what shows up in-game. A users `username` can never be changed.'
	// **DEPRECATED:** VRChat API no longer return usernames of other users. [See issue by Tupper for more information](https://github.com/pypy-vrc/VRCX/issues/429).
	Username string `json:"username,omitempty"`

	// WorldId WorldID be "offline" on User profiles if you are not friends with that user.
	WorldId WorldId `json:"worldId,omitempty"`
}

type UpdateUserRequest struct {
	AcceptedTosVersion float64  `json:"acceptedTOSVersion,omitempty"`
	Bio                string   `json:"bio,omitempty"`
	BioLinks           []string `json:"bioLinks,omitempty"`
	Birthday           string   `json:"birthday,omitempty"`
	Email              string   `json:"email,omitempty"`
	IsBoopingEnabled   bool     `json:"isBoopingEnabled,omitempty"`
	Pronouns           string   `json:"pronouns,omitempty"`

	// Status Defines the User's current status, for example "ask me", "join me" or "offline. This status is a combined indicator of their online activity and privacy preference.
	Status            UserStatus `json:"status,omitempty"`
	StatusDescription string     `json:"statusDescription,omitempty"`

	// Tags
	Tags []Tag `json:"tags,omitempty"`

	// UserIcon MUST be a valid VRChat /file/ url.
	UserIcon string `json:"userIcon,omitempty"`
}

type LimitedUserGroups struct {
	BannerId          string             `json:"bannerId,omitempty"`
	BannerUrl         string             `json:"bannerUrl,omitempty"`
	Description       string             `json:"description,omitempty"`
	Discriminator     GroupDiscriminator `json:"discriminator,omitempty"`
	GroupId           GroupId            `json:"groupId,omitempty"`
	IconId            string             `json:"iconId,omitempty"`
	IconUrl           string             `json:"iconUrl,omitempty"`
	Id                GroupMemberId      `json:"id,omitempty"`
	IsRepresenting    bool               `json:"isRepresenting,omitempty"`
	LastPostCreatedAt time.Time          `json:"lastPostCreatedAt,omitempty"`
	LastPostReadAt    time.Time          `json:"lastPostReadAt,omitempty"`
	MemberCount       int64              `json:"memberCount,omitempty"`
	MemberVisibility  string             `json:"memberVisibility,omitempty"`
	MutualGroup       bool               `json:"mutualGroup,omitempty"`
	Name              string             `json:"name,omitempty"`

	// OwnerId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId   UserId         `json:"ownerId,omitempty"`
	Privacy   string         `json:"privacy,omitempty"`
	ShortCode GroupShortCode `json:"shortCode,omitempty"`
}

type RepresentedGroup struct {
	BannerId         string              `json:"bannerId,omitempty"`
	BannerUrl        string              `json:"bannerUrl,omitempty"`
	Description      string              `json:"description,omitempty"`
	Discriminator    GroupDiscriminator  `json:"discriminator,omitempty"`
	GroupId          GroupId             `json:"groupId,omitempty"`
	IconId           string              `json:"iconId,omitempty"`
	IconUrl          string              `json:"iconUrl,omitempty"`
	IsRepresenting   bool                `json:"isRepresenting,omitempty"`
	MemberCount      int64               `json:"memberCount,omitempty"`
	MemberVisibility GroupUserVisibility `json:"memberVisibility,omitempty"`
	Name             string              `json:"name,omitempty"`

	// OwnerId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	OwnerId   UserId         `json:"ownerId,omitempty"`
	Privacy   GroupPrivacy   `json:"privacy,omitempty"`
	ShortCode GroupShortCode `json:"shortCode,omitempty"`
}

type LimitedUnityPackage struct {
	// Platform This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	Platform     Platform `json:"platform"`
	UnityVersion string   `json:"unityVersion"`
}

type LimitedWorld struct {
	// AuthorId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId   UserId    `json:"authorId"`
	AuthorName string    `json:"authorName"`
	Capacity   int64     `json:"capacity"`
	CreatedAt  time.Time `json:"created_at"`
	Favorites  int64     `json:"favorites"`
	Heat       int64     `json:"heat"`

	// Id WorldID be "offline" on User profiles if you are not friends with that user.
	Id                  WorldId       `json:"id"`
	ImageUrl            string        `json:"imageUrl"`
	LabsPublicationDate string        `json:"labsPublicationDate"`
	Name                string        `json:"name"`
	Occupants           int64         `json:"occupants"`
	Organization        string        `json:"organization"`
	Popularity          int64         `json:"popularity"`
	PreviewYoutubeId    string        `json:"previewYoutubeId,omitempty"`
	PublicationDate     string        `json:"publicationDate"`
	RecommendedCapacity int64         `json:"recommendedCapacity,omitempty"`
	ReleaseStatus       ReleaseStatus `json:"releaseStatus"`

	// Tags
	Tags              []Tag           `json:"tags"`
	ThumbnailImageUrl string          `json:"thumbnailImageUrl"`
	UdonProducts      []UdonProductId `json:"udonProducts,omitempty"`

	// UnityPackages
	UnityPackages []LimitedUnityPackage `json:"unityPackages"`
	UpdatedAt     time.Time             `json:"updated_at"`
	Visits        int64                 `json:"visits,omitempty"`
}

type CreateWorldRequest struct {
	AssetUrl     string `json:"assetUrl"`
	AssetVersion int64  `json:"assetVersion,omitempty"`

	// AuthorId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId    UserId `json:"authorId,omitempty"`
	AuthorName  string `json:"authorName,omitempty"`
	Capacity    int64  `json:"capacity,omitempty"`
	Description string `json:"description,omitempty"`

	// Id WorldID be "offline" on User profiles if you are not friends with that user.
	Id       WorldId `json:"id,omitempty"`
	ImageUrl string  `json:"imageUrl"`
	Name     string  `json:"name"`

	// Platform This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	Platform      Platform      `json:"platform,omitempty"`
	ReleaseStatus ReleaseStatus `json:"releaseStatus,omitempty"`

	// Tags
	Tags            []Tag  `json:"tags,omitempty"`
	UnityPackageUrl string `json:"unityPackageUrl,omitempty"`
	UnityVersion    string `json:"unityVersion,omitempty"`
}

type UpdateWorldRequest struct {
	AssetUrl     string `json:"assetUrl,omitempty"`
	AssetVersion string `json:"assetVersion,omitempty"`

	// AuthorId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	AuthorId    UserId `json:"authorId,omitempty"`
	AuthorName  string `json:"authorName,omitempty"`
	Capacity    int64  `json:"capacity,omitempty"`
	Description string `json:"description,omitempty"`
	ImageUrl    string `json:"imageUrl,omitempty"`
	Name        string `json:"name,omitempty"`

	// Platform This can be `standalonewindows` or `android`, but can also pretty much be any random Unity verison such as `2019.2.4-801-Release` or `2019.2.2-772-Release` or even `unknownplatform`.
	Platform      Platform      `json:"platform,omitempty"`
	ReleaseStatus ReleaseStatus `json:"releaseStatus,omitempty"`

	// Tags
	Tags            []Tag  `json:"tags,omitempty"`
	UnityPackageUrl string `json:"unityPackageUrl,omitempty"`
	UnityVersion    string `json:"unityVersion,omitempty"`
}

type WorldMetadata struct {
	// Id WorldID be "offline" on User profiles if you are not friends with that user.
	Id       WorldId `json:"id"`
	Metadata any     `json:"metadata"`
}

type WorldPublishStatus struct {
	CanPublish bool `json:"canPublish"`
}

type NotificationDetailInvite struct {
	InviteMessage string `json:"inviteMessage,omitempty"`

	// WorldId WorldID be "offline" on User profiles if you are not friends with that user.
	WorldId   WorldId `json:"worldId"`
	WorldName string  `json:"worldName"`
}

type NotificationDetailInviteResponse struct {
	InResponseTo    NotificationId `json:"inResponseTo"`
	ResponseMessage string         `json:"responseMessage"`
}

type NotificationDetailRequestInvite struct {
	// Platform TODO: Does this still exist?
	Platform string `json:"platform,omitempty"`

	// RequestMessage Used when using InviteMessage Slot.
	RequestMessage string `json:"requestMessage,omitempty"`
}

type NotificationDetailRequestInviteResponse struct {
	InResponseTo NotificationId `json:"inResponseTo"`

	// RequestMessage Used when using InviteMessage Slot.
	RequestMessage string `json:"requestMessage,omitempty"`
}

type NotificationDetailVoteToKick struct {
	// InitiatorUserId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	InitiatorUserId UserId `json:"initiatorUserId"`

	// UserToKickId A users unique ID, usually in the form of `usr_c1644b5b-3ca4-45b4-97c6-a2a0de70d469`. Legacy players can have old IDs in the form of `8JoV9XEdpo`. The ID can never be changed.
	UserToKickId UserId `json:"userToKickId"`
}

// UserExistsResponse Status object representing if a queried user by username or userId exists or not. This model is primarily used by the `/auth/exists` endpoint, which in turn is used during registration. Please see the documentation on that endpoint for more information on usage.
type UserExistsResponse UserExists

type MissingParameterError Error

type CurrentUserLoginResponse CurrentUser

type MissingCredentialsError Error

type Verify2FaResponse Verify2FaResult

type Verify2FaEmailCodeResponse Verify2FaEmailCodeResult

type VerifyAuthTokenResponse VerifyAuthTokenResult

type LogoutSuccess Success

type DeleteUserResponse CurrentUser

type AvatarResponse Avatar

type AvatarSeeOtherUserCurrentAvatarError Error

type AvatarListResponse []Avatar

type FeaturedSetNotAdminError Error

type AvatarNotFoundError Error

type CurrentUserResponse CurrentUser

type AvatarNotTaggedAsFallbackError Error

type AvatarSeeOtherUserFavoritesError Error

type TransactionListResponse []Transaction

type TransactionResponse Transaction

type UserSubscriptionListResponse []UserSubscription

type SubscriptionListResponse []Subscription

type LicenseGroupListResponse []LicenseGroup

type LicenseGroupResponse LicenseGroup

type FavoriteListResponse []Favorite

type FavoriteResponse Favorite

type FavoriteAddAlreadyFavoritedError Error

type FavoriteAddNotFriendsError Error

type FavoriteNotFoundError Error

type FavoriteRemovedSuccess Success

type FavoriteGroupListResponse []FavoriteGroup

type FavoriteGroupResponse FavoriteGroup

type FavoriteGroupClearedSuccess Success

type FileListResponse []File

type FileResponse File

type FileNotFoundError Error

type FileDeletedError Error

// RawFileResponse Raw file
type RawFileResponse any

type FileVersionDeleteInitialError Error

type FileVersionDeleteMiddleError Error

type FileUploadUrlResponse FileUploadUrl

type FileUploadAlreadyFinishedError Error

type FileVersionUploadStatusResponse FileVersionUploadStatus

type LimitedUserListResponse []LimitedUser

type NotificationResponse Notification

type FriendBadRequestError Error

type UserDoesntExistError Error

type DeleteFriendSuccess Success

type DeleteFriendRequestError Error

type FriendStatusResponse FriendStatus

type UnfriendSuccess Success

type NotFriendsError Error

type LimitedGroupListResponse []LimitedGroup

type GroupResponse Group

type GroupNotFoundError Error

type DeleteGroupSuccess Success

type GroupAnnouncementResponse GroupAnnouncement

type DeleteGroupAnnouncementSuccess Success

type GroupAuditLogListResponse PaginatedGroupAuditLogEntryList

type GroupMemberListResponse []GroupMember

type NoPermission Error

type GroupMemberResponse GroupMember

type BanGroupMemberBadRequestError Error

type GroupGalleryResponse GroupGallery

type GroupGalleryImageListResponse []GroupGalleryImage

type DeleteGroupGallerySuccess Success

type GroupGalleryImageResponse GroupGalleryImage

type DeleteGroupGalleryImageSuccess Success

type GroupGalleryImageDeleteForbiddenError Error

type GroupInstanceListResponse []GroupInstance

type GroupNotMemberError Error

type GroupInviteBadRequestError Error

type GroupInviteForbiddenError Error

type DeleteGroupInviteBadRequestError Error

type GroupAlreadyMemberError Error

type UsersInvalidSearchError Error

type GroupLimitedMemberResponse GroupLimitedMember

// GroupRoleIdListResponse
type GroupRoleIdListResponse GroupRoleIdList

type GroupPermissionListResponse []GroupPermission

type GroupPostResponse GroupPost

type GroupPostResponseSuccess Success

type GroupJoinRequestResponseBadRequestError Error

type GroupRoleListResponse []GroupRole

type GroupRoleResponse GroupRole

type SendNotificationResponse SentNotification

type InviteMustBeFriendsError Error

type InstanceNotFoundError Error

type InviteResponse400Error Error

type InviteMessageListResponse []InviteMessage

type InviteMessageInvalidSlotNumberError Error

type NotAuthorizedActionError Error

type InviteMessageResponse InviteMessage

type InviteMessageGetNegativeSlotError Error

type InviteMessageGetTooHighSlotError Error

type InviteMessageUpdateRateLimitError Error

type InviteMessageNoEntryForSlotError Error

// InstanceResponse * `hidden` field is only present if InstanceType is `hidden` aka "Friends+", and is instance creator.
// * `friends` field is only present if InstanceType is `friends` aka "Friends", and is instance creator.
// * `private` field is only present if InstanceType is `private` aka "Invite" or "Invite+", and is instance creator.
type InstanceResponse Instance

type InstanceCloseForbiddenError Error

type InstanceSelfInviteSuccess Success

type NotificationListResponse []Notification

type FriendSuccess Success

type AcceptFriendRequestError Error

type ClearNotificationsSuccess Success

type PermissionListResponse []Permission

type InvalidAdminCredentialsError Error

type PermissionResponse Permission

type PlayerModerationListResponse []PlayerModeration

type PlayerModerationResponse PlayerModeration

type PlayerModerationClearAllSuccess Success

type PlayerModerationNotFoundError Error

type PlayerModerationRemovedSuccess Success

type PlayerModerationDeleteOthersError Error

type PlayerModerationUnmoderatedSuccess Success

type ApiConfigResponse ApiConfig

type InfoPushListResponse []InfoPush

type DownloadSourceCodeAccessError Error

type ApiHealthResponse ApiHealth

// CurrentOnlineUsersResponse Number of online users
type CurrentOnlineUsersResponse int64

// SystemTimeResponse Does not return millisecond precision. Always returns time in UTC.
type SystemTimeResponse time.Time

type UserResponse User

type LimitedUserGroupListResponse []LimitedUserGroups

type GroupListResponse []Group

type LimitedWorldListResponse []LimitedWorld

type WorldResponse World

type WorldCreateNotAllowedYetError Error

type WorldSeeOtherUserFavoritesError Error

type WorldSeeOtherUserRecentsError Error

type WorldNotFoundError Error

type WorldMetadataResponse WorldMetadata

type WorldPublishStatusResponse WorldPublishStatus
