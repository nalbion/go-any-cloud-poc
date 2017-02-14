FROM gcr.io/google_appengine/golang

ENV GOPATH /go

RUN apt-get update \
    && apt-get install -y make git unzip \
    && go get -u github.com/kardianos/govendor \
    && go get -u github.com/jstemmer/go-junit-report \
    && go get -u github.com/axw/gocov/gocov \
    && go get -u github.com/marinbek/gocov-xml \
    && go get -u github.com/wadey/gocovmerge

# install the Google Go SDK
RUN curl -sSL https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.40.zip -q -o go_appengine_sdk.zip && \
    unzip -q go_appengine_sdk.zip -d /usr/lib/ && \
    rm go_appengine_sdk.zip
#ENV PATH /usr/lib/go_appengine:/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin