PROJECT_NAME={{.ProjectName}}
MODULE_NAME={{.ModuleName}}
WORKING_DIRECTORY={{.WorkingDirectory}}
PROJECT_FULL_PATH={{.ProjectFullPath}}

.DEFAULT_GOAL := build

{{ if eq .Project "simple"}}
build:
	@go build .
{{- end -}}

{{ if eq .Project "default"}}
build:
	@go build -o {{.ProjectName}} ./cmd/main.go
{{- end -}}

dependencies:
	@go mod download

{{ if .isDocker -}}
docker-build:
	docker build -t {{.ProjectName}}:latest .
{{- end }}

{{ if .isKubernetes -}}
apply-kubernetes:
	kubectl apply -f app-deployment.yaml
{{- end }}