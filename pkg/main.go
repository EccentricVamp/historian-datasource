package main

import (
	"bhw-csgit/cs/grafana-historian/pkg/plugin"
	"os"

	"github.com/grafana/grafana-plugin-sdk-go/backend/datasource"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
)

func main() {
	err := datasource.Manage("clevelandcliffs-historian-datasource", plugin.NewDatasource, datasource.ManageOpts{})

	if err != nil {
		log.DefaultLogger.Error(err.Error())
		os.Exit(1)
	}
}
