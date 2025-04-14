pipeline {
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
                    sh "docker run -d -p 8082:8082 --name myapp myapp:latest"
                    
                    sh "sleep 30"

                    // // Test the application by curling the server
                    // sh "curl -f http://localhost:8082 || exit 1"
                    
                    sh "docker stop myapp"
                }
            }
        }
    }
}
