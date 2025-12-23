import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import type { Product, ProductFormData } from '@/types/product'
import { productService } from '@/services/api'

export const useProductStore = defineStore('product', () => {
  const products = ref<Product[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)
  const selectedProduct = ref<Product | null>(null)

  const productCount = computed(() => products.value.length)

  async function fetchProducts() {
    loading.value = true
    error.value = null
    try {
      products.value = await productService.getAllProducts()
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch products'
      console.error('Error fetching products:', err)
    } finally {
      loading.value = false
    }
  }

  async function fetchProductById(id: string) {
    loading.value = true
    error.value = null
    try {
      selectedProduct.value = await productService.getProductById(id)
      return selectedProduct.value
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to fetch product'
      console.error('Error fetching product:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function createProduct(data: ProductFormData) {
    loading.value = true
    error.value = null
    try {
      const newProduct = await productService.createProduct(data)
      products.value.push(newProduct)
      return newProduct
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to create product'
      console.error('Error creating product:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function updateProduct(id: string, data: ProductFormData) {
    loading.value = true
    error.value = null
    try {
      const updatedProduct = await productService.updateProduct(id, data)
      const index = products.value.findIndex(p => p.ID === id)
      if (index !== -1) {
        products.value[index] = updatedProduct
      }
      if (selectedProduct.value?.ID === id) {
        selectedProduct.value = updatedProduct
      }
      return updatedProduct
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to update product'
      console.error('Error updating product:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  async function deleteProduct(id: string) {
    loading.value = true
    error.value = null
    try {
      await productService.deleteProduct(id)
      products.value = products.value.filter(p => p.ID !== id)
      if (selectedProduct.value?.ID === id) {
        selectedProduct.value = null
      }
    } catch (err: any) {
      error.value = err.response?.data?.error || 'Failed to delete product'
      console.error('Error deleting product:', err)
      throw err
    } finally {
      loading.value = false
    }
  }

  function clearError() {
    error.value = null
  }

  function setSelectedProduct(product: Product | null) {
    selectedProduct.value = product
  }

  return {
    products,
    loading,
    error,
    selectedProduct,
    productCount,
    fetchProducts,
    fetchProductById,
    createProduct,
    updateProduct,
    deleteProduct,
    clearError,
    setSelectedProduct,
  }
})

