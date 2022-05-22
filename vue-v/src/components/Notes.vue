<script setup lang="ts">
import axios from "axios";
import { onMounted, Ref, ref } from "vue";
import { Post, Response } from "@/typings/response";

let posts: Ref<Post[]> = ref([]);

onMounted(async () => {
  let { data } = await axios.get<Response>("http://localhost:8000");
  posts.value = data.data;
});
</script>

<template>
  <main className="p-1">
    <h2 className="font-medium">Notes Made</h2>
    <ul>
      <li v-if="posts.length > 0" v-for="post in posts" :key="post.ID">
        #{{ post.ID }} {{ post.Body }}
      </li>
      <p v-else>No posts</p>
    </ul>
  </main>
</template>
