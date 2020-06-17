pipeline {
  agent {
    docker {
      image 'golang:buster'
    }

  }
  stages {
    stage('deps') {
      steps {
        sh '''export GOCACHE=.cache
export GOENV=.env
export GO111MODULE=on

go mod init github.com/kdihalas/divan
rm -rf Gopkg.*
go get -u -v'''
      }
    }

  }
}