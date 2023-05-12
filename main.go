package main

import (
	"fmt"
	"github.com/compose-spec/compose-go/dotenv"
	"github.com/shono-io/go-shono/shono"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	LogLevelEnv = "LOG_LEVEL"

	ShonoUrlEnv = "SHONO_URL"

	ShonoOrgEnv    = "SHONO_ORG"
	ShonoKeyEnv    = "SHONO_KEY"
	ShonoSecretEnv = "SHONO_SECRET"
)

func main() {
	if err := dotenv.Load(); err != nil {
		logrus.Panicf("failed to load .env file: %v", err)
	}

	ll := os.Getenv(LogLevelEnv)
	if ll != "" {
		lv, err := logrus.ParseLevel(ll)
		if err != nil {
			logrus.Panicf("failed to parse log level: %v", err)
		} else {
			logrus.SetLevel(lv)
		}
	}

	sc, err := shono.NewClient(
		fmt.Sprintf("agent-%s", os.Getenv(ShonoOrgEnv)),
		shono.WithCredentials(os.Getenv(ShonoKeyEnv), os.Getenv(ShonoSecretEnv)),
		shono.WithOrg(os.Getenv(ShonoOrgEnv)),
		shono.WithUrl(os.Getenv(ShonoUrlEnv)),
	)
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	fmt.Println("connected to the shono stream")
}
