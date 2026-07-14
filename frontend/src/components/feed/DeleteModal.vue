<script setup lang="ts">
import type { Feed } from '@/types/feed'
import { toRef } from 'vue'

const props = defineProps<{ feed: Feed }>()
const feed = toRef(props, 'feed')

const emits = defineEmits<{
  close: []
  save: [Feed]
}>()

function close() {
  emits('close')
}
function save() {
  emits('save', feed.value!)
  close()
}
</script>

<template>
  <div class="modal-overlay" @click="close">
    <div class="modal" @click.stop>
      <h3>Delete Feed</h3>
      <p class="ellipsis">
        Are you sure you want to delete <strong>{{ feed!.title }}</strong
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
