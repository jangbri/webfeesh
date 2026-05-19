<script setup lang="ts"></script>

<template>
  <nav class="sidebar">
    <h2 class="logo">Webbing</h2>
    <router-link class="nav-link" to="/">Feeds</router-link>
    <router-link class="nav-link" to="/email">Email</router-link>
  </nav>
</template>

<style scoped>
.sidebar {
  width: 4vw;
  min-width: 120px;
  height: 100%;

  background: green;
  color: white;

  display: flex;
  flex-direction: column;
  padding: 1rem;

  overflow: hidden;
}

.logo {
  margin-bottom: 2rem;
}

.nav-link {
  color: white;
  text-decoration: none;
  padding: 0.75rem 1rem;
  border-radius: 8px;
  transition: background 0.2s;
}

.nav-link:hover {
  background: #334155;
}

/* Active route */
.router-link-active {
  background: #475569;
}
</style>
