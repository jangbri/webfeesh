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

.layout {
  display: flex;
}

.content {
  background: #f8fafc;
}

.ellipses {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
