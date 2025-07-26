<template>
  <div>
    <input v-model="newDept" placeholder="Название отдела" />
    <button @click="addDepartment">➕ Добавить</button>

    <ul>
      <li v-for="dept in departments" :key="dept.ID">
        <!-- Название отдела -->
        <span v-if="!dept.editing">{{ dept.Name }}</span>
        <input v-else v-model="dept.editName" />

        <!-- Кнопки -->
        <template v-if="!dept.editing">
          <button @click="startEdit(dept)"> ✏️</button>
          <button @click="deleteDepartment(dept.ID)"> 🗑</button>
        </template>
        <template v-else>
          <button @click="saveDepartment(dept)"> 💾</button>
          <button @click="cancelEdit(dept)"> ✖</button>
        </template>
      </li>
    </ul>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const departments = ref([])
const newDept = ref("")

const fetchDepartments = async () => {
  const res = await axios.get("/api/departments")
  if (res.data.success) {
    departments.value = res.data.data.map(d => ({
      ...d,
      editing: false,
      editName: d.Name
    }))
  } else {
    alert(res.data.error || "Ошибка загрузки отделов")
  }
}

const addDepartment = async () => {
  if (!newDept.value.trim()) return
  const res = await axios.post("/api/departments", { Name: newDept.value })
  if (res.data.success) {
    newDept.value = ""
    fetchDepartments()
  } else {
    alert(res.data.error || "Ошибка при добавлении отдела")
  }
}

const deleteDepartment = async id => {
  if (confirm("Удалить отдел?")) {
    const res = await axios.delete(`/api/departments/${id}`)
    if (res.data.success || res.status === 204) {
      fetchDepartments()
    } else {
      alert(res.data.error || "Ошибка при удалении отдела")
    }
  }
}

const startEdit = (dept) => {
  dept.editing = true
}
const cancelEdit = (dept) => {
  dept.editing = false
  dept.editName = dept.Name
}
const saveDepartment = async (dept) => {
  const res = await axios.put(`/api/departments/${dept.ID}`, {
    Name: dept.editName
  })
  if (res.data.success) {
    dept.Name = dept.editName
    dept.editing = false
    fetchDepartments()
  } else {
    alert(res.data.error || "Ошибка при сохранении")
  }
}

onMounted(fetchDepartments)
</script>
