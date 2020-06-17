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

go get -u github.com/golang/dep/cmd/dep'''
      }
    }

  }
}