pipeline{
    environment{
        DOCKER_CREDENTIALS = credentials("docker-credentials")
        registry = "csye712504/webapp-hello-world"
    }
    agent any
    stages{
        stage("Image Building and pushing"){
            steps{
                script{
                    sh 'echo ${DOCKER_CREDENTIALS_PSW} | docker login -u ${DOCKER_CREDENTIALS_USR} --password-stdin'
                    sh 'docker buildx create --use --name imagebuilder'
                    sh "docker buildx build --platform linux/amd64,linux/arm64 -t ${registry} --push ."
                    sh 'docker buildx rm imagebuilder'

                }
            }
        }
    }

    post {
        failure {
            echo "Build failed. Cleaning up builder instance."
            sh 'docker buildx rm imagebuilder || true'
        }
        always {
            echo "Pipeline execution completed."
        }
    }
}
