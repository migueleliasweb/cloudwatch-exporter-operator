# Vars
OWNER=https://github.com/migueleliasweb

# Golang vars
MODULE_NAME=cloudwatchexporteroperator

# Docker vars
DOCKER_RUN=docker run -it --rm
DOCKER_IMAGE_NAME=cloudwatch-exporter-operator
KUBEBUILDER_DOCKER_IMAGE_NAME=${DOCKER_IMAGE_NAME}-kubebuilder

# Kubernetes var
CRD_API_GROUP=cloudwatch
CRD_API_VERSION=v1
CRD_KIND=CloudwatchExporterOperator

# Docker

build-kubebuilder-image:
	docker build -f kubebuilder.Dockerfile -t ${KUBEBUILDER_DOCKER_IMAGE_NAME} .

# Kubebuilder

kubebuilder-run: build-kubebuilder-image
	${DOCKER_RUN} ${KUBEBUILDER_DOCKER_IMAGE_NAME} kubebuilder ${COMMAND}

kubebuilder-init:
	COMMAND="init --license apache2 --domain --owner ${OWNER}" $(MAKE) kubebuilder-run

kubebuilder-create:
	COMMAND="create api --group ${CRD_API_GROUP} --version ${CRD_API_VERSION} --kind ${CRD_KIND}" $(MAKE) kubebuilder-run