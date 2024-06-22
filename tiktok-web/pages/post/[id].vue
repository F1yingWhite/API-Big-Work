<template>
    <div
        id="PostPage"
        class="fixed lg:flex justify-between z-50 top-0 left-0 w-full h-full bg-black lg:overflow-hidden overflow-auto"
    >
        <div v-if="$generalStore.selectedPost" class="lg:w-[calc(100%-540px)] h-full relative">
            <NuxtLink
                :href="$generalStore.isBackUrl"
                class="absolute z-20 m-5 rounded-full bg-gray-700 p-1.5 hover:bg-gray-800"
            >
                <Icon name="material-symbols:close" color="#FFFFFF" size="27"/>
            </NuxtLink>

            <div v-if=true>
                <button
                    :disabled="!isLoaded"
                    @click="loopThroughPostsUp()"
                    class="absolute z-20 right-4 top-4 flex items-center justify-center rounded-full bg-gray-700 px-1.5 hover:bg-gray-800"
                >
                    <Icon name="mdi:chevron-up" size="30" color="#FFFFFF" />
                </button>

                <button
                    :disabled="!isLoaded"
                    @click="loopThroughPostsDown()"
                    class="absolute z-20 right-4 top-20 flex items-center justify-center rounded-full bg-gray-700 px-1.5 hover:bg-gray-800"
                >
                    <Icon name="mdi:chevron-down" size="30" color="#FFFFFF" />
                </button>
            </div>

            <img 
                class="absolute top-[18px] left-70[px] rounded-full lg:mx-0 mx-auto"
                width="45" 
                src="~/assets/images/tiktok-logo-small.png"
            >

            <video 
                v-if="$generalStore.selectedPost.path"
                class="absolute object-cover w-full my-auto z-[-1] h-screen" 
                :src="`http://127.0.0.1:8888/api/movie/noauth/${$generalStore.selectedPost.path}`" 
            />

            <div
                v-if="!isLoaded"
                class="flex items-center justify-center bg-black bg-opacity-70 h-screen lg:min-w-[480px]"
            >
                <Icon class="animate-spin ml-1" name="mingcute:loading-line" size="100" color="#FFFFFF"/>
            </div>

            <div class="bg-black bg-opacity-70 xlg:min-w-[480px]">
                <video
                v-if="$generalStore.selectedPost.path"
                ref="video"
                loop
                class="h-screen mx-auto"
                />
            </div>
        </div>

        <div
            id="InfoSection"
            v-if="$generalStore.selectedPost" 
            class="lg:max-w-[550px] relative w-full h-full bg-white"
        >
            <div class="py-7" />

            <div class="flex items-center justify-between px-8">
                <div class="flex items-center">
                    <NuxtLink :href="`/profile/${$generalStore.selectedPost.author}`">
                        <img
                            class="rounded-full lg:mx-0 mx-auto"
                            width="40" 
                            src="https://picsum.photos/id/8/300/320"
                        >
                    </NuxtLink>
                    <div class="ml-3 pt-0.5">
                        <div class="text-[17px] font-semibold">
                            {{ $generalStore.allLowerCaseNoCaps($generalStore.selectedPost.author) }}
                        </div>
                        <div class="text-[13px] -mt-5 font-light">
                            {{ $generalStore.selectedPost.author }}
                            <span class="relative -top-[2px] text-[30px] pr-0.5 ">.</span>
                            <span class="font-medium">{{ $generalStore.selectedPost.CreatedAt }}</span>
                        </div>
                    </div>
                </div>

                <Icon
                    v-if="$userStore.id === $generalStore.selectedPost.author"
                    @click="$event => deletePost()"
                    class="cursor-pointer"
                    name="material-symbols:delete-outline-sharp"
                    size="25"
                />
            </div>

            <div class="px-8 mt-4 text-sm">some post text</div>

            <div class="px-8 mt-4 text-sm font-bold">
                <Icon name="mdi:music" size="17"/>
                original sound - {{ $generalStore.allLowerCaseNoCaps($generalStore.selectedPost.author) }}
            </div>

            <div class="flex items-center px-8 mt-8">
                <div class="pb-4 text-center flex items-center">
                    <button 
                        @click="toggleLike($generalStore.selectedPost)" 
                        class="rounded-full bg-gray-200 p-2 cursor-pointer"
                    >
                        <Icon
                            name="mdi:heart"
                            size="25"
                            :color="isLiked ? '#F02C56' : ''"
                        />
                    </button>
                    <span class="text-xs pl-2 pr-4 text-gray-800 font-semibold">
                        {{ likeCount }}
                    </span>
                </div>

                <div class="pb-4 text-center flex items-center">
                    <div class="rounded-full bg-gray-200 p-2 cursor-pointer">
                        <Icon
                            name="bx:bxs-message-rounded-dots"
                            size="25"
                        />
                    </div>
                    <span class="text-xs pl-2 pr-4 text-gray-800 font-semibold">
                        123
                    </span>
                </div>
            </div>

            <div
                id="Comments"
                class="bg-[#F8F8F8] z-0 w-full h-[calc(100% - 273px)] border-t-2 overflow-auto"
            >

            <div class="pt-2"/>

            <div
                v-if="true"
                class="text-center mt-6 text-xl text-gray-500"
            >
                没有评论...
            </div>

            <div
                v-else
                class="flex items-center justify-between px-8 mt-4"
            >
                <div class="flex items-center relative w-full">
                    <NuxtLink to="/">
                        <img
                            class="absolute top-0 rounded-full lg:mx-0 mx-auto"
                            width="40"
                            src="https://picsum.photos/id/8/300/320"
                        >
                    </NuxtLink>
                    <div class="ml-14 pt-0.5 w-full">
                        <div class="text-[18px] font-semibold flex items-center justify-between ">
                            用户名
                            <Icon
                                v-if="true"
                                @click="$event => deleteComment()"
                                class="cursor-pointer"
                                name="material-symbols:delete-outline-sharp"
                                size="25"
                            />
                        </div>
                        <div class="text-[15px] font-light">
                            正直之人的道路…四处被利己者的不公及恶人的暴政所笼罩，以仁慈和善良之名佑护他，
                            引领弱小穿越黑暗之谷，因他确是其兄弟的守护者及迷路孩童的发现者，
                            而我将以极大复仇及激烈怒火击倒汝等，那些意图毒害和毁灭我兄弟之人，
                            而你将知晓我的名即是天主，待我降下仇恨于汝等之时！
                        </div>
                    </div>
                </div>
            </div>

            <div class="mb-28"/>

            </div>

            <div
                id="CreateComment"
                v-if="$userStore.id"
                class="absolute flex items-center justify-between bottom-0 bg-white h-[85px] lg:max-w-[550px] w-full py-5 px-8 border-t-2"
            >
                <div
                    :class="inputFocused ? 'border-2 border-gray-400' : 'border-2 border-[#F1F1F2]'"
                    class="bg-[#F1F1F2] flex items-center rounded-lg w-full lg:max-w-[420px]"
                >
                    <input 
                        v-model="comment"
                        @focus="$event => inputFocused = true"
                        @blur="$event => inputFocused = false"
                        class="bg-[#F1F1F2] text-[14px] focus:outline-none w-full lg:max-w-[420px] p-2 rounded-lg"
                        type="text" 
                        placeholder="留下你的评论..."
                    >
                </div>
                <button
                    :disabled="!comment"
                    @click="$event => addComment()"
                    :class="comment ? 'text-[#F02C56] cursor-pointer' : 'text-gray-400'"
                    class="font-semibold text-sm ml-5 pr-1"
                >
                    评论
                </button>
            </div>
        </div>
    </div>
