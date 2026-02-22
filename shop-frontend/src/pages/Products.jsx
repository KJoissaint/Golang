import { useState, useEffect } from 'react'
import Navbar from '../components/Navbar'
import { useAuth } from '../context/AuthContext'
import { productsAPI } from '../services/api'
import './Products.css'

const Products = () => {
  const { isSuperAdmin } = useAuth()
  const [products, setProducts] = useState([])
  const [showModal, setShowModal] = useState(false)
  const [editingProduct, setEditingProduct] = useState(null)
  const [viewingProduct, setViewingProduct] = useState(null)
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    category: '',
    purchase_price: '',
    selling_price: '',
    stock: '',
    image_url: ''
  })

  useEffect(() => {
    loadProducts()
  }, [])

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
    // 🔥 Convert numeric fields properly
    const payload = {
      name: formData.name,
      description: formData.description,
      category: formData.category,
      purchase_price: Number(formData.purchase_price),
      selling_price: Number(formData.selling_price),
      stock: Number(formData.stock),
      image_url: formData.image_url
    }

    console.log("📦 Product payload:", payload)

    if (editingProduct) {
      await productsAPI.update(editingProduct.id, payload)
    } else {
      await productsAPI.create(payload)
    }

    setShowModal(false)
    resetForm()
    loadProducts()

  } catch (error) {
    console.error("❌ Full error:", error.response)
    alert(error.response?.data?.error || 'Error saving product')
  }
}

  const handleDelete = async (id) => {
    if (!confirm('Are you sure you want to delete this product?')) return
    try {
      await productsAPI.delete(id)
      loadProducts()
    } catch (error) {
      alert(error.response?.data?.error || 'Error deleting product')
    }
  }

  const openModal = (product = null) => {
    if (product) {
      setEditingProduct(product)
      setFormData({
        name: product.name,
        description: product.description,
        category: product.category,
        purchase_price: product.purchase_price,
        selling_price: product.selling_price,
        stock: product.stock,
        image_url: product.image_url || ''
      })
    } else {
      resetForm()
    }
    setShowModal(true)
  }

  const resetForm = () => {
    setEditingProduct(null)
    setFormData({
      name: '',
      description: '',
      category: '',
      purchase_price: '',
      selling_price: '',
      stock: '',
      image_url: ''
    })
  }

  const formatPrice = (price) => {
    return new Intl.NumberFormat('fr-MA').format(price) + ' DH'
  }

  return (
    <div>
      <Navbar />

      <div className="container dashboard">
        <div className="dashboard-header">
          <h2>Products</h2>
          <button className="btn btn-primary" onClick={() => openModal()}>
            + Add Product
          </button>
        </div>

        <div className="product-grid">
          {products.map(product => (
            <div
              key={product.id}
              className="product-card clickable"
              onClick={() => setViewingProduct(product)}
            >
              <div className="product-image">
                {product.image_url ? (
                  <img src={product.image_url} alt={product.name} />
                ) : (
                  <span className="product-icon">📦</span>
                )}
              </div>
              <div className="product-details">
                <h3>{product.name}</h3>
                <span className="category">{product.category}</span>
                <p className="description">{product.description}</p>

                {isSuperAdmin() && (
                  <div className="price-row">
                    <small>Purchase: {formatPrice(product.purchase_price)}</small>
                  </div>
                )}

                <div className="price">{formatPrice(product.selling_price)}</div>

                {isSuperAdmin() && (
                  <div className="profit">
                    Profit: {formatPrice(product.selling_price - product.purchase_price)}
                  </div>
                )}

                <div className="stock">Stock: {product.stock} items</div>

                <div className="actions">
                  <button
                    onClick={(e) => {
                      e.stopPropagation()
                      openModal(product)
                    }}
                    className="btn btn-sm btn-primary"
                  >
                    Edit
                  </button>

                  <button
                    onClick={(e) => {
                      e.stopPropagation()
                      handleDelete(product.id)
                    }}
                    className="btn btn-sm btn-danger"
                  >
                    Delete
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>

        {showModal && (
          <div className="modal" onClick={() => setShowModal(false)}>
            <div className="modal-content" onClick={(e) => e.stopPropagation()}>
              <div className="modal-header">
                <h3>{editingProduct ? 'Edit Product' : 'Add Product'}</h3>
                <button onClick={() => setShowModal(false)} className="close">&times;</button>
              </div>

              <form onSubmit={handleSubmit}>
                <div className="form-group">
                  <label>Product Name</label>
                  <input
                    type="text"
                    value={formData.name}
                    onChange={(e) => setFormData({...formData, name: e.target.value})}
                    required
                  />
                </div>

                <div className="form-group">
                  <label>Description</label>
                  <textarea
                    value={formData.description}
                    onChange={(e) => setFormData({...formData, description: e.target.value})}
                    rows="3"
                    required
                  />
                </div>

                <div className="form-group">
                  <label>Category</label>
                  <input
                    type="text"
                    value={formData.category}
                    onChange={(e) => setFormData({...formData, category: e.target.value})}
                    required
                  />
                </div>

                <div className="form-group">
                  <label>Purchase Price (DH)</label>
                  <input
                    type="number"
                    step="0.01"
                    value={formData.purchase_price}
                    onChange={(e) => setFormData({...formData, purchase_price: e.target.value})}
                    required
                  />
                </div>

                <div className="form-group">
                  <label>Selling Price (DH)</label>
                  <input
                    type="number"
                    step="0.01"
                    value={formData.selling_price}
                    onChange={(e) => setFormData({...formData, selling_price: e.target.value})}
                    required
                  />
                </div>

                <div className="form-group">
                  <label>Stock</label>
                  <input
                    type="number"
                    value={formData.stock}
                    onChange={(e) => setFormData({...formData, stock: e.target.value})}
                    required
                  />
                </div>

                <div className="form-group">
                  <label>Image URL</label>
                  <input
                    type="url"
                    value={formData.image_url}
                    onChange={(e) => setFormData({...formData, image_url: e.target.value})}
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
        {viewingProduct && (
          <div className="modal" onClick={() => setViewingProduct(null)}>
            <div className="modal-content" onClick={(e) => e.stopPropagation()}>
              <div className="modal-header">
                <h3>Product Details</h3>
                <button onClick={() => setViewingProduct(null)} className="close">
                  &times;
                </button>
              </div>

              <div className="product-details-view">
                {viewingProduct.image_url ? (
                  <img
                    src={viewingProduct.image_url}
                    alt={viewingProduct.name}
                    className="details-image"
                  />
                ) : (
                  <div className="details-image-placeholder">📦</div>
                )}

                <h2>{viewingProduct.name}</h2>
                <p><strong>Category:</strong> {viewingProduct.category}</p>
                <p><strong>Description:</strong> {viewingProduct.description}</p>
                <p><strong>Selling Price:</strong> {formatPrice(viewingProduct.selling_price)}</p>

                {isSuperAdmin() && (
                  <>
                    <p><strong>Purchase Price:</strong> {formatPrice(viewingProduct.purchase_price)}</p>
                    <p>
                      <strong>Profit per Unit:</strong>{" "}
                      {formatPrice(viewingProduct.selling_price - viewingProduct.purchase_price)}
                    </p>
                  </>
                )}

                <p><strong>Stock:</strong> {viewingProduct.stock} items</p>

                {viewingProduct.stock < 5 && (
                  <div className="low-stock-warning">⚠ Low stock!</div>
                )}
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  )
}

export default Products
