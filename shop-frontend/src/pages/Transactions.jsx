import { useState, useEffect } from 'react'
import Navbar from '../components/Navbar'
import { transactionsAPI, productsAPI } from '../services/api'
import './Transactions.css'

const Transactions = () => {
  const [transactions, setTransactions] = useState([])
  const [products, setProducts] = useState([])
  const [showModal, setShowModal] = useState(false)
  const [formData, setFormData] = useState({
    type: '',
    product_id: '',
    quantity: '',
    amount: ''
  })

  useEffect(() => {
    loadTransactions()
    loadProducts()
  }, [])

  const loadTransactions = async () => {
    try {
      const response = await transactionsAPI.getAll()
      setTransactions(response.data || [])
    } catch (error) {
      console.error('Error loading transactions:', error)
    }
  }

  const loadProducts = async () => {
    try {
      const response = await productsAPI.getAll()
      setProducts(response.data || [])
    } catch (error) {
      console.error('Error loading products:', error)
    }
  }

  const handleSubmit = async (e) => {
  e.preventDefault()

  try {
    const data = {
      type: formData.type,
      quantity: Number(formData.quantity),
      amount: Number(formData.amount)
    }

    if (formData.type === 'Sale') {
      data.product_id = Number(formData.product_id)
    }

    console.log("📦 Payload being sent:", data)

    await transactionsAPI.create(data)

    setShowModal(false)
    resetForm()
    loadTransactions()
  } catch (error) {
    console.error(" Full error response:", error.response)
    alert(error.response?.data?.error || 'Error creating transaction')
  }
}


  const resetForm = () => {
    setFormData({
      type: '',
      product_id: '',
      quantity: '',
      amount: ''
    })
  }

  const formatCurrency = (amount) => {
    return new Intl.NumberFormat('fr-MA').format(amount) + ' DH'
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

  return (
    <div>
      <Navbar />

      <div className="container dashboard">
        <div className="dashboard-header">
          <h2>Transactions</h2>
          <button className="btn btn-primary" onClick={() => setShowModal(true)}>
            + Add Transaction
          </button>
        </div>

        <div className="table-container">
          <table>
            <thead>
              <tr>
                <th>ID</th>
                <th>Type</th>
                <th>Product ID</th>
                <th>Quantity</th>
                <th>Amount</th>
                <th>Date</th>
              </tr>
            </thead>
            <tbody>
              {transactions.length === 0 ? (
                <tr><td colSpan="6" style={{ textAlign: 'center' }}>No transactions yet</td></tr>
              ) : (
                transactions.map(t => (
                  <tr key={t.id}>
                    <td>{t.id}</td>
                    <td><span className={`badge badge-${t.type.toLowerCase()}`}>{t.type}</span></td>
                    <td>{t.product_id || 'N/A'}</td>
                    <td>{t.quantity}</td>
                    <td>{formatCurrency(t.amount)}</td>
                    <td>{formatDate(t.created_at)}</td>
                  </tr>
                ))
              )}
            </tbody>
          </table>
        </div>

        {showModal && (
          <div className="modal" onClick={() => setShowModal(false)}>
            <div className="modal-content" onClick={(e) => e.stopPropagation()}>
              <div className="modal-header">
                <h3>Add Transaction</h3>
                <button onClick={() => setShowModal(false)} className="close">&times;</button>
              </div>

              <form onSubmit={handleSubmit}>
                <div className="form-group">
                  <label>Transaction Type</label>
                  <select
                    value={formData.type}
                    onChange={(e) => setFormData({...formData, type: e.target.value})}
                    required
                  >
                    <option value="">Select type...</option>
                    <option value="Sale">Sale</option>
                    <option value="Expense">Expense</option>
                    <option value="Withdrawal">Withdrawal</option>
                  </select>
                </div>

                {formData.type === 'Sale' && (
                  <div className="form-group">
                    <label>Product</label>
                    <select
                      value={formData.product_id}
                      onChange={(e) => setFormData({...formData, product_id: e.target.value})}
                      required
                    >
                      <option value="">Select product...</option>
                      {products.map(p => (
                        <option key={p.id} value={p.id}>
                          {p.name} (Stock: {p.stock})
                        </option>
                      ))}
                    </select>
                  </div>
                )}

                <div className="form-group">
                  <label>Quantity</label>
                  <input
                    type="number"
                    min="1"
                    value={formData.quantity}
                    onChange={(e) => setFormData({...formData, quantity: e.target.value})}
                    required
                  />
                </div>

                <div className="form-group">
                  <label>Amount (DH)</label>
                  <input
                    type="number"
                    step="0.01"
                    min="0"
                    value={formData.amount}
                    onChange={(e) => setFormData({...formData, amount: e.target.value})}
                    required
                  />
                </div>

                <div className="modal-actions">
                  <button type="submit" className="btn btn-primary">Save</button>
                  <button type="button" onClick={() => setShowModal(false)} className="btn btn-secondary">
                    Cancel
                  </button>
                </div>
              </form>
            </div>
          </div>
        )}
      </div>
    </div>
  )
}

export default Transactions
