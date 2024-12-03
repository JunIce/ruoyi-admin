<script>
import { h, createApp } from 'vue'
import { debounce } from 'lodash-es'
export default {
    name: 'VResizer',
    data() {
        return {
            width: 0,
            renderChildrens: []
        }
    },
    methods: {
        observeParentWidth() {
            this.resizeObserver = new ResizeObserver(debounce(() => {
                this.width = this.$el.parentElement.clientWidth
                this.calculateWidth()
            }, 100)).observe(this.$el.parentElement)
        },
        calculateWidth() {
            let children = this.renderChildrens

            let totalWidth = 0

            for (let i = 0; i < children.length; i++) {
                let child = children[i]
                let el = child.el
                let rect = el.getBoundingClientRect()
                if (rect.width + totalWidth > this.width) {
                    child.props = Object.assign(child.props || {}, {
                        style: {
                            display: 'none'
                        }
                    })
                }
                totalWidth += rect.width
            }

            this.$forceUpdate()
        },
        getWidth() {
            return this.$el.parentElement.clientWidth
        }
    },
    mounted() {
        this.$nextTick(() => {
            this.renderChildrens = this.$slots.default()
            this.observeParentWidth()
        })
    },
    beforeDestroy() {
        this.resizeObserver.disconnect()
    },
    render() {
        return this.renderChildrens
    }
}
</script>
