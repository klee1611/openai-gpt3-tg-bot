.PHONY: test deploy

test:
	@cd openaichatbot && \
		go test -v -coverprofile=coverage.out -covermode=atomic -coverpkg=../openaichatbot ./... && \
		go tool cover -html=coverage.out -o coverage.html

deploy:
	@bash deploy.sh
