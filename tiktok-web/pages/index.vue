<template>
    <MainLayout>
        <div class="pt-[80px] w-[calc(100%-90px)] max-w-[690px]">
            <div v-if="show" v-for="post in $generalStore.posts" :key="post.ID">
                <PostMain v-if="post" :post="post" />
            </div>
        </div>
    </MainLayout>
</template>

<script setup>
import MainLayout from "~/layouts/MainLayout.vue";
const { $generalStore } = useNuxtApp()

onMounted(async () => {
    try {
        // 获取所有的视频和用户
        $generalStore.getAllUsersAndPosts()
        console.log("Posts:" + $generalStore.posts)
    } catch (error) {
        console.log(error)
    }
})

</script>
