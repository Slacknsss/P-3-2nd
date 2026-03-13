# Go API - Todos
### > Ce projet a été réalisé dans le cadre d'un P3. Il va un peu plus loin que le sujet de base en ajoutant une vraie base de données PostgreSQL et une authentification par token JWT.
## Prérequis
- Go 1.21+
- PostgreSQL

## Installation

### 1. Cloner le projet
```bash
git clone git@github.com:Slacknsss/P-3-2nd.git
cd P-3-2nd
```

### 2. Installer les dépendances
```bash
go mod tidy
```

### 3. Créer le fichier .env
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=ton_user
DB_PASSWORD=
DB_NAME=go_api_db
JWT_SECRET=supersecret123
```

### 4. Créer la base de données
```bash
psql postgres -c "CREATE DATABASE go_api_db;"
```

### 5. Lancer l'API
```bash
go run main.go
```

## Routes

### Register
```bash
curl -X POST http://localhost:3000/register \
  -H "Content-Type: application/json" \
  -d '{"username":"simon","email":"simon@test.com","password":"123456"}'
```

### Login
```bash
curl -X POST http://localhost:3000/login \
  -H "Content-Type: application/json" \
  -d '{"email":"simon@test.com","password":"123456"}'
```

### Ajouter un todo
```bash
curl -X POST http://localhost:3000/addtodo \
  -H "Content-Type: application/json" \
  -d '{"title":"Ma tâche"}'
```

### Supprimer un todo
```bash
curl -X DELETE http://localhost:3000/suptodo/1
```
