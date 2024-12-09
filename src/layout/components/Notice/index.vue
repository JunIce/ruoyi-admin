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
// import { Bell } from "@element-plus/icons-vue"
import List from "./List.vue"

const notifyData = [
    {
        avatar: "https://gw.alipayobjects.com/zos/rmsportal/OKJXDXrmkNshAMvwtvhu.png",
        title: "V3 Admin Vite 上线啦",
        datetime: "两年前",
        description: "一个免费开源的中后台管理系统基础解决方案，基于 Vue3、TypeScript、Element Plus、Pinia 和 Vite 等主流技术"
    },
    {
        avatar: "https://gw.alipayobjects.com/zos/rmsportal/OKJXDXrmkNshAMvwtvhu.png",
        title: "V3 Admin 上线啦",
        datetime: "三年前",
        description: "一个中后台管理系统基础解决方案，基于 Vue3、TypeScript、Element Plus 和 Pinia"
    }
]

const messageData = [
    {
        avatar: "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png",
        title: "来自楚门的世界",
        description: "如果再也不能见到你，祝你早安、午安和晚安",
        datetime: "1998-06-05"
    },
    {
        avatar: "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png",
        title: "来自大话西游",
        description: "如果非要在这份爱上加上一个期限，我希望是一万年",
        datetime: "1995-02-04"
    },
    {
        avatar: "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png",
        title: "来自龙猫",
        description: "心存善意，定能途遇天使",
        datetime: "1988-04-16"
    }
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
            this.$success(`跳转到${this.activeName}历史页面`)
        }
    }
}
</script>
<style lang="scss" scoped>
.notify {
    margin-right: 10px;
}

.notify-history {
    text-align: center;
    padding-top: 12px;
    border-top: 1px solid var(--el-border-color);
}
</style>