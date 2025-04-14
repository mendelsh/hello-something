
pipline{
    agent any

    stages {
        stage('Build') {
            steps {
                script {
                    sh "docker build -t myapp:latest ."
                }
            }
        }
        stage('Test') {
            steps {
                script {
                    sh "docker run --rm -p 8082:8082 myapp:latest test"
                    sh "sleep 30"
                    sh "docker stop myapp"
                }
            }
        }
    }
}