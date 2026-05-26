<script setup lang="ts">
import type { Collection } from '@/types/collection'
import type { Feed } from '@/types/feed'
import { ref, toRef, type Ref } from 'vue'

const title: Ref<string> = ref<string>('')
const link: Ref<string> = ref<string>('')
const selectedCollection: Ref<Collection | null> = ref<Collection | null>(null)

const props = defineProps<{ collections: Collection[] }>()

const collections = toRef(props, 'collections')

const emits = defineEmits<{
  close: []
  save: [Feed]
}>()

function close() {
  emits('close')
}
function save() {
  emits('save', {
    id: -1,
    collection_id: selectedCollection.value!.id,
    title: title.value,
    link: link.value,
  })
  close()
}
</script>

<template>
  <div class="modal-overlay" @click="close">
    <div class="modal" @click.stop>
      <h2>Add New Feed</h2>
      <input v-model="title" placeholder="Feed Name" />
      <input v-model="link" placeholder="https://example.com/rss" />
      <select v-model="selectedCollection">
        <option v-for="collection in collections" :key="collection.id" :value="collection">
          {{ collection.name }}
        </option>
      </select>
      <div class="actions">
        <button @click="save" :disabled="title.trim() === ''">Save</button>
        <button @click="close">Cancel</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.modal-overlay {
  position: fixed;
  inset: 0;

  background: rgba(0, 0, 0, 0.5);

  display: flex;
  align-items: center;
  justify-content: center;
}

.modal {
  width: 400px;
  padding: 20px;

  background: blue;
  border-radius: 8px;
}

.actions {
  margin-top: 16px;

  display: flex;
  justify-content: flex-end;
  gap: 8px;
}
</style>
