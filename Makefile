all: build-generate-resume

build-generate-resume:
	cd ./src/cmd/generate_resume && go build -o generate_resume
	cd ./src/cmd/generate_resume && mv generate_resume ../../../