FROM google/cloud-sdk:alpine

ENV GOLANG_VERSION=1.9.6
ENV GOLANG_DOWNLOAD_SHA256=d1eb07f99ac06906225ac2b296503f06cc257b472e7d7817b8f822fe3766ebfe

RUN mkdir -p /data/www/gaeapp
RUN mkdir -p /data/www/src
WORKDIR /data/www/gaeapp

ENV GOPATH /data/www
ENV PATH $PATH:/usr/local/go/bin:/data/www/bin

# install git
RUN apk add --update --no-cache \
		sudo \
		git
		
# install go
RUN 	curl -o go.tgz -sSL "https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz" && \
	echo "${GOLANG_DOWNLOAD_SHA256} *go.tgz" | sha256sum -c - && \
	tar -C /usr/local -xzf go.tgz && \
	rm go.tgz 
	
# install GAE for Go SDK
RUN gcloud components install app-engine-go

RUN chmod +x /google-cloud-sdk/platform/google_appengine/go* && \
        chmod +x /google-cloud-sdk/platform/google_appengine/*.py
        
# for saving gcloud config
RUN gcloud config set core/disable_usage_reporting true && \
    gcloud config set component_manager/disable_update_check true && \
    gcloud config set metrics/environment github_docker_image


