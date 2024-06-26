package server

import (
	"io"
	"net/http"

	"github.com/coolestowl/edge-collector/configs"
	"github.com/coolestowl/edge-collector/db"
	"github.com/coolestowl/edge-collector/env"
)

func MonitorInit() error {
	resp, err := http.Get(env.CONFIG_URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := configs.Init(data); err != nil {
		return err
	}

	if len(env.CLICKHOUSE_DSN) > 0 {
		db.InitClickHouse(env.CLICKHOUSE_DSN)
	}

	return nil
}
