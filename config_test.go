package httpconfig

import (
	"testing"
	"time"
)

func TestGetConfig(t *testing.T) {
	expected := serverConfig{
		Host:    "127.0.0.1",
		Port:    ":8080",
		Timeout: 5 * time.Second,
	}

	t.Run("Get from example.yaml", func(t *testing.T) {
		cfg, err := GetConfig("./example.yaml")
		if err != nil {
			t.Fatal("err is not nil", err)
		}

		if cfg.Server != expected {
			t.Fatal("wrong config", *cfg)
		}
	})
}
