FROM golang:1.14-alpine

ENV KUBEBUILDER_VERSION=2.3.0

RUN apk --update add curl

RUN export OS=$(go env GOOS) \
    && export ARCH=$(go env GOARCH) \
    && curl -L -o - https://go.kubebuilder.io/dl/${KUBEBUILDER_VERSION}/${OS}/${ARCH} | tar -xz -C /tmp/ \
    && mv /tmp/kubebuilder_${KUBEBUILDER_VERSION}_${OS}_${ARCH} /usr/local/kubebuilder

ENV PATH=$PATH:/usr/local/kubebuilder/bin