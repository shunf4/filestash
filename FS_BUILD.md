See .drone.yml.

Build frontend:
	make build_frontend

Build backend:
	docker start -ai filestash_build_backend_new
	cd filestash
	./build_backend.sh

Release & Update:
	(Quit container)
	./release.sh
	cd docker
	./build.sh
