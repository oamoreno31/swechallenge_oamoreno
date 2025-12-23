import axios from 'axios'
import type { Product, ProductFormData } from '@/types/product'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api'

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export const productService = {
  async getAllProducts(): Promise<Product[]> {
    const response = await api.get<Product[]>('/products')
    console.log("Fetching all products ", response)
    return response.data
  },

  async getProductById(id: string): Promise<Product> {
    const response = await api.get<Product>(`/products/${id}`)
    return response.data
  },

  async createProduct(data: ProductFormData): Promise<Product> {
    const response = await api.post<Product>('/products', data)
    return response.data
  },

  async updateProduct(id: string, data: ProductFormData): Promise<Product> {
    const response = await api.put<Product>(`/products/${id}`, data)
    return response.data
  },

  async deleteProduct(id: string): Promise<void> {
    await api.delete(`/products/${id}`)
  },
}

