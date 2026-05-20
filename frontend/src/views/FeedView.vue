<script setup lang="ts">
import { fetchCollectionFeeds } from "@/api/collections";
import CollectionsPanel from "@/components/collectionsPanel.vue";
import type { Collection } from "@/types/collection";
import type { Feed } from "@/types/feed";
import { ref, type Ref } from "vue";

const selectedCollection: Ref<Collection | null> = ref<Collection | null>(null);
const feeds: Ref<Feed[]> = ref<Feed[]>([]);

async function onSelectCollection(collection: Collection | null) {
  if (!collection) {
    feeds.value = [];
    return;
  }
  selectedCollection.value = collection;

  const response = await fetchCollectionFeeds(collection.id);
  feeds.value = response;
}
</script>

<template>
  <div class="page">
    <section class="panel left">
      <CollectionsPanel @select_collection="onSelectCollection" />
    </section>

    <section class="panel middle">
      <FeedsPanel :feeds="feeds" />
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
