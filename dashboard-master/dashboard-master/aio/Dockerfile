# Copyright 2017 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Main Dockerfile of the project. It creates an image that serves the application. This image should
# be built from the dist directory.

ARG BUILDPLATFORM

FROM --platform=$BUILDPLATFORM busybox:1.35 as builder

# Create a nonroot user for final image
RUN adduser -D -u 10001 nonroot

# Use a distroless image as the base for the final image.
FROM gcr.io/distroless/static-debian11 as final

# Add all files from current working directory to the root of the image, i.e., copy dist directory
# layout to the root directory.
ADD . /

# Copy nonroot user
COPY --from=builder /etc/passwd /etc/passwd
USER nonroot

# The port that the application listens on.
EXPOSE 9090 8443

HEALTHCHECK --interval=30s --timeout=3s \
  CMD curl --insecure https://localhost:8443/ || exit 1

ENTRYPOINT ["/dashboard", "--insecure-bind-address=0.0.0.0", "--bind-address=0.0.0.0"]
