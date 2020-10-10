// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"actionflow/config"
)

type Response struct {
	Code   int    `json:"code"`
	Expire string `json:"expire"`
	Token  string `json:"token"`
}

func TestRouter(t *testing.T) {
	var err error
	var rec *httptest.ResponseRecorder
	var req *http.Request

	r := Router{}

	err = r.initAuth()
	assert.Equal(t, nil, err)

	err = r.initRoute()
	assert.Equal(t, nil, err)

	err = r.setRoute()
	assert.Equal(t, nil, err)

	rec = httptest.NewRecorder()
	data := url.Values{}
	data.Set("username", "admin")
	data.Set("password", "admin")
	req, _ = http.NewRequest("POST", "/auth/login", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	r.engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEqual(t, nil, rec.Body.String())

	var resp Response
	decoder := json.NewDecoder(rec.Body)
	err = decoder.Decode(&resp)
	assert.Equal(t, nil, err)

	token := resp.Token

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/auth/refresh", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	r.engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEqual(t, nil, rec.Body.String())

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/accounts/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	r.engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEqual(t, nil, rec.Body.String())

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/accounts/?q=admin", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	r.engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEqual(t, nil, rec.Body.String())

	rec = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/config/server/version", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	r.engine.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "\""+config.Version+"-build-"+config.Build+"\"", rec.Body.String())
}
