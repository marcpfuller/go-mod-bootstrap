//
// Copyright (C) 2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/go-mod-bootstrap/v3/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	defaultCORSAllowCredentials = "true"
	defaultCORSAllowedOrigin    = "https://localhost"
	defaultCORSAllowedMethods   = "GET, POST, PUT, PATCH, DELETE"
	defaultCORSAllowedHeaders   = "Authorization, Accept, Accept-Language, Content-Language, Content-Type, X-Correlation-ID"
	defaultCORSExposeHeaders    = "Cache-Control, Content-Language, Content-Length, Content-Type, Expires, Last-Modified, Pragma, X-Correlation-ID"
	defaultCORSMaxAge           = "3600"
)

var defaultCORSInfo = config.CORSConfigurationInfo{CORSAllowCredentials: true, CORSAllowedOrigin: defaultCORSAllowedOrigin, CORSAllowedMethods: defaultCORSAllowedMethods, CORSAllowedHeaders: defaultCORSAllowedHeaders, CORSExposeHeaders: defaultCORSExposeHeaders, CORSMaxAge: 3600}

func TestProcessCORS(t *testing.T) {
	corsInfo := defaultCORSInfo

	simpleHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	tests := []struct {
		Name                                  string
		EnableCORS                            bool
		HttpMethod                            string
		Origin                                string
		AccessControlRequestMethod            string
		ExpectedAccessControlExposeHeaders    string
		ExpectedAccessControlAllowOrigin      string
		ExpectedAccessControlAllowCredentials string
		ExpectedVary                          string
	}{
		{"not enable CORS", false, http.MethodGet, "http://test.com", http.MethodGet, "", "", "", ""},
		{"enable CORS without Origin header", true, http.MethodGet, "", http.MethodGet, "", "", "", ""},
		{"enable CORS and receive a preflight request", true, http.MethodOptions, "http://test.com", http.MethodGet, "", defaultCORSAllowedOrigin, defaultCORSAllowCredentials, Origin},
		{"enable CORS and receive an actual request", true, http.MethodGet, "http://test.com", "", defaultCORSExposeHeaders, defaultCORSAllowedOrigin, defaultCORSAllowCredentials, Origin},
	}

	for _, testCase := range tests {
		t.Run(testCase.Name, func(t *testing.T) {
			corsInfo.EnableCORS = testCase.EnableCORS
			corsMiddleware := ProcessCORS(corsInfo)
			handler := corsMiddleware(simpleHandler)

			req, err := http.NewRequest(testCase.HttpMethod, "/", http.NoBody)
			require.NoError(t, err)

			if len(testCase.Origin) > 0 {
				req.Header.Set(Origin, testCase.Origin)
			}
			if len(testCase.AccessControlRequestMethod) > 0 {
				req.Header.Set(AccessControlRequestMethod, testCase.AccessControlRequestMethod)
			}

			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, req)
			resp := recorder.Result()

			require.Equal(t, http.StatusOK, resp.StatusCode, "http status code is not as expected")
			assert.Equal(t, testCase.ExpectedAccessControlExposeHeaders, resp.Header.Get(AccessControlExposeHeaders), "http header Access-Control-Expose-Headers is not as expected")
			assert.Equal(t, testCase.ExpectedAccessControlAllowOrigin, resp.Header.Get(AccessControlAllowOrigin), "http header Access-Control-Allow-Origin is not as expected")
			assert.Equal(t, testCase.ExpectedAccessControlAllowCredentials, resp.Header.Get(AccessControlAllowCredentials), "http header Access-Control-Expose-Headers is not as expected")
			assert.Equal(t, testCase.ExpectedVary, resp.Header.Get(Vary), "http header Vary is not as expected")
		})
	}
}

func TestHandlePreflight(t *testing.T) {
	corsInfo := defaultCORSInfo

	tests := []struct {
		Name                              string
		EnableCORS                        bool
		Origin                            string
		ExpectedAccessControlAllowMethods string
		ExpectedAccessControlAllowHeaders string
		ExpectedAccessControlMaxAge       string
	}{
		{"not enable CORS", false, "http://test.com", "", "", ""},
		{"enable CORS without Origin header", true, "", "", "", ""},
		{"enable CORS and receive a preflight request", true, "http://test.com", defaultCORSAllowedMethods, defaultCORSAllowedHeaders, defaultCORSMaxAge},
	}

	for _, testCase := range tests {
		t.Run(testCase.Name, func(t *testing.T) {
			corsInfo.EnableCORS = testCase.EnableCORS
			handler := http.HandlerFunc(HandlePreflight(corsInfo))

			req, err := http.NewRequest(http.MethodGet, "/", http.NoBody)
			require.NoError(t, err)

			if len(testCase.Origin) > 0 {
				req.Header.Set(Origin, testCase.Origin)
			}

			recorder := httptest.NewRecorder()
			handler.ServeHTTP(recorder, req)
			resp := recorder.Result()

			require.Equal(t, http.StatusOK, resp.StatusCode, "http status code is not as expected")
			assert.Equal(t, testCase.ExpectedAccessControlAllowMethods, resp.Header.Get(AccessControlAllowMethods), "http header Access-Control-Allow-Methods is not as expected")
			assert.Equal(t, testCase.ExpectedAccessControlAllowHeaders, resp.Header.Get(AccessControlAllowHeaders), "http header Access-Control-Allow-Headers is not as expected")
			assert.Equal(t, testCase.ExpectedAccessControlMaxAge, resp.Header.Get(AccessControlMaxAge), "http header Access-Control-Max-Age is not as expected")
		})
	}
}
