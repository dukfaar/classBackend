node {
    checkout scm
        
    stage('Docker Build') {
        docker.build('dukfaar/classbackend')
    }

    stage('Update Service') {
        sh 'docker service update --force classbackend_classbackend'
    }
}
