package config

var (
  GithubDefaultOwner = getEnv("CLOUDBOT_GH_DEFAULT_OWNER", "rocky-linux")
)
