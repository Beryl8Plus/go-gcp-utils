FROM google/cloud-sdk

RUN apt-get update && apt-get install -y \
	git

RUN wget https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz
#COPY go1.6.linux-amd64.tar.gz go1.6.linux-amd64.tar.gz
RUN tar xvf go1.6.linux-amd64.tar.gz && chown -R root:root ./go && mv go /usr/local && rm go1.6.linux-amd64.tar.gz

RUN wget https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.54.zip
#COPY go_appengine_sdk_linux_amd64-1.9.54.zip go_appengine_sdk_linux_amd64-1.9.54.zip
RUN unzip go_appengine_sdk_linux_amd64-1.9.54.zip && chown -R root:root ./go_appengine && mv go_appengine /usr/local && rm go_appengine_sdk_linux_amd64-1.9.54.zip

ENV GOROOT /usr/local/go
ENV GOPATH /go/
ENV GOAPPPATH /usr/local/go_appengine/
ENV PATH ${PATH}:$GOROOT/bin:$GOAPPPATH

RUN go get google.golang.org/appengine
RUN go get cloud.google.com/go/logging

ENV HOME /home/default
ENV GCLOUD_PROJECTID chatbot-v1
ENV PROJECT_PATH go/src/github.com/Beryl8Plus/go-gcp-utils

EXPOSE 8080 8000 8081
WORKDIR $PROJECT_PATH

ADD bashrc /home/default/.bashrc
CMD bash
