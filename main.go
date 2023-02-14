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
	"fmt"
	"log"
	"net/http"
	"os"
)

// PostgresEvents receives and processes a Cloud Audit Log event with Postgres audit data.
func PostgresEvents(w http.ResponseWriter, r *http.Request) {
	s := fmt.Sprintf("Detected change in Postgres database: %s", string(r.Header.Get("Ce-Subject")))
	log.Printf(s)
	_, _ = fmt.Fprintln(w, s)
}

func main() {
	http.HandleFunc("/", PostgresEvents)
	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Start HTTP server.
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
