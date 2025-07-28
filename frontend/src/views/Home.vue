<template>
  <div>
    <input v-model="search" placeholder="Поиск по ФИО..." />
    <table>
      <thead>
        <tr><th>ID</th><th>Имя</th><th>Должность</th><th>Отдел</th><th>Действия</th></tr>
      </thead>
      <tbody>
        <tr v-for="emp in filteredEmployees" :key="emp.ID">
          <td>{{ emp.ID }}</td>


          <td>
            <span v-if="!emp.editing">{{ emp.Name }}</span>
            <input v-else v-model="emp.editName" />
          </td>


          <td>
            <span v-if="!emp.editing">{{ emp.Position }}</span>
            <input v-else v-model="emp.editPosition" />
          </td>


          <td>
            <span v-if="!emp.editing && emp.Department.IsDeleted">{{'Отдел удален' }}</span>
            <span v-else-if="!emp.editing">{{ emp.Department.Name }}</span>
            <select v-else v-model="emp.editDeptID">
              <option v-for="d in departments" :key="d.ID" :value="Number(d.ID)">{{ d.Name }}</option>
            </select>
          </td>


            <td>
            <template v-if="!emp.editing">
                <button @click="startEdit(emp)">✏️</button>
                <button @click="deleteEmployee(emp.ID)">🗑</button>
            </template>
            <template v-else>
                <button @click="saveEmployee(emp)">💾</button>
                <button @click="cancelEdit(emp)">✖</button>
            </template>
            </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'

const employees = ref([])
const departments = ref([])
const search = ref("")

const fetchData = async () => {
  const [empRes, deptRes] = await Promise.all([
    axios.get("/api/employees"),
    axios.get("/api/departments")
  ])

  if (empRes.data.success && deptRes.data.success) {
    employees.value = empRes.data.data.map(emp => ({
      ...emp,
      editing: false,
      editName: emp.Name,
      editPosition: emp.Position,
      editDeptID: Number(emp.DepartmentID)
    }))
    .sort((a, b) => a.ID - b.ID)
    departments.value = deptRes.data.data
  } else {
    console.error("Ошибка загрузки данных", empRes.data.error || deptRes.data.error)
  }

}

const startEdit = (emp) => {
  emp.editing = true
}
const cancelEdit = (emp) => {
  emp.editing = false
  emp.editName = emp.Name
  emp.editPosition = emp.Position
  emp.editDeptID = emp.DepartmentID
}
const saveEmployee = async (emp) => {
  const payload = {}

  if (emp.editName && emp.editName !== emp.Name) {
    payload.Name = emp.editName
  }
  if (emp.editPosition && emp.editPosition !== emp.Position) {
    payload.Position = emp.editPosition
  }

  if (emp.editDeptID && Number(emp.editDeptID) !== Number(emp.DepartmentID)) {
    payload.DepartmentID = Number(emp.editDeptID)
  }

  if (Object.keys(payload).length === 0) {
    console.log("Нет изменений для сохранения.")
    emp.editing = false
    return
  }
  
  try {
    console.log("Payload отправки:", payload)

    const res = await axios.put(`/api/employees/${emp.ID}`, payload)

    if (res.data.success) {
      const updated = res.data.data

      emp.Name = updated.Name
      emp.Position = updated.Position
      emp.DepartmentID = updated.DepartmentID
      emp.Department = updated.Department

      emp.editing = false
    } else {
      console.error("Ошибка сохранения:", res.data.error)
      alert("Ошибка: " + res.data.error)
    }
  } catch (err) {
    console.error("Сетевая ошибка при сохранении сотрудника:", err)

    const errorMsg = err.response?.data?.error || "Неизвестная ошибка"
    alert("Ошибка: " + errorMsg)
  }

}

const deleteEmployee = async (id) => {
   if (confirm("Удалить сотрудника?")) {
      const res = await axios.delete(`/api/employees/${id}`)
      if (res.data.success || res.status === 204) {
        fetchData()
      } else {
        console.error("Ошибка удаления:", res.data.error)
      }
  }
}

const filteredEmployees = computed(() =>
  employees.value.filter(e =>
    e.Name.toLowerCase().includes(search.value.toLowerCase())
  )
)

onMounted(fetchData)
</script>
