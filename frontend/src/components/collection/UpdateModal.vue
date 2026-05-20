<script setup lang="ts">
import type { Collection } from "@/types/collection";
import { ref, type Ref } from "vue";

const props = defineProps<{ collection: Collection }>();
const emits = defineEmits<{
  close: [];
  save: [Collection];
}>();

const localName: Ref<string> = ref(props.collection.name);

function close() {
  emits("close");
}
function save() {
  emits("save", {
    ...props.collection,
    name: localName.value,
  });
  close();
}
</script>

<template>
  <div class="modal-overlay" @click="close">
    <div class="modal" @click.stop>
      <h2>Edit Collection</h2>
      <input v-model="localName" placeholder="Collection Name" />
      <div class="actions">
        <button @click="save" :disabled="localName.trim() === ''">Save</button>
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

  background: green;
  border-radius: 8px;
}

.actions {
  margin-top: 16px;

  display: flex;
  justify-content: flex-end;
  gap: 8px;
}
</style>
