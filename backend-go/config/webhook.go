package config

type Webhook struct {
	BaseURL                 string `mapstructure:"API_WEBHOOK_BASE_URL"`
	APIKey                  string `mapstructure:"API_WEBHOOK_KEY"`
	WHAddressActivityURL    string `mapstructure:"API_WEBHOOK_ADDRESS_ACTIVITY_URL"`
	WHNFTMetadataUpdatesURL string `mapstructure:"API_WEBHOOK_NFT_METADATA_UPDATES_URL"`
}
