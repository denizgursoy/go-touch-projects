PROJECT_NAME={{.ProjectName}}
MODULE_NAME={{.ModuleName}}
WORKING_DIRECTORY={{.WorkingDirectory}}
PROJECT_FULL_PATH={{.ProjectFullPath}}

.DEFAULT_GOAL := build

.PHONY: build
build:
	@go build .

.PHONY: fmt
fmt:
	@go fmt ./...

.PHONY: test
test:
	@go test -v ./...

.PHONY: get
get:
	@go mod download

{{ if .isDocker -}}
.PHONY: docker
docker:
	docker build -t {{.ProjectName}}:latest .
{{- end }}

{{ if .isKubernetes -}}
.PHONY: kubernetes
kubernetes:
	kubectl apply -f app-deployment.yaml
{{- end }}