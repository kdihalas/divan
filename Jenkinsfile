pipeline {
  agent {
    docker {
      image 'golang:buster'
    }

  }
  stages {
    stage('create workspace') {
      steps {
        ws(dir: '${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/') {
          sh '''export GOCACHE=$(pwd)/.cache
export GOENV=$(pwd)/.env
export GO111MODULE=on

go mod init github.com/kdihalas/divan
go get -u -v'''
        }

      }
    }

  }
}