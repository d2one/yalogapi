package yalogapi

type Types struct {
	YmSCounterID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSCounterID"`
	YmSWatchIDs struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSWatchIDs"`
	YmSDateTime struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSDateTime"`
	YmSDateTimeUTC struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSDateTimeUTC"`
	YmSIsNewUser struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSIsNewUser"`
	YmSStartURL struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSStartURL"`
	YmSEndURL struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSEndURL"`
	YmSPageViews struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPageViews"`
	YmSVisitDuration struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSVisitDuration"`
	YmSBounce struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSBounce"`
	YmSIPAddress struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSIpAddress"`
	YmSParams struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSParams"`
	YmSGoalsID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSGoalsID"`
	YmSGoalsSerialNumber struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSGoalsSerialNumber"`
	YmSGoalsDateTime struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSGoalsDateTime"`
	YmSGoalsPrice struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSGoalsPrice"`
	YmSGoalsOrder struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSGoalsOrder"`
	YmSGoalsCurrency struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSGoalsCurrency"`
	YmSClientID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSClientID"`
	YmSLastTrafficSource struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastTrafficSource"`
	YmSLastAdvEngine struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastAdvEngine"`
	YmSLastReferalSource struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastReferalSource"`
	YmSLastSearchEngineRoot struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastSearchEngineRoot"`
	YmSLastSearchEngine struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastSearchEngine"`
	YmSLastSocialNetwork struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastSocialNetwork"`
	YmSLastSocialNetworkProfile struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastSocialNetworkProfile"`
	YmSReferer struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSReferer"`
	YmSLastDirectClickOrder struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastDirectClickOrder"`
	YmSLastDirectBannerGroup struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastDirectBannerGroup"`
	YmSLastDirectClickBanner struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastDirectClickBanner"`
	YmSLastDirectPhraseOrCond struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastDirectPhraseOrCond"`
	YmSLastDirectPlatformType struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastDirectPlatformType"`
	YmSLastDirectPlatform struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastDirectPlatform"`
	YmSLastDirectSearchPhrase struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastDirectSearchPhrase"`
	YmSLastDirectConditionType struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastDirectConditionType"`
	YmSLastCurrencyID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastCurrencyID"`
	YmSFrom struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSFrom"`
	YmSUTMCampaign struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSUTMCampaign"`
	YmSUTMContent struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSUTMContent"`
	YmSUTMMedium struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSUTMMedium"`
	YmSUTMSource struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSUTMSource"`
	YmSUTMTerm struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSUTMTerm"`
	YmSOpenstatAd struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOpenstatAd"`
	YmSOpenstatCampaign struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOpenstatCampaign"`
	YmSOpenstatService struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOpenstatService"`
	YmSOpenstatSource struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOpenstatSource"`
	YmSHasGCLID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSHasGCLID"`
	YmSRegionCountry struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSRegionCountry"`
	YmSRegionCity struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSRegionCity"`
	YmSBrowserLanguage struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSBrowserLanguage"`
	YmSBrowserCountry struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSBrowserCountry"`
	YmSClientTimeZone struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSClientTimeZone"`
	YmSDeviceCategory struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSDeviceCategory"`
	YmSMobilePhone struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSMobilePhone"`
	YmSMobilePhoneModel struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSMobilePhoneModel"`
	YmSOperatingSystemRoot struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOperatingSystemRoot"`
	YmSOperatingSystem struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOperatingSystem"`
	YmSBrowser struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSBrowser"`
	YmSBrowserMajorVersion struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSBrowserMajorVersion"`
	YmSBrowserMinorVersion struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSBrowserMinorVersion"`
	YmSBrowserEngine struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSBrowserEngine"`
	YmSBrowserEngineVersion1 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSBrowserEngineVersion1"`
	YmSBrowserEngineVersion2 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSBrowserEngineVersion2"`
	YmSBrowserEngineVersion3 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSBrowserEngineVersion3"`
	YmSBrowserEngineVersion4 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSBrowserEngineVersion4"`
	YmSCookieEnabled struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSCookieEnabled"`
	YmSJavascriptEnabled struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSJavascriptEnabled"`
	YmSFlashMajor struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSFlashMajor"`
	YmSFlashMinor struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSFlashMinor"`
	YmSScreenFormat struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSScreenFormat"`
	YmSScreenColors struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSScreenColors"`
	YmSScreenOrientation struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSScreenOrientation"`
	YmSScreenWidth struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSScreenWidth"`
	YmSScreenHeight struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSScreenHeight"`
	YmSPhysicalScreenWidth struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPhysicalScreenWidth"`
	YmSPhysicalScreenHeight struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPhysicalScreenHeight"`
	YmSWindowClientWidth struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSWindowClientWidth"`
	YmSWindowClientHeight struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSWindowClientHeight"`
	YmSPurchaseID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPurchaseID"`
	YmSPurchaseDateTime struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPurchaseDateTime"`
	YmSPurchaseAffiliation struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPurchaseAffiliation"`
	YmSPurchaseRevenue struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPurchaseRevenue"`
	YmSPurchaseTax struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPurchaseTax"`
	YmSPurchaseShipping struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPurchaseShipping"`
	YmSPurchaseCoupon struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPurchaseCoupon"`
	YmSPurchaseCurrency struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPurchaseCurrency"`
	YmSPurchaseProductQuantity struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSPurchaseProductQuantity"`
	YmSProductsPurchaseID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsPurchaseID"`
	YmSProductsName struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsName"`
	YmSProductsBrand struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsBrand"`
	YmSProductsCategory struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsCategory"`
	YmSProductsCategory1 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsCategory1"`
	YmSProductsCategory2 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsCategory2"`
	YmSProductsCategory3 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsCategory3"`
	YmSProductsCategory4 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsCategory4"`
	YmSProductsCategory5 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsCategory5"`
	YmSProductsVariant struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsVariant"`
	YmSProductsPosition struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsPosition"`
	YmSProductsPrice struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsPrice"`
	YmSProductsCurrency struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsCurrency"`
	YmSProductsCoupon struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsCoupon"`
	YmSProductsQuantity struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSProductsQuantity"`
	YmSImpressionsURL struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsURL"`
	YmSImpressionsDateTime struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsDateTime"`
	YmSImpressionsProductID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductID"`
	YmSImpressionsProductName struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductName"`
	YmSImpressionsProductBrand struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductBrand"`
	YmSImpressionsProductCategory struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductCategory"`
	YmSImpressionsProductCategory1 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductCategory1"`
	YmSImpressionsProductCategory2 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductCategory2"`
	YmSImpressionsProductCategory3 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductCategory3"`
	YmSImpressionsProductCategory4 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductCategory4"`
	YmSImpressionsProductCategory5 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductCategory5"`
	YmSImpressionsProductVariant struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductVariant"`
	YmSImpressionsProductPrice struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductPrice"`
	YmSImpressionsProductCurrency struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductCurrency"`
	YmSImpressionsProductCoupon struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSImpressionsProductCoupon"`
	YmSLastDirectClickOrderName struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastDirectClickOrderName"`
	YmSLastClickBannerGroupName struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastClickBannerGroupName"`
	YmSLastDirectClickBannerName struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastDirectClickBannerName"`
	YmSNetworkType struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSNetworkType"`
	YmSVisitID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSVisitID"`
	YmSDate struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSDate"`
	YmSRegionCountryID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSRegionCountryID"`
	YmSRegionCityID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSRegionCityID"`
	YmSLastGCLID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastGCLID"`
	YmSFirstGCLID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSFirstGCLID"`
	YmSLastSignificantGCLID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSLastSignificantGCLID"`
	YmSOfflineCallTalkDuration struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOfflineCallTalkDuration"`
	YmSOfflineCallHoldDuration struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOfflineCallHoldDuration"`
	YmSOfflineCallMissed struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOfflineCallMissed"`
	YmSOfflineCallTag struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOfflineCallTag"`
	YmSOfflineCallFirstTimeCaller struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOfflineCallFirstTimeCaller"`
	YmSOfflineCallURL struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmSOfflineCallURL"`
	YmPvOfflineCallURL struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvOfflineCallURL"`
	YmPvCounterID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvCounterID"`
	YmPvDateTime struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvDateTime"`
	YmPvTitle struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvTitle"`
	YmPvURL struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvURL"`
	YmPvReferer struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvReferer"`
	YmPvUTMCampaign struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvUTMCampaign"`
	YmPvUTMContent struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvUTMContent"`
	YmPvUTMMedium struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvUTMMedium"`
	YmPvUTMSource struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvUTMSource"`
	YmPvUTMTerm struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvUTMTerm"`
	YmPvBrowser struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvBrowser"`
	YmPvBrowserMajorVersion struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvBrowserMajorVersion"`
	YmPvBrowserMinorVersion struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvBrowserMinorVersion"`
	YmPvBrowserCountry struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvBrowserCountry"`
	YmPvBrowserEngine struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvBrowserEngine"`
	YmPvBrowserEngine1 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvBrowserEngine1"`
	YmPvBrowserEngine2 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvBrowserEngine2"`
	YmPvBrowserEngine3 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvBrowserEngine3"`
	YmPvBrowserEngine4 struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvBrowserEngine4"`
	YmPvBrowserLanguage struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvBrowserLanguage"`
	YmPvClientTimeZone struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvClientTimeZone"`
	YmPvCookieEnabled struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvCookieEnabled"`
	YmPvDeviceCategory struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvDeviceCategory"`
	YmPvFlashMajor struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvFlashMajor"`
	YmPvFlashMinor struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvFlashMinor"`
	YmPvFrom struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvFrom"`
	YmPvIPAddress struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvIpAddress"`
	YmPvJavascriptEnabled struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvJavascriptEnabled"`
	YmPvMobilePhone struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvMobilePhone"`
	YmPvMobilePhoneModel struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvMobilePhoneModel"`
	YmPvOpenstatAd struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvOpenstatAd"`
	YmPvOpenstatCampaign struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvOpenstatCampaign"`
	YmPvOpenstatService struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvOpenstatService"`
	YmPvOpenstatSource struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvOpenstatSource"`
	YmPvOperatingSystem struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvOperatingSystem"`
	YmPvOperatingSystemRoot struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvOperatingSystemRoot"`
	YmPvPhysicalScreenHeight struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvPhysicalScreenHeight"`
	YmPvPhysicalScreenWidth struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvPhysicalScreenWidth"`
	YmPvRegionCity struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvRegionCity"`
	YmPvRegionCountry struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvRegionCountry"`
	YmPvScreenColors struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvScreenColors"`
	YmPvScreenFormat struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvScreenFormat"`
	YmPvScreenHeight struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvScreenHeight"`
	YmPvScreenOrientation struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvScreenOrientation"`
	YmPvScreenWidth struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvScreenWidth"`
	YmPvWindowClientHeight struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvWindowClientHeight"`
	YmPvWindowClientWidth struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvWindowClientWidth"`
	YmPvParams struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvParams"`
	YmPvLastTrafficSource struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvLastTrafficSource"`
	YmPvLastSearchEngine struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvLastSearchEngine"`
	YmPvLastSearchEngineRoot struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvLastSearchEngineRoot"`
	YmPvLastAdvEngine struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvLastAdvEngine"`
	YmPvArtificial struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvArtificial"`
	YmPvPageCharset struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvPageCharset"`
	YmPvLink struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvLink"`
	YmPvDownload struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvDownload"`
	YmPvNotBounce struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvNotBounce"`
	YmPvEvent struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvEvent"`
	YmPvLastSocialNetwork struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvLastSocialNetwork"`
	YmPvHTTPError struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvHttpError"`
	YmPvClientID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvClientID"`
	YmPvNetworkType struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvNetworkType"`
	YmPvLastSocialNetworkProfile struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvLastSocialNetworkProfile"`
	YmPvGoalsID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvGoalsID"`
	YmPvShareService struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvShareService"`
	YmPvShareURL struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvShareURL"`
	YmPvShareTitle struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvShareTitle"`
	YmPvIFrame struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvIFrame"`
	YmPvDate struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvDate"`
	YmPvGCLID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvGCLID"`
	YmPvRegionCityID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvRegionCityID"`
	YmPvRegionCountryID struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvRegionCountryID"`
	YmPvIsPageView struct {
		Name  string `mapstructure:"name"`
		Value string `mapstructure:"value"`
	} `mapstructure:"YmPvIsPageView"`
}
