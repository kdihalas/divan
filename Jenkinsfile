pipeline {
  agent {
    node {
      def root = tool name: 'default', type: 'go'
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
