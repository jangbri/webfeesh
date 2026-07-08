<script setup lang="ts">
import { createFeed, deleteFeed, updateFeed } from '@/api/feeds'
import type { Feed } from '@/types/feed'
import FeedCreateModal from '@/components/feed/CreateModal.vue'
import FeedDeleteModal from '@/components/feed/DeleteModal.vue'
import FeedUpdateModal from '@/components/feed/UpdateModal.vue'
import { computed, ref, toRef, type Ref } from 'vue'
import type { Collection } from '@/types/collection'
import { apiURL } from '@/api/client'

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

const feedLink = computed(() => {
  return `${apiURL}/collections/${props.collection.id}/rss`
})

async function copyText(text: string) {
  await navigator.clipboard.writeText(text)
}

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
    <h2 class="ellipsis" style="text-align: center">{{ props.collection.name }}</h2>
    <div class="snippet">
      <pre>{{ feedLink }}</pre>
      <button @click="copyText(feedLink)">Copy</button>
    </div>
    <!-- List -->
    <div class="list">
      <div class="feed add-feed" @click="createFeedRef = true">
        <span>New Feed</span>
      </div>
      <div class="feed" v-for="feed in feeds" :key="feed.id" @click="selectFeed(feed)">
        <span class="ellipsis">{{ feed.title }}</span>
        <div class="actions">
          <button class="update-btn" @click.stop="updateFeedRef = feed">Change</button>
          <button class="delete-btn" @click.stop="deleteFeedRef = feed">Delete</button>
        </div>
      </div>
    </div>

    <!-- Modal -->
    <FeedCreateModal
      v-if="createFeedRef"
      :collection="collection"
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
  height: 100%;
}

.snippet {
  display: flex;
  justify-content: space-between;

  background: #1e1e1e;
  color: #fff;

  border-radius: 8px;
  margin: 0rem 4rem;
  padding: 16px;
}

.snippet pre {
  font-size: 12px;
  position: relative;
  margin: 0;
  white-space: pre-wrap;
  overflow-x: auto;
}

.snippet button {
  background: yellow;
}

.list {
  display: flex;
  flex-direction: column;

  overflow-y: auto;

  padding: 0.5rem;
}

.add-feed {
  background-color: darkgoldenrod;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.feed {
  display: flex;
  flex-direction: row;

  padding: 2px 8px;

  margin: 0.2rem;
  width: calc(100% - 0.5rem);

  border-radius: 5px;
}

.feed:hover {
  background-color: lightslategray;
  cursor: pointer;
}

.actions {
  display: none;
  pointer-events: none;
}

.feed:hover .actions,
.feed:focus-within .actions {
  display: flex;
  pointer-events: auto;
}

.feed span {
  flex: 1 1 100%;
  min-width: 100px;
}
</style>
