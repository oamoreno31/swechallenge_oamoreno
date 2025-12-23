import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '@/views/Dashboard.vue'
import ProductList from '@/views/ProductList.vue'
import CreateProduct from '@/views/CreateProduct.vue'
import UpdateProduct from '@/views/UpdateProduct.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'dashboard',
      component: Dashboard,
    },
    {
      path: '/products',
      name: 'products',
      component: ProductList,
    },
    {
      path: '/products/create',
      name: 'create-product',
      component: CreateProduct,
    },
    {
      path: '/products/edit/:id',
      name: 'edit-product',
      component: UpdateProduct,
      props: true,
    },
  ],
})

export default router

