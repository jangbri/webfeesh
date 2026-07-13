<script setup lang="ts">
import type { Collection } from '@/types/collection'
import type { Feed } from '@/types/feed'
import { ref, toRef, type Ref } from 'vue'

const props = defineProps<{ collection: Collection; collections: Collection[] }>()
const collections = toRef(props, 'collections')

const title: Ref<string> = ref<string>('')
const link: Ref<string> = ref<string>('')
const selectedCollection: Ref<Collection | null> = ref<Collection | null>(props.collection)

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
      <h3>Add New Feed</h3>
      <input class="modal-textbox" v-model="title" placeholder="Feed Name" />
      <input class="modal-textbox" v-model="link" placeholder="https://example.com/rss" />
      <select class="modal-textbox" v-model="selectedCollection">
        <option v-for="collection in collections" :key="collection.id" :value="collection">
          {{ collection.name }}
        </option>
      </select>
      <div class="actions">
        <button
          class="btn save"
          @click="save"
          :disabled="title.trim() === '' || link.trim() === ''"
        >
          Save
        </button>
        <button class="btn cancel" @click="close">Cancel</button>
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

  background: antiquewhite;
  border-radius: 8px;
}

.modal-textbox {
  width: 100%;

  padding: 4px;
}

p {
  font-size: 12px;
}

.actions {
  margin-top: 16px;

  display: flex;
  justify-content: flex-end;
  gap: 8px;
}

.btn {
  padding: 8px 16px;

  border: 1px solid gray;
  border-radius: 8px;

  font-size: 14px;
  cursor: pointer;

  transition: background 0.15s;
}

.btn:disabled {
  background: lightslategray;
  color: black;
}

.save {
  background-color: black;
  color: white;
}

.cancel {
  background: transparent;
  color: black;
}

.cancel:hover {
  background-color: lightslategray;
  color: white;
}
</style>
