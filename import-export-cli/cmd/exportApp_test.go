/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package cmd

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/renstrom/dedent"
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

func TestExportApp(t *testing.T) {
	var server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected 'GET', got '%s'\n", r.Method)
		}

		if r.Header.Get(utils.HeaderAccept) != utils.HeaderValueApplicationZip {
			t.Errorf("Expected '"+utils.HeaderValueApplicationZip+"', got '%s'\n",
				r.Header.Get(utils.HeaderContentType))
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set(utils.HeaderContentType, utils.HeaderValueApplicationJSON)
		w.Header().Set(utils.HeaderContentEncoding, utils.HeaderValueGZIP)
		w.Header().Set(utils.HeaderTransferEncoding, utils.HeaderValueChunked)

		body := dedent.Dedent(`
		`)

		w.Write([]byte(body))
	}))
	defer server.Close()

	resp, err := getExportAppResponse("testApp", "admin", server.URL, "")
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode(), "Should return correct status code")
}

func TestWriteApplicationToZip(t *testing.T) {
	name := "sampleApp"
	owner := "admin"
	//environment := "dev"
	response := new(resty.Response)

	exportDirectory, err := ioutil.TempDir("", "")
	assert.Nil(t, err, "should be able to create temp directory")
	defer os.RemoveAll(exportDirectory)
	WriteApplicationToZip(name, owner, exportDirectory, response)
}
