<template>
    <UploadError :errorType="errorType" />

    <UploadLayout>
        <div class="w-full mt-[80px] mb-[40px] bg-white shadow-lg rounded-md py-6 md:px-10 px-4">
            <div>
                <div class="text-[23px] font-semibold">上传视频</div>
                <div class="text-gray-400 mt-1">在你的账号上发布一个视频</div>
            </div>

            <div class="mt-8 md:flex gap-6">
                <label
                    v-if="!fileDisplay"
                    for="fileInput"
                    @drop.prevent="onDrop"
                    @dragover.prevent=""
                    class="
                        md:mx-0
                        mx-auto
                        mt-4
                        mb-6
                        flex
                        flex-col
                        items-center
                        justify-center
                        w-full
                        max-w-[260px]
                        h-[470px]
                        text-center
                        p-3
                        border-2
                        border-dashed
                        border-gray-300
                        rounded-lg
                        hover:bg-gray-100
                        cursor-pointer    
                    "
                >
                    <Icon name="majesticons:cloud-upload" size="40" color="#b3b3b1"/>
                    <div class="mt-4 text-[17px]">选择一个视频上传</div>
                    <div class="mt-1.5 text-gray-500 text-[13px]">或者拖放一个文件</div>
                    <div class="mt-12 text-gray-400 text-sm">MP4</div>
                    <div class="mt-2 text-gray-400 text-[13px]">最长30分钟</div>
                    <div class="mt-2 text-gray-400 text-[13px]">文件大小要低于2GB</div>
                    <div class="px-2 py-1.5 mt-8 text-white text-[15x] w-[80%] bg-[#F02C56] rounded-sm">
                        选择文件
                    </div>

                    <input 
                        ref="file"
                        type="file"
                        id="fileInput"
                        @input="onChange"
                        hidden
                        accept=".mp4"
                    >
                </label>

                <div
                    v-else
                    class="
                        md:mx-0
                        mx-auto
                        mt-4
                        md:mb-12
                        mb-16
                        flex
                        items-center
                        justify-center
                        w-full
                        max-w-[260px]
                        h-[540px]
                        text-center
                        p-3
                        rounded-2xl
                        cursor-pointer
                        relative   
                    "
                >
                    <div class="bg-black h-full w-full"/>

                    <img
                        class="absolute z-20 pointer-events-none" 
                        src="~/assets/images/mobile-case.png"
                    >
                    <img
                        class="absolute right-4 bottom-6 z-20"
                        width="90"
                        src="~/assets/images/tiktok-logo-white.png"
                    >
                    <video
                        autoplay
                        loop
                        muted
                        class="absolute rounded-xl object-cover z-10 p-[13px] w-full h-full" 
                        :src="fileDisplay"
                    />

                    <div class="absolute -bottom-12 flex items-center justify-between z-50 rounded-xl border w-full p-2 border-gray-300">
                        <div class="flex items-center truncate">
                            <Icon name="clarity:success-standard-line" size="16" class="min-w-[16px]" />
                            <div class="text-[11px] pl-1 truncate text-ellipsis">{{ fileData.name }}</div>
                        </div>
                        <button 
                            @click="clearVideo"
                            class="text-[11px] ml-2 font-semibold"
                        >
                            更改
                        </button>
                    </div>
                </div>

                <div class="mt-4 mb-6">
                    <div class="flex bg-[#F8F8F8] py-4 px-6">
                        <div>
                            <Icon class="mr-4" size="20" name="mdi:box-cutter-off" />
                        </div>

                        <div>
                            <div class="text-semibold text-[15px] mb-1.5">分割视频并剪辑</div>
                            <div class="text-semibold text-[13px] text-gray-400">
                                您可以快速将视频分成多个部分，删除冗余部分，并将景观视频变成肖像视频
                            </div>
                        </div>

                        <div class="flex justify-end max-w-[130px] w-full h-full text-center my-auto">
                            <button class="px-8 py-1.5 text-white text-[15px] bg-[#F02C56] rounded-sm">
                                编辑
                            </button>
                        </div>
                    </div>

                    <div class="mt-5">
                        <div class="flex items-center justify-bewtween">
                            <div class="mb-1 text-[15px]">标题</div>
                            <div class="text-gray-400 text-[12px]">{{ caption.length }}/150</div>
                        </div>
                        <input
                            v-model="caption"
                            maxlength="150"
                            type="text"
                            class="
                                w-full
                                border
                                p-2.5
                                rounded-md
                                focus:outline-none
                            "
                        >
                    </div>

                    <div class="flex gap-3">
                        <button
                            @click="discard" 
                            class="px-10 py-2.5 mt-8 border text-[16px] hover:bg-gray-100 rounded-sm"
                        >
                            重置
                        </button>
                        <button class="px-10 py-2.5 mt-8 border text-[16px] text-white bg-[#F02C56] rounded-sm">
                            发布
                        </button>
                    </div>
                </div>
            </div>
        </div>
    </UploadLayout>
</template>

<script setup>
import UploadLayout from "~/layouts/UploadLayout.vue"

let file = ref(null)
let fileDisplay = ref(null)
let errorType = ref(null)
let caption = ref('')
let fileData = ref(null)
let errors = ref(null)
let isUploading = ref(false)

watch (() => caption.value,(caption) => {
    if (caption.length >= 150){
        errorType.value = 'caption'
        return
    }
    errorType.value = null
})

const onChange = () =>{
    fileDisplay.value = URL.createObjectURL(file.value.files[0])
    fileData.value = file.value.files[0]
}

const onDrop = (e) =>{
    errorType.value = ''
    file.value = e.dataTransfer.files[0];
    fileData.value = e.dataTransfer.files[0];

    let extention = file.value.name.substring(file.value.name.lastIndexOf('.') + 1)

    if (extention !== 'mp4'){
        errorType.value = 'file'
        return
    }

    fileDisplay.value = URL.createObjectURL(e.dataTransfer.files[0])
}

const discard = () =>{
    file.value = null
    fileDisplay.value = null
    fileData.value = null
    caption.value = ''
}

const clearVideo = () =>{
    file.value = null
    fileDisplay.value = null
    fileData.value = null
}
</script>
