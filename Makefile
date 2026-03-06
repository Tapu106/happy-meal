STAGE_DEV=dev
REGION?=us-east-1
FUNC?=handler

.PHONY: build clean deploy-dev logs-dev run-dev run-prod

build:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags="-s -w" -o bootstrap ./cmd/lambda

clean:
	rm -f bootstrap

deploy-dev: build
	npx sls deploy --stage $(STAGE_DEV) --region $(REGION)

logs-dev:
	sls logs -f $(FUNC) --tail --stage $(STAGE_DEV) --region $(REGION)

run-local:
	IS_OFFLINE=true go run ./cmd/lambda/main.go

	

