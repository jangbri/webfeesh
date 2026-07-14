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
      <h3>Delete Collection</h3>
      <p class="ellipsis">
        Are you sure you want to delete <strong>{{ props.collection.name }}</strong
        >?
      </p>

      <div class="actions">
        <button class="btn delete" @click="save">Delete</button>
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

.ellipsis {
  font-size: 14px;
}

.modal {
  width: 400px;
  padding: 20px;

  background: antiquewhite;
  border-radius: 8px;
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

.delete {
  background-color: red;
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
