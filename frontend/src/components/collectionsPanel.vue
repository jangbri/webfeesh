<script setup lang="ts">
import { createCollection, deleteCollection, updateCollection } from '@/api/collections'
import type { Collection } from '@/types/collection'
import CollectionCreateModal from '@/components/collection/CreateModal.vue'
import CollectionDeleteModal from '@/components/collection/DeleteModal.vue'
import CollectionUpdateModal from '@/components/collection/UpdateModal.vue'
import { ref, toRef, type Ref } from 'vue'

const props = defineProps<{ collections: Collection[] }>()
const collections = toRef(props, 'collections')

const loading: Ref<boolean> = ref(false)

const createCollectionRef: Ref<boolean> = ref(false)
const updateCollectionRef: Ref<Collection | null> = ref<Collection | null>(null)
const deleteCollectionRef: Ref<Collection | null> = ref<Collection | null>(null)

const emits = defineEmits<{
  'collection-selected': [Collection]
  'collections-changed': []
}>()

async function selectCollection(collection: Collection) {
  emits('collection-selected', collection)
}

async function handleCreateRequest(created: Collection) {
  loading.value = true

  try {
    await createCollection(created)
  } catch (error) {
    console.error('Axios error:', error)
  } finally {
    loading.value = false
  }

  emits('collections-changed')
}

async function handleUpdateRequest(updated: Collection) {
  loading.value = true

  try {
    await updateCollection(updated)
  } catch (error) {
    console.log('Axios error:', error)
  } finally {
    loading.value = false
  }

  emits('collections-changed')
}

async function handleDeleteRequest(deleted: Collection) {
  loading.value = true

  try {
    await deleteCollection(deleted)
  } catch (error) {
    console.log('Axios error:', error)
  } finally {
    loading.value = false
  }

  emits('collections-changed')
}
</script>

<template>
  <div class="collection-panel">
    <h2 style="text-align: center">Collections</h2>
    <!-- List -->
    <div class="list">
      <div class="collection-item add-collection" @click="createCollectionRef = true">
        <p>+ New Collection</p>
      </div>
      <div
        class="collection-item"
        v-for="collection in collections"
        :key="collection.id"
        @click="selectCollection(collection)"
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

<style scoped>
.collection-panel {
  display: flex;
  flex-direction: column;

  height: 100%;
  width: 100%;
}

.list {
  display: flex;
  flex-direction: column;

  overflow-y: auto;

  width: 100%;

  justify-content: center;
  align-items: center;

  padding: 0.5rem;
}

.add-collection {
  background-color: darkgoldenrod;
}

.collection-item {
  display: flex;
  flex-direction: row;

  margin: 0.25rem;
  width: calc(100% - 0.5rem);

  border-radius: 5px;
}

.collection-item:hover {
  background-color: antiquewhite;
  cursor: pointer;
}

.collection-item span {
  flex: 1;
  min-width: 0;

  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
