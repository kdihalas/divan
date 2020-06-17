pipeline {
  agent {
    docker {
      image 'golang:buster'
    }

  }
  stages {
    stage('deps') {
      steps {
        sh '''export GOPATH=$(pwd)
export GOCACHE=$GOPATH/.cache
export GOENV=$GOPATH/.env
export GO111MODULE=on

go mod init'''
      }
    }

    stage('mod') {
      steps {
        sh '''export GOPATH=$(pwd)
export GOCACHE=$GOPATH/.cache
export GOENV=$GOPATH/.env
export GO111MODULE=on

go get -u -v ./..'''
      }
    }

  }
}