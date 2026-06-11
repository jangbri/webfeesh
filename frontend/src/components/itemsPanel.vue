<script setup lang="ts">
import type { Feed } from '@/types/feed'
import type { Item } from '@/types/item'
import { formatDateTime } from '@/util/datetime'
import { toRef } from 'vue'

const props = defineProps<{ feed: Feed; items: Item[] }>()
const items = toRef(props, 'items')
</script>
<template>
  <div class="items-panel">
    <h2 class="ellipses" style="text-align: center">{{ feed.title }}</h2>
    <div class="list">
      <div class="item" v-for="item in items" :key="item.id">
        <a :href="item.link" class="link" target="_blank">
          <span class="title ellipses">{{ item.title }}</span>
          <span class="timestamp">Uploaded {{ formatDateTime(item.date_updated) }}</span>
        </a>
      </div>
    </div>
  </div>
</template>

<style scoped>
.items-panel {
  display: flex;
  flex-direction: column;
}

.list {
  display: flex;
  flex-direction: column;

  overflow-y: auto;

  padding: 0.5rem;
}

.item {
  display: flex;
  flex-direction: row;

  margin: 0.25rem;
  width: calc(100% - 0.5rem);

  border: 1px solid;
  border-radius: 5px;
  font-size: 1.5rem;
  padding: 0.25rem;
  cursor: default;
  transition: background 0.2s ease;
}

.item:hover {
  background-color: antiquewhite;
  cursor: pointer;
}

.item span {
  min-width: 0;
}

.link {
  width: 100%;

  display: flex;
  justify-content: space-between;

  align-items: center;
  text-decoration: none;
  color: inherit;
}

.title {
  flex: 1;
  font-size: 1rem;
  font-weight: 500;
  line-height: 1.3;
}

.timestamp {
  flex: 0 0 auto;
  font-size: 0.8rem;
  color: #888;
  margin-top: 0.25rem;
  margin-left: 1rem;
}
</style>
