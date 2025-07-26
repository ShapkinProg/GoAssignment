<template>
  <div class="container">
    <h1>Управление персоналом</h1>
    <nav>
      <router-link to="/">Сотрудники</router-link> |
      <router-link to="/departments">Отделы</router-link> |
      <router-link to="/employees">Добавить сотрудника</router-link>
    </nav>
    <router-view />
  </div>
</template>

