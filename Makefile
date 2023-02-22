.PHONY: test deploy

test:
	@cd src && \
		go test -v -coverprofile=coverage.out -covermode=atomic -coverpkg=../src ./... && \
		go tool cover -html=coverage.out -o coverage.html

deploy:
	@bash deploy.sh
