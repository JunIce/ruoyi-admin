<template>
    <div class="page-table-wrapper h-full" ref="pageTableWrapper">
        <slot :height="pageTableWrapperHeight"></slot>
    </div>
</template>
<script>
import { throttle } from 'lodash-es'
export default {
    name: 'VPageWrapper',
    data() {
        return {
            pageTableWrapper: null,
            pageTableWrapperHeight: 0
        }
    },
    mounted() {
        // this.init()
        this.observeHeight()
    },
    beforeDestroy() {
        this.unobserveHeight()
    },
    methods: {
        init() {
            this.getPageTableWrapperHeight()
        },
        getPageTableWrapperHeight() {
            this.pageTableWrapperHeight = this.$refs.pageTableWrapper?.clientHeight
        },
        events() {
            this.$bus.$on('update:pageTableHeight', this.getPageTableWrapperHeight)
        },
        observeHeight() {
            let parentElement = this.$refs.pageTableWrapper
            const resizeObserver = new ResizeObserver(throttle((entries) => {
                for (let entry of entries) {
                    if (entry.target === parentElement) {
                        this.getPageTableWrapperHeight()
                    }
                }
            }, 66));
            parentElement && resizeObserver.observe(parentElement);
            this.resizeObserver = resizeObserver
        },
        unobserveHeight() {
            this.resizeObserver.disconnect()
        }
    }
}
</script>
<style scoped lang="scss">
.page-table-wrapper {}
</style>
