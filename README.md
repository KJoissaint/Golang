# Employee API

API HTTP simple en Go pour gérer des employés.

## Structure du projet

```
employee-api/
├── go.mod
├── main.go
├── models/
│   ├── employee.go
│   └── manager.go
├── services/
│   └── employee_service.go
└── handlers/
    └── employee_handler.go
```

## Installation et lancement

```bash
# Aller dans le répertoire
cd employee-api

# Lancer le serveur
go run main.go
```

## Tester l'API

```bash
# Récupérer tous les employés
curl http://localhost:8080/employees
```

## Réponse attendue

```json
[
  {"ID":1,"Name":"Alice","Salary":5000},
  {"ID":2,"Name":"Bob","Salary":7000},
  {"ID":3,"Name":"Charlie","Salary":6000}
]
```
## POST /employees


```bash
curl -X POST http://localhost:8080/employees \
  -H "Content-Type: application/json" \
  -d '{"name": "David", "salary": 5500}'
```
## bODY attendue

```json
{
  "name": "David",
  "salary": 5500
}
```
## Réponse attendue

```json
{"ID":4,"Name":"David","Salary":5500}
```
## PUT /employees/raise
```bash
curl -X PUT http://localhost:8080/employees/raise \
  -H "Content-Type: application/json" \
  -d '{"id": 1, "percent": 10}'
```
## BODY attendue
```json
{
  "id": 1,
  "percent": 10
}
```
## Réponse attendue
```json
{"message": "Salaire augmenté avec succès"}
```
## Architecture

- **models/** : Définition des structures de données (Employee, Manager)
- **services/** : Logique métier avec interface et implémentation
- **handlers/** : Gestion des requêtes HTTP (pas de logique métier)
- **main.go** : Point d'entrée qui connecte tous les composants
