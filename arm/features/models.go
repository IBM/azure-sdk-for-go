package features

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.11.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/azure/go-autorest/autorest"
)

// List of previewed features.
type FeatureOperationsListResult struct {
	autorest.Response `json:"-"`
	Value             []FeatureResult `json:"value,omitempty"`
	NextLink          string          `json:"nextLink,omitempty"`
}

// Previewed feature information.
type FeatureResult struct {
	autorest.Response `json:"-"`
	Name              string `json:"name,omitempty"`
	Properties        struct {
		State string `json:"state,omitempty"`
	} `json:"properties,omitempty"`
	Id   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}