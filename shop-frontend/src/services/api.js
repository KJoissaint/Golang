import axios from 'axios'

const API_BASE_URL = 'http://localhost:8081'

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Add auth token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Auth API
export const authAPI = {
  login: (email, password) => api.post('/login', { email, password }),
  register: (data) => api.post('/register', data),
}

// Public API
export const publicAPI = {
  getProducts: (shopId) => api.get(`/public/${shopId}/products`),
  getShops: () => api.get('/shops'),
}

// Products API
export const productsAPI = {
  getAll: () => api.get('/products'),
  create: (data) => api.post('/products', data),
  update: (id, data) => api.put(`/products/${id}`, data),
  delete: (id) => api.delete(`/products/${id}`),
}

// Transactions API
export const transactionsAPI = {
  getAll: () => api.get('/transactions'),
  create: (data) => api.post('/transactions', data),
}

// Dashboard API
export const dashboardAPI = {
  getStats: () => api.get('/reports/dashboard'),
}

// Shop API
export const shopAPI = {
  updateWhatsApp: (whatsappNumber) => api.put('/shops/whatsapp', { whatsapp_number: whatsappNumber }),
  getAll: () => api.get('/shops'),
}

export default api
