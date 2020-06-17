pipeline {
  agent any
  stages {
    stage('install go') {
      steps {
        sh '''WORKROOT=$(pwd)
mkdir ${WORKROOT}/tmp
cd ${WORKROOT}/tmp

# unzip go environment
go_env="go1.13.12.linux-amd64.tar.gz"
wget -c https://dl.google.com/go/go1.13.12.linux-amd64.tar.gz
tar -zxf $go_env
if [ $? -ne 0 ];
then
    echo "fail in extract go"
    exit 1
fi
echo "OK for extract go"
rm -rf $go_env

# prepare PATH, GOROOT and GOPATH
export PATH=$(pwd)/go/bin:$PATH
export GOROOT=$(pwd)/go
export GOPATH=$(pwd)'''
      }
    }

    stage('build') {
      steps {
        sh 'go get ./...'
      }
    }

  }
}