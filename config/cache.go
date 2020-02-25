package config

import "github.com/koushamad/goro-app/app/helper"

var (
	Cache = map[string]string{
		"driver": helper.Env("CACHE_DRIVER", "file"),
	}

	Redis = map[string]string{
		"host": helper.Env("REDIS_HOST", ""),
		"port": helper.Env("REDIS_PORT", ""),
	}
)
