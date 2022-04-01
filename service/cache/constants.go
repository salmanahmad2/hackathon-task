package cache

import "time"

const (
	AllUsersKey            = "allUsers"
	AllUsersKey_Exp        = 15 * time.Minute
	UserDataKey            = "userData:"
	UserDataKey_Exp        = 5 * time.Second
	SearchProfileKey       = "searchProfile:"
	SearchProfileKey_Exp   = 5 * time.Second
	ReportsSuicideKey      = "reports:Suicide"
	ReportsSuicideKey_Exp  = 15 * time.Minute
	ReportsNudityKey       = "reports:Nudity"
	ReportsNudityKey_Exp   = 15 * time.Minute
	ReportsPropertyKey     = "reports:Property"
	ReportsPropertyKey_Exp = 15 * time.Minute
	AllReportsKey          = "reports:All"
	AllReportsKey_Exp      = 15 * time.Minute
	YourAssetsKey          = "userAssets:"
	YourAssetsKey_Exp      = 5 * time.Second
	AssetDataKey           = "assetData:"
	AssetDataKey_Exp       = 5 * time.Second
	FavouriteAssetsKey     = "favouritedAssets:"
	FavouriteAssetsKey_Exp = 5 * time.Second
	RefreshTokenKey        = "refreshToken:"
	RefreshToken_Exp       = (30 * 24) * time.Hour
	UserSignupKey          = "userSignup:"
	UserSignupKey_Exp      = 3 * time.Minute
	ForgetPasswordKey      = "forgetPassword:"
	ForgetPasswordKey_Exp  = 3 * time.Minute
	ListedAuctionsKey      = "auctions:Listed"
	ListedAuctionsKey_Exp  = 1 * time.Minute
	OpenedAuctionsKey      = "auctions:Opened"
	OpenedAuctionsKey_Exp  = 10 * time.Second
	ClosedAuctionsKey      = "auctions:Closed"
	ClosedAuctionsKey_Exp  = 5 * time.Minute
	AuctionDataKey         = "auctionData:"
	AuctionDataKey_Exp     = 5 * time.Second
	AuctionBidsKey         = "auctionBids:"
	AuctionBidsKey_Exp     = 5 * time.Second
	AssetAuctionsKey       = "assetAuctions:"
	AssetAuctionsKey_Exp   = 5 * time.Second
)
