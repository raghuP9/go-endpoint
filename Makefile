BASEDIR := $(shell pwd)
BUILDDIR := ${BASEDIR}/build
ORG := rpsraghu
REPO := go-endpoint
REPO_VERSION := 0.1.0
SRCPATH := github.com/${ORG}/${REPO}

# Docker commands
DRUN := docker run
DBUILD := docker build
DREM := docker rm
DREM_IMAGE := docker rmi
RM = rm -rf

#.DEFAULT_GOAL := build

build: setup compile image

all: setup compile image test

setup:
  # Setup containerized build environment
	${DBUILD} -t ${ORG}/tools-compile ${BUILDDIR}/compile

compile:
  # Build application (go-endpoint)
	${DRUN} --rm -i -e ORG=${ORG} -e REPO=${REPO} --name ${REPO}-build -v ${BASEDIR}:/go/src/${SRCPATH} -v ${BUILDDIR}/output:/output ${ORG}/tools-compile

image:
  # Create application container
	${DBUILD} -t ${ORG}/${REPO}:${REPO_VERSION} -f ${BUILDDIR}/runtime/Dockerfile ${BASEDIR}

test:
  # Test application using go test
	${DRUN} --rm -i -e ORG=${ORG} -e REPO=${REPO} --name ${REPO}-build -v ${BASEDIR}:/go/src/${SRCPATH} -v ${BUILDDIR}/output:/output ${ORG}/tools-compile go test -v ${SRCPATH}/... ${SRCPATH}
clean:
  # Remove all docker images created during build process
	-${DREM_IMAGE} -f ${ORG}/${REPO}:${REPO_VERSION}
	-${DREM_IMAGE} -f ${ORG}/tools-compile
	# Remove generated binaries
	-${RM} ${BUILDDIR}/output
