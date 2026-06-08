<script setup lang="ts"></script>

<template>
  <div class="layout">
    <main class="content">
      <router-view />
    </main>
  </div>
</template>

<style>
* {
  box-sizing: border-box;
}

html,
body,
#app {
  margin: 0;
  height: 100%;
  width: 100%;
  overflow: hidden;
}

.layout {
  display: flex;
  height: 100vh;
  width: 100%;
}

.content {
  width: 100%;

  background: #f8fafc;
  padding: 0.25rem;
}
</style>
