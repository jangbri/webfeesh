<script setup lang="ts">
import { fetchCollectionFeeds, fetchCollections } from "@/api/collections";
import CollectionsPanel from "@/components/collectionsPanel.vue";
import FeedsPanel from "@/components/feedsPanel.vue";
import type { Collection } from "@/types/collection";
import type { Feed } from "@/types/feed";
import { onMounted, ref, type Ref } from "vue";

const loading: Ref<boolean> = ref<boolean>(false);

const collections: Ref<Collection[]> = ref<Collection[]>([]);
const feeds: Ref<Feed[]> = ref<Feed[]>([]);

const selectedCollection: Ref<Collection | null> = ref<Collection | null>(null);
const selectedFeed: Ref<Feed | null> = ref<Feed | null>(null);

async function updateCollections() {
  loading.value = true;

  try {
    collections.value = await fetchCollections();
  } catch (error) {
    console.error("Axios error:", error);
  } finally {
    loading.value = false;
  }
}

async function collectionSelectedTrigger(collection: Collection) {
  loading.value = true;
  if (!collection) {
    feeds.value = [];
    return;
  }

  selectedCollection.value = collection;
  try {
    feeds.value = await fetchCollectionFeeds(collection.id);
  } catch (error) {
    console.error("Axios error:", error);
  } finally {
    loading.value = false;
  }
}

async function updateFeeds() {
  loading.value = true;

  try {
    feeds.value = await fetchCollectionFeeds(selectedCollection.value!.id);
  } catch (error) {
    console.error("Axios error:", error);
  } finally {
    loading.value = false;
  }
}

async function feedSelectedTrigger(f: Feed) {
  loading.value = true;

  try {
    selectedFeed.value = f;
    // TODO: retrieve items associated with this feed
  } catch (error) {
    console.error("Axios error:", error);
  } finally {
    loading.value = false;
  }
}

onMounted(async () => {
  await updateCollections();
});
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
        :collections="collections"
        :feeds="feeds"
        @feeds-changed="updateFeeds"
        @feed-selected="feedSelectedTrigger"
      />
    </section>

    <section class="panel right">
      <h1>Right</h1>
    </section>
  </div>
</template>

<style scoped>
.page {
  display: flex;

  min-height: 0;
  height: 100%;

  width: 100%;
}

.panel {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;

  box-sizing: border-box;

  color: white;
  font-size: 2rem;
  font-family: Arial, sans-serif;

  margin: 0;
  min-width: 0;
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
