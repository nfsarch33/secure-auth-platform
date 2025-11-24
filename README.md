### 🧪 Running E2E Tests Locally

To run the Playwright E2E tests locally using Docker:

1.  Ensure you have `.env` file (copy from `.env.example`):
    ```bash
    cp .env.example .env
    ```

2.  Run the E2E tests:
    ```bash
    make test-e2e-docker
    ```

    This command will:
    *   Start the full stack (Postgres, Backend, Frontend) using Docker Compose.
    *   Run Playwright tests in a separate container connected to the stack.
    *   Clean up resources afterwards.

### ☸️ Kubernetes Deployment

(Instructions for K8s deployment...)
