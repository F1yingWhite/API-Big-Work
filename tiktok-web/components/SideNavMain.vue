<template>
    <div
        id="SideNavMain"
        :class="route.fullPath === '/' ? 'lg:w-[310px]' : 'lg:w-[220px]'"
        class="fixed z-20 bg-white pt-[70px] h-full lg:border-r-0 border-r w-[75px] overflow-auto"    
    >
        <div class="lg:w-full w-[55px] mx-auto">
            <NuxtLink to="/">
                <MenuItem iconString="推荐" colorString="#F02C56" sizeString="30"/>
            </NuxtLink>
            <MenuItem iconString="关注" colorString="#000000" sizeString="27"/>
            <MenuItem iconString="直播" colorString="#000000" sizeString="27"/>

            <div class="border-b lg:ml-2 mt-2" />
            
            <div class="lg:block hidden text-xs text-gray-600 font-semibold pt-4 pb-2 px-2">
                推荐账号
            </div>

            <div class="lg:hidden block pt-3"/>
            
            <div class="cursor-pointer">
                <MenuItemFollow />
            </div>

            <button class="lg:block hidden text-[#F02C56] pt-1.5 pl-2 text-[13px]">
                查看所有
            </button>

            <div class="border-b lg:ml-2 mt-2" />
            
            <div class="lg:block hidden text-xs text-gray-600 font-semibold pt-4 pb-2 px-2">
                推荐账号
            </div>

            <div class="lg:hidden block pt-3"/>
            
            <div class="cursor-pointer">
                <MenuItemFollow />
            </div>

            <button class="lg:block hidden text-[#F02C56] pt-1.5 pl-2 text-[13px]">
                查看更多
            </button>

            <div class="lg:blcok hidden border-b lg:ml-2 mt-2" />

            <div class="lg:block hidden text-[11px] text-gray-500">
                <div class="pt-4 px-2">About Newsroom TikTok Shop Contact Careers ByteDance</div>
                <div class="pt-4 px-2">TikTok for Good Advertise Developers Transparency TikTok Rewards TikTok Browse TikTok Embeds</div>
                <div class="pt-4 px-2">Help Safety Terms Privacy Creator Portal Community Guidelines</div>
                <div class="pt-4 px-2">© 2023 TikTok</div>
            </div>

            <div class="pb-14"></div>
        </div>
    </div>
</template>

<script setup>
const { $generalStore, $userStore } = useNuxtApp()
const route = useRoute()
const router = useRouter()

const isLoggedIn = (fol) => {
    if (!$userStore.id) {
        $generalStore.isLoginOpen = true
        return
    }
    // setTimeout(() => router.push(`/profile/${fol.id}`), 200)
    setTimeout(() => router.push(`/profile/test`), 200)
}
</script>