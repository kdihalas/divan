pipeline {
  agent {
    docker {
      image 'golang'
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