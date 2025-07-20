import { useState, useEffect } from 'react'
import './App.css'

// Example component demonstrating the monorepo structure
function App() {
  const [health, setHealth] = useState<string>('Checking...')

  useEffect(() => {
    // Simple health check to test backend connectivity
    fetch('http://localhost:8080/api/v1/health')
      .then(res => res.json())
      .then(data => setHealth(data.status))
      .catch(() => setHealth('Backend not available'))
  }, [])

  return (
    <div className="App">
      <header className="App-header">
        <h1>🚀 My Monorepo</h1>
        <p>Go + MongoDB + React + pnpm + Turborepo</p>

        <div className="status-card">
          <h3>Backend Status</h3>
          <p style={{
            color: health === 'healthy' ? 'green' :
                   health === 'Checking...' ? 'orange' : 'red'
          }}>
            {health}
          </p>
        </div>

        <div className="features">
          <h3>✨ Features</h3>
          <ul>
            <li>🔥 Type-safe API with OpenAPI</li>
            <li>⚡ Hot reload development</li>
            <li>📦 Shared packages</li>
            <li>🏗️ Turborepo cache</li>
            <li>🗄️ MongoDB ready</li>
          </ul>
        </div>

        <div className="next-steps">
          <h3>🎯 Next Steps</h3>
          <ol>
            <li>Start MongoDB server</li>
            <li>Run <code>pnpm dev</code> in root</li>
            <li>Edit <code>packages/oas/openapi.yaml</code></li>
            <li>Run <code>pnpm generate:api</code></li>
            <li>Use generated RTK Query hooks!</li>
          </ol>
        </div>
      </header>
    </div>
  )
}

export default App
