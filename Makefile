RUN_NAME="go_schedule"

build:
	gofmt -w .
	chmod +x build.sh
	sh build.sh

relog:
	rm -rf output/${RUN_NAME}_log
	mkdir output/${RUN_NAME}_log

clean:
	rm -rf output

run:
	cd ./output && ./${RUN_NAME}
