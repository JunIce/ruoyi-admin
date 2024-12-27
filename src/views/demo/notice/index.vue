<template>
  <div class="notice-demo w-full h-full flex flex-col items-center justify-center">
    <div>
      <p class="text-2xl font-bold">动态渲染Notification,并且需要手动关闭时的情况</p>
    </div>
    <el-button @click="handleFetchNotice" type="primary">测试获取未读消息</el-button>
  </div>
</template>

<script setup>
import { useNotice } from "@/composable/useNotice";
const { fetchNotice } = useNotice();

let list = [{ id: 1, content: "测试消息1", status: 0 }, { id: 2, content: "测试消息2", status: 0 }]

const handleFetchNotice = () => {
  fetchNotice(list);
}

</script>
