import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";

function App() {
  return (
    <main className="min-h-screen bg-gray-950 text-white flex items-center justify-center">
      <div className="text-center">
        <h1 className="text-5xl font-bold">
          Perpetual Trading System
        </h1>

        <p className="mt-4 text-gray-400">
          React + Go + PostgreSQL
        </p>

        <button className="mt-8 rounded-lg bg-blue-600 px-6 py-3 font-medium hover:bg-blue-500">
          Get Started
        </button>
      </div>
    </main>
  );
}

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <App />
  </StrictMode>
);
