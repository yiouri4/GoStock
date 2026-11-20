
<script setup>
import { ref, computed, onMounted } from 'vue'

// Sidebar navigation
const activeSection = ref('dashboard')

// Products
const products = ref([])


async function loadProducts() {
  try {
    const response = await fetch('http://localhost:8080/api/products')

    if (!response.ok) {
      throw new Error('Failed to load products')
    }

    const data = await response.json()

    products.value = data

    console.log('Products loaded from Go:', products.value)
  } catch (error) {
    console.error('Error loading products:', error)
    alert('Could not connect to the Go backend.')
  }
}

onMounted(() => {
  loadProducts()
})


// Dashboard calculations
const totalProducts = computed(() => {
  return products.value.length
})

const totalStockUnits = computed(() => {
  return products.value.reduce((total, product) => {
    return total + product.quantity
  }, 0)
})

const inventoryValue = computed(() => {
  return products.value.reduce((total, product) => {
    return total + (product.price * product.quantity)
  }, 0)
})

const lowStock = computed(() => {
  return products.value.filter(product => product.quantity < 10).length
})

// Search
const searchQuery = ref('')

const filteredProducts = computed(() => {
  return products.value.filter(product =>
    product.name.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

// Form fields
const newProductName = ref('')
const newProductPrice = ref('')
const newProductQuantity = ref('')

// Add product

async function addProduct() {
  const price = Number(newProductPrice.value)
  const quantity = Number(newProductQuantity.value)

  if (
    newProductName.value.trim() === '' ||
    newProductPrice.value === '' ||
    newProductQuantity.value === '' ||
    !Number.isFinite(price) ||
    !Number.isFinite(quantity) ||
    price < 0 ||
    quantity < 0
  ) {
    alert('Please enter a valid name, price and quantity.')
    return
  }

  const nextId = products.value.length > 0
    ? Math.max(...products.value.map(product => product.id)) + 1
    : 1

  const product = {
    id: nextId,
    name: newProductName.value.trim(),
    price: price,
    quantity: quantity
  }

  try {
    const response = await fetch('http://localhost:8080/api/products', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(product)
    })

    if (!response.ok) {
      const errorMessage = await response.text()
      throw new Error(errorMessage)
    }

    await loadProducts()

    newProductName.value = ''
    newProductPrice.value = ''
    newProductQuantity.value = ''

    alert('Product added successfully!')
  } catch (error) {
    console.error('Error adding product:', error)
    alert(error.message)
  }
}



// Edit product
const editingProductId = ref(null)

function startEditing(product) {
  editingProductId.value = product.id

  newProductName.value = product.name
  newProductPrice.value = product.price
  newProductQuantity.value = product.quantity

  activeSection.value = 'dashboard'

  document.getElementById('add-product-section')?.scrollIntoView({
    behavior: 'smooth',
    block: 'center'
  })
}

async function updateProduct() {
  const product = products.value.find(
    product => product.id === editingProductId.value
  )

  if (!product) return

  const updatedProduct = {
    id: product.id,
    name: newProductName.value.trim(),
    price: Number(newProductPrice.value),
    quantity: Number(newProductQuantity.value)
  }

  if (
    updatedProduct.name === '' ||
    newProductPrice.value === '' ||
    newProductQuantity.value === '' ||
    !Number.isFinite(updatedProduct.price) ||
    !Number.isFinite(updatedProduct.quantity) ||
    updatedProduct.price < 0 ||
    updatedProduct.quantity < 0
  ) {
    alert('Please enter valid product details.')
    return
  }

  try {
    const response = await fetch('http://localhost:8080/api/products', {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(updatedProduct)
    })

    if (!response.ok) {
      const errorMessage = await response.text()
      throw new Error(errorMessage)
    }

    await loadProducts()

    editingProductId.value = null
    newProductName.value = ''
    newProductPrice.value = ''
    newProductQuantity.value = ''

    alert('Product updated successfully!')
  } catch (error) {
    console.error('Error updating product:', error)
    alert(error.message)
  }
}

// Delete product
async function deleteProduct(id) {
  const confirmed = confirm('Are you sure you want to delete this product?')
  if (!confirmed) return

  try {
    const response = await fetch(
      `http://localhost:8080/api/products?id=${id}`,
      {
        method: 'DELETE'
      }
    )

    if (!response.ok) {
      const errorMessage = await response.text()
      throw new Error(errorMessage)
    }

    await loadProducts()

    alert('Product deleted successfully!')
  } catch (error) {
    console.error('Error deleting product:', error)
    alert(error.message)
  }
}

// Update stock
function updateStock(id, amount) {
  const product = products.value.find(
    product => product.id === id
  )

  if (!product) {
    return
  }

  if (product.quantity + amount < 0) {
    return
  }

  product.quantity += amount
}
</script>

<template>
  <div class="app-container">

    <!-- Sidebar -->
    <aside class="sidebar">
      <h2>GoStock</h2>

      <nav>
        <a
          href="#dashboard"
          :class="{ active: activeSection === 'dashboard' }"
          @click="activeSection = 'dashboard'"
        >
          Dashboard
        </a>

        <a
          href="#products"
          :class="{ active: activeSection === 'products' }"
          @click="activeSection = 'products'"
        >
          Products
        </a>

        <a
          href="#stock-management"
          :class="{ active: activeSection === 'stock-management' }"
          @click="activeSection = 'stock-management'"
        >
          Stock Management
        </a>
      </nav>
    </aside>

    <!-- Main content -->
    <main id="dashboard" class="main-content">

      <h1>Dashboard</h1>
      <p>Welcome to your inventory management system.</p>

      <!-- Dashboard cards -->
      <div class="dashboard-cards">

        <div class="card">
          <div class="card-header">
            <h3>Total Products</h3>
            <span class="card-icon blue">📦</span>
          </div>
          <p class="card-number">{{ totalProducts }}</p>
          <p class="card-description">Different products in inventory</p>
        </div>

        <div class="card">
          <div class="card-header">
            <h3>Inventory Value</h3>
            <span class="card-icon green">💰</span>
          </div>
          <p class="card-number">€{{ inventoryValue.toFixed(2) }}</p>
          <p class="card-description">Total value of current stock</p>
        </div>

        <div class="card">
          <div class="card-header">
            <h3>Low Stock</h3>
            <span class="card-icon orange">⚠️</span>
          </div>
          <p class="card-number">{{ lowStock }}</p>
          <p class="card-description">Products with fewer than 10 units</p>
        </div>

      </div>

      <!-- Add / Edit Product Form -->
      <div id="add-product-section" class="add-product-section">

        <h2>
          {{ editingProductId === null ? 'Add New Product' : 'Edit Product' }}
        </h2>

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

      <!-- Products Section -->
      <div id="products" class="products-section">

        <h2>Products</h2>

        <input
          v-model="searchQuery"
          type="text"
          placeholder="Search products..."
          class="search-input"
        />

        <!-- Stock Management -->
        <h3 id="stock-management" class="stock-management-heading">
          Stock Management
        </h3>

        <div class="table-container">
          <table>

            <thead>
              <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Price</th>
                <th>Quantity</th>
                <th>Actions</th>
              </tr>
            </thead>

            <tbody>
              <tr v-for="product in filteredProducts" :key="product.id">

                <td>{{ product.id }}</td>

                <td>{{ product.name }}</td>

                <td>€{{ product.price.toFixed(2) }}</td>

                <td>
                  <span :class="product.quantity < 10 ? 'low-stock-badge' : 'stock-badge'">
                    {{ product.quantity }}
                  </span>
                </td>

                <td>
                  <button
                    class="edit-btn"
                    @click="startEditing(product)"
                  >
                    Edit
                  </button>

                  <button
                    class="delete-btn"
                    @click="deleteProduct(product.id)"
                  >
                    Delete
                  </button>

                  <button
                    class="increase-btn"
                    @click="updateStock(product.id, 1)"
                  >
                    +1
                  </button>

                  <button
                    class="decrease-btn"
                    @click="updateStock(product.id, -1)"
                  >
                    -1
                  </button>
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

* {
  box-sizing: border-box;
}

.app-container {
  display: flex;
  min-height: 100vh;
  background: #f8fafc;
  font-family: Inter, Arial, sans-serif;
}

/* =========================
   SIDEBAR
========================= */

.sidebar {
  width: 240px;
  min-width: 240px;
  background: #102a56;
  color: white;
  padding: 28px 20px;
}

.sidebar h2 {
  font-size: 23px;
  font-weight: 700;
  margin-bottom: 40px;
  letter-spacing: -0.5px;
}

.sidebar nav {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.sidebar a {
  color: #b8c7df;
  text-decoration: none;
  padding: 12px 16px;
  border-radius: 9px;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.2s ease;
}

.sidebar a:hover {
  background: #1c3b6c;
  color: white;
}

.sidebar a.active {
  background: #2563eb;
  color: white;
}

/* =========================
   MAIN CONTENT
========================= */

.main-content {
  flex: 1;
  min-width: 0;
  padding: 36px;
  background: #f8fafc;
}

.main-content h1 {
  color: #0f172a;
  font-size: 28px;
  font-weight: 700;
  letter-spacing: -0.7px;
}

/* =========================
   DASHBOARD CARDS
========================= */

.dashboard-cards {
  display: flex;
  gap: 22px;
  margin-top: 30px;
}

.card {
  flex: 1;
  min-width: 0;
  background: white;
  padding: 26px;
  border: 1px solid #e8edf4;
  border-radius: 14px;
  box-shadow: 0 2px 8px rgba(15, 23, 42, 0.03);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.card:hover {
  transform: translateY(-3px);
  box-shadow: 0 8px 20px rgba(15, 23, 42, 0.06);
}

.card h3 {
  color: #64748b;
  font-size: 14px;
  font-weight: 500;
  margin: 0;
}

.card-number {
  color: #0f172a;
  font-size: 30px;
  font-weight: 700;
  margin-top: 14px;
  letter-spacing: -0.8px;
}

/* =========================
   PRODUCTS SECTION
========================= */

.products-section {
  margin-top: 42px;
}

.products-section h2 {
  color: #0f172a;
  font-size: 21px;
  font-weight: 650;
  margin-bottom: 20px;
}

/* =========================
   TABLE CONTAINER
========================= */

.table-container {
  background: white;
  border: 1px solid #e8edf4;
  border-radius: 14px;
  padding: 8px;
  overflow-x: auto;
  box-shadow: 0 2px 8px rgba(15, 23, 42, 0.03);
}

table {
  width: 100%;
  min-width: 650px;
  border-collapse: collapse;
  background: white;
}

th {
  background: #f8fafc;
  color: #64748b;
  font-size: 12px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  text-align: left;
  padding: 16px;
}

td {
  padding: 16px;
  border-bottom: 1px solid #f1f5f9;
  color: #334155;
  font-size: 14px;
}

tbody tr {
  transition: background-color 0.2s ease;
}

tbody tr:hover {
  background: #f8fafc;
}

tbody tr:last-child td {
  border-bottom: none;
}

/* =========================
   ADD / EDIT PRODUCT FORM
========================= */

.add-product-section {
  margin-top: 40px;
  padding: 28px;
  background: white;
  border: 1px solid #e8edf4;
  border-radius: 14px;
  box-shadow: 0 2px 8px rgba(15, 23, 42, 0.03);
}

.add-product-section h2 {
  margin-top: 0;
  margin-bottom: 24px;
  color: #0f172a;
  font-size: 21px;
}

.form-group {
  display: flex;
  flex-direction: column;
  margin-bottom: 18px;
}

.form-group label {
  margin-bottom: 8px;
  color: #334155;
  font-size: 14px;
  font-weight: 600;
}

.form-group input {
  padding: 12px 14px;
  border: 1px solid #dbe3ed;
  border-radius: 9px;
  font-size: 14px;
  color: #0f172a;
  background: white;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.form-group input:focus {
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.12);
}

.form-group input::placeholder {
  color: #94a3b8;
}

/* Form submit button */

.add-product-section button {
  padding: 11px 20px;
  background: #2563eb;
  color: white;
  border: none;
  border-radius: 9px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
}

.add-product-section button:hover {
  background: #1d4ed8;
  transform: translateY(-1px);
}

/* =========================
   PRODUCT ACTION BUTTONS
========================= */

.edit-btn,
.delete-btn,
.increase-btn,
.decrease-btn {
  border: none;
  border-radius: 8px;
  padding: 8px 12px;
  font-size: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  margin: 3px;
}

.edit-btn {
  background: #eff6ff;
  color: #2563eb;
}

.delete-btn {
  background: #fef2f2;
  color: #dc2626;
}

.increase-btn {
  background: #f0fdf4;
  color: #16a34a;
}

.decrease-btn {
  background: #fff7ed;
  color: #ea580c;
}

.edit-btn:hover {
  background: #dbeafe;
}

.delete-btn:hover {
  background: #fee2e2;
}

.increase-btn:hover {
  background: #dcfce7;
}

.decrease-btn:hover {
  background: #ffedd5;
}

.edit-btn:hover,
.delete-btn:hover,
.increase-btn:hover,
.decrease-btn:hover {
  transform: translateY(-2px);
}

/* =========================
   RESPONSIVE DESIGN
========================= */

@media (max-width: 900px) {
  .sidebar {
    width: 190px;
    min-width: 190px;
    padding: 22px 14px;
  }

  .main-content {
    padding: 24px;
  }

  .dashboard-cards {
    flex-wrap: wrap;
  }

  .card {
    min-width: 200px;
  }
}

@media (max-width: 600px) {
  .app-container {
    flex-direction: column;
  }

  .sidebar {
    width: 100%;
    min-width: 0;
    padding: 18px;
  }

  .sidebar h2 {
    margin-bottom: 18px;
  }

  .sidebar nav {
    flex-direction: row;
    flex-wrap: wrap;
  }

  .main-content {
    padding: 20px;
  }

  .dashboard-cards {
    flex-direction: column;
  }

  .card {
    width: 100%;
  }
}

/* Dashboard card header */
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.card-header h3 {
  margin: 0;
}

/* Dashboard icons */
.card-icon {
  width: 42px;
  height: 42px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 10px;
  font-size: 20px;
}

.card-icon.blue {
  background: #eff6ff;
}

.card-icon.green {
  background: #f0fdf4;
}

.card-icon.orange {
  background: #fff7ed;
}

/* Dashboard descriptions */
.card-description {
  color: #94a3b8;
  font-size: 13px;
  margin-top: 8px;
}



/* Add Product Form Layout */
.add-product-section {
  margin-top: 30px;
}

.add-product-section .form-group {
  max-width: 100%;
}

@media (min-width: 900px) {
  .add-product-section {
    display: flex;
    flex-wrap: wrap;
    align-items: flex-end;
    gap: 18px;
  }

  .add-product-section h2 {
    width: 100%;
  }

  .add-product-section .form-group {
    flex: 1;
    min-width: 150px;
    margin-bottom: 0;
  }

  .add-product-section button {
    height: 42px;
    margin-bottom: 0;
  }
}


/* Stock quantity badges */
.stock-badge,
.low-stock-badge {
  display: inline-block;
  min-width: 35px;
  padding: 5px 10px;
  text-align: center;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
}

.stock-badge {
  background: #f0fdf4;
  color: #16a34a;
}

.low-stock-badge {
  background: #fff7ed;
  color: #ea580c;
}


/* Products search bar */
.search-input {
  width: 100%;
  max-width: 350px;
  padding: 12px 16px;
  margin-bottom: 18px;
  border: 1px solid #dbe3ed;
  border-radius: 9px;
  background: white;
  color: #0f172a;
  font-size: 14px;
  outline: none;
  transition: border-color 0.2s, box-shadow 0.2s;
}

.search-input:focus {
  border-color: #3b82f6;
  box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.12);
}

.search-input::placeholder {
  color: #94a3b8;
}


/* Smooth navigation */
html {
  scroll-behavior: smooth;
  scroll-padding-top: 25px;
}

/* Keep sidebar visible while scrolling */
.sidebar {
  position: sticky;
  top: 0;
  height: 100vh;
  align-self: flex-start;
}

/* Stock management heading */
.stock-management-heading {
  color: #0f172a;
  font-size: 17px;
  margin: 20px 0 15px;
}

</style>
