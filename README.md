# 🛍️ Electronic Shop Management System (Full-Stack)

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![React](https://img.shields.io/badge/React-18-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)
![Vite](https://img.shields.io/badge/Vite-B73BFE?style=for-the-badge&logo=vite&logoColor=FFD62E)
![JWT](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens)
![License](https://img.shields.io/badge/License-MIT-green.svg?style=for-the-badge)

Solution complète (Full-Stack) de gestion multi-boutiques pour la vente de matériel électronique. Ce projet intègre une **API robuste en Go** avec une architecture propre, et une **interface moderne en React**. Il garantit une isolation stricte des données entre les boutiques (Multi-tenant), une gestion granulaire des rôles et une intégration fluide avec WhatsApp pour les clients publics.

---

## 📑 Sommaire

1. [✨ Fonctionnalités Clés](#-fonctionnalités-clés)
2. [🛠️ Stack Technique](#️-stack-technique)
3. [🧱 Architecture du Projet](#-architecture-du-projet)
4. [🚀 Démarrage Rapide](#-démarrage-rapide)
5. [🧪 Comptes de Test](#-comptes-de-test)
6. [🌐 Aperçu de l'API](#-aperçu-de-lapi)
7. [🔐 Sécurité & Rôles](#-sécurité--rôles)
8. [🗺️ Roadmap & Améliorations](#️-roadmap--améliorations)

---

## ✨ Fonctionnalités Clés

- 🏢 **Architecture Multi-Tenant :** Isolation stricte des données. Un administrateur ne voit et ne gère que les produits et transactions de sa propre boutique.
- 👥 **Gestion des Rôles Avancée :**
    - `SuperAdmin` : Vue globale, dashboard financier, gestion de la marge de profit (prix d'achat vs prix de vente).
    - `Admin` : Gestion de sa boutique (CRUD produits, transactions, stock).
    - `Public/Guest` : Consultation des produits sans authentification.
- 💬 **Intégration WhatsApp :** Génération automatique de liens cliquables pour permettre aux clients de contacter directement le vendeur pour un produit spécifique.
- 📦 **Gestion de Stock Temps Réel :** Déduction automatique du stock lors des ventes et alertes visuelles de stock faible (< 5 articles).
- 📊 **Dashboard Analytique :** Calcul automatique du chiffre d'affaires, des dépenses et du bénéfice net (Réservé au SuperAdmin).

---

## 🛠️ Stack Technique

### Backend (API)

- **Langage :** Go (Golang) 1.21+
- **Sécurité :** JWT (JSON Web Tokens) & Bcrypt (Hashage des mots de passe)
- **Routage :** `net/http` (Standard Library)
- **Architecture :** Modèle en couches (Handlers, Services, Middlewares)

### Frontend (Client)

- **Framework :** React 18
- **Build Tool :** Vite ⚡
- **Routage :** React Router v6
- **Requêtes HTTP :** Axios (avec intercepteurs pour JWT)
- **Styling :** CSS3 natif (variables CSS, flexbox/grid, design responsive)

---

## 🧱 Architecture du Projet

Le projet est divisé en deux parties distinctes pour une séparation claire des responsabilités :

```text
📁 electronic-shop/
├── 📁 backend/                # API REST en Go
│   ├── config/              # Configuration globale (JWT, Ports)
│   ├── handlers/            # Contrôleurs HTTP (traitement des requêtes)
│   ├── middleware/          # Vérification JWT, Rôles, Multi-tenant
│   ├── models/              # Structures de données (Shop, User, Product...)
│   ├── services/            # Logique métier et persistance (In-memory)
│   └── main.go              # Point d'entrée de l'API (Port 8081)
│
└── 📁 frontend/               # Interface Utilisateur React
    ├── src/
    │   ├── components/      # Composants réutilisables (Navbar, PrivateRoute)
    │   ├── context/         # AuthContext (Gestion d'état global)
    │   ├── pages/           # Vues de l'application (Dashboard, Login...)
    │   ├── services/        # Appels API (Axios setup)
    │   └── App.jsx          # Routeur principal
    └── vite.config.js       # Configuration Vite (Port 3000)
```

---

## 🚀 Démarrage Rapide

### Prérequis

- Node.js (v16 ou supérieur)
- Go (v1.21 ou supérieur)

### 1️⃣ Lancer le Backend (Go)

```bash
# Se placer dans le dossier backend
cd backend

# Télécharger les dépendances
go mod download

# Démarrer le serveur
go run main.go
```

L'API démarrera sur `http://localhost:8081`

### 2️⃣ Lancer le Frontend (React)

Dans un nouveau terminal :

```bash
# Se placer dans le dossier frontend
cd frontend

# Installer les dépendances
npm install

# Démarrer le serveur de développement Vite
npm run dev
```

L'application sera accessible sur `http://localhost:3000`

---

## 🧪 Comptes de Test

La base de données en mémoire est pré-peuplée avec les comptes suivants pour faciliter les tests :

| Email | Mot de passe | Rôle | Shop Assigné | Privilèges |
|---|---|---|---|---|
| super@shop1.com | admin123 | SuperAdmin | Shop 1 | Accès total, Dashboard, Prix d'achat |
| admin@shop1.com | admin123 | Admin | Shop 1 | Gestion produits/ventes, Stock |
| (Aucun) | (Aucun) | Public | - | Navigation catalogue, Redirection WhatsApp |

---

## 🌐 Aperçu de l'API

L'API utilise des conventions RESTful claires. Note : Tu peux tester toutes les routes via la collection Postman incluse dans le projet.

### 🔓 Routes Publiques

- `POST /login` : Authentification et récupération du Token JWT.
- `POST /register` : Création de compte.
- `GET /public/:shopID/products` : Récupération du catalogue (Prix d'achat masqué, lien WhatsApp inclus).

### 🔒 Routes Privées (Nécessite `Authorization: Bearer <token>`)

- `GET /products` : Liste des produits de la boutique de l'utilisateur.
- `POST /products` : Ajouter un produit.
- `PUT /products/:id` : Modifier un produit.
- `DELETE /products/:id` : Supprimer un produit.
- `GET /transactions` : Historique des ventes/dépenses (Admin+).
- `POST /transactions` : Enregistrer une transaction (Vente, Dépense, Retrait).
- `GET /reports/dashboard` : Statistiques financières (SuperAdmin uniquement).

---

## 🔐 Sécurité & Rôles

- **Protection du Prix d'Achat :** Le champ `purchase_price` est strictement censuré par le backend. Seul un profil SuperAdmin recevra cette donnée dans la réponse JSON.
- **Isolation JWT (Multi-Tenant) :** Lors de chaque requête, le backend lit le `ShopID` directement depuis le token JWT signé, et non depuis le corps de la requête. Un admin du "Shop 1" ne peut physiquement pas requêter les produits du "Shop 2".
- **Mots de passe Hashés :** Utilisation de l'algorithme Bcrypt avec un coût (cost) standard.

---

## 🗺️ Roadmap & Améliorations

Ce projet est actuellement conçu avec une base de données en mémoire (In-Memory) à des fins éducatives et de démonstration rapide. Les prochaines étapes pour une mise en production :

- **Base de données persistante :** Remplacer le stockage en mémoire par PostgreSQL (utilisation de GORM ou sqlx).
- **Gestion des médias :** Upload réel des images produits (via AWS S3 ou stockage local) au lieu de simples URLs.
- **Pagination & Filtres :** Ajouter la pagination sur la route `GET /products` et des filtres par catégories.
- **Déploiement Docker :** Création d'un `docker-compose.yml` pour unifier le backend, le frontend et la future base de données.
- **Tests Unitaires :** Ajouter des tests Go (`testing` package) pour les services métier.

---

## 👨‍💻 Auteur & Licence

Projet Éducatif - Electronic Shop Management System  
Distribué sous la licence MIT. Libre d'utilisation et de modification.