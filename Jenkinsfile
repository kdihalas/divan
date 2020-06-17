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

go get -u github.com/golang/dep/cmd/dep
ls -lha'''
      }
    }

  }
}