</template>

<script setup>
const { $generalStore, $userStore } = useNuxtApp()

const route = useRoute()
const router = useRouter()

let video = ref(null)
let isLoaded = ref(false)
let comment = ref(null)
let inputFocused = ref(false)
let isLiked = ref(false);
let likeCount = ref(0);

onMounted(async () => {
    $generalStore.selectedPost = null
    try {
        await $generalStore.getPostById(route.params.id)
        likeCount.value = $generalStore.selectedPost.like;

        const timestamp = new Date().getTime();
        const videoPath = $generalStore.selectedPost.path;
        const token = localStorage.getItem('token');
        const response = await fetch(`http://127.0.0.1:8888/api/movie/${videoPath}?timestamp=${timestamp}`, {
            headers: {
                'Authorization': token
            }
        });
        const videoBlob = await response.blob();
        video.value.src = URL.createObjectURL(videoBlob);

        video.value.addEventListener('loadeddata', () => {
            setTimeout(() => {
                isLoaded.value = true;
            }, 500);
        });
    } catch (error) {
        if (error && error.response && error.response.status === 400) {
            router.push('/');
        }
    }
});

const loopThroughPostsUp = async () => {
    try {
        if (!isLoaded.value) return;
        isLoaded.value = false;
        const previousVideo = await $generalStore.fetchPreviousVideo(route.params.id);
        if (previousVideo) {
            await router.push(`/post/${previousVideo.ID}`);
            console.log(previousVideo)
        }
    } catch (error) {
        console.error('Error fetching next video:', error);
    } finally {
        isLoaded.value = true;
    }
};

const loopThroughPostsDown = async () => {
    try {
        if (!isLoaded.value) return;
        isLoaded.value = false;
        const nextVideo = await $generalStore.fetchNextVideo(route.params.id);
        if (nextVideo) {
            await router.push(`/post/${nextVideo.ID}`);
            console.log(nextVideo)
        }
    } catch (error) {
        console.error('Error fetching previous video:', error);
    } finally {
        isLoaded.value = true;
    }
};

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


onBeforeUnmount(() => {
    if (video.value) {
        video.value.pause();
        video.value.currentTime = 0;
        video.value.src = '';
    }
});

watch(() => isLoaded.value, () =>{
    if (isLoaded.value) {
        setTimeout(() => video.value.play(), 500)
    }
})
</script>
