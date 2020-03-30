.EXPORT_ALL_VARIABLES:

# Vars
OWNER=https://github.com/migueleliasweb
DOMAIN=migueleliasweb.github.io/cloudwatch-exporter-operator

# Docker vars
DOCKER_RUN=docker run -it --rm -v ${PWD}:/cloudwatch-exporter-operator --workdir=/cloudwatch-exporter-operator
DOCKER_IMAGE_NAME=cloudwatch-exporter-operator
KUBEBUILDER_DOCKER_IMAGE_NAME=${DOCKER_IMAGE_NAME}-kubebuilder

# Kubernetes vars
CRD_API_GROUP=cloudwatch
CRD_API_VERSION=v1
CRD_KIND=CloudwatchExporterOperator

# Docker

.PHONY: build-kubebuilder-image
build-kubebuilder-image:
	docker build -f kubebuilder.Dockerfile -t ${KUBEBUILDER_DOCKER_IMAGE_NAME} .

# Kubebuilder

.PHONY: kubebuilder-run
kubebuilder-run: build-kubebuilder-image
	${DOCKER_RUN} --entrypoint ${ENTRYPOINT} ${KUBEBUILDER_DOCKER_IMAGE_NAME} ${COMMAND}

.PHONY: kubebuilder-init
kubebuilder-init: ENTRYPOINT="kubebuilder"
kubebuilder-init: COMMAND="init --license apache2 --domain ${DOMAIN} --owner ${OWNER}"
kubebuilder-init: kubebuilder-run

.PHONY: kubebuilder-create
kubebuilder-create:
	COMMAND="create api --group ${CRD_API_GROUP} --version ${CRD_API_VERSION} --kind ${CRD_KIND}" $(MAKE) kubebuilder-run