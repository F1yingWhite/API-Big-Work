<template>
  <MainLayout>
      <div class="pt-[80px] w-[calc(100%-90px)] max-w-[690px]">
          <div v-for="post in posts" :key="post.ID">
              <PostMain v-if="post" :post="post" />
          </div>

          <div class="pagination flex justify-center items-center mt-4 space-x-4">
              <button 
                @click="previousPage" 
                :disabled="currentPage === 1"
                class="px-4 py-2 bg-blue-400 text-white rounded disabled:bg-gray-300 disabled:text-gray-500"
              >
                上一页
              </button>
              <span class="text-gray-700">
                第 {{ currentPage }} 页
              </span>
              <button 
                @click="nextPage" 
                :disabled="currentPage === totalPages"
                class="px-4 py-2 bg-blue-400 text-white rounded disabled:bg-gray-300 disabled:text-gray-500"
              >
                下一页
              </button>
          </div>
      </div>
  </MainLayout>
</template>

<script setup>
import MainLayout from "~/layouts/MainLayout.vue";
const { $generalStore, $userStore } = useNuxtApp();
import { storeToRefs } from 'pinia';
const { posts } = storeToRefs($generalStore);

const currentPage = ref(1);
const pageSize = ref(5);

const loadPosts = async () => {
  try {
    if ($userStore.isAuthenticated) {
      // 给登录用户推荐视频
      await $generalStore.getAuthPosts(currentPage.value, pageSize.value);
      console.log('Posts updated:', posts);
    } else {
      // 冷启动
      await $generalStore.getNoauthPosts();
      console.log('Default Posts:', posts);
    }
  } catch (error) {
    console.error(error);
  }
};

const previousPage = () => {
  if (currentPage.value > 1) {
    currentPage.value--;
    loadPosts();
  }
};

const nextPage = () => {
  if (currentPage.value < 20) {
    currentPage.value++;
    loadPosts();
  }
};

onMounted(loadPosts);

watch(
  () => $userStore.isAuthenticated,
    async (isAuthenticated) => {
      // 登录时重新设置这个值
      currentPage.value = 1; 
      if (isAuthenticated) {
        await $generalStore.getAuthPosts(currentPage.value, pageSize.value);
        console.log('Posts updated:', posts);
      } else {
        await $generalStore.getNoauthPosts(currentPage.value, pageSize.value);
        console.log('Default Posts:', posts);
      }
  }
);
</script>
