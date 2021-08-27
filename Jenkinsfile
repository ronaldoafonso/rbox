pipeline {
    agent any

    stages {
        stage('build dev') {
            steps {
                sh 'docker network create --driver bridge net_rbox'
                sh 'docker image build --target devel --tag rbox-dev:0.0.1 .'
            }
        }
        stage('test dev') {
            steps {
                sh 'docker container run --rm --name rbox_dev_unit_test        --network net_rbox --mount source=z3nbox,target=/home/rbox/.ssh rbox-dev:0.0.1 bash -c "cd rbox; go test -v ./..."'
                sh 'docker container run --rm --name rbox_dev_integration_test --network net_rbox --mount source=z3nbox,target=/home/rbox/.ssh rbox-dev:0.0.1 bash -c "bats tests.bats"'
            }
        }
        stage('build prod') {
            steps {
                sh 'docker image build --tag rbox:0.0.1 .'
            }
        }
    }

    post {
        always {
            sh 'docker network rm net_rbox'
        }
    }
}
