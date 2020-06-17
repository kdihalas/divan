pipeline {
  node {
      // Install the desired Go version
      def root = tool name: 'default', type: 'go'

      // Export environment variables pointing to the directory where Go was installed
      withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
          sh 'go version'
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
