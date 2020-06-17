pipeline {
  agent {
    docker {
      image 'golang:buster'
    }

  }
  stages {
    stage('build') {
      steps {
        sh 'go get ./...'
      }
    }

  }
}