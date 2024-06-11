<template>
    <MainLayout>
        <div class="pt-[80px] w-[calc(100%-90px)] max-w-[690px]">
            <div v-for="post in posts" :key="post.author">
                <PostMain v-if="post" :post="post" />
            </div>
        </div>
    </MainLayout>
</template>

<script setup>
import MainLayout from "~/layouts/MainLayout.vue";
const { $generalStore,$userStore } = useNuxtApp()
import { storeToRefs } from 'pinia';
const { posts } = storeToRefs($generalStore)

onMounted(async () => {
  try {
    if ($userStore.isAuthenticated) {
      // 登录后推荐视频
      await $generalStore.getAuthPosts()
      console.log('Posts updated:', posts)
    } else {
      // 冷启动
      await $generalStore.getNoauthPosts()
      console.log('Default Posts:', posts)
    }
  } catch (error) {
    console.error(error)
  }
})

</script>
