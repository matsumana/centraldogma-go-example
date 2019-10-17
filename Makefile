setup-centraldogma:
	$(eval CENTRAL_DOGMA_CONTAINER := $(shell docker container ls | grep 'line/centraldogma:' | awk '{print $$1}'))
	docker exec -it $(CENTRAL_DOGMA_CONTAINER) curl -X POST -H 'authorization: bearer anonymous' -H 'Content-Type: application/json' -d '{"name": "fukuoka-go"}' http://localhost:36462/api/v1/projects
	docker exec -it $(CENTRAL_DOGMA_CONTAINER) curl -X POST -H 'authorization: bearer anonymous' -H 'Content-Type: application/json' -d '{"name": "demo"}' http://localhost:36462/api/v1/projects/fukuoka-go/repos
	docker exec -it $(CENTRAL_DOGMA_CONTAINER) curl -X POST -H 'authorization: bearer anonymous' -H 'Content-Type: application/json' -d '{"commitMessage": {"summary": "Add initial data", "detail": {"content": "", "markup": "PLAINTEXT"}}, "file": {"name": "config.json", "type": "TEXT", "content": "{\"config1\": \"foo\"}", "path": "/config.json"}}' http://localhost:36462/api/v0/projects/fukuoka-go/repositories/demo/files/revisions/head
