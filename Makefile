generate:
	oapi-codegen --package=api --generate types,client \
		pkg/eventor/api/openapi.yaml \
		> pkg/eventor/api/api.go
