// Copyright 2019 Tobias Guggenmos
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	kitlog "github.com/go-kit/log"
	"github.com/prometheus-community/promql-langserver/config"
	promClient "github.com/prometheus-community/promql-langserver/prometheus"
	"github.com/prometheus/common/route"

	"github.com/prometheus-community/promql-langserver/langserver"
	"github.com/prometheus-community/promql-langserver/rest"
)

func main() {
	configFilePath := flag.String("config-file", "", "Configuration file for the language server. If unset, the configuration will be retrieved from environment")

	flag.Parse()

	conf, err := config.ReadConfig(*configFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading config file:", err.Error())
		os.Exit(1)
	}
	if conf.RESTAPIPort != 0 {
		fmt.Fprintln(os.Stderr, "REST API: Listening on port ", conf.RESTAPIPort)
		prometheusClient, err := promClient.NewClient(conf.PrometheusURL, time.Duration(conf.MetadataLookbackInterval))
		if err != nil {
			log.Fatal(err)
		}

		var logger kitlog.Logger

		switch conf.LogFormat {
		case config.JSONFormat:
			logger = kitlog.NewJSONLogger(os.Stderr)
		case config.TextFormat:
			logger = kitlog.NewLogfmtLogger(os.Stderr)
		default:
			log.Fatalf(`invalid log format: "%s"`, conf.LogFormat)
		}

		logger = kitlog.NewSyncLogger(logger)
		// create the http router
		router := route.New()
		// create the api
		api, err := rest.NewLangServerAPI(context.Background(), prometheusClient, logger, true)
		if err != nil {
			log.Fatal(err)
		}
		// register the different route
		api.Register(router, "")
		// start the http server
		err = http.ListenAndServe(fmt.Sprint(":", conf.RESTAPIPort), router)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, s := langserver.StdioServer(context.Background(), conf)
		if err := s.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
