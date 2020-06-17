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

        stage('create workspace') {
          steps {
            ws(dir: '${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/') {
              sh '''export GOCACHE=$(pwd)/.cache
export GOENV=$(pwd)/.env
export GO111MODULE=on

go mod init github.com/kdihalas/divan
rm -rf Gopkg.*
go get -u -v'''
            }

          }
        }

      }
    }

  }
}