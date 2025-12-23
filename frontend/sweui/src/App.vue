<template>
  <div id="app">
    <nav class="bg-primary-900 border-b border-primary-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex">
            <div class="flex-shrink-0 flex items-center">
              <router-link to="/" class="text-2xl font-bold text-white">
                SWE Challenge
              </router-link>
            </div>
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <router-link
                to="/"
                class="border-transparent text-primary-200 hover:border-primary-500 hover:text-white inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                active-class="border-primary-500 text-white"
              >
                Dashboard
              </router-link>
              <router-link
                to="/products"
                class="border-transparent text-primary-200 hover:border-primary-500 hover:text-white inline-flex items-center px-1 pt-1 border-b-2 text-sm font-medium"
                active-class="border-primary-500 text-white"
              >
                Products
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <router-view />
  </div>
</template>

<script setup lang="ts">
</script>

<style>
@import './assets/main.css';
</style>
