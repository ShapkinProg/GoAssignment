import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Departments from '../views/Departments.vue'
import AddEmployee from '../views/AddEmployee.vue'

const routes = [
  { path: '/', component: Home },
  { path: '/departments', component: Departments },
  { path: '/employees', component: AddEmployee }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
