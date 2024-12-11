<template>
    <div class="notify">
        <el-popover placement="bottom" :width="popoverWidth" trigger="click">
            <template #reference>
                <el-badge :value="badgeValue" :max="badgeMax" :hidden="badgeValue === 0">
                    <el-tooltip effect="dark" content="消息通知" placement="bottom">
                        <el-icon :size="20">
                            <Bell />
                        </el-icon>
                    </el-tooltip>
                </el-badge>
            </template>
            <template #default>
                <el-tabs v-model="activeName" class="demo-tabs" stretch>
                    <el-tab-pane v-for="(item, index) in data" :key="index" :name="item.name">
                        <template #label>
                            {{ item.name }}
                            <el-badge :value="item.list.length" :max="badgeMax" :type="item.type" />
                        </template>
                        <el-scrollbar height="400px">
                            <List :data="item.list" />
                        </el-scrollbar>
                    </el-tab-pane>
                </el-tabs>
                <div class="notify-history">
                    <el-button link @click="handleHistory">
                        查看{{ activeName }}历史
                    </el-button>
                </div>
            </template>
        </el-popover>
    </div>
</template>
<script>
import { ElMessage } from "element-plus"
import List from "./List.vue"

const notifyData = [
    {
        avatar: "https://iconfont.alicdn.com/p/illus_3d/file/ZsWruISgVCKK/6b56e558-c360-43e2-8e76-d897146c954f.png",
        title: "ruoyi-admin",
        datetime: "两年前",
        description: "一个免费开源的中后台管理系统基础解决方案，基于 Vue3、Element Plus 和 Pinia"
    },
    {
        avatar: "https://iconfont.alicdn.com/p/illus_3d/file/ZsWruISgVCKK/6b56e558-c360-43e2-8e76-d897146c954f.png",
        title: "ruoyi-admin",
        datetime: "三年前",
        description: "一个中后台管理系统基础解决方案，基于 Vue3、Element Plus 和 Pinia"
    }
]

const messageData = [
    {
        avatar: "https://iconfont.alicdn.com/p/illus_3d/file/ZsWruISgVCKK/0a782364-de90-4fc2-b832-898efbd579d5.png",
        title: "ABCD",
        description: "生活就像这个通知组件，有时会收到好消息，有时会收到坏消息，还有待办的任务。但重要的是，我们要学会像这个组件一样，把所有的信息都妥善管理起来，分类处理，及时响应。不要让通知堆积成山，也不要对重要的信息视而不见。生活中的每一个提醒，都是让我们成长的机会",
        datetime: "2024-12-11"
    },
]

const todoData = [
    {
        title: "任务名称",
        description: "这家伙很懒，什么都没留下",
        extra: "未开始",
        status: "info"
    },
    {
        title: "任务名称",
        description: "这家伙很懒，什么都没留下",
        extra: "进行中",
        status: "primary"
    },
    {
        title: "任务名称",
        description: "这家伙很懒，什么都没留下",
        extra: "已超时",
        status: "danger"
    }
]

export default {
    name: "NoticeBar",
    components: {
        List
    },
    data() {
        return {
            badgeMax: 99,
            popoverWidth: 350,
            activeName: '通知',
            data: [
                // 通知数据
                {
                    name: "通知",
                    type: "primary",
                    list: notifyData
                },
                // 消息数据
                {
                    name: "消息",
                    type: "danger",
                    list: messageData
                },
                // 待办数据
                {
                    name: "待办",
                    type: "warning",
                    list: todoData
                }
            ]
        }
    },
    computed: {
        badgeValue() {
            return this.data.reduce((sum, item) => sum + item.list.length, 0)
        }
    },
    methods: {
        handleHistory() {
            ElMessage.success(`跳转到${this.activeName}历史页面`)
            this.$router.push({
                path: "/system/notice",
                query: {
                    type: this.activeName
                }
            })
        }
    }
}
</script>
<style lang="scss" scoped>
.notify {
    padding: 0 8px;
    cursor: pointer;
}

.notify-history {
    text-align: center;
    padding-top: 12px;
    border-top: 1px solid var(--el-border-color);
}
</style>