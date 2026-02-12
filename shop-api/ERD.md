# 📊 Entity Relationship Diagram (ERD)

```
┌─────────────────────────────────────────────────────────────────────┐
│                         SHOP                                         │
├─────────────────────────────────────────────────────────────────────┤
│ PK  id               : int                                           │
│     name             : string                                        │
│     active           : bool                                          │
│     whatsapp_number  : string                                        │
│     created_at       : timestamp                                     │
└──────────────────┬──────────────────────────────────────────────────┘
                   │
                   │ 1 Shop has many Users
                   │
        ┌──────────┴──────────┐
        │                     │
        ▼                     ▼
┌──────────────────┐  ┌──────────────────┐
│      USER        │  │     PRODUCT      │
├──────────────────┤  ├──────────────────┤
│ PK  id           │  │ PK  id           │
│     name         │  │     name         │
│     email        │  │     description  │
│     password     │  │     category     │
│     role         │  │     purchase_pr. │
│ FK  shop_id      │  │     selling_pr.  │
│     created_at   │  │     stock        │
└──────────────────┘  │     image_url    │
                      │ FK  shop_id      │
                      │     created_at   │
                      └──────┬───────────┘
                             │
                             │ 1 Product has many Transactions
                             │
                             ▼
                      ┌─────────────────┐
                      │  TRANSACTION    │
                      ├─────────────────┤
                      │ PK  id          │
                      │     type        │
                      │ FK  product_id  │
                      │     quantity    │
                      │     amount      │
                      │ FK  shop_id     │
                      │     created_at  │
                      └─────────────────┘
```

## Relationships:

1. **Shop ─< User** (One-to-Many)
   - One Shop has many Users
   - Each User belongs to one Shop

2. **Shop ─< Product** (One-to-Many)
   - One Shop has many Products
   - Each Product belongs to one Shop

3. **Shop ─< Transaction** (One-to-Many)
   - One Shop has many Transactions
   - Each Transaction belongs to one Shop

4. **Product ─< Transaction** (One-to-Many, Optional)
   - One Product can have many Transactions
   - Transactions can exist without Products (expenses, withdrawals)

## Data Types:

### Shop
- `id`: int (Primary Key)
- `name`: string
- `active`: boolean
- `whatsapp_number`: string (e.g., "212600000001")
- `created_at`: timestamp

### User
- `id`: int (Primary Key)
- `name`: string
- `email`: string (unique)
- `password`: string (bcrypt hashed)
- `role`: enum ("SuperAdmin", "Admin")
- `shop_id`: int (Foreign Key → Shop)
- `created_at`: timestamp

### Product
- `id`: int (Primary Key)
- `name`: string
- `description`: string
- `category`: string
- `purchase_price`: float (hidden from Admin, public)
- `selling_price`: float
- `stock`: int
- `image_url`: string
- `shop_id`: int (Foreign Key → Shop)
- `created_at`: timestamp

### Transaction
- `id`: int (Primary Key)
- `type`: enum ("Sale", "Expense", "Withdrawal")
- `product_id`: int (Foreign Key → Product, nullable)
- `quantity`: int
- `amount`: float
- `shop_id`: int (Foreign Key → Shop)
- `created_at`: timestamp

## Business Rules:

1. **Multi-Tenant Isolation**
   - All queries MUST filter by shop_id
   - Users can only see data from their own shop
   - shop_id is extracted from JWT, not user input

2. **Role Permissions**
   - SuperAdmin: Full access, can see purchase_price
   - Admin: CRUD operations, cannot see purchase_price
   - Guest: Read-only public access, no purchase_price

3. **Data Visibility**
   - Public API: name, description, category, selling_price, stock, image_url
   - Admin API: All except purchase_price
   - SuperAdmin API: All fields

4. **WhatsApp Integration**
   - Each Shop has a whatsapp_number
   - Products generate dynamic WhatsApp links
   - Format: https://wa.me/{number}?text={message}

5. **Transaction Types**
   - Sale: Requires product_id, decreases stock
   - Expense: No product_id, general expense
   - Withdrawal: No product_id, cash withdrawal

6. **Stock Management**
   - Stock updated on sales
   - Low stock alert when stock < 5
   - Products with stock = 0 still visible

## Indexes (for future database implementation):

```sql
-- Users
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_shop_id ON users(shop_id);

-- Products
CREATE INDEX idx_products_shop_id ON products(shop_id);
CREATE INDEX idx_products_category ON products(category);
CREATE INDEX idx_products_stock ON products(stock);

-- Transactions
CREATE INDEX idx_transactions_shop_id ON transactions(shop_id);
CREATE INDEX idx_transactions_product_id ON transactions(product_id);
CREATE INDEX idx_transactions_type ON transactions(type);
CREATE INDEX idx_transactions_created_at ON transactions(created_at);
```

## Sample Data:

### Shops:
```
| ID | Name                    | Active | WhatsApp      |
|----|-------------------------|--------|---------------|
| 1  | TechStore Casablanca    | true   | 212600000001  |
| 2  | ElectroShop Rabat       | true   | 212600000002  |
```

### Users:
```
| ID | Name         | Email              | Role        | ShopID |
|----|--------------|--------------------|-------------|--------|
| 1  | Super Admin  | super@shop1.com    | SuperAdmin  | 1      |
| 2  | Admin        | admin@shop1.com    | Admin       | 1      |
```

### Products:
```
| ID | Name            | Category     | Purchase | Selling | Stock | ShopID |
|----|-----------------|--------------|----------|---------|-------|--------|
| 1  | iPhone 14 Pro   | Smartphones  | 8000     | 10000   | 15    | 1      |
| 2  | MacBook Pro M2  | Laptops      | 15000    | 18000   | 8     | 1      |
| 3  | Galaxy S23      | Smartphones  | 6000     | 7500    | 20    | 2      |
| 4  | AirPods Pro     | Accessories  | 1500     | 2000    | 3     | 1      |
```

### Transactions:
```
| ID | Type    | ProductID | Quantity | Amount | ShopID |
|----|---------|-----------|----------|--------|--------|
| 1  | Sale    | 1         | 2        | 20000  | 1      |
| 2  | Expense | NULL      | 1        | 5000   | 1      |
```
