<template>
    <div class="video-container h-full p-8">
        <div class="flex flex-row gap-10 h-full">
            <div class="flex flex-col gap-4 flex-1">

                <div class="flex">
                    <el-input v-model="videoUrl" placeholder="请输入视频地址" />
                    <el-button type="primary" class="ml-2" @click="handleConvert" :loading="loading">转换</el-button>
                </div>

                <video ref="videoRef" class="w-full" controls></video>
            </div>

            <div class="flex flex-col w-1/3">
                <h2>Logs</h2>
                <div class="overflow-y-auto h-full p-2 b b-gray">
                    <div v-for="item in log" :key="item">
                        {{ item }}
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script setup>
import { ref } from 'vue';
import { ffmpegConvert } from '@/utils/ffmpeg';

const videoRef = ref(null);
const videoUrl = ref('http://110.40.165.31:9001/public-data/sample_640x360.mpg');
const loading = ref(false);

const log = ref([]);

const handleConvert = async () => {
    try {
        loading.value = true;
        const data = await ffmpegConvert(videoUrl.value, (msg) => {
            log.value.push(msg);
        });
        videoRef.value.src = URL.createObjectURL(new Blob([data.buffer], { type: 'video/mp4' }));
    } catch (error) {
        console.error(error);
    } finally {
        loading.value = false;
    }
};
</script>
