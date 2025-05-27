# Real-Time Quiz Backend

## Features
- Join quiz session by quiz ID
- Multiple users can join simultaneously
- Submit answers, real-time score update
- Real-time leaderboard via WebSocket
- Scalable, stateless API

## Tech Stack
- Go
- Redis (Sorted Set, Pub/Sub)
- Gin (REST API)
- Gorilla WebSocket

## Run

1. Copy `.env` and update config if needed
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Start the server:
   ```sh
   go run cmd/main.go
   ```
   Or use the VS Code task: **Run Real-Time Quiz Backend**

## API Endpoints

- `POST /api/join` — Join a quiz session
- `POST /api/answer` — Submit answer (body: quiz_id, user_id, question_id, answer)
- `GET /api/leaderboard?quiz_id=...` — Get leaderboard
- `GET /ws/leaderboard?quiz_id=...` — WebSocket for real-time leaderboard

## Test API directly in VS Code

### 1. Install extensions:
- **REST Client** (`humao.rest-client`) to test API via `.http` file
- **WebSocket Client** (`qinjiahao.vscode-websocket-client`) to test realtime leaderboard

### 2. Test API with `test.http` file
- Open `test.http` in your workspace
- Click "Send Request" above each block to send join quiz, submit answer, and get leaderboard requests for multiple users
- Example:
  ```http
  ### Join quiz user1
  POST http://localhost:8080/api/join
  Content-Type: application/json
  {
    "quiz_id": "quiz1",
    "user_id": "user1"
  }
  ### Submit answer user1 - Q1
  POST http://localhost:8080/api/answer
  Content-Type: application/json
  {
    "quiz_id": "quiz1",
    "user_id": "user1",
    "question_id": "1",
    "answer": "Paris"
  }
  ### Get leaderboard
  GET http://localhost:8080/api/leaderboard?quiz_id=quiz1
  ```

### 3. Test realtime leaderboard with WebSocket
- Open Command Palette (Cmd+Shift+P), select `WebSocket: Connect`
- Enter URL: `ws://localhost:8080/ws/leaderboard?quiz_id=quiz1`
- When a user submits a correct answer, the leaderboard will be pushed in real time, for example:
  ```json
  {
    "leaderboard": [
      { "user": "user1", "score": 30, "rank": 1 },
      { "user": "user2", "score": 20, "rank": 2 }
    ]
  }
  ```

## Notes
- The sample question set is available in the code (`internal/common/demo_questions.go`)
- You can extend with an API to get questions, add validation, or store results in another DB if needed.
