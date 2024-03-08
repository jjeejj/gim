CMDDIR=$(shell pwd)/cmd
bus:
	cd $(CMDDIR)/business || exit && go run main.go

logic:
	cd $(CMDDIR)/logic || exit && go run main.go

connect:
	cd $(CMDDIR)/connect || exit && go run main.go


run_all:
	make bus
	make logic
	make connect

build_business:
	GOOS=linux GOARCH=amd64 go build -o ./bin/gim_business_srv $(CMDDIR)/business/main.go -gcflags=-trimpath=${CMDDIR} -asmflags=-trimpath=${CMDDIR}

build_logic:
	GOOS=linux GOARCH=amd64 go build -o ./bin/gim_logic_srv $(CMDDIR)/logic/main.go -gcflags=-trimpath=${CMDDIR} -asmflags=-trimpath=${CMDDIR}
	#mkdir bin
#	serverName=gim_logic_srv
#    echo "build $(serverName) start"

#  	buildVersion="$(git rev-parse --abbrev-ref HEAD)_$(git rev-parse --short HEAD)";
#  	buildTime="$(date -u '+%Y-%m-%dT%H:%M:%SZ')";
#  	buildCommit="$(git rev-parse --short HEAD)";
#  	buildGoVersion="$(go version)";
#  	GOOS=linux GOARCH=amd64 go build -o ./bin/${serverName} -ldflags "-w -X 'main.BuildGoVersion=${buildGoVersion}' -X 'main.BuildVersion=${buildVersion}' -X 'main.BuildTime=${buildTime}' -X 'main.BuildCommit=${buildCommit}'" $(CMDDIR)/logic/main.go;

build_connect:
	GOOS=linux GOARCH=amd64 go build -o ./bin/gim_connect_srv $(CMDDIR)/connect/main.go -gcflags=-trimpath=${CMDDIR} -asmflags=-trimpath=${CMDDIR}

build_all:
	make build_business
	make build_logic
	make build_connect