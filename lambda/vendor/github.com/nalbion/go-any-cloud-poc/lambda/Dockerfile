FROM eawsy/aws-lambda-go-shim

RUN yum -q -e 0 -y install git \
    && go get -u github.com/kardianos/govendor \
    && go get -u github.com/jstemmer/go-junit-report \
    && go get -u github.com/axw/gocov/gocov \
    && go get -u github.com/marinbek/gocov-xml \
    && go get -u github.com/wadey/gocovmerge
    && go get -u -d github.com/eawsy/aws-lambda-go-core/... \
    && go get -u -d github.com/eawsy/aws-lambda-go-event/...

# Makefile calls `mkdir -p lambda/_gopath/src && cp -r /usr/lib/go_appengine/goroot/src/appengine* lambda/_gopath/src/`
COPY _gopath/src /go/src/

# Deploy
RUN curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py && python get-pip.py && rm get-pip.py
RUN pip install --upgrade awscli
#&& cd /root/.local/lib/python2.7/site-packages/awscli && python setup.py install
