<template>
  <div class="min-h-screen bg-gradient-to-br from-primary-900 via-primary-800 to-primary-700 p-8">
    <div class="max-w-7xl mx-auto">
      <div class="mb-8">
        <h1 class="text-4xl font-bold text-white mb-2">Dashboard</h1>
        <p class="text-primary-200">Overview of your financial products</p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <!-- Stats Card -->
        <div class="bg-white/10 backdrop-blur-lg rounded-lg p-6 border border-primary-600">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-primary-200 text-sm font-medium">Total Products</p>
              <p class="text-3xl font-bold text-white mt-2">{{ productCount }}</p>
            </div>
            <div class="bg-primary-500/20 p-3 rounded-lg">
              <svg class="w-8 h-8 text-primary-300" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Quick Actions Card -->
        <div class="bg-white/10 backdrop-blur-lg rounded-lg p-6 border border-primary-600">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-primary-200 text-sm font-medium">Quick Actions</p>
              <div class="mt-4 space-y-2">
                <router-link
                  to="/products/create"
                  class="block w-full bg-primary-500 hover:bg-primary-600 text-white font-medium py-2 px-4 rounded-lg transition-colors text-center"
                >
                  Create Product
                </router-link>
                <router-link
                  to="/products"
                  class="block w-full bg-primary-400/20 hover:bg-primary-400/30 text-white font-medium py-2 px-4 rounded-lg transition-colors text-center border border-primary-400"
                >
                  View All Products
                </router-link>
              </div>
            </div>
          </div>
        </div>

        <!-- Status Card -->
        <div class="bg-white/10 backdrop-blur-lg rounded-lg p-6 border border-primary-600">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-primary-200 text-sm font-medium">System Status</p>
              <div class="mt-4 flex items-center">
                <div class="w-3 h-3 bg-green-400 rounded-full mr-2 animate-pulse"></div>
                <p class="text-white font-medium">All Systems Operational</p>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Recent Products -->
      <div class="bg-white/10 backdrop-blur-lg rounded-lg p-6 border border-primary-600">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-2xl font-bold text-white">Recent Products</h2>
          <router-link
            to="/products"
            class="text-primary-300 hover:text-primary-200 font-medium text-sm"
          >
            View All →
          </router-link>
        </div>

        <div v-if="loading" class="text-center py-8">
          <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-white"></div>
          <p class="text-primary-200 mt-2">Loading products...</p>
        </div>

        <div v-else-if="error" class="text-center py-8">
          <p class="text-red-400">{{ error }}</p>
          <button
            @click="fetchProducts"
            class="mt-4 bg-primary-500 hover:bg-primary-600 text-white px-4 py-2 rounded-lg"
          >
            Retry
          </button>
        </div>

        <div v-else-if="recentProducts.length === 0" class="text-center py-8">
          <p class="text-primary-200">No products yet. Create your first product!</p>
        </div>

        <div v-else class="overflow-x-auto">
          <table class="w-full text-left">
            <thead>
              <tr class="border-b border-primary-600">
                <th class="pb-3 text-primary-200 font-medium">Ticker</th>
                <th class="pb-3 text-primary-200 font-medium">Company</th>
                <th class="pb-3 text-primary-200 font-medium">Action</th>
                <th class="pb-3 text-primary-200 font-medium">Rating</th>
                <th class="pb-3 text-primary-200 font-medium">Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="product in recentProducts"
                :key="product.ID"
                class="border-b border-primary-700/50 hover:bg-white/5 transition-colors"
              >
                <td class="py-3 text-white font-semibold">{{ product.Ticker }}</td>
                <td class="py-3 text-primary-100">{{ product.Company }}</td>
                <td class="py-3">
                  <span class="bg-primary-500/20 text-primary-200 px-2 py-1 rounded text-sm">
                    {{ product.Action }}
                  </span>
                </td>
                <td class="py-3 text-primary-100">{{ product.Rating_from }} → {{ product.Rating_to }}</td>
                <td class="py-3">
                  <router-link
                    :to="`/products/edit/${product.ID}`"
                    class="text-primary-300 hover:text-primary-200 mr-3"
                  >
                    Edit
                  </router-link>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { useProductStore } from '@/stores/productStore'

const productStore = useProductStore()
const { products, loading, error, productCount, fetchProducts } = productStore

const recentProducts = computed(() => {
  return [...products].slice(0, 5)
})

onMounted(() => {
  fetchProducts()
})
</script>

<style scoped></style>

