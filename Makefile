# Makefile for foo

.PHONY: build deploy clean

# Replace with your GCP Project ID and Region
GCP_PROJECT_ID ?= $(GOOGLE_CLOUD_PROJECT)
GCP_REGION ?= us-central1

SERVICE_NAME = foo
BINARY_NAME = foo

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) main.go

deploy: build
	@echo "Deploying $(SERVICE_NAME) to Cloud Run..."
	@mkdir -p .gcloud_config
	@CLOUDSDK_CONFIG=./.gcloud_config gcloud run deploy $(SERVICE_NAME) --source . --project $(GCP_PROJECT_ID) --region $(GCP_REGION)

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
