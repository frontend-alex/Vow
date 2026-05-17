# Vow App

Monorepo for the Vow backend API and Flutter client.

## Apps

- `server/`: Go HTTP API, PostgreSQL access, migrations, SQL queries, and backend deployment assets.
- `mobile/`: Flutter client app, screens, routing, API client, local storage, and mobile app state.

## Common Commands

```bash
make server-run
make server-test
make mobile-run
make mobile-test
```

Run app-specific commands from the app directory when you need lower-level control.
