package redis

const (
	Prefix = "scaffold:" //项目key前缀
)

func getRedisKey(key string) string {
	return Prefix + key
}
