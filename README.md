# Full Stack Todo App with Go and React

Made by following the @burakorkmez [video](https://www.youtube.com/watch?v=lNd7XlXwlho) for examples, with some modifications and additions.

A modern, full-stack todo application built with Go (Fiber) backend and React (Vite) frontend, featuring a clean UI with Chakra UI and MongoDB for data persistence.

![Tech Stack](https://skillicons.dev/icons?i=go,react,typescript,mongodb,vite,docker)

## 🚀 Features

- ⚡️ Lightning-fast Go Fiber backend
- 🎨 Modern React frontend with Chakra UI
- 📦 MongoDB for data persistence
- 🔄 Real-time updates with React Query
- 🌙 Dark/Light mode support
- 🔐 Environment-based configuration
- 🐳 Docker support (optional)

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://go.dev/doc/install) (1.24.0 or later)
- [Node.js](https://nodejs.org/) (v18 or later)
- [MongoDB](https://www.mongodb.com/try/download/community) (local or Atlas account)
- [Air](https://github.com/cosmtrek/air) for hot-reload during development:

```bash
  go install github.com/cosmtrek/air@latest
```

## 🛠️ Setup & Installation

### Backend Setup

1. Navigate to the backend directory:

   ```bash
   cd go-react-tutorial-with-db
   ```

2. Install Go dependencies:

   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the backend directory:

   ```env
   PORT=5000
   MONGODB_URI=your_mongodb_connection_string
   ENV=development
   ```

4. Start the backend server with hot-reload:

   ```bash
   air
   ```

   Or without hot-reload:

   ```bash
   go run main.go
   ```

### Frontend Setup

1. Navigate to the UI directory:

   ```bash
   cd ui
   ```

2. Install dependencies:

   ```bash
   npm install
   ```

3. Start the development server:

```bash
   npm run dev
```

The frontend will be available at `http://localhost:5173` and will proxy API requests to the backend at `http://localhost:5000`.

## 🌍 Environment Variables

### Backend (.env)

```env
PORT=5000                           # Server port
MONGODB_URI=your_connection_string  # MongoDB connection string
ENV=development                     # Environment (development/production)
```

### Frontend

The frontend uses Vite's built-in env variable system. Create a `.env.local` file in the `ui` directory if you need to customize the API URL:

```env
VITE_API_URL=http://localhost:5000  # Backend API URL
```

## 🚢 Deployment

This application can be easily deployed using [Railway](https://railway.app/):

1. Create a new project in Railway
2. Connect your GitHub repository
3. Add the following environment variables in Railway:
   - `PORT`
   - `MONGODB_URI`
   - `ENV=production`
4. Deploy both services:
   - Backend: Use the Dockerfile in the root directory
   - Frontend: Use the build command `npm run build` and start command `npm run preview`

Railway will automatically build and deploy your application, providing you with a public URL.

## 📁 Project Structure

```
.
├── go-react-tutorial-with-db/     # Backend directory
│   ├── main.go                    # Entry point
│   └── .env                       # Environment variables
│   └── go.mod                     # Go modules
│   └── go.sum                     # Go dependencies
│   └── air.toml                   # Air configuration
│
└── ui/                            # Frontend directory
    ├── src/
    │   ├── components/           # React components
    │   ├── chakra/              # Chakra UI configuration
    │   └── App.tsx              # Main application component
    └── package.json             # Frontend dependencies
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
