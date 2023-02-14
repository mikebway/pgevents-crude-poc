// This file was derived from source code provided by Google under the following license:
//
// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// The original code was found here on February 14, 2023:
//
// https://github.com/GoogleCloudPlatform/golang-samples/tree/main/eventarc/audit_storage

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestPostgresEvents(t *testing.T) {
	tests := []struct {
		subject string
		want    string
	}{
		{subject: "don't know what this should be yet", want: "Detected change in Postgres database: don't know what this should be yet\n"},
	}
	for _, test := range tests {
		r, w, _ := os.Pipe()
		log.SetOutput(w)
		defer log.SetOutput(os.Stderr)

		originalFlags := log.Flags()
		defer log.SetFlags(originalFlags)
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

		payload := strings.NewReader("{}")
		req := httptest.NewRequest("POST", "/", payload)
		req.Header.Set("ce-subject", test.subject)
		rr := httptest.NewRecorder()
		PostgresEvents(rr, req)

		w.Close()

		if code := rr.Result().StatusCode; code == http.StatusBadRequest {
			t.Errorf("PostgresEvents(%q) invalid input, status code (%q)", test.subject, code)
		}

		out, err := ioutil.ReadAll(r)
		if err != nil {
			t.Fatalf("ReadAll: %v", err)
		}
		if got := string(out); got != test.want {
			t.Errorf("PostgresEvents(%q): got %q, want %q", test.subject, got, test.want)
		}
	}
}
