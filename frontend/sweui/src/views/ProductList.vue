<template>
  <div class="min-h-screen bg-gradient-to-br from-primary-900 via-primary-800 to-primary-700 p-8">
    <div class="max-w-7xl mx-auto">
      <div class="mb-8 flex items-center justify-between">
        <div>
          <h1 class="text-4xl font-bold text-white mb-2">Products</h1>
          <p class="text-primary-200">Manage your financial products</p>
        </div>
        <router-link
          to="/products/create"
          class="bg-primary-500 hover:bg-primary-600 text-white font-medium py-2 px-6 rounded-lg transition-colors"
        >
          + Create Product
        </router-link>
      </div>

      <div class="bg-white/10 backdrop-blur-lg rounded-lg p-6 border border-primary-600">
        <div v-if="loading" class="text-center py-12">
          <div class="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-white"></div>
          <p class="text-primary-200 mt-4">Loading products...</p>
        </div>

        <div v-else-if="error" class="text-center py-12">
          <p class="text-red-400 mb-4">{{ error }}</p>
          <button
            @click="fetchProducts"
            class="bg-primary-500 hover:bg-primary-600 text-white px-6 py-2 rounded-lg"
          >
            Retry
          </button>
        </div>

        <div v-else-if="products.length === 0" class="text-center py-12">
          <p class="text-primary-200 mb-4">No products found.</p>
          <router-link
            to="/products/create"
            class="inline-block bg-primary-500 hover:bg-primary-600 text-white px-6 py-2 rounded-lg"
          >
            Create Your First Product
          </router-link>
        </div>

        <div v-else class="overflow-x-auto">
          <table class="w-full text-left">
            <thead>
              <tr class="border-b border-primary-600">
                <th class="pb-3 text-primary-200 font-medium">ID</th>
                <th class="pb-3 text-primary-200 font-medium">Ticker</th>
                <th class="pb-3 text-primary-200 font-medium">Company</th>
                <th class="pb-3 text-primary-200 font-medium">Target</th>
                <th class="pb-3 text-primary-200 font-medium">Action</th>
                <th class="pb-3 text-primary-200 font-medium">Brokerage</th>
                <th class="pb-3 text-primary-200 font-medium">Rating</th>
                <th class="pb-3 text-primary-200 font-medium">Time</th>
                <th class="pb-3 text-primary-200 font-medium text-right">Actions</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="product in products"
                :key="String(product.ID)"
                class="border-b border-primary-700/50 hover:bg-white/5 transition-colors"
              >
                <td class="py-4 text-white font-semibold">#{{ String(product.ID) }}</td>
                <td class="py-4 text-white font-semibold">{{ product.Ticker }}</td>
                <td class="py-4 text-primary-100">{{ product.Company }}</td>
                <td class="py-4 text-primary-100">
                  {{ product.Target_from }} - {{ product.Target_to }}
                </td>
                <td class="py-4">
                  <span class="bg-primary-500/20 text-primary-200 px-2 py-1 rounded text-sm">
                    {{ product.Action }}
                  </span>
                </td>
                <td class="py-4 text-primary-100">{{ product.Brokerage }}</td>
                <td class="py-4 text-primary-100">
                  {{ product.Rating_from }} → {{ product.Rating_to }}
                </td>
                <td class="py-4 text-primary-200 text-sm">{{ formatTime(product.Time) }}</td>
                <td class="py-4 text-right">
                  <div class="flex items-center justify-end gap-2">
                    <router-link
                      :to="`/products/edit/${product.ID}`"
                      class="text-primary-300 hover:text-primary-200 font-medium"
                    >
                      Edit
                    </router-link>
                    <button
                      @click="handleDelete(product.ID)"
                      :disabled="deletingId === product.ID"
                      class="text-red-400 hover:text-red-300 font-medium disabled:opacity-50"
                    >
                      {{ deletingId === product.ID ? 'Deleting...' : 'Delete' }}
                    </button>
                  </div>
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
import { onMounted, ref } from 'vue'
import { useProductStore } from '@/stores/productStore'
import type { Product } from '@/types/product'

const productStore = useProductStore()
const { products, loading, error, fetchProducts, deleteProduct } = productStore
const deletingId = ref<number | null>(null)

const handleDelete = async (id: number) => {
  if (!confirm('Are you sure you want to delete this product?')) {
    return
  }

  deletingId.value = id
  try {
    await deleteProduct(id)
  } catch (err) {
    console.error('Error deleting product:', err)
  } finally {
    deletingId.value = null
  }
}

const formatTime = (time: string) => {
  try {
    return new Date(time).toLocaleDateString()
  } catch {
    return time
  }
}

onMounted(() => {
  fetchProducts()
})
</script>

<style scoped></style>

