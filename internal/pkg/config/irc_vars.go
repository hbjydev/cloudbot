package config

var (
	IrcNick    = getEnv("CLOUDBOT_IRC_NICK", "baka")
	IrcChannel = getEnv("CLOUDBOT_IRC_CHANNEL", "cloudbot-devel")
)
