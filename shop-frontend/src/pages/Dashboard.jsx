import { useState, useEffect } from 'react'
import Navbar from '../components/Navbar'
import { useAuth } from '../context/AuthContext'
import { dashboardAPI, transactionsAPI, shopAPI, productsAPI } from '../services/api'
import './Dashboard.css'

const Dashboard = () => {
  const { user, isSuperAdmin } = useAuth()
  const [stats, setStats] = useState(null)
  const [transactions, setTransactions] = useState([])
  const [shopName, setShopName] = useState('')
  const [products, setProducts] = useState([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    loadData()
  }, [])

  const loadData = async () => {
  try {
    if (isSuperAdmin()) {
      const statsResponse = await dashboardAPI.getStats()
      setStats(statsResponse.data)
    }

    const transResponse = await transactionsAPI.getAll()
    setTransactions(transResponse.data?.slice(0, 5) || [])

    const productsResponse = await productsAPI.getAll()
    setProducts(productsResponse.data || [])

    const shopsResponse = await shopAPI.getAll()
    const shop = shopsResponse.data.find(s => s.id === user.shop_id)
    if (shop) setShopName(shop.name)

  } catch (error) {
    console.error('Error loading dashboard:', error)
  }

  setLoading(false)
}


  const formatCurrency = (amount) => {
    return new Intl.NumberFormat('fr-MA', {
      style: 'decimal',
      minimumFractionDigits: 2
    }).format(amount) + ' DH'
  }

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('fr-FR', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit'
    })
  }

  if (loading) {
    return (
      <div>
        <Navbar />
        <div className="container" style={{ padding: '2rem', textAlign: 'center' }}>
          Loading...
        </div>
      </div>
    )
  }
  const totalProducts = products.length
  const totalStock = products.reduce((sum, p) => sum + p.stock, 0)
  const lowStockCount = products.filter(p => p.stock < 5).length

  return (
    <div>
      <Navbar />

      <div className="container dashboard">
        <div className="dashboard-header">
          <div>
            <h2>Welcome, {user?.name}</h2>
            <p>Role: <strong>{user?.role}</strong> | Shop: <strong>{shopName}</strong></p>
          </div>
        </div>

        {isSuperAdmin() && stats && (
          <div className="stats-grid">
            <div className="stat-card success">
              <div className="stat-label">Total Sales</div>
              <div className="stat-value">{formatCurrency(stats.total_sales || 0)}</div>
            </div>
            <div className="stat-card danger">
              <div className="stat-label">Total Expenses</div>
              <div className="stat-value">{formatCurrency(stats.total_expenses || 0)}</div>
            </div>
            <div className="stat-card primary">
              <div className="stat-label">Net Profit</div>
              <div className="stat-value">{formatCurrency(stats.net_profit || 0)}</div>
            </div>
            <div className="stat-card warning">
              <div className="stat-label">Low Stock Items</div>
              <div className="stat-value">{stats.low_stock_count || 0}</div>
            </div>
          </div>
        )}
        <div className="stats-grid">
          <div className="stat-card primary">
            <div className="stat-label">Total Products</div>
            <div className="stat-value">{totalProducts}</div>
          </div>

          <div className="stat-card success">
            <div className="stat-label">Total Stock Units</div>
            <div className="stat-value">{totalStock}</div>
          </div>

          <div className="stat-card warning">
            <div className="stat-label">Low Stock Items</div>
            <div className="stat-value">{lowStockCount}</div>
          </div>
        </div>

        <div className="section">
          <div className="section-header">
            <h3>Recent Transactions</h3>
            <button
              className="btn btn-primary btn-sm"
              onClick={() => window.location.href = '/transactions'}
            >
              + Add Transaction
            </button>
          </div>
          <div className="table-container">
            <table>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Type</th>
                  <th>Amount</th>
                  <th>Quantity</th>
                  <th>Date</th>
                </tr>
              </thead>
              <tbody>
                {transactions.length === 0 ? (
                  <tr><td colSpan="5" style={{ textAlign: 'center' }}>No transactions yet</td></tr>
                ) : (
                  transactions.map(t => (
                    <tr key={t.id}>
                      <td>{t.id}</td>
                      <td><span className="badge">{t.type}</span></td>
                      <td>{formatCurrency(t.amount)}</td>
                      <td>{t.quantity}</td>
                      <td>{formatDate(t.created_at)}</td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Dashboard
