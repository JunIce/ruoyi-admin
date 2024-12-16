<template>
    <div class="dark-mode">
        <el-switch :model-value="darkMode" @change="handleDarkModeChange">
            <template #active-action>
                <el-icon :size="16">
                    <Moon />
                </el-icon>
            </template>
            <template #inactive-action>
                <el-icon :size="16">
                    <Sunny />
                </el-icon>
            </template>
        </el-switch>
    </div>
</template>
<script>
import { Sunny, Moon } from '@element-plus/icons-vue'
import { useSysmode } from '@/composable/useSysmode'

export default {
    components: {
        Sunny,
        Moon
    },
    setup() {
        const { darkMode, handleDarkModeChange } = useSysmode()
        return {
            darkMode,
            handleDarkModeChange
        }
    },
}
</script>
<style scoped>
.dark-mode {
    margin-right: 10px;
}
</style>
