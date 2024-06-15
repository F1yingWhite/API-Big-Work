<template>
    <div
        @mouseenter="$event => isHover(true)"
        @mouseleave="$event => isHover(false)"
        class="relative brightness-90 hover:brightness-[1.1] cursor-pointer"    
    >
        <div
            v-if="!isLoaded"
            class="absolute flex items-center justify-center top-0 left-0 aspect-[3/4] w-full object-cover rounded-md bg-black"
        >
            <Icon class="animate-spin ml-1" name="mingcute:loading-line" size="100" color="#FFFFFF"/>
        </div>
        <div>
            <video 
                ref="video"
                @click="$event => displayPost(post)"
                muted
                loop
                class="aspect-[3/4] object-cover rounded-md"
                :src="`http://127.0.0.1:8888/api/movie/noauth/${post.path}`"
            />
        </div>

        <div class="px-1">
            <div class="text-gray-700 text-[15px] pt-1 break-words">
                {{ post.title }}
            </div>
            <div class="flex items-center -ml-1 text-gray-600 font-bold text-xs">
                <Icon name="gg:loadbar-sound" size="20"/>
                {{ post.like }}%
                <Icon class="ml-1" name="tabler:alert-circle" size="16"/>
                <div class="text-[18px] font-semibold flex items-center justify-between ">
                    <Icon
                        v-if="true"
                        @click="deleteVideo(post.ID)"
                        class="cursor-pointer"
                        name="material-symbols:delete-outline-sharp"
                        size="25"
                    />
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
defineProps(['post'])
const { $generalStore } = useNuxtApp()

const route = useRoute()
const router = useRouter()

let video = ref(null)
let isLoaded = ref(false)

onMounted(() => {
    if (video.value) {
        video.value.addEventListener('loadeddata', (e) => {
            if (e.target) {
                setTimeout(() => {
                    isLoaded.value = true
                }, 200)
            }
        });
    }
})

onBeforeUnmount(() => {
    if (video.value) {
        video.value.pause()
        video.value.currentTime = 0
        video.value.src = ''
    }
})

const deleteVideo = async (id) => {
    try {
        await $generalStore.deleteVideo(id);
        console.log(`Video with ID ${id} deleted successfully.`);
        router.go(0); // 刷新页面
        // setTimeout(() => router.push(`/profile/${route.params.id}`), 300)
    } catch (error) {
        console.error('Error deleting video:', error);
    }
}

const displayPost = (post) => {
    $generalStore.setBackUrl("/profile/" + route.params.id)
    $generalStore.selectedPost = null
    setTimeout(() => router.push(`/post/${post.ID}`), 300)
}

const isHover = (bool) => {
    if (bool) {
        video.value.play()
    } else {
        video.value.pause()
    }
}

</script>
