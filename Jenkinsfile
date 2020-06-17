pipeline {
  agent {
    docker {
      image 'golang:buster'
    }

  }
  stages {
    stage('deps') {
      parallel {
        stage('deps') {
          steps {
            sh '''export GOCACHE=$(pwd)/.cache
export GOENV=$(pwd)/.env
export GO111MODULE=on

go mod init github.com/kdihalas/divan
rm -rf Gopkg.*
go get -u -v'''
          }
        }

        stage('env') {
          steps {
            sh 'env'
          }
        }

      }
    }

  }
}