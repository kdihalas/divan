pipeline {
  agent {
    docker {
      image 'golang:buster'
    }

  }
  stages {
    stage('Build') {
      steps {
        sh '''GOCACHE=$(HOME)/.cache/go-build
GOENV=$(HOME)/.config/go/env
PROJECT_PATH=github.com/kdihalas/divan

mkdir -p ${GOPATH}/src/${PROJECT_PATH}

cp -r ${WORKSPACE}/* ${GOPATH}/src/${PROJECT_PATH}

cd ${GOPATH}/src/${PROJECT_PATH}

go get -u github.com/golang/dep/cmd/dep

dep install
go build cmd/main.go'''
      }
    }

  }
}