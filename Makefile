create:
	docker build -t movies .

run:
	@docker run -d -p 8080:8080 --rm --name movies_container movies