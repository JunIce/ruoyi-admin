<template>
    <div class="page-table-wrapper h-full" ref="pageTableWrapper">
        <div ref="pageSearchWrapperRef">
            <slot name="header"></slot>
        </div>
        <slot name="default" :height="pageTableWrapperHeight"></slot>
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
    beforeUnmounted() {
        this.unobserveHeight()
    },
    methods: {
        init() {
            this.getPageTableWrapperHeight()
        },
        getPageTableWrapperHeight() {
            this.pageTableWrapperHeight = this.$refs.pageTableWrapper?.clientHeight - (this.$refs.pageSearchWrapperRef?.clientHeight || 0)
        },
        events() {
            this.$bus.$on('update:pageTableHeight', this.getPageTableWrapperHeight)
        },
        observeHeight() {
            let parentElement = this.$refs.pageTableWrapper
            let searchElement = this.$refs.pageSearchWrapperRef
            const resizeObserver = new ResizeObserver(throttle((entries) => {
                for (let entry of entries) {
                    if (entry.target === parentElement) {
                        this.getPageTableWrapperHeight()
                    }
                    if (entry.target === searchElement) {
                        this.getPageTableWrapperHeight()
                    }
                }
            }, 66));
            if (parentElement) {
                resizeObserver.observe(parentElement);
            }
            if (searchElement) {
                resizeObserver.observe(searchElement);
            }
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
