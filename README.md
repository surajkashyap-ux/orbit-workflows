Orbit Workflows

This project is a self-contained serverless demo that runs fully on Vercel. It includes a Next.js UI and Go API functions using only in-memory storage.

Local run
cd ui
npm install
npm run dev
npm install -g vercel
vercel dev

Health: http://localhost:3000/api/health
Metrics: http://localhost:3000/api/metrics

Production path
Move state to Postgres and Redis, deploy a worker for background execution, and keep the control plane on Vercel or Kubernetes. Add secrets through your platform and apply migrations under api/migrations.

© 2025 Orbit Workflows
