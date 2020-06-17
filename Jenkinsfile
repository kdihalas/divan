pipeline {
  agent {
    docker {
      image 'golang:latest'
    }

  }
  stages {
    stage('deps') {
      steps {
        sh 'go get -u ./..'
      }
    }

  }
}