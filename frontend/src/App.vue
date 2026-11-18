
<script setup>

import { ref, computed } from 'vue'

const totalProducts = ref(125)
const inventoryValue = ref(8450)
const lowStock = ref(8)
const products = ref([
  { id: 1, name: 'Laptop', price: 899.99, quantity: 5 },
  { id: 2, name: 'Mouse', price: 24.99, quantity: 25 },
  { id: 3, name: 'Keyboard', price: 59.99, quantity: 8 },
  { id: 4, name: 'Monitor', price: 249.99, quantity: 3 },
])
const searchQuery = ref('')
const filteredProducts = computed(() => {
  return products.value.filter(product =>
    product.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

const newProductName = ref('')
const newProductPrice = ref('')
const newProductQuantity = ref('')

function addProduct() {
  const product = {
    id: Date.now(),
    name: newProductName.value,
    price: Number(newProductPrice.value),
    quantity: Number(newProductQuantity.value)
  }

  products.value.push(product)

  newProductName.value = ''
  newProductPrice.value = ''
  newProductQuantity.value = ''
}


const editingProductId = ref(null)

function startEditing(product) {
  editingProductId.value = product.id

  newProductName.value = product.name
  newProductPrice.value = product.price
  newProductQuantity.value = product.quantity
}

function updateProduct() {
  const product = products.value.find(
    product => product.id === editingProductId.value
  )

  if (!product) {
    return
  }

  product.name = newProductName.value
  product.price = Number(newProductPrice.value)
  product.quantity = Number(newProductQuantity.value)

  editingProductId.value = null

  newProductName.value = ''
  newProductPrice.value = ''
  newProductQuantity.value = ''
}
</script>


<template>
  <div class="app-container">

    <!-- Sidebar -->
    <aside class="sidebar">
      <h2>GoStock</h2>

      <nav>
        <a href="#" class="active">Dashboard</a>
        <a href="#">Products</a>
        <a href="#">Stock Management</a>
      </nav>
    </aside>

    <!-- Main content -->
    <main class="main-content">
     <h1>Dashboard</h1>
<p>Welcome to your inventory management system.</p>
<div class="add-product-section">
  <h2>Add New Product</h2>

  <div class="form-group">
    <label>Product Name</label>
    <input
      v-model="newProductName"
      type="text"
      placeholder="Enter product name"
    />
  </div>

  <div class="form-group">
    <label>Price (€)</label>
    <input
      v-model="newProductPrice"
      type="number"
      placeholder="Enter price"
    />
  </div>

  <div class="form-group">
    <label>Quantity</label>
    <input
      v-model="newProductQuantity"
      type="number"
      placeholder="Enter quantity"
    />
  </div>

  <button @click="editingProductId === null ? addProduct() : updateProduct()">
    {{ editingProductId === null ? 'Add Product' : 'Update Product' }}
  </button>
</div>
<div class="dashboard-cards">

  <div class="card">
    <h3>Total Products</h3>
    <p class="card-number">{{ totalProducts }}</p>
  </div>

  <div class="card">
    <h3>Inventory Value</h3>
    <p class="card-number">€{{ inventoryValue }}</p>
  </div>

  <div class="card">
    <h3>Low Stock</h3>
    <p class="card-number">{{ lowStock }}</p>
  </div>
<div class="products-section">
  <h2>Products</h2>
  <input
  v-model="searchQuery"
  type="text"
  placeholder="Search products..."
  class="search-input"
/>

  <table>
    <thead>
      <tr>
        <th>ID</th>
        <th>Name</th>
        <th>Price</th>
        <th>Quantity</th>
      </tr>
    </thead>

    <tbody>
      <tr v-for="product in filteredProducts" :key="product.id">
        <td>{{ product.id }}</td>
        <td>{{ product.name }}</td>
        <td>€{{ product.price.toFixed(2) }}</td>
        <td>{{ product.quantity }}</td>
        <td>
        <button @click="startEditing(product)">Edit</button>
        </td>
      </tr>
    </tbody>
  </table>
  
</div>  
</div>
    </main>

  </div>
</template>

<style scoped>
.app-container {
  display: flex;
  min-height: 100vh;
}

.sidebar {
  width: 240px;
  background-color: #0e3f93;
  color: white;
  padding: 24px;
}

.main-content {
  flex: 1;
  padding: 32px;
  background-color: #f4f6f9;
}

.sidebar h2 {
  font-size: 24px;
  margin-bottom: 40px;
}

.sidebar nav {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.sidebar a {
  color: #cbd5e1;
  text-decoration: none;
  padding: 12px 16px;
  border-radius: 8px;
  transition: background-color 0.2s;
}

.sidebar a:hover {
  background-color: #193e75;
  color: white;
}

.sidebar a.active {
  background-color: #2563eb;
  color: white;
}
.dashboard-cards {
  display: flex;
  gap: 20px;
  margin-top: 30px;
}

.card {
  flex: 1;
  background-color: white;
  padding: 24px;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.card h3 {
  color: #64748b;
  font-size: 14px;
  font-weight: normal;
}

.card-number {
  color: #17243a;
  font-size: 28px;
  font-weight: bold;
  margin-top: 12px;
}
.products-section {
  margin-top: 40px;
}

.products-section h2 {
  color: #17243a;
  margin-bottom: 20px;
}

table {
  width: 100%;
  border-collapse: collapse;
  background-color: white;
  border-radius: 12px;
  overflow: hidden;
}

th,
td {
  padding: 16px;
  text-align: left;
  border-bottom: 1px solid #e2e8f0;
}

th {
  background-color: #f8fafc;
  color: #64748b;
  font-weight: 600;
}

td {
  color: #334155;
}

.add-product-section {
  margin-top: 40px;
  padding: 24px;
  background-color: white;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.add-product-section h2 {
  margin-bottom: 20px;
  color: #17243a;
}

.form-group {
  display: flex;
  flex-direction: column;
  margin-bottom: 16px;
}

.form-group label {
  margin-bottom: 8px;
  color: #334155;
  font-weight: 600;
}

.form-group input {
  padding: 12px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 14px;
}

.add-product-section button {
  padding: 12px 20px;
  background-color: #2563eb;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
}

.add-product-section button:hover {
  background-color: #1d4ed8;
}
</style>
