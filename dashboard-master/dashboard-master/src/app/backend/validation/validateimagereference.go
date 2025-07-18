// Copyright 2017 The Kubernetes Authors.
//
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

package validation

import (
	distributionref "github.com/distribution/reference"
)

// ImageReferenceValiditySpec is a specification of an image reference validation request.
type ImageReferenceValiditySpec struct {
	// Reference of the image
	Reference string `json:"reference"`
}

// ImageReferenceValidity describes validity of the image reference.
type ImageReferenceValidity struct {
	// True when the image reference is valid.
	Valid bool `json:"valid"`
	// Error reason when image reference is valid
	Reason string `json:"reason"`
}

// ValidateImageReference validates image reference.
func ValidateImageReference(spec *ImageReferenceValiditySpec) (*ImageReferenceValidity, error) {
	s := spec.Reference
	ref, err := distributionref.ParseNormalizedNamed(s)
	if err != nil {
		return &ImageReferenceValidity{Valid: false, Reason: err.Error()}, nil
	}

	// Check if the reference has a domain, otherwise it's not a valid image reference.
	if distributionref.Domain(ref) == "" {
		return &ImageReferenceValidity{Valid: false, Reason: "Image reference must contain a domain"}, nil
	}

	// Check if the reference has a path, otherwise it's not a valid image reference.
	if distributionref.Path(ref) == "" {
		return &ImageReferenceValidity{Valid: false, Reason: "Image reference must contain a path"}, nil
	}

	return &ImageReferenceValidity{Valid: true}, nil
}
