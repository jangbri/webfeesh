<script setup lang="ts">
import type { Collection } from '@/types/collection'

const props = defineProps<{ collection: Collection }>()
const emits = defineEmits<{
  close: []
  save: [Collection]
}>()

function close() {
  emits('close')
}
function save() {
  emits('save', {
    ...props.collection,
  })
  close()
}
</script>

<template>
  <div class="modal-overlay" @click="close">
    <div class="modal" @click.stop>
      <h2>Delete Collection</h2>
      <p class="ellipses">{{ props.collection.name }}</p>
      <div class="actions">
        <button @click="save">Delete</button>
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

  background: red;
  border-radius: 8px;
}

.actions {
  margin-top: 16px;

  display: flex;
  justify-content: flex-end;
  gap: 8px;
}
</style>
