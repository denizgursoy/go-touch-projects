PROJECT_NAME={{.ProjectName}}
MODULE_NAME={{.ModuleName}}
WORKING_DIRECTORY={{.WorkingDirectory}}
PROJECT_FULL_PATH={{.ProjectFullPath}}

.DEFAULT_GOAL := build

build:
	@go build .

fmt:
	@go fmt ./...

test:
	@go test -v ./...

get:
	@go mod download

{{ if .isDocker -}}
docker-build:
	docker build -t {{.ProjectName}}:latest .
{{- end }}

{{ if .isKubernetes -}}
apply-kubernetes:
	kubectl apply -f app-deployment.yaml
{{- end }}