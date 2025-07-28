<template>
  <form @submit.prevent="submit">
    <input v-model="name" placeholder="ФИО (3 слова)" />
    <input v-model="position" placeholder="Должность" />
    <select v-model="departmentID">
      <option v-for="dept in departments" :value="dept.ID" :key="dept.ID">
        {{ dept.Name }}
      </option>
    </select>
    <button>Добавить</button>
    <p v-if="error" style="color: red">{{ error }}</p>
  </form>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const name = ref("")
const position = ref("")
const departmentID = ref("")
const departments = ref([])
const error = ref("")

onMounted(async () => {
  const res = await axios.get("/api/departments")
  if (res.data.success) {
    departments.value = res.data.data
  } else {
    error.value = res.data.error || "Не удалось загрузить отделы"
  }
})

const submit = async () => {
  error.value = ""
  const parts = name.value.trim().split(/\s+/)
  if (parts.length !== 3) {
    error.value = "ФИО должно содержать 3 слова"
    return
  }
  const trimmedPosition = position.value.trim()

  if (trimmedPosition.length === 0) {
    error.value = "Должность не может быть пустой"
    return
  } else if (trimmedPosition.length < 3) {
    error.value = "Должность слишком короткая, минимум 3 символа"
    return
  } 
  if (!departmentID.value || !departments.value.find(d => d.ID === departmentID.value)) {
    error.value = "Пожалуйста, выберите отдел"
    return
  }

  const res = await axios.post("/api/employees", {
    Name: name.value.trim(),
    Position: position.value,
    DepartmentID: parseInt(departmentID.value)
  })

  if (res.data.success) {
    alert("Пользователь был добавлен")
    name.value = ""
    position.value = ""
  } else {
    error.value = res.data.error || "Ошибка при добавлении сотрудника"
  }
}
</script>
