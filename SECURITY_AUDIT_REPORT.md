# Security Audit Report for near-rt-ric

This report summarizes the findings of a security audit performed on the `near-rt-ric` project. The audit focused on the backend (Go), frontend (Angular), and deployment configurations.

**NOTE:** Due to persistent timeout issues in the execution environment, it was not possible to use any security scanning tools (`gosec`, `govulncheck`, `eslint`, `hadolint`, `kube-score`). The findings in this report are based on a manual review of the source code and configuration files.

## High Priority

### 1. Inability to Perform a Comprehensive Security Scan

*   **Description:** The execution environment provided for this audit was not functional. All attempts to install and run security scanning tools resulted in timeouts. This means that the audit is incomplete and there may be undiscovered vulnerabilities in the codebase.
*   **Remediation:** The execution environment must be fixed to allow for the use of standard security scanning tools. Once the environment is functional, a full security audit should be performed.

## Medium Priority

### 1. Dockerfile Best Practices

*   **Description:** The Dockerfiles in the project have some areas for improvement.
    *   `dashboard-master/dashboard-master/aio/Dockerfile`: The `ADD . /` command copies the entire build context into the image. This could include sensitive files if they are not properly excluded with a `.dockerignore` file.
    *   `xAPP_dashboard-master/Dockerfile`:
        *   Uses `node:latest` as the build image. It's better to use a specific version of Node.js to ensure reproducible builds and to avoid unexpected changes.
        *   The final image is `nginx:alpine`. While Alpine is a small image, it's not as minimal as `scratch`. It would be better to use a more minimal image if possible.
        *   The `COPY . .` command copies the entire source code into the build stage. This could include sensitive files if they are not properly excluded with a `.dockerignore` file.
        *   The final stage runs as root by default. It would be better to create a non-root user and run Nginx as that user.
*   **Remediation:**
    *   Review and update the `.dockerignore` files to exclude any sensitive files.
    *   Use specific versions for all base images (e.g., `node:16`).
    *   Create a non-root user in the `xAPP_dashboard-master/Dockerfile` and run Nginx as that user.

### 2. Kubernetes Deployment Configuration

*   **Description:** The Kubernetes deployment for the dashboard has a toleration that allows it to be deployed on the master node. This could be a security risk, as the master node is a critical component of the cluster.
*   **Remediation:** Remove the toleration for `node-role.kubernetes.io/master` from the dashboard deployment. The dashboard should be deployed on a worker node.

## Low Priority

### 1. Hardcoded Secret Names

*   **Description:** The Go backend code contains hardcoded secret names (`kubernetes-dashboard-csrf` and `kubernetes-dashboard-certs`). This is not a direct vulnerability, but it makes the application less flexible and could make it more difficult to deploy in some environments.
*   **Remediation:** The secret names should be configurable, for example, through command-line flags or environment variables.

## Conclusion

This security audit was severely limited by the non-functional execution environment. The findings in this report are based on a manual review of the code and configuration files, and should not be considered a comprehensive security assessment. It is strongly recommended that a full security audit be performed as soon as the environment issues are resolved.
