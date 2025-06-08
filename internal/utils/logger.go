package utils

import log "github.com/sirupsen/logrus"

func SetupLogger() {
	log.SetFormatter(&log.TextFormatter{
		// DisableColors: true,
		FullTimestamp: true,
	})

}
