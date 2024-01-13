build-all: build-generate-resume build-generate-personal-history

build-generate-resume:
	cd ./src/cmd/generate_resume && go build -o generate_resume
	cd ./src/cmd/generate_resume && mv generate_resume ../../../

build-generate-personal-history:
	cd ./src/cmd/generate_personal_history && go build -o generate_personal_history
	cd ./src/cmd/generate_personal_history && mv generate_personal_history ../../../