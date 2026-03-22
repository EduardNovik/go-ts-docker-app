import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  server: {
    host: true, // ← expose to all interfaces, not just localhost (required for Docker)
    port: 3000,
    proxy: {
      "/api": {
        target: "http://localhost:8080",
        rewrite: (path) => path.replace(/^\/api/, ""), // removes the /api prefix
      },
    },
  },
  plugins: [react()],
});
