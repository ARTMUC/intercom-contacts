package intercom

type ApiErrorReponse struct {
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
	Type string `json:"type"`
}

type Location struct {
	Type          string `json:"type"`
	Country       string `json:"country"`
	Region        string `json:"region"`
	City          string `json:"city"`
	CountryCode   string `json:"country_code"`
	ContinentCode string `json:"continent_code"`
}

type SocialProfiles struct {
	Type string     `json:"type"`
	Data []struct{} `json:"data"`
}

type ListData struct {
	Type       string     `json:"type"`
	Data       []struct{} `json:"data"`
	URL        string     `json:"url"`
	TotalCount int        `json:"total_count"`
	HasMore    bool       `json:"has_more"`
}

type Contact struct {
	Type                      string            `json:"type"`
	ID                        string            `json:"id"`
	WorkspaceID               string            `json:"workspace_id"`
	ExternalID                string            `json:"external_id"`
	Role                      string            `json:"role"`
	Email                     string            `json:"email"`
	Phone                     *string           `json:"phone"`
	Name                      string            `json:"name"`
	Avatar                    *string           `json:"avatar"`
	OwnerID                   *string           `json:"owner_id"`
	SocialProfiles            SocialProfiles    `json:"social_profiles"`
	HasHardBounced            bool              `json:"has_hard_bounced"`
	MarkedEmailAsSpam         bool              `json:"marked_email_as_spam"`
	UnsubscribedFromEmails    bool              `json:"unsubscribed_from_emails"`
	CreatedAt                 int64             `json:"created_at"`
	UpdatedAt                 int64             `json:"updated_at"`
	SignedUpAt                int64             `json:"signed_up_at"`
	LastSeenAt                int64             `json:"last_seen_at"`
	LastRepliedAt             *int64            `json:"last_replied_at"`
	LastContactedAt           *int64            `json:"last_contacted_at"`
	LastEmailOpenedAt         *int64            `json:"last_email_opened_at"`
	LastEmailClickedAt        *int64            `json:"last_email_clicked_at"`
	LanguageOverride          *string           `json:"language_override"`
	Browser                   string            `json:"browser"`
	BrowserVersion            string            `json:"browser_version"`
	BrowserLanguage           string            `json:"browser_language"`
	OS                        string            `json:"os"`
	Location                  Location          `json:"location"`
	AndroidAppName            *string           `json:"android_app_name"`
	AndroidAppVersion         *string           `json:"android_app_version"`
	AndroidDevice             *string           `json:"android_device"`
	AndroidOSVersion          *string           `json:"android_os_version"`
	AndroidSDKVersion         *string           `json:"android_sdk_version"`
	AndroidLastSeenAt         *int64            `json:"android_last_seen_at"`
	IosAppName                *string           `json:"ios_app_name"`
	IosAppVersion             *string           `json:"ios_app_version"`
	IosDevice                 *string           `json:"ios_device"`
	IosOSVersion              *string           `json:"ios_os_version"`
	IosSDKVersion             *string           `json:"ios_sdk_version"`
	IosLastSeenAt             *int64            `json:"ios_last_seen_at"`
	CustomAttributes          map[string]string `json:"custom_attributes"`
	Tags                      ListData          `json:"tags"`
	Notes                     ListData          `json:"notes"`
	Companies                 ListData          `json:"companies"`
	OptedOutSubscriptionTypes ListData          `json:"opted_out_subscription_types"`
	OptedInSubscriptionTypes  ListData          `json:"opted_in_subscription_types"`
	UtmCampaign               *string           `json:"utm_campaign"`
	UtmContent                *string           `json:"utm_content"`
	UtmMedium                 *string           `json:"utm_medium"`
	UtmSource                 *string           `json:"utm_source"`
	UtmTerm                   *string           `json:"utm_term"`
	Referrer                  string            `json:"referrer"`
	SMSConsent                bool              `json:"sms_consent"`
	UnsubscribedFromSMS       bool              `json:"unsubscribed_from_sms"`
}

type ContactList struct {
	Type string    `json:"type"`
	Data []Contact `json:"data"`
}
