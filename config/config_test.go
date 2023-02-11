// Copyright 2020 The Prometheus Authors
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

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalENV(t *testing.T) {
	testSuites := []struct {
		title     string
		variables map[string]string
		expected  *Config
	}{
		{
			title:     "empty config",
			variables: map[string]string{},
			expected: &Config{
				ActivateRPCLog:           false,
				LogFormat:                TextFormat,
				MetadataLookbackInterval: defaultInterval,
			},
		},
		{
			title: "full config",
			variables: map[string]string{
				"LANGSERVER_ACTIVATERPCLOG":           "true",
				"LANGSERVER_PROMETHEUSURL":            "http://localhost:9090",
				"LANGSERVER_RESTAPIPORT":              "8080",
				"LANGSERVER_LOGFORMAT":                "json",
				"LANGSERVER_METADATALOOKBACKINTERVAL": "1w",
			},
			expected: &Config{
				ActivateRPCLog:           true,
				PrometheusURL:            "http://localhost:9090",
				RESTAPIPort:              8080,
				LogFormat:                JSONFormat,
				MetadataLookbackInterval: 604800000000000,
			},
		},
	}
	for _, testSuite := range testSuites {
		//nolint
		t.Run(testSuite.title, func(t *testing.T) {
			os.Clearenv()
			//nolint
			for k, v := range testSuite.variables {
				os.Setenv(k, v)
			}
			conf, err := ReadConfig("")
			assert.NoError(t, err)
			assert.Equal(t, testSuite.expected, conf)
		})
	}
}
