<template>
    <div :id="`PostMain-${post.ID}`" class="flex border-b py-6">
        <div
            @click = "$event => isLoggedIn(route.params.id)"
            class="cursor-pointer"
        >
            <img class="rounded-full max-h-[60px]" width="60" src="https://picsum.photos/id/8/300/320">
        </div>

        <div class="pl-3 w-full px-4">
            <div class="flex items-center justify-between ph-0.5">
                <button @click="isLoggedIn(route.params.id)">
                    <span class="font-bold hover:underline cursor-pointer">
                        {{ post.author }}
                    </span>

                    <span class="text-[13px] text-light text-gray-500 pl-1 cursor-pointer">
                        {{ post.author }}
                    </span>
                </button>

                <button class="border text-[15px] px-[21px] py-0.5 border-[#F02C56] text-[#F02C56] hover:bg-[#ffeef2] font-semibold rounded-md">
                    关注
                </button>
            </div>

            <div class="text-[15px] pb-0.5 break-words md:max-w-[400px] max-w-[300px]">
                简介
            </div>

            <div class="text-[14px] text-gray-500 pb-0.5">
                #fun #cool #SuperAwesome
            </div>

            <div class="text-[14px] pb-0.5 flex items-center font-semibold">
                <Icon name="mdi:music" size="17"/>
                <div class="px-1">original sound - Qige</div>
                <Icon name="mdi:heart" size="20"/>
            </div>

            <div class="mt-2.5 flex">
                <div 
                    @click = "displayPost(post)"  
                    class="relative min-h-[480px] max-h-[580px] max-w-[260px] flex items-center bg-black rounded-xl cursor-pointer"
                >
                    <video
                        v-if="`http://127.0.0.1:8888/api/movie/noauth/${post.path}`" 
                        ref="video"
                        loop
                        muted
                        class="rounded-xl object-cover mx-auto h-full" 
                        :src="`http://127.0.0.1:8888/api/movie/noauth/${post.path}`" 
                    />

                    <img 
                        class="absolute right-2 bottom-14"
                        width="90"
                        src="~/assets/images/tiktok-logo-white.png"
                    >
                </div>

                <div class="relative mr-[75px]">
                    <div class="absolute bottom-0 pl-2">
                        <div class="pb-4 text-center">
                            <button
                                @click="toggleLike(post)" 
                                class="rounded-full bg-gray-200 p-2 cursor-pointer"
                            >
                                <Icon 
                                    name="mdi:heart" 
                                    size="25"
                                    :color="isLiked ? '#F02C56' : ''"
                                />
                            </button>

                            <span class="text-xs text-gray-800 font-semibold">{{ likeCount }}</span>
                        </div>

                        <div class="pb-4 text-center">
                            <button class="rounded-full bg-gray-200 p-2 cursor-pointer">
                                <Icon name="bx:bxs-message-rounded-dots" size="25" />
                            </button>
                            <span class="text-xs text-gray-800 font-semibold">34</span>
                        </div>

                        <div class="pb-4 text-center">
                            <button class="rounded-full bg-gray-200 p-2 cursor-pointer">
                                <Icon name="ri:share-forward-fill" size="25" />
                            </button>
                            <span class="text-xs text-gray-800 font-semibold">56</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
const { $generalStore, $userStore } = useNuxtApp();
const props = defineProps(['post']);
const { post } = toRefs(props);

const router = useRouter();
const route = useRoute();

let video = ref(null);
let isLiked = ref(false);
let likeCount = ref(post.value.like);

console.log("likeCount:" + likeCount)

onMounted(() => {
  let observer = new IntersectionObserver(function(entries) {
    entries.forEach(entry => {
      if (entry.isIntersecting) {
        console.log('Element is playing ' + post.ID);
        entry.target.play();
      } else {
        console.log('Element is paused ' + post.ID);
        entry.target.pause();
      }
    });
  }, { threshold: [0.6] });

  observer.observe(video.value);

  if (isLiked.value) {
    likeCount.value++;
  }
});

onBeforeUnmount(() => {
    video.value.pause()
    video.value.currentTime = 0
    video.value.src = ''
})

const toggleLike = async (post) => {
    if (!$userStore.id) {
        $generalStore.isLoginOpen = true;
        return;
    }
    try {
        await $userStore.like(post);
        isLiked.value = !isLiked.value;
        likeCount.value += isLiked.value ? 1 : -1;
    } catch (error) {
        console.error(error);
    }
};

const isLoggedIn = (id) => {
    if (id) {
        $generalStore.isLoginOpen = true
        return
    }
    setTimeout(() => router.push(`/profile/${id}`), 200)
}

// TODO:这里的路由应该要改
const displayPost = (post) => {
    if (!$userStore.id) {
        $generalStore.isLoginOpen = true
        return
    }

    $generalStore.setBackUrl('/')
    $generalStore.selectedPost = null
    setTimeout(() => router.push(`/post/${post.id}`), 200)
}
</script>
