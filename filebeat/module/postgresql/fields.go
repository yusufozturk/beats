// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package postgresql

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "postgresql", asset.ModuleFieldsPri, AssetPostgresql); err != nil {
		panic(err)
	}
}

// AssetPostgresql returns asset data.
// This is the base64 encoded gzipped contents of module/postgresql.
func AssetPostgresql() string {
	return "eJyclEFv2zAMhe/5FQ85tUNrdNhhgDf0UrTAgHZb194DxaJtYbKoSnTX7NcPspsldZwsHk+GhcfvkRR1jp+0yuE5ShUoPtkZIEYs5Zh/738+3N/OZ4CmWATjxbDLcTkDgDvWrSWUHOBViMZVkJqw0cFyhdJYitkMiDUHWRTsSlPlkNDSDCgNWR3zLt85nGpo4CaFrDzlqAK3/vXPiJs+brp8KAM3AyOdhxTbyG2s5Wor0S7zIPct+VAb1jG0sW1FTENRVOPfnCa8D1QoIZ3jY/Yhu9g532svxWNNm9Qbp8meNY6yUS8FB1oYPUjWt8eyq6ZZuOJAGGRbk7QStVSRBiJ6UY3vbmSz0sv5NN5X1RC4HM+9Bj+1FFZ7qQ/Xt9dXj3iHmx/f7tBGCvHTRBf3CYAoSqghJ+ON7lwsotBw6hsrac1oIvthTUVKjV81ObTdstKLkNOkezJ8YOGCLU7YdT376/es2286w9I4DQ6gFypaodMDdaTv/XVoXkRpZHHRx/v/mWplnslBGOq1gH+XluFLCSMwEfPPrUse9eX8LJXkWOADxVTuDk9qE/uFTVJTOQ6kN1McbQOFwCErWA/bcNzqXCc5khyBpA2ONJar9asScWJKKLc6PWwibftvduMWlDUqDk68kjoHPadbukfcmCqo3uTrIz7CrQMpve/Z2A/2gQuKMfM7ymOoaTcnApMkG7mrx+AsPZOdyLNcZWO6Y3gNxaiqqbMcVw15fwIAAP//ui0d+Q=="
}