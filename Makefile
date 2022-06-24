prisma-generate:
	go get github.com/prisma/prisma-client-go && go run github.com/prisma/prisma-client-go generate
prisma-sync:
	go run github.com/prisma/prisma-client-go db push
prisma-format:
	prisma format
prisma-migrate:
	go install github.com/prisma/prisma-client-go
	go run github.com/prisma/prisma-client-go migrate dev --name init
prisma-seed:
	chmod +x seed.sh
	./seed.sh