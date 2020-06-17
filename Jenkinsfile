pipeline {
  agent {
    docker {
      image 'golang:latest'
    }

  }
  stages {
    stage('install go') {
      steps {
        sh 'go get -u ./...'
      }
    }

  }
}