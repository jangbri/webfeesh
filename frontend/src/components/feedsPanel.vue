<script setup lang="ts">
import { createFeed, deleteFeed, updateFeed } from '@/api/feeds'
import type { Feed } from '@/types/feed'
import FeedCreateModal from '@/components/feed/CreateModal.vue'
import FeedDeleteModal from '@/components/feed/DeleteModal.vue'
import FeedUpdateModal from '@/components/feed/UpdateModal.vue'
import { ref, toRef, type Ref } from 'vue'
import type { Collection } from '@/types/collection'

const loading: Ref<boolean> = ref(false)

const props = defineProps<{
  collection: Collection
  collections: Collection[]
  feeds: Feed[]
}>()

const collections = toRef(props, 'collections')
const feeds = toRef(props, 'feeds')

const createFeedRef: Ref<boolean> = ref(false)
const updateFeedRef: Ref<Feed | null> = ref<Feed | null>(null)
const deleteFeedRef: Ref<Feed | null> = ref<Feed | null>(null)

const emits = defineEmits<{
  'feed-selected': [Feed]
  'feeds-changed': []
}>()

async function selectFeed(f: Feed) {
  emits('feed-selected', f)
}

async function handleCreateRequest(created: Feed) {
  loading.value = true

  try {
    await createFeed(created)
  } catch (error) {
    console.error('Axios error:', error)
  } finally {
    loading.value = false
  }

  emits('feeds-changed')
}

async function handleUpdateRequest(updated: Feed) {
  loading.value = true

  try {
    await updateFeed(updated)
  } catch (error) {
    console.log('Axios error:', error)
  } finally {
    loading.value = false
  }

  emits('feeds-changed')
}

async function handleDeleteRequest(deleted: Feed) {
  loading.value = true

  try {
    await deleteFeed(deleted)
  } catch (error) {
    console.log('Axios error:', error)
  } finally {
    loading.value = false
  }

  emits('feeds-changed')
}
</script>

<template>
  <div class="feed-panel">
    <h2 class="ellipses" style="text-align: center">{{ props.collection.name }}</h2>
    <!-- List -->
    <div class="list">
      <div class="feed-item add-feed" @click="createFeedRef = true">
        <p>+ New Feed</p>
      </div>
      <div class="feed-item" v-for="feed in feeds" :key="feed.id" @click="selectFeed(feed)">
        <span class="ellipses">{{ feed.title }}</span>
        <button class="update-btn" @click.stop="updateFeedRef = feed">Change</button>
        <button class="delete-btn" @click.stop="deleteFeedRef = feed">Delete</button>
      </div>
    </div>

    <!-- Modal -->
    <FeedCreateModal
      v-if="createFeedRef"
      :collections="collections"
      @close="createFeedRef = false"
      @save="handleCreateRequest"
    />
    <FeedUpdateModal
      v-if="updateFeedRef"
      :collections="collections"
      :feed="updateFeedRef"
      @close="updateFeedRef = null"
      @save="handleUpdateRequest"
    />
    <FeedDeleteModal
      v-if="deleteFeedRef"
      :feed="deleteFeedRef"
      @close="deleteFeedRef = null"
      @save="handleDeleteRequest"
    />
  </div>
</template>

<style scoped>
.feed-panel {
  display: flex;
  flex-direction: column;
}

.list {
  display: flex;
  flex-direction: column;

  overflow-y: auto;

  padding: 0.5rem;
}

.add-feed {
  background-color: darkgoldenrod;
}

.feed-item {
  display: flex;
  flex-direction: row;

  margin: 0.25rem;
  width: calc(100% - 0.5rem);

  border-radius: 5px;
}

.feed-item:hover {
  background-color: antiquewhite;
  cursor: pointer;
}

.feed-item span {
  flex: 1;
  min-width: 0;
}
</style>
