<script setup lang="ts"></script>

<template>
  <div class="page">
    <section class="panel left">
      <h1>Left</h1>
    </section>

    <section class="panel middle">
      <h1>Middle</h1>
    </section>

    <section class="panel right">
      <h1>Right</h1>
    </section>
  </div>
</template>

<style scoped>
.page {
  display: flex;

  min-height: 0;
  height: 100%;

  width: 100%;
}

.panel {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;

  box-sizing: border-box;

  color: white;
  font-size: 2rem;
  font-family: Arial, sans-serif;

  margin: 0;
  min-width: 0;
}

.left {
  background: #1e293b;
}

.middle {
  background: #334155;
}

.right {
  background: #475569;
}
</style>
