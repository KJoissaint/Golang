import { useState, useEffect } from 'react'
import { useParams } from 'react-router-dom'
import Navbar from '../components/Navbar'
import { publicAPI } from '../services/api'
import './PublicShop.css'

const PublicShop = () => {
  const { shopId } = useParams()
  const [products, setProducts] = useState([])
  const [shops, setShops] = useState([])
  const [selectedShop, setSelectedShop] = useState(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    loadShops()
  }, [])

  useEffect(() => {
    if (shopId) {
      loadProducts(shopId)
    }
  }, [shopId])

  const loadShops = async () => {
    try {
      const response = await publicAPI.getShops()
      setShops(response.data)
    } catch (error) {
      console.error('Error loading shops:', error)
    }
  }

  const loadProducts = async (id) => {
    setLoading(true)
    try {
      const response = await publicAPI.getProducts(id)
      setProducts(response.data || [])
      const shop = shops.find(s => s.id == id)
      setSelectedShop(shop)
    } catch (error) {
      console.error('Error loading products:', error)
    }
    setLoading(false)
  }

  const formatPrice = (price) => {
    return new Intl.NumberFormat('fr-MA', {
      style: 'decimal',
      minimumFractionDigits: 2
    }).format(price) + ' DH'
  }

  const getStockStatus = (stock) => {
    if (stock === 0) return { text: '❌ Out of Stock', class: 'out' }
    if (stock < 5) return { text: `⚠️ Low Stock: ${stock} items`, class: 'low' }
    return { text: `✅ In Stock: ${stock} items`, class: 'in' }
  }

  return (
    <div>
      <Navbar />

      <div className="container" style={{ padding: '2rem 20px' }}>
        <h1 className="page-title">Browse Our Shops</h1>

        {selectedShop && (
          <div className="shop-info">
            <h3>{selectedShop.name}</h3>
            <p>📞 Contact us on WhatsApp: <strong>{selectedShop.whatsapp_number}</strong></p>
          </div>
        )}

        {loading ? (
          <div className="loading">Loading products...</div>
        ) : products.length === 0 ? (
          <div className="empty-state">No products available in this shop.</div>
        ) : (
          <div className="product-grid">
            {products.map(product => {
              const stockStatus = getStockStatus(product.stock)
              return (
                <div key={product.id} className="product-card">
                  <div className="product-image">
                    {product.image_url ? (
                      <img src={product.image_url} alt={product.name} />
                    ) : (
                      <span className="product-icon">📱</span>
                    )}
                  </div>
                  <div className="product-details">
                    <h3>{product.name}</h3>
                    <span className="category">{product.category}</span>
                    <p className="description">{product.description}</p>
                    <div className="price">{formatPrice(product.selling_price)}</div>
                    <div className={`stock ${stockStatus.class}`}>{stockStatus.text}</div>
                    <a
                      href={product.whatsapp_link}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="btn btn-whatsapp"
                    >
                      💬 Contact on WhatsApp
                    </a>
                  </div>
                </div>
              )
            })}
          </div>
        )}
      </div>
    </div>
  )
}

export default PublicShop
