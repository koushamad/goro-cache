package configProvider

import (
	"github.com/koushamad/goro-cache/config"
	"github.com/koushamad/goro-core/app/conf"
)

func Load() {
	conf.Add("Catch", config.Cache)
}
