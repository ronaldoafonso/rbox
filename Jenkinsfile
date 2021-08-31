pipeline {
    agent any

    environment {
        RBOX_VERSION = '0.0.4'
    }

    stages {
        stage('build dev') {
            steps {
                sh "docker image build --target devel --tag rbox-dev:${env.RBOX_VERSION} ."
            }
        }
        stage('unit test dev') {
            steps {
                sh "docker container run --rm --name rbox_dev_unit_test \
                    --mount source=z3nbox,target=/home/rbox/.ssh \
                    rbox-dev:${env.RBOX_VERSION} bash -c 'cd rbox; go test -v ./...'"
            }
        }
        stage('build prod') {
            steps {
                sh "docker image build --tag z3nz3n/rbox:${env.RBOX_VERSION} ."
            }
        }
        stage('integration test prod') {
            steps {
                sh "minikube start"
                sh "minikube kubectl -- apply -f k8s/rbox-namespace.yml"
                sh "minikube kubectl -- apply -f /etc/rbox/rbox-configmap.yml"
                sh "minikube kubectl -- apply -f k8s/rbox-deployment.yml"
                sh "minikube kubectl -- apply -f k8s/rbox-service.yml"
                sh "bash -c 'bats tests.bats'"
            }
        }
        stage('push prod') {
            steps {
                sh "docker image push z3nz3n/rbox:${env.RBOX_VERSION}"
            }
        }
    }

    post {
        always {
            sh "minikube kubectl -- delete -f k8s/rbox-service.yml"
            sh "minikube kubectl -- delete -f k8s/rbox-deployment.yml"
            sh "minikube kubectl -- delete -f /etc/rbox/rbox-configmap.yml"
            sh "minikube kubectl -- delete -f k8s/rbox-namespace.yml"
            sh "minikube stop"
        }
    }
}
