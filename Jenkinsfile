pipeline {
  agent any
  stages {
    stage('deps') {
      steps {
        sh '''wget https://dl.google.com/go/go1.13.12.linux-amd64.tar.gz
tar -zxvf go1.13.12.linux-amd64.tar.gz'''
      }
    }

  }
}