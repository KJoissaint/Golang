# 🛍️ Electronic Shop Management API

Backend en Go pour un système de gestion multi-boutiques d'électronique avec isolation complète des shops, gestion des rôles, et page publique pour les clients.

## 🎯 Fonctionnalités

- ✅ **Multi-tenant** : Isolation complète entre les boutiques
- ✅ **Authentication JWT** : Sécurité avec tokens
- ✅ **Gestion des rôles** : SuperAdmin, Admin
- ✅ **API publique** : Accès sans authentification pour les clients
- ✅ **Redirection WhatsApp** : Liens dynamiques pour contact direct
- ✅ **Dashboard** : Statistiques et profits (SuperAdmin)
- ✅ **Gestion du stock** : Suivi en temps réel

## 📁 Structure du Projet

```
shop-api/
├── main.go                 # Point d'entrée
├── go.mod                  # Dépendances
├── config/
│   └── config.go          # Configuration JWT et serveur
├── models/
│   ├── shop.go            # Modèle Shop
│   ├── user.go            # Modèle User
│   ├── product.go         # Modèle Product
│   ├── transaction.go     # Modèle Transaction
│   └── whatsapp.go        # Génération liens WhatsApp
├── services/
│   ├── shop_service.go
│   ├── user_service.go
│   ├── product_service.go
│   └── transaction_service.go
├── handlers/
│   ├── auth_handler.go
│   ├── product_handler.go
│   ├── transaction_handler.go
│   └── shop_handler.go
├── middleware/
│   └── auth.go            # JWT et validation rôles
└── utils/
    ├── jwt.go             # Génération/validation JWT
    └── password.go        # Hashage bcrypt
```

## 🧱 Modèles de Données

### 1️⃣ Shop
```go
{
  "id": 1,
  "name": "TechStore Casablanca",
  "active": true,
  "whatsapp_number": "212600000001",
  "created_at": "2026-02-12T10:00:00Z"
}
```

### 2️⃣ User
```go
{
  "id": 1,
  "name": "Super Admin",
  "email": "super@shop1.com",
  "role": "SuperAdmin",  // SuperAdmin ou Admin
  "shop_id": 1,
  "created_at": "2026-02-12T10:00:00Z"
}
```

### 3️⃣ Product
```go
{
  "id": 1,
  "name": "iPhone 14 Pro",
  "description": "Latest iPhone",
  "category": "Smartphones",
  "purchase_price": 8000,  // Visible SuperAdmin uniquement
  "selling_price": 10000,
  "stock": 15,
  "image_url": "https://example.com/iphone14.jpg",
  "shop_id": 1,
  "created_at": "2026-02-12T10:00:00Z"
}
```

### 4️⃣ Transaction
```go
{
  "id": 1,
  "type": "Sale",  // Sale, Expense, Withdrawal
  "product_id": 1,
  "quantity": 2,
  "amount": 20000,
  "shop_id": 1,
  "created_at": "2026-02-12T10:00:00Z"
}
```

## 🚀 Installation et Lancement

### 1. Installer les dépendances

```bash
cd shop-api
go mod download
```

### 2. Lancer le serveur

```bash
go run main.go
```

Le serveur démarre sur `http://localhost:8080`

## 🌐 API Routes

### 🔓 Routes Publiques

#### POST /register
Créer un nouveau compte utilisateur

```bash
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123",
    "role": "Admin",
    "shop_id": 1
  }'
```

**Réponse:**
```json
{
  "id": 3,
  "name": "John Doe",
  "email": "john@example.com",
  "role": "Admin",
  "shop_id": 1,
  "created_at": "2026-02-12T10:00:00Z"
}
```

#### POST /login
Connexion et récupération du token JWT

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "super@shop1.com",
    "password": "admin123"
  }'
```

**Réponse:**
```json
{
  "user": {
    "id": 1,
    "name": "Super Admin 1",
    "email": "super@shop1.com",
    "role": "SuperAdmin",
    "shop_id": 1,
    "created_at": "2026-02-12T10:00:00Z"
  },
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### GET /public/:shopID/products
Liste des produits pour les clients (sans authentification)

```bash
curl http://localhost:8080/public/1/products
```

**Réponse:**
```json
[
  {
    "id": 1,
    "name": "iPhone 14 Pro",
    "description": "Latest iPhone with advanced camera system",
    "category": "Smartphones",
    "selling_price": 10000,
    "stock": 15,
    "image_url": "https://example.com/iphone14.jpg",
    "whatsapp_link": "https://wa.me/212600000001?text=Bonjour%20je%20veux%20plus%20d%27information%20sur%20iPhone%2014%20Pro"
  }
]
```

⚠️ **Note:** `purchase_price` n'est jamais exposé dans l'API publique

### 🔒 Routes Privées (Authentification requise)

#### GET /products
Liste des produits (filtrés par shop de l'utilisateur)

```bash
curl http://localhost:8080/products \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

#### POST /products
Créer un nouveau produit

```bash
curl -X POST http://localhost:8080/products \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPad Pro",
    "description": "Professional tablet",
    "category": "Tablets",
    "purchase_price": 6000,
    "selling_price": 7500,
    "stock": 10,
    "image_url": "https://example.com/ipad.jpg"
  }'
```

#### PUT /products/:id
Mettre à jour un produit

```bash
curl -X PUT http://localhost:8080/products/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "iPhone 14 Pro Max",
    "description": "Updated description",
    "category": "Smartphones",
    "purchase_price": 8500,
    "selling_price": 11000,
    "stock": 12,
    "image_url": "https://example.com/iphone14.jpg"
  }'
