<script setup lang="ts"></script>

<template>
  <div class="layout">
    <main class="content">
      <router-view />
    </main>
  </div>
</template>

<style>
*,
*::before,
*::after {
  box-sizing: border-box;
}

html,
body,
#app {
  margin: 0;
}

.content {
  background: #f8fafc;
}

.ellipsis {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
