FROM golang:stretch as builder

WORKDIR /opt/app/

COPY config/Config.yaml /opt/app/config/Config.yaml

# RUN cd focusapi && GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -o /opt/app/focusapi main.go

RUN apt-get install git && \
	git clone https://github.com/codeandcode0x/goAlltime.git && \
	cd goAlltime/go-web/focusapi/ && \
    CGO_ENABLED=0 GOOS=linux go build -o /opt/app/focusapi-cgi main.go

FROM ubuntu:16.04 as prod

WORKDIR /root/

COPY --from=0 /opt/app/config/Config.yaml  /root/config/Config.yaml
COPY --from=0 /opt/app/focusapi-cgi  /root/focusapi-cgi

RUN ls -all
RUN chmod 777 focusapi-cgi

CMD ["/bin/sh", "-c", "/root/focusapi-cgi"]