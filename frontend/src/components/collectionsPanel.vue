<script setup lang="ts">
import { fetchCollections } from "@/api/collections";
import type { Collection } from "@/types/collection";
import { onMounted, ref, type Ref } from "vue";

const loading: Ref<boolean> = ref(false);
const collections: Ref<Collection[]> = ref([]);

async function retrieveCollections() {
  loading.value = true;

  try {
    collections.value = await fetchCollections();
  } catch (error) {
    console.error("Axios error:", error);
  } finally {
    loading.value = false;
  }
}

onMounted(async () => {
  retrieveCollections();
  console.log(collections.value);
});
</script>

<template>
  <div>
    <!-- Header -->
    <div class="header">
      <h2>Collections</h2>
      <button @click="">+ Add New Collection</button>
    </div>

    <!-- List -->
    <div class="list">
      <div
        v-for="collection in collections"
        :key="collection.id"
        class="collection"
        @click="retrieveCollections()"
      >
        <span>{{ collection.name }}</span>
        <button class="edit-btn" @click.stop="">Edit</button>
      </div>
    </div>

    <!-- Modal -->
  </div>
</template>

<style scoped></style>
