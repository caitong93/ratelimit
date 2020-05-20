package mocks

//go:generate mockgen -destination ./runtime/snapshot/snapshot.go github.com/lyft/goruntime/snapshot IFace
//go:generate mockgen -destination ./runtime/loader/loader.go github.com/lyft/goruntime/loader IFace
//go:generate mockgen -destination ./config/config.go github.com/envoyproxy/ratelimit/src/config RateLimitConfig,RateLimitConfigLoader
//go:generate mockgen -destination ./redis/redis.go github.com/envoyproxy/ratelimit/src/redis RateLimitCache,Client,TimeSource,JitterRandSource
