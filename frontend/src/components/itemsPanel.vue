<script setup lang="ts">
import type { Item } from '@/types/item'
import { formatDateTime } from '@/util/datetime'
import { toRef } from 'vue'

const props = defineProps<{ items: Item[] }>()
const items = toRef(props, 'items')
</script>
<template>
  <div>
    <div class="header">
      <h2>Items</h2>
    </div>
    <div class="list">
      <div class="item" v-for="item in items" :key="item.id">
        <a :href="item.link">{{ item.title }} | {{ formatDateTime(item.date_updated) }}</a>
      </div>
    </div>
  </div>
</template>

<style scoped>
.list {
  height: 100%;
  width: 100%;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: 1 rem;
  border: 1px solid #ddd;
  scrollbar-width: thin;
  scrollbar-color: #c1c1c1 transparent;
}

.item {
  min-height: 75px;
  max-height: 75px;
  font-size: 1.5rem;
  padding: 12px 14px;
  border-bottom: 1px solid;
  cursor: default;
  transition: background 0.2s ease;
}

.item:last-child {
  border-bottom: none;
}
</style>
