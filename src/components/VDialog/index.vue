<template>
    <el-dialog modal-class="v-dialog" v-bind="dialogAttrs" v-on="$listeners">
        <template #header>
            <div class="v-dialog-header">
                <template v-if="$slots.header">
                    <slot name="header"></slot>
                </template>
                <div v-else class="v-dialog-header-inner">
                    <span class="v-dialog-header-title">{{ $attrs.title }}</span>
                    <div class="v-dialog-header-actions">
                        <slot name="header-actions"></slot>

                        <svg-icon v-if="!fullscreen" icon-class="fullscreen1" class="v-dialog-header-actions-icon"
                            @click="openFullscreen" />

                        <svg-icon v-else icon-class="exit-fullscreen1" class="v-dialog-header-actions-icon"
                            @click="exitFullscreen" />

                        <svg-icon icon-class="close" class="v-dialog-header-actions-icon" @click="close" />
                    </div>
                </div>
            </div>
        </template>

        <div class="v-dialog-body">
            <slot></slot>
        </div>
        <template #footer>
            <div v-if="$slots.footer" class="v-dialog-footer">
                <slot name="footer"></slot>
            </div>
        </template>
    </el-dialog>
</template>
<script>
export default {
    name: 'HtDialog',
    data() {
        return {
            fullscreen: this.fullscreen,
            visible: this.visible
        }
    },
    computed: {
        dialogAttrs() {
            return {
                appendToBody: true,
                draggable: true,
                ...this.$attrs,
                title: null,
                showClose: false,
                fullscreen: this.fullscreen,
                modelValue: this.visible
            }
        }
    },
    methods: {
        close() {
            this.visible = false;
            this.$emit('update:modelValue', false);
        },
        openFullscreen() {
            this.fullscreen = true;
        },
        exitFullscreen() {
            this.fullscreen = false;
        }
    }
}
</script>
<style lang="scss">
.v-dialog {
    --el-dialog-padding-primary: 0;

    .el-dialog {
        padding: 0;
        display: flex;
        flex-direction: column;
    }

    .el-dialog__header {
        padding-bottom: 0;
    }

    .el-dialog__footer {
        padding-top: 0;
    }

    .el-dialog__body {
        max-height: 80vh;
        overflow-y: auto;
        flex: 1;
    }

    .is-fullscreen {
        .el-dialog__body {
            max-height: unset;
        }
    }

    .v-dialog-header {
        padding: 10px;
        border-bottom: 1px solid #e5e5e5;

        &-inner {
            display: flex;
            align-items: center;
            justify-content: space-between;
            height: 25px;
        }

        &-title {
            font-size: 16px;
            font-weight: 500;
            color: #333;
        }

        &-actions {
            display: flex;
            align-items: center;
            gap: 5px;
            font-size: 20px;
        }

        &-actions-icon {
            cursor: pointer;
            color: #666;
            padding: 4px;
            width: 30px;
            height: 30px;

            &:hover {
                background-color: #f0f0f0;
                border-radius: 50px;
            }
        }
    }

    .v-dialog-body {
        padding: 15px;
        overflow-y: auto;
    }

    .v-dialog-footer {
        padding: 10px;
        border-top: 1px solid #e5e5e5;
    }
}
</style>