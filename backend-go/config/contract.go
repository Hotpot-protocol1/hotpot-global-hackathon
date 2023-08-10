package config

type Contract struct {
	ProxyContract       string `mapstructure:"API_PROXY_CONTRACT_ADDRESS"`
	MarketplaceContract string `mapstructure:"API_MARKETPLACE_CONTRACT_ADDRESS"`
	OperatorPrivKey     string `mapstructure:"API_OPERATOR_PRIV_KEY"`
}
