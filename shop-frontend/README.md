# 🎨 Electronic Shop Management - React Frontend

## Overview
Modern React frontend application for the Electronic Shop Management API built with **React 18**, **React Router**, **Axios**, and **Vite**.

## 🚀 Tech Stack

- **React 18** - Modern React with Hooks
- **React Router v6** - Client-side routing
- **Axios** - HTTP client for API calls
- **Vite** - Lightning-fast build tool
- **CSS3** - Custom responsive styling (no frameworks)

## 📁 Project Structure

```
shop-frontend/
├── index.html                 # HTML template
├── package.json               # Dependencies
├── vite.config.js             # Vite configuration
└── src/
    ├── main.jsx               # Entry point
    ├── App.jsx                # Main app component with routing
    ├── index.css              # Global styles
    │
    ├── components/            # Reusable components
    │   ├── Navbar.jsx        # Navigation bar
    │   ├── Navbar.css
    │   └── PrivateRoute.jsx  # Protected route wrapper
    │
    ├── context/               # React Context
    │   └── AuthContext.jsx   # Authentication state management
    │
    ├── services/              # API services
    │   └── api.js            # Axios setup and API calls
    │
    └── pages/                 # Page components
        ├── Home.jsx          # Landing page
        ├── Home.css
        ├── Login.jsx         # Login page
        ├── Register.jsx      # Registration page
        ├── Auth.css          # Shared auth styles
        ├── PublicShop.jsx    # Public product browsing
        ├── PublicShop.css
        ├── Dashboard.jsx     # Admin dashboard
        ├── Dashboard.css
        ├── Products.jsx      # Product management
        ├── Products.css
        ├── Transactions.jsx  # Transaction management
        └── Transactions.css
```

## 🔧 Installation

### Prerequisites
- Node.js 16+ installed
- Backend API running on http://localhost:8081

### Setup

```bash
cd C:\Users\tamim\GolandProjects\Golang\shop-frontend
npm install
```

## ▶️ Running the Application

### Development Mode (Recommended)
```bash
npm run dev
```
The app will run on **http://localhost:3000**

### Build for Production
```bash
npm run build
```

### Preview Production Build
```bash
npm run preview
```

## 🌐 Features

### Public Features (No Login Required)
✅ Browse products from different shops  
✅ View product details and stock  
✅ WhatsApp integration (one-click contact)  
✅ Real-time stock status indicators

### Admin Features (Login Required)
✅ Product CRUD operations  
✅ Transaction management (Sale/Expense/Withdrawal)  
✅ View stock levels  
✅ Create and edit products

### SuperAdmin Features
✅ Full dashboard with statistics  
✅ View purchase prices and profit margins  
✅ Access to all financial data  
✅ Complete analytics

## 🔐 Authentication

The app uses **JWT-based authentication**:

1. User logs in with email/password
2. Backend returns JWT token
3. Token stored in localStorage
4. Token sent in Authorization header for protected routes
5. Context API manages auth state globally

### Test Accounts

**SuperAdmin:**
- Email: `super@shop1.com`
- Password: `admin123`

**Admin:**
- Email: `admin@shop1.com`
- Password: `admin123`

## 📱 Pages Overview

### 1. Home (`/`)
- Landing page with features showcase
- Links to public shop and login
- Responsive hero section
- Feature cards

### 2. Public Shop (`/shop/:shopId`)
- Browse products without login
- View prices and stock
- WhatsApp contact buttons
- No purchase prices shown (security)

### 3. Login (`/login`)
- Email/password authentication
- JWT token storage
- Redirect to dashboard on success
- Test credentials displayed

### 4. Register (`/register`)
- Create new admin accounts
- Select role (Admin/SuperAdmin)
- Assign to shop
- Input validation

### 5. Dashboard (`/dashboard`)
- Protected route (login required)
- Statistics cards (SuperAdmin only)
- Recent transactions table
- User info display

### 6. Products (`/products`)
- Protected route
- Product grid view
- Add/edit/delete products
- Modal for product form
- Role-based display (SuperAdmin sees purchase prices)

### 7. Transactions (`/transactions`)
- Protected route
- Transaction history table
- Add new transactions
- Type selection (Sale/Expense/Withdrawal)
- Product selection for sales

## 🎨 Design Features

- **Responsive Design** - Works on desktop, tablet, mobile
- **Modern UI** - Clean, professional interface
- **Color-Coded** - Different colors for transaction types, stock levels
- **Loading States** - Visual feedback for async operations
- **Error Handling** - User-friendly error messages
- **Modals** - For forms (products, transactions)
- **Protected Routes** - Automatic redirect to login if not authenticated

