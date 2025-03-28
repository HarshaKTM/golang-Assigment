@echo off
echo Starting Book API (Backend and Frontend)...

:: Start backend in a new window
start cmd /k "go run main.go -port 5001 -seed"

:: Wait a bit for backend to initialize
timeout /t 5

:: Start frontend in a new window
start cmd /k "cd frontend && npm start"

echo Both services should be starting now.
echo Backend API: http://localhost:5001
echo Frontend: http://localhost:3000 