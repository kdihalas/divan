pipeline {
  agent {
    node {
      label 'golang'
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