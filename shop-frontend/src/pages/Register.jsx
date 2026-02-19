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
    confirmPassword: '',
    role: 'Admin',
    shop_id: 1
  })

  const [showPassword, setShowPassword] = useState(false)
  const [showConfirmPassword, setShowConfirmPassword] = useState(false)
  const [passwordError, setPasswordError] = useState('')
  const [error, setError] = useState('')
  const [success, setSuccess] = useState(false)
  const [loading, setLoading] = useState(false)

  const { register } = useAuth()
  const navigate = useNavigate()

  const handleChange = (e) => {
    const { name, value } = e.target

    setFormData({
      ...formData,
      [name]: value
    })

    // Live password match validation
    if (name === 'password' || name === 'confirmPassword') {
      if (
        name === 'confirmPassword'
          ? value !== formData.password
          : formData.confirmPassword && value !== formData.confirmPassword
      ) {
        setPasswordError('Passwords do not match')
      } else {
        setPasswordError('')
      }
    }
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError('')

    if (passwordError) {
      setError('Please fix password errors.')
      return
    }

    setLoading(true)

    const result = await register({
      ...formData,
      confirmPassword: undefined // don’t send this to backend
    })

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
          {success && <div className="alert alert-success">Registration successful! Redirecting...</div>}

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

            {/* Password Field */}
            <div className="form-group password-wrapper">
              <label>Password</label>
              <div className="password-container">
                <input
                  type={showPassword ? 'text' : 'password'}
                  name="password"
                  value={formData.password}
                  onChange={handleChange}
                  required
                  placeholder="Minimum 6 characters"
                />
                <span
                  className="eye-icon"
                  onClick={() => setShowPassword(!showPassword)}
                >
                  {showPassword ? '🙈' : '👁️'}
                </span>
              </div>
            </div>

            {/* Confirm Password */}
            <div className="form-group password-wrapper">
              <label>Confirm Password</label>
              <div className="password-container">
                <input
                  type={showConfirmPassword ? 'text' : 'password'}
                  name="confirmPassword"
                  value={formData.confirmPassword}
                  onChange={handleChange}
                  required
                  placeholder="Re-enter password"
                />
                <span
                  className="eye-icon"
                  onClick={() => setShowConfirmPassword(!showConfirmPassword)}
                >
                  {showConfirmPassword ? '🙈' : '👁️'}
                </span>
              </div>
              {passwordError && (
                <div className="field-error">{passwordError}</div>
              )}
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
            </div>

            <button
              type="submit"
              className="btn btn-primary btn-full"
              disabled={loading || passwordError}
            >
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
