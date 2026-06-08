<script setup lang="ts">
import { fetchCollectionFeeds, fetchCollections } from '@/api/collections'
import { fetchFeedItems } from '@/api/feeds'
import CollectionsPanel from '@/components/collectionsPanel.vue'
import FeedsPanel from '@/components/feedsPanel.vue'
import ItemsPanel from '@/components/itemsPanel.vue'
import type { Collection } from '@/types/collection'
import type { Feed } from '@/types/feed'
import type { Item } from '@/types/item'
import { computed, onMounted, ref, type Ref } from 'vue'

const loading: Ref<boolean> = ref<boolean>(false)

const collections: Ref<Collection[]> = ref<Collection[]>([])
const feeds: Ref<Feed[]> = ref<Feed[]>([])
const items: Ref<Item[]> = ref<Item[]>([])

const selectedCollectionID: Ref<number | null> = ref<number | null>(null)
const selectedFeedID: Ref<number | null> = ref<number | null>(null)

const selectedCollection = computed(
  () => collections.value.find((c) => c.id === selectedCollectionID.value) ?? null,
)
const selectedFeed = computed(() => feeds.value.find((f) => f.id === selectedFeedID.value) ?? null)

async function updateCollections() {
  loading.value = true

  try {
    collections.value = await fetchCollections()
  } catch (error) {
    console.error('Axios error:', error)
  } finally {
    loading.value = false
  }
}

async function collectionSelectedTrigger(collection: Collection) {
  loading.value = true
  if (!collection) {
    feeds.value = []
    items.value = []
    return
  }

  selectedCollectionID.value = collection.id
  selectedFeedID.value = null
  try {
    feeds.value = await fetchCollectionFeeds(collection.id)
  } catch (error) {
    console.error('Axios error:', error)
  } finally {
    loading.value = false
  }
}

async function updateFeeds() {
  loading.value = true

  items.value = []

  try {
    feeds.value = await fetchCollectionFeeds(selectedCollection.value!.id)
  } catch (error) {
    console.error('Axios error:', error)
  } finally {
    loading.value = false
  }
}

async function feedSelectedTrigger(f: Feed) {
  loading.value = true

  selectedFeedID.value = f.id
  try {
    items.value = await fetchFeedItems(f)
  } catch (error) {
    console.error('Axios error:', error)
  } finally {
    loading.value = false
  }
}

onMounted(async () => {
  await updateCollections()
})
</script>

<template>
  <div class="page">
    <section class="panel left">
      <CollectionsPanel
        :collections="collections"
        @collections-changed="updateCollections"
        @collection-selected="collectionSelectedTrigger"
      />
    </section>

    <section class="panel middle">
      <FeedsPanel
        v-if="selectedCollection"
        :collection="selectedCollection"
        :collections="collections"
        :feeds="feeds"
        @feeds-changed="updateFeeds"
        @feed-selected="feedSelectedTrigger"
      />
    </section>

    <section class="panel right">
      <ItemsPanel v-if="selectedFeed" :feed="selectedFeed" :items="items" />
    </section>
  </div>
</template>

<style scoped>
.page {
  display: flex;
  flex-direction: column;
}

.panel {
  width: 100vw;

  color: white;

  font-size: 2rem;
  font-family: system-ui, sans-serif;
}

@media (min-width: 1024px) {
  .page {
    flex-direction: row;

    min-height: 100vh;
    max-height: 100vh;
  }

  .panel {
    width: calc(100vw / 3);
  }
}

.left {
  background: #1e293b;
}

.middle {
  background: #334155;
}

.right {
  background: #475569;
}
</style>