## 🔌 API Integration

All API calls are configured in `src/services/api.js`:

```javascript
// Example API calls
authAPI.login(email, password)
publicAPI.getProducts(shopId)
productsAPI.getAll()
productsAPI.create(data)
transactionsAPI.create(data)
dashboardAPI.getStats()
```

### API Endpoints Used

- `POST /login` - User authentication
- `POST /register` - User registration
- `GET /public/:shopID/products` - Public products
- `GET /products` - Private products list
- `POST /products` - Create product
- `PUT /products/:id` - Update product
- `DELETE /products/:id` - Delete product
- `GET /transactions` - List transactions
- `POST /transactions` - Create transaction
- `GET /reports/dashboard` - Dashboard stats (SuperAdmin)
- `GET /shops` - List all shops

## 🔒 Security Features

1. **JWT Authentication** - All private routes protected
2. **Role-Based Access** - Different views for Admin vs SuperAdmin
3. **Purchase Price Protection** - Never shown in public routes
4. **Auto Token Refresh** - Token included in all requests
5. **Protected Routes** - Automatic redirect to login
6. **Input Validation** - Form validation on client side

## 🎯 Advantages Over Vanilla JS Frontend

### 1. **Better Code Organization**
- Component-based architecture
- Separation of concerns
- Reusable components

### 2. **State Management**
- React Context for global state
- No prop drilling
- Centralized auth state

### 3. **Performance**
- Virtual DOM for efficient updates
- Component memoization
- Automatic re-renders

### 4. **Developer Experience**
- Hot Module Replacement (HMR)
- Fast refresh
- Better debugging with React DevTools

### 5. **Maintainability**
- Easy to add new features
- Clear component hierarchy
- Type-safe with PropTypes (optional)

### 6. **Routing**
- Client-side routing with React Router
- Protected routes
- URL parameters
- Programmatic navigation

## 🛠️ Customization

### Change API URL
Edit `src/services/api.js`:
```javascript
const API_BASE_URL = 'http://your-backend-url:port'
```

### Change Colors
Edit CSS variables in `src/index.css`:
```css
:root {
  --primary-color: #2563eb;
  --secondary-color: #7c3aed;
  /* ... */
}
```

### Add New Pages
1. Create component in `src/pages/`
2. Add route in `src/App.jsx`
3. Add link in `src/components/Navbar.jsx`

## 📊 Project Comparison

### Vanilla JS Frontend (old)
- ❌ Manual DOM manipulation
- ❌ No state management
- ❌ Harder to maintain
- ✅ No build step
- ✅ Smaller bundle

### React Frontend (new)
- ✅ Declarative UI
- ✅ Component reusability
- ✅ Better state management
- ✅ Easier to scale
- ✅ Modern dev tools
- ✅ Better developer experience

## 🚀 Deployment

### Build for Production
```bash
npm run build
```

This creates a `dist/` folder with optimized files.

### Deploy Options
- **Vercel** - `vercel deploy`
- **Netlify** - Drag and drop `dist/` folder
- **GitHub Pages** - Push `dist/` to gh-pages branch
- **Any static host** - Upload `dist/` contents

### Environment Variables
For production, update the API URL in `src/services/api.js` or use environment variables:

```javascript
const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8081'
```

Create `.env` file:
```
VITE_API_URL=https://your-production-api.com
```

## 📝 Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run preview` - Preview production build

## 🎓 For Your Presentation

### Demo Flow

1. **Show Modern Architecture**
   - Component-based structure
   - Clean separation of concerns
   - React Context for state management

2. **Public Access**
   - Browse products without login
   - WhatsApp integration
   - Responsive design

3. **Authentication**
   - Login flow
   - JWT token handling
   - Protected routes

4. **Admin Features**
   - Product management with modals
   - Transaction creation
   - Real-time updates

5. **SuperAdmin Features**
   - Full dashboard
   - Statistics
   - Purchase prices

## ✅ Assignment Compliance

This React frontend provides:
- ✅ Professional, modern UI
- ✅ Separate from backend (better architecture)
- ✅ All required functionality
- ✅ Better maintainability
- ✅ Enhanced user experience
- ✅ Production-ready code

## 🎉 Conclusion

The React frontend is a **significant improvement** over vanilla JavaScript:

- **Better organized** - Component-based architecture
- **More maintainable** - Easy to add features
- **Better UX** - Smooth transitions and updates
- **Professional** - Modern development practices
- **Scalable** - Easy to extend and modify

Perfect for your assignment demonstration! 🚀

---

**Frontend URL**: http://localhost:3000  
**Backend URL**: http://localhost:8081

**Happy Coding!** 💻
