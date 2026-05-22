<script setup lang="ts">
import type { Feed } from "@/types/feed";
import { toRef } from "vue";

const props = defineProps<{ feed: Feed }>();
const feed = toRef(props, "feed");

const emits = defineEmits<{
  (e: "close"): void;
  (e: "save", f: Feed): void;
}>();

function close() {
  emits("close");
}
function save() {
  emits("save", feed.value!);
  close();
}
</script>

<template>
  <div class="modal-overlay" @click="close">
    <div class="modal" @click.stop>
      <h2>Delete Feed</h2>
      <p>{{ feed!.title }}</p>
      <div class="actions">
        <button @click="save">Confirm</button>
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
