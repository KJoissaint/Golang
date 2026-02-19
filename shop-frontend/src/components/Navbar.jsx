import { Link, useNavigate } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'
import './Navbar.css'

const Navbar = () => {
  const { isAuthenticated, user, logout } = useAuth()
  const navigate = useNavigate()

  const handleLogout = () => {
    logout()
    navigate('/login')
  }

  return (
    <nav className="navbar">
      <div className="container">
        <Link to="/" className="nav-brand">
          <h1>🏪 Electronic Shop</h1>
        </Link>

        <div className="nav-links">
          <Link to="/">Home</Link>
          <Link to="/shop/1">Browse Shops</Link>

          {isAuthenticated ? (
            <>
              <Link to="/dashboard">Dashboard</Link>
              <Link to="/products">Products</Link>
              <Link to="/transactions">Transactions</Link>
              <span className="user-info">👤 {user?.name}</span>
              <button onClick={handleLogout} className="btn btn-sm">Logout</button>
            </>
          ) : (
            <>
              <Link to="/login">Login</Link>
              <Link to="/register">Register</Link>
            </>
          )}
        </div>
      </div>
    </nav>
  )
}

export default Navbar
