<template>
  <div class="min-h-screen bg-gradient-to-br from-primary-900 via-primary-800 to-primary-700 p-8">
    <div class="max-w-4xl mx-auto">
      <div class="mb-8">
        <router-link
          to="/products"
          class="text-primary-300 hover:text-primary-200 mb-4 inline-block"
        >
          ← Back to Products
        </router-link>
        <h1 class="text-4xl font-bold text-white mb-2">Update Product</h1>
        <p class="text-primary-200">Edit product information</p>
      </div>

      <div v-if="loading && !product" class="text-center py-12">
        <div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-white"></div>
        <p class="text-primary-200 mt-4">Loading product...</p>
      </div>

      <div v-else-if="error && !product" class="bg-red-500/20 border border-red-500 rounded-lg p-6">
        <p class="text-red-300 mb-4">{{ error }}</p>
        <router-link
          to="/products"
          class="text-primary-300 hover:text-primary-200"
        >
          ← Back to Products
        </router-link>
      </div>

      <div v-else class="bg-white/10 backdrop-blur-lg rounded-lg p-8 border border-primary-600">
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-primary-200 font-medium mb-2">
                Ticker <span class="text-red-400">*</span>
              </label>
              <input
                v-model="formData.Ticker"
                type="text"
                required
                class="w-full bg-white/10 border border-primary-600 rounded-lg px-4 py-2 text-white placeholder-primary-400 focus:outline-none focus:ring-2 focus:ring-primary-500"
                placeholder="e.g., AAPL"
              />
            </div>

            <div>
              <label class="block text-primary-200 font-medium mb-2">
                Company <span class="text-red-400">*</span>
              </label>
              <input
                v-model="formData.Company"
                type="text"
                required
                class="w-full bg-white/10 border border-primary-600 rounded-lg px-4 py-2 text-white placeholder-primary-400 focus:outline-none focus:ring-2 focus:ring-primary-500"
                placeholder="Company Name"
              />
            </div>

            <div>
              <label class="block text-primary-200 font-medium mb-2">
                Target From <span class="text-red-400">*</span>
              </label>
              <input
                v-model="formData.Target_from"
                type="text"
                required
                class="w-full bg-white/10 border border-primary-600 rounded-lg px-4 py-2 text-white placeholder-primary-400 focus:outline-none focus:ring-2 focus:ring-primary-500"
                placeholder="e.g., 150"
                name="target_from"
                v-on:change="handleOnChange"
              />
            </div>

            <div>
              <label class="block text-primary-200 font-medium mb-2">
                Target To <span class="text-red-400">*</span>
              </label>
              <input
                v-model="formData.Target_to"
                type="text"
                required
                class="w-full bg-white/10 border border-primary-600 rounded-lg px-4 py-2 text-white placeholder-primary-400 focus:outline-none focus:ring-2 focus:ring-primary-500"
                placeholder="e.g., 180"
                name="target_to"
                v-on:change="handleOnChange"
              />
            </div>

            <div>
              <label class="block text-primary-200 font-medium mb-2">
                Action <span class="text-red-400">*</span>
              </label>
              <input
                v-model="formData.Action"
                type="text"
                required
                class="w-full bg-white/10 border border-primary-600 rounded-lg px-4 py-2 text-white placeholder-primary-400 focus:outline-none focus:ring-2 focus:ring-primary-500"
                placeholder="e.g., Buy, Sell, Hold"
              />
            </div>

            <div>
              <label class="block text-primary-200 font-medium mb-2">
                Brokerage <span class="text-red-400">*</span>
              </label>
              <input
                v-model="formData.Brokerage"
                type="text"
                required
                class="w-full bg-white/10 border border-primary-600 rounded-lg px-4 py-2 text-white placeholder-primary-400 focus:outline-none focus:ring-2 focus:ring-primary-500"
                placeholder="Brokerage Firm"
              />
            </div>

            <div>
              <label class="block text-primary-200 font-medium mb-2">
                Rating From <span class="text-red-400">*</span>
              </label>
              <input
                v-model="formData.Rating_from"
                type="text"
                required
                class="w-full bg-white/10 border border-primary-600 rounded-lg px-4 py-2 text-white placeholder-primary-400 focus:outline-none focus:ring-2 focus:ring-primary-500"
                placeholder="e.g., Hold"
              />
            </div>

            <div>
              <label class="block text-primary-200 font-medium mb-2">
                Rating To <span class="text-red-400">*</span>
              </label>
              <input
                v-model="formData.Rating_to"
                type="text"
                required
                class="w-full bg-white/10 border border-primary-600 rounded-lg px-4 py-2 text-white placeholder-primary-400 focus:outline-none focus:ring-2 focus:ring-primary-500"
                placeholder="e.g., Buy"
              />
            </div>

            <div class="md:col-span-2">
              <label class="block text-primary-200 font-medium mb-2">
                Time <span class="text-red-400">*</span>
              </label>
              <input
                v-model="formData.Time"
                type="datetime-local"
                required
                class="w-full bg-white/10 border border-primary-600 rounded-lg px-4 py-2 text-white placeholder-primary-400 focus:outline-none focus:ring-2 focus:ring-primary-500"
              />
            </div>
          </div>

          <div v-if="error" class="bg-red-500/20 border border-red-500 rounded-lg p-4">
            <p class="text-red-300">{{ error }}</p>
          </div>

          <div class="flex items-center justify-end gap-4 pt-4">
            <router-link
              to="/products"
              class="px-6 py-2 text-primary-200 hover:text-white transition-colors"
            >
              Cancel
            </router-link>
            <button
              type="submit"
              :disabled="loading"
              class="bg-primary-500 hover:bg-primary-600 disabled:opacity-50 disabled:cursor-not-allowed text-white font-medium px-6 py-2 rounded-lg transition-colors"
            >
              {{ loading ? 'Updating...' : 'Update Product' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useProductStore } from '@/stores/productStore'
import type { ProductFormData } from '@/types/product'

const router = useRouter()
const route = useRoute()
const productStore = useProductStore()
const { loading, error, selectedProduct, fetchProductById, updateProduct } = productStore

const productId = computed(() => String(route.params.id))
const product = computed(() => selectedProduct)

const handleOnChange = ((element:any) => {
  console.log("HandleOnChange ","name: ", element.target.name, " Value ", element.target.value)
  let opposite_field = element.target.name == "target_from" ? "target_to" : "target_from"
  console.log("opposite_field", opposite_field)
  console.log("Target field value ", document.getElementsByName(opposite_field)[0])
})

const formData = ref<ProductFormData>({
  Ticker: '',
  Target_from: '',
  Target_to: '',
  Company: '',
  Action: '',
  Brokerage: '',
  Rating_from: '',
  Rating_to: '',
  Time: new Date().toISOString().slice(0, 16),
})

const loadProduct = async () => {
  try {
    const product = await fetchProductById(productId.value)
    if (product) {
      formData.value = {
        Ticker: product.Ticker,
        Target_from: product.Target_from,
        Target_to: product.Target_to,
        Company: product.Company,
        Action: product.Action,
        Brokerage: product.Brokerage,
        Rating_from: product.Rating_from,
        Rating_to: product.Rating_to,
        Time: product.Time ? new Date(product.Time).toISOString().slice(0, 16) : new Date().toISOString().slice(0, 16),
      }
    }
  } catch (err) {
    console.error('Error loading product:', err)
  }
}

const handleSubmit = async () => {
  try {
    await updateProduct(productId.value, formData.value)
    router.push('/products')
  } catch (err) {
    console.error('Error updating product:', err)
  }
}

onMounted(() => {
  productStore.clearError()
  loadProduct()
})
</script>

<style scoped></style>

