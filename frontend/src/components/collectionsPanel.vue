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

const emit = defineEmits<{
  (e: 'collection-selected', c: Collection): void
  (e: 'collections-changed'): void
}>()

async function selectCollection(collection: Collection) {
  emit('collection-selected', collection)
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

  emit('collections-changed')
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

  emit('collections-changed')
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

  emit('collections-changed')
}
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
        class="collection"
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

<style scoped></style>
