<script setup lang="ts">
import { createFeed, deleteFeed, updateFeed } from "@/api/feeds";
import type { Feed } from "@/types/feed";
import FeedCreateModal from "@/components/feed/CreateModal.vue";
import { ref, toRef, type Ref } from "vue";
import type { Collection } from "@/types/collection";

const loading: Ref<boolean> = ref(false);

const props = defineProps<{ collections: Collection[]; feeds: Feed[] }>();

const collections = toRef(props, "collections");
const feeds = toRef(props, "feeds");

const createFeedRef: Ref<boolean> = ref(false);
const updateFeedRef: Ref<Feed | null> = ref<Feed | null>(null);
const deleteFeedRef: Ref<Feed | null> = ref<Feed | null>(null);

const emit = defineEmits<{
  (e: "feed-selected", f: Feed): void;
  (e: "feeds-changed"): void;
}>();

async function selectFeed(f: Feed) {
  emit("feed-selected", f);
}

async function handleCreateRequest(created: Feed) {
  loading.value = true;

  try {
    feeds.value = await createFeed(created);
  } catch (error) {
    console.error("Axios error:", error);
  } finally {
    loading.value = false;
  }

  emit("feeds-changed");
}

async function handleUpdateRequest(updated: Feed) {
  loading.value = true;

  try {
    feeds.value = await updateFeed(updated);
  } catch (error) {
    console.log("Axios error:", error);
  } finally {
    loading.value = false;
  }

  emit("feeds-changed");
}

async function handleDeleteRequest(deleted: Feed) {
  loading.value = true;

  try {
    feeds.value = await deleteFeed(deleted);
  } catch (error) {
    console.log("Axios error:", error);
  } finally {
    loading.value = false;
  }

  emit("feeds-changed");
}
</script>

<template>
  <div>
    <!-- Header -->
    <div class="header">
      <h2>Feeds</h2>
      <button @click="createFeedRef = true">+ Add New Feed</button>
    </div>

    <!-- List -->
    <div class="list">
      <div class="feed" v-for="feed in feeds" :key="feed.id" @click="selectFeed(feed)">
        <span>{{ feed.title }}</span>
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
      :collections="collections"
      :feed="deleteFeedRef"
      @close="deleteFeedRef = null"
      @save="handleDeleteRequest"
    />
  </div>
</template>

<style scoped></style>
