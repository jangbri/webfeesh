<script setup lang="ts">
import type { Collection } from '@/types/collection'
import { ref, type Ref } from 'vue'

const name: Ref<string> = ref<string>('')
const emits = defineEmits<{
  close: []
  save: [Collection]
}>()

function close() {
  emits('close')
}
function save() {
  emits('save', {
    id: -1,
    name: name.value,
  })
  close()
}
</script>

<template>
  <div class="modal-overlay" @click="close">
    <div class="modal" @click.stop>
      <h3>Add Collection</h3>
      <input class="modal-textbox" v-model="name" placeholder="Collection Name" />
      <div class="actions">
        <button class="btn save" @click="save" :disabled="name.trim() === ''">Save</button>
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
