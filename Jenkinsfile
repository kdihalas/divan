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

go mod init github.com/kdihalas/divan
rm -rf Gopkg.*
go get -u -v'''
      }
    }

  }
}