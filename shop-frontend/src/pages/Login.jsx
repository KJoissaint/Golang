import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import { useAuth } from '../context/AuthContext'
import Navbar from '../components/Navbar'
import './Auth.css'

const Login = () => {
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [emailError, setEmailError] = useState('')
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  const { login } = useAuth()
  const navigate = useNavigate()
  const [showPassword, setShowPassword] = useState(false)


  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

  const handleEmailChange = (e) => {
    const value = e.target.value
    setEmail(value)

    // Live validation
    if (value.length === 0) {
      setEmailError('')
    } else if (!emailRegex.test(value)) {
      setEmailError('Invalid email format')
    } else {
      setEmailError('')
    }
  }

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError('')

    if (emailError || !email) {
      setError('Please fix the email before submitting.')
      return
    }

    setLoading(true)

    const result = await login(email, password)

    if (result.success) {
      navigate('/dashboard')
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
          <h2>Admin Login</h2>

          {error && <div className="alert alert-error">{error}</div>}

          <form onSubmit={handleSubmit}>
            <div className="form-group">
              <label>Email</label>
              <input
                type="email"
                value={email}
                onChange={handleEmailChange}
                required
                placeholder="admin@shop.com"
                className={emailError ? 'input-error' : ''}
              />
              {emailError && (
                <div className="field-error">{emailError}</div>
              )}
            </div>

            <div className="form-group password-wrapper">
              <label>Password</label>
              <div className="password-container">
                <input
                  type={showPassword ? 'text' : 'password'}
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  required
                  placeholder="Enter your password"
                />
                <span
                  className="eye-icon"
                  onClick={() => setShowPassword(!showPassword)}
                >
                  {showPassword ? '🙈' : '👁️'}
                </span>
              </div>
            </div>


            <button
              type="submit"
              className="btn btn-primary btn-full"
              disabled={loading || emailError}
            >
              {loading ? 'Logging in...' : 'Login'}
            </button>
          </form>

          <p className="auth-link">
            Don't have an account? <Link to="/register">Register here</Link>
          </p>
          <div className="test-credentials">
            <h4>Test Credentials:</h4>
            <p><strong>SuperAdmin:</strong> super@shop1.com / admin123</p>
            <p><strong>Admin:</strong> admin@shop1.com / admin123</p>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Login
