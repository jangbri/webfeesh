<script setup lang="ts">
import {
  createCollection,
  deleteCollection,
  fetchCollectionFeeds,
  fetchCollections,
  updateCollection,
} from "@/api/collections";
import type { Collection } from "@/types/collection";
import CollectionCreateModal from "@/components/collection/CreateModal.vue";
import CollectionDeleteModal from "@/components/collection/DeleteModal.vue";
import CollectionUpdateModal from "@/components/collection/UpdateModal.vue";
import { onMounted, ref, type Ref } from "vue";

const loading: Ref<boolean> = ref(false);
const collections: Ref<Collection[]> = ref([]);

const createCollectionRef: Ref<boolean> = ref(false);
const updateCollectionRef: Ref<Collection | null> = ref<Collection | null>(null);
const deleteCollectionRef: Ref<Collection | null> = ref<Collection | null>(null);

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

async function handleCreateRequest(created: Collection) {
  loading.value = true;

  try {
    collections.value = await createCollection(created);
  } catch (error) {
    console.error("Axios error:", error);
  } finally {
    loading.value = false;
  }
}

async function handleUpdateRequest(updated: Collection) {
  loading.value = true;

  try {
    collections.value = await updateCollection(updated.id, updated);
  } catch (error) {
    console.log("Axios error:", error);
  } finally {
    loading.value = false;
  }
}

async function handleDeleteRequest(deleted: Collection) {
  loading.value = true;

  try {
    collections.value = await deleteCollection(deleted.id);
  } catch (error) {
    console.log("Axios error:", error);
  } finally {
    loading.value = false;
  }
}

onMounted(async () => {
  retrieveCollections();
});
</script>

<template>
  <div>
    <!-- Header -->
    <div class="header">
      <h2>Collections</h2>
      <button @click="createCollectionRef = true">+ Add New Collection</button>
    </div>

    <!-- List -->
    <div class="list">
      <div
        v-for="collection in collections"
        :key="collection.id"
        class="collection"
        @click="fetchCollectionFeeds(collection.id)"
      >
        <span>{{ collection.name }}</span>
        <button class="update-btn" @click.stop="updateCollectionRef = collection">Change</button>
        <button class="delete-btn" @click.stop="deleteCollectionRef = collection">Delete</button>
      </div>
    </div>

    <!-- Modal -->
    <CollectionCreateModal
      v-if="createCollectionRef"
      @close="createCollectionRef = false"
      @save="handleCreateRequest"
    />
    <CollectionUpdateModal
      v-if="updateCollectionRef"
      :collection="updateCollectionRef"
      @close="updateCollectionRef = null"
      @save="handleUpdateRequest"
    />
    <CollectionDeleteModal
      v-if="deleteCollectionRef"
      :collection="deleteCollectionRef"
      @close="deleteCollectionRef = null"
      @save="handleDeleteRequest"
    />
  </div>
</template>

<style scoped></style>