```

#### DELETE /products/:id
Supprimer un produit

```bash
curl -X DELETE http://localhost:8080/products/1 \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

### 👥 Routes Admin

#### GET /transactions
Liste des transactions

```bash
curl http://localhost:8080/transactions \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

#### POST /transactions
Créer une transaction

```bash
curl -X POST http://localhost:8080/transactions \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "Sale",
    "product_id": 1,
    "quantity": 2,
    "amount": 20000
  }'
```

Types de transactions: `Sale`, `Expense`, `Withdrawal`

### 👑 Routes SuperAdmin

#### GET /reports/dashboard
Dashboard avec statistiques et profits

```bash
curl http://localhost:8080/reports/dashboard \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

**Réponse:**
```json
{
  "total_sales": 20000,
  "total_expenses": 5000,
  "net_profit": 11000,
  "low_stock_count": 1,
  "total_revenue": 20000,
  "total_cost": 4000,
  "products_sold": 2
}
```

#### PUT /shops/whatsapp
Modifier le numéro WhatsApp du shop

```bash
curl -X PUT http://localhost:8080/shops/whatsapp \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "whatsapp_number": "212611223344"
  }'
```

#### GET /shops
Liste de tous les shops

```bash
curl http://localhost:8080/shops \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 🔐 Gestion des Rôles

### 👑 SuperAdmin
**Peut:**
- ✅ CRUD produits
- ✅ Voir `purchase_price`
- ✅ Voir profits et dashboard
- ✅ Gérer utilisateurs
- ✅ Modifier WhatsApp du shop

### 🧑‍💼 Admin
**Peut:**
- ✅ CRUD produits
- ✅ CRUD transactions
- ✅ Voir `selling_price`
- ✅ Voir stock

**Ne peut PAS:**
- ❌ Voir `purchase_price`
- ❌ Voir profit
- ❌ Modifier WhatsApp

### 👥 Guest (Client)
- ✅ Voir produits disponibles
- ✅ Voir stock
- ✅ Cliquer pour demander info (WhatsApp)
- ❌ Aucun compte requis

## 📱 Redirection WhatsApp

Le backend génère automatiquement des liens WhatsApp formatés:

**Format:**
```
https://wa.me/<WhatsAppNumber>?text=Bonjour%20je%20veux%20plus%20d'information%20sur%20<NomProduit>
```

**Exemple:**
```
https://wa.me/212600000001?text=Bonjour%20je%20veux%20plus%20d%27information%20sur%20iPhone%2014%20Pro
```

## 🧠 Logique Métier

### Multi-tenant Strict
- Chaque utilisateur ne voit que les données de son shop
- ShopID extrait automatiquement du JWT
- Validation stricte des permissions

### Gestion du Stock
- Produits avec `stock = 0` restent visibles avec mention "Out of stock"
- Déduction automatique lors des ventes
- Alertes pour stock faible (<5)

### Sécurité
- ✅ Passwords hashés avec bcrypt
- ✅ JWT avec expiration (7 jours)
- ✅ `purchase_price` jamais exposé publiquement
- ✅ Validation des rôles via middleware
- ✅ Isolation multi-tenant

## 🧪 Comptes de Test

| Email | Password | Rôle | Shop |
|-------|----------|------|------|
| super@shop1.com | admin123 | SuperAdmin | 1 |
| admin@shop1.com | admin123 | Admin | 1 |

## 📊 Testing avec cURL

### Workflow complet

1. **Login**
```bash
TOKEN=$(curl -s -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"email":"super@shop1.com","password":"admin123"}' \
  | jq -r '.token')
```

2. **Voir produits privés**
```bash
curl http://localhost:8080/products \
  -H "Authorization: Bearer $TOKEN"
```

3. **Voir produits publics**
```bash
curl http://localhost:8080/public/1/products
```

4. **Dashboard**
```bash
curl http://localhost:8080/reports/dashboard \
  -H "Authorization: Bearer $TOKEN"
```

## 🛠️ Technologies Utilisées

- **Go 1.21+**
- **JWT** (golang-jwt/jwt/v5)
- **Bcrypt** (golang.org/x/crypto)
- **HTTP Standard Library** (net/http)

## 📝 Notes de Développement

### Architecture
- **Models**: Structures de données pures
- **Services**: Logique métier et persistance (in-memory)
- **Handlers**: Gestion HTTP et validation
- **Middleware**: Authentication et authorization
- **Utils**: Fonctions utilitaires (JWT, password)

### Persistance
Actuellement en mémoire (in-memory). Pour production:
- Remplacer par PostgreSQL/MySQL
- Ajouter GORM ou sqlx
- Implémenter connection pool

### Améliorations Futures
- [ ] Base de données réelle
- [ ] Upload d'images
- [ ] Pagination
- [ ] Filtres et recherche
- [ ] Logs structurés
- [ ] Tests unitaires
- [ ] Docker
- [ ] CI/CD

## 📄 License

MIT

## 👨‍💻 Auteur

Projet éducatif - Electronic Shop Management System
