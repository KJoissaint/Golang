import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'
import Navbar from '../components/Navbar'
import './Auth.css'

const Register = () => {
  const [formData, setFormData] = useState({
    name: '',
    email: '',
    password: '',
    role: 'Admin',
    shop_id: 1
  })
  const [error, setError] = useState('')
  const [success, setSuccess] = useState(false)
  const [loading, setLoading] = useState(false)
  const { register } = useAuth()
  const navigate = useNavigate()

  const handleChange = (e) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value
    })
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError('')
    setLoading(true)

    const result = await register(formData)

    if (result.success) {
      setSuccess(true)
      setTimeout(() => {
        navigate('/login')
      }, 2000)
    } else {
      setError(result.error)
    }

    setLoading(false)
  }

  return (
    <div>
      <Navbar />

      <div className="auth-container">
        <div className="auth-card">
          <h2>Create Account</h2>

          {error && <div className="alert alert-error">{error}</div>}
          {success && <div className="alert alert-success">Registration successful! Redirecting to login...</div>}

          <form onSubmit={handleSubmit}>
            <div className="form-group">
              <label>Full Name</label>
              <input
                type="text"
                name="name"
                value={formData.name}
                onChange={handleChange}
                required
                placeholder="John Doe"
              />
            </div>

            <div className="form-group">
              <label>Email</label>
              <input
                type="email"
                name="email"
                value={formData.email}
                onChange={handleChange}
                required
                placeholder="john@example.com"
              />
            </div>

            <div className="form-group">
              <label>Password</label>
              <input
                type="password"
                name="password"
                value={formData.password}
                onChange={handleChange}
                required
                placeholder="Minimum 6 characters"
              />
            </div>

            <div className="form-group">
              <label>Role</label>
              <select name="role" value={formData.role} onChange={handleChange}>
                <option value="Admin">Admin</option>
                <option value="SuperAdmin">Super Admin</option>
              </select>
            </div>

            <div className="form-group">
              <label>Shop ID</label>
              <input
                type="number"
                name="shop_id"
                value={formData.shop_id}
                onChange={handleChange}
                required
              />
              <small>Enter 1 for TechStore Casablanca, 2 for ElectroShop Rabat</small>
            </div>

            <button type="submit" className="btn btn-primary btn-full" disabled={loading}>
              {loading ? 'Registering...' : 'Register'}
            </button>
          </form>

          <p className="auth-link">
            Already have an account? <Link to="/login">Login here</Link>
          </p>
        </div>
      </div>
    </div>
  )
}

export default Register
