## Project Structure Backend

- `cmd/server/main.go`: The entry point of the API server.
- `domain/`: Contains core domain logic.
- `repository/`: Data access layer.
- `server/`: Server configuration and setup.
- `routes/`: Routing configurations.
- `users/`: Business logic for user operations.

## Project Structure Frontend

- `src/router/` Contains routing logic
- `src/utils/` Contains Rest APIs and Backend URL
- `src/views/` Contains Views
- `src/App.vue` Entry Point
- `src/main.ts` Setup

### Changing Url for backend in frontend

To run the frontend that interacts with this API, update the backend URL in `/src/utils/network.ts`:

```typescript
const BACKEND_URL: string = 'http://localhost:8081'
```

## Running the application in docker

```bash
docker-compose build
docker-compose up
```
