PROJECT_NAME={{.ProjectName}}
MODULE_NAME={{.ModuleName}}
WorkingDirectory={{.WorkingDirectory}}
ProjectFullPath={{.ProjectFullPath}}


{{ if .isDocker -}}
docker-build:
	docker build . -t {{.ProjectName}}:latest
{{- end }}

{{ if .isKubernetes -}}
apply-kubernetes:
	kubectl apply -f app-deployment.yaml
{{- end }}