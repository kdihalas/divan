pipeline {
  agent {
    node {
      def root = tool name: 'default', type: 'go'
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
