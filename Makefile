GOFLAGS = GOOS=windows

output-dir: bin
	mkdir -p bin

bin/starter.exe: output-dir
	$(GOFLAGS) go build cmd/starter/main.go && mv main.exe bin/starter.exe

bin/stopper.exe: output-dir
	$(GOFLAGS) go build cmd/stopper/main.go && mv main.exe bin/stopper.exe

all: bin/starter.exe bin/stopper.exe