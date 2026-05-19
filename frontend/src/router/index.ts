import FeedView from "@/views/FeedView.vue";
import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "feeds",
      component: FeedView,
    },
  ],
});

export default router;
