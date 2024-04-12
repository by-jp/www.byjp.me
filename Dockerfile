FROM golang:1.22-alpine

ARG TARGETPLATFORM

ENV PAGEFIND_VERSION=1.1.0
ENV TASK_VERSION=3.36.0

RUN apk update && \
    apk add ca-certificates wget git hugo && \
    update-ca-certificates && \
    rm -rf /var/cache/apk/*

RUN wget -q -O pagefind.tar.gz "https://github.com/CloudCannon/pagefind/releases/download/v${PAGEFIND_VERSION}/pagefind-v${PAGEFIND_VERSION}-$(if [ "$TARGETPLATFORM" = "linux/arm64" ]; then echo "aarch64"; else echo "x86_64"; fi)-unknown-linux-musl.tar.gz" \
    && tar -xpf pagefind.tar.gz pagefind \
    && mv pagefind /usr/local/bin/pagefind \
    && rm pagefind.tar.gz

RUN wget -q -O task.tar.gz "https://github.com/go-task/task/releases/download/v${TASK_VERSION}/task_${TARGETPLATFORM//\//_}.tar.gz" \
    && tar -xpf task.tar.gz task \
    && mv task /usr/local/bin/task \
    && rm task.tar.gz

WORKDIR /build

CMD ["task", "build"]
