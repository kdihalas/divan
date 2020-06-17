pipeline {
  agent any
  stages {
    stage('install go') {
      steps {
        sh '''wget https://dl.google.com/go/go1.13.12.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.13.12.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin'''
      }
    }

  }
}