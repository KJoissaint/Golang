import { Link } from 'react-router-dom'
import Navbar from '../components/Navbar'
import './Home.css'

const Home = () => {
  return (
    <div>
      <Navbar />

      <div className="hero">
        <div className="container">
          <h1>Welcome to Electronic Shop Management System</h1>
          <p>Multi-shop platform for electronics retailers</p>
          <div className="hero-buttons">
            <Link to="/shop/1" className="btn btn-primary">Browse Products</Link>
            <Link to="/login" className="btn btn-secondary">Admin Login</Link>
          </div>
        </div>
      </div>

      <div className="features">
        <div className="container">
          <h2>Features</h2>
          <div className="feature-grid">
            <div className="feature-card">
              <div className="feature-icon">🛍️</div>
              <h3>Public Shopping</h3>
              <p>Browse products from multiple shops without login</p>
            </div>
            <div className="feature-card">
              <div className="feature-icon">💬</div>
              <h3>WhatsApp Integration</h3>
              <p>Contact shops directly via WhatsApp for inquiries</p>
            </div>
            <div className="feature-card">
              <div className="feature-icon">👑</div>
              <h3>Multi-Role Management</h3>
              <p>SuperAdmin and Admin roles for shop management</p>
            </div>
            <div className="feature-card">
              <div className="feature-icon">📊</div>
              <h3>Analytics Dashboard</h3>
              <p>Track sales, expenses, and profits in real-time</p>
            </div>
            <div className="feature-card">
              <div className="feature-icon">🏢</div>
              <h3>Multi-Tenant</h3>
              <p>Complete isolation between different shops</p>
            </div>
            <div className="feature-card">
              <div className="feature-icon">🔒</div>
              <h3>Secure & Safe</h3>
              <p>JWT authentication and role-based access control</p>
            </div>
          </div>
        </div>
      </div>

      <footer className="footer">
        <div className="container">
          <p>&copy; 2026 Electronic Shop Management System. Built with Go & React.</p>
        </div>
      </footer>
    </div>
  )
}

export default Home
