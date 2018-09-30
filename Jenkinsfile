node {
    checkout scm
        
    stage('Docker Build') {
        docker.build('dukfaar/namespacebackend')
    }

    stage('Update Service') {
        sh 'docker service update --force namespacebackend_namespacebackend'
    }
}
