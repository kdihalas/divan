pipeline {
  agent {
    docker {
      image 'golang:buster'
    }

  }
  stages {
    stage('Build') {
      steps {
        sh '''PROJECT_PATH=github.com/kdihalas/divan

cd ${GOPATH}/src
mkdir -p ${GOPATH}/src/${PROJECT_PATH}
cp -r ${WORKSPACE}/* ${GOPATH}/src/${PROJECT_PATH}
go build'''
      }
    }

  }
}