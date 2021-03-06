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

package impl

import (
	"strings"

	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

type newUserRequestBody struct {
	UserID   string `json:"userId"`
	Password string `json:"password"`
	IsAdmin  string `json:"isAdmin"`
}

// AddMIUser adds a new user to the micro integrator in a given environment
func AddMIUser(env, userName, password, isAdmin string) (interface{}, error) {
	isAdmin = resolveIsAdmin(isAdmin)
	body := newUserRequestBody{
		UserID:   userName,
		Password: password,
		IsAdmin:  isAdmin,
	}
	url := utils.GetMIManagementEndpointOfResource(utils.MiManagementUserResource, env, utils.MainConfigFilePath)
	return addNewMIUser(env, url, body)
}

// DeleteMIUser deletes a user from a micro integrator in a given environment
func DeleteMIUser(env, userName string) (interface{}, error) {
	url := utils.GetMIManagementEndpointOfResource(utils.MiManagementUserResource, env, utils.MainConfigFilePath) + "/" + userName
	return deleteMIUser(url, env)
}

func addNewMIUser(env, url string, body interface{}) (string, error) {
	resp, err := invokePOSTRequestWithRetry(env, url, body)
	return handleResponse(resp, err, url, "status", "Error")
}

func deleteMIUser(url, env string) (string, error) {
	resp, err := invokeDELETERequestWithRetry(url, env)
	return handleResponse(resp, err, url, "status", "Error")
}

func resolveIsAdmin(isAdminConsoleInput string) string {
	if len(strings.TrimSpace(isAdminConsoleInput)) == 0 {
		return "false"
	}
	yesResponses := []string{"y", "yes"}
	if containsString(yesResponses, strings.TrimSpace(isAdminConsoleInput)) {
		return "true"
	}
	return "false"
}

func containsString(slice []string, element string) bool {
	for _, elem := range slice {
		// case in-sensitive comparison
		if strings.EqualFold(elem, element) {
			return true
		}
	}
	return false
}
