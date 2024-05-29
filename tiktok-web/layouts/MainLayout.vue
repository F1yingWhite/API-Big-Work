<template>
    <TopNav/>
    <div
        :class="route.fullPath === '/' ? 'max-w-[1140px]' : ''"
        class="flex justify-between mx-audio w-full lg:px-2.5 px-0"
    >
        <div>
            <SideNavMain />
        </div>

        <slot/>
    </div>
</template>

<script setup>
const route = useRoute()
</script>