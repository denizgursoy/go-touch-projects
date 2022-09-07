PROJECT_NAME={{.ProjectName}}
MODULE_NAME={{.ModuleName}}
WORKING_DIRECTORY={{.WorkingDirectory}}
PROJECT_FULL_PATH={{.ProjectFullPath}}


{{ if .isDocker -}}
docker-build:
	docker build . -t {{.ProjectName}}:latest
{{- end }}

{{ if .isKubernetes -}}
apply-kubernetes:
	kubectl apply -f app-deployment.yaml
{{- end }}