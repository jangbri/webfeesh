<script setup lang="ts">
import type { Collection } from "@/types/collection";
import { ref, toRef, type Ref } from "vue";

const props = defineProps<{ collections: Collection[] }>();

const collections = toRef(props, "collections");

const name: Ref<string> = ref<string>("");
const emits = defineEmits<{
  close: [];
  save: [Collection];
}>();

function close() {
  emits("close");
}
function save() {
  emits("save", {
    id: -1,
    name: name.value,
  });
  close();
}
</script>

<template>
  <div class="modal-overlay" @click="close">
    <div class="modal" @click.stop>
      <h2>Update Feed</h2>
      <input v-model="name" placeholder="Feed Name" />
      <div class="actions">
        <button @click="save" :disabled="name.trim() === ''">Save</button>
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
