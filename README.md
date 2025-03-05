# Go App with Redis - DevOps Challenge

This repository contains the fixed and enhanced version of a Go web application that uses Redis for caching visitor counts. It includes:
- A revised Dockerfile and Go source code with necessary bug fixes and image size optimizations.
- Kubernetes YAML files to deploy the application (as a stateless workload in the `app` namespace) and Redis (as a stateful workload in the `db` namespace with persistent storage and proper networking).
- A Docker Compose file for local multi-container testing (Go app + Redis).

## Challenge Overview

**DevOps Challenge: Fix & Deploy Go App with Redis**

- **Part 1: Fix the Dockerfile and Go Application**
  - Troubleshoot and fix issues in the provided Dockerfile and Go code so that the app can build and run successfully.
  - The Go app uses Redis to cache the number of visitors.
  - **Bonus:** Optimize the Go app image size.

- **Part 2: Deploy to Kubernetes**
  - Create Kubernetes YAML files to deploy:
    - The Go application as a stateless workload (namespace: `app`).
    - Redis as a stateful workload (namespace: `db`) with persistent storage and a headless service for internal communication.
  - Use Kubernetes-native variable management.
  - Expose the Go app using a NodePort or LoadBalancer service.
  - Redis is exposed internally.

## Project Structure

```
.
├── Dockerfile
├── docker-compose.yml
├── k8s
│   ├── go-app-deployment.yaml
│   ├── redis-statefulset.yaml
│   └── services.yaml
└── app
    ├── main.go
    ├── go.mod
    └── go.sum
```

## Fixes & Enhancements

### Dockerfile
- **Base Image Optimization:** Switched to a multi-stage build using a lightweight base image (e.g., `alpine` or `scratch`) to reduce the final image size.
- **Dependency Management:** Ensured `go.mod` and `go.sum` files are copied first for layer caching.
- **Correct Build & CMD:** Fixed build commands and the executable path in the final stage from `CMD ["/bin/myapp"]` to `CMD ["./main"]`.
![Image size](Images/image_size.png)

### Go Application
- **Error Handling:** Added proper error handling when converting visit count values and setting data in Redis.
- **Default Values:** Provided default values using helper functions to ensure the app runs even if environment variables are missing.

### Kubernetes Deployment
- **Namespaces:** The application is deployed in two namespaces:
  - `app` – for the Go application.
  - `db` – for Redis.

![Image size](Images/ns.png)
  
- **Persistent Storage:** Redis StatefulSet uses `volumeClaimTemplates` to dynamically provision persistent storage.
- **Headless Service:** A headless service (`redis-headless`) is used to allow direct pod-to-pod communication for Redis.
- **InitContainer:** The Go app deployment includes an initContainer that waits for Redis connectivity before starting the main container.

## How to Run Locally

### Using Docker Compose

1. **Build and Run the Stack:**

   ```sh
   docker-compose up --build
   ```

2. **Access the Application:**
   - Open your browser and navigate to `http://localhost:8080` to see the visitor count.

## Kubernetes Deployment Instructions

1. **Apply the Kubernetes Manifests:**

   ```sh
   kubectl apply -f K8S/namespace.yaml
   ```
  ![Image size](Images/ns.png)
  
   ```sh
   kubectl apply -f K8S/redis-statefulset.yaml
   ```
  ![Image size](Images/ns.png)

2. **Check the Status of Deployments:**

   ```sh
   kubectl get pods --namespace app
   kubectl get pods --namespace db
   ```

3. **Access the Go App:**
   - Get the NodePort assigned:

   ```sh
   kubectl get services --namespace app
   ```

   - Open your browser and navigate to `http://<NodeIP>:<NodePort>`.

## Evidence of Work

- **Docker Image Size:** The optimized image size can be verified using:

  ```sh
  docker images
  ```

- **Kubernetes Logs:** Check logs for both Go app and Redis:

  ```sh
  kubectl logs <go-app-pod-name> --namespace app
  kubectl logs <redis-pod-name> --namespace db
  ```

## Conclusion

This project demonstrates the ability to troubleshoot, optimize, and deploy a Go web application using Redis in a Kubernetes environment. The enhancements made ensure better performance, error handling, and efficient resource management.
