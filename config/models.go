package config

type ConfigModel struct {
	RabbitMq    RabbitMqConfig `json:"rabbitMq"`
	ChannelName string         `json:"channelName"`
}

type RabbitMqConfig struct {
	Server   string `json:"server"`
	User     string `json:"user"`
	Password string `json:"password"`
}
