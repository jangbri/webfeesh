<script setup lang="ts">
import type { Collection } from "@/types/collection";
import type { Feed } from "@/types/feed";
import { ref, toRef, type Ref } from "vue";

const props = defineProps<{
  collections: Collection[];
  feed: Feed;
}>();

const collections = toRef(props, "collections");
const feed: Ref<Feed> = ref({ ...props.feed });

const initialC: Collection = collections.value.find((c) => c.id === feed.value.collection_id)!;
const selectedCollection: Ref<Collection> = ref<Collection>(initialC);

const emits = defineEmits<{
  close: [];
  save: [Feed];
}>();

function close() {
  emits("close");
}
function save() {
  emits("save", {
    id: feed.value.id,
    collection_id: selectedCollection.value.id,
    title: feed.value.title,
    link: feed.value.link,
  });
  close();
}
</script>

<template>
  <div class="modal-overlay" @click="close">
    <div class="modal" @click.stop>
      <h2>Modify Feed</h2>
      <input v-model="feed.title" placeholder="Feed Name" />
      <input v-model="feed.link" placeholder="https://example.com/rss" />
      <select v-model="selectedCollection">
        <option v-for="collection in collections" :key="collection.id" :value="collection">
          {{ collection.name }}
        </option>
      </select>
      <div class="actions">
        <button @click="save" :disabled="feed.title.trim() === ''">Save</button>
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
