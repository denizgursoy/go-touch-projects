PROJECT_NAME={{.ProjectName}}


{{- if .isDocker }}
docker-build:
	docker build . -t {{.ProjectName}}:latest
{{- end }}

{{- if .isKubernetes }}
apply-kubernetes:
	kubectl apply -f app-deployment.yaml
{{- end }}