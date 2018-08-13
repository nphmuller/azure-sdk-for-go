// Package contentmoderator implements the Azure ARM Contentmoderator service API version 1.0.
//
// You use the API to scan your content as it is generated. Content Moderator then processes your content and sends the
// results along with relevant information either back to your systems or to the built-in review tool. You can use this
// information to take decisions e.g. take it down, send to human judge, etc.
//
// When using the API, images need to have a minimum of 128 pixels and a maximum file size of 4MB.
// Text can be at most 1024 characters long.
// If the content passed to the text API or the image API exceeds the size limits, the API will return an error code
// that informs about the issue.
package contentmoderator

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

// BaseClient is the base client for Contentmoderator.
type BaseClient struct {
	autorest.Client
	Endpoint string
}

// New creates an instance of the BaseClient client.
func New(endpoint string) BaseClient {
	return NewWithoutDefaults(endpoint)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(endpoint string) BaseClient {
	return BaseClient{
		Client:   autorest.NewClientWithUserAgent(UserAgent()),
		Endpoint: endpoint,
	}
}
