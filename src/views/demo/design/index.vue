<template>
    <div class="w-full h-full">
        <div class="flex designer-wrapper">
            <div class="action-btns">
                <ul>
                    <li v-for="menu in menus" :key="menu.value" @click="handleCommand(menu.value)">{{ menu.label }}</li>
                </ul>
            </div>

            <div class="design-content" @contextmenu.prevent>
                <div ref="chart" class="design-stage"></div>
            </div>

            <div class="right-side h-full p-2">
                <el-tabs tabPosition="bottom">
                    <el-tab-pane label="属性">
                        <el-scrollbar>
                            <DesignForm v-model:modelValue="form" :uploader="uploadImage" :columns="columns"
                                @update:modelValue="handleInputEvent" />
                        </el-scrollbar>
                    </el-tab-pane>
                </el-tabs>
            </div>
        </div>
        <div class="menu-container">
            <div v-for="icon in icons" :key="icon.name" class="icon-item" @click="handleCommand(icon.command)">
                <el-tooltip :content="icon.name">
                    <img :src="icon.icon" alt="icon" />
                </el-tooltip>
            </div>
        </div>
    </div>
</template>
<script>
import Designer from './core';
import DesignForm from './components/form'
import icons from './core/icons';
import { uploadImage } from '@/api/tool/templateDesign';
import beforeLoad from './beforeLoad';

export default {
    components: {
        DesignForm
    },
    data() {
        return {
            icons,
            menus: [
                {
                    label: '文本',
                    value: 'Text'
                },
                {
                    label: '图片',
                    value: 'Image'
                },
                {
                    label: '线条',
                    value: 'Line'
                },
                {
                    label: '矩形',
                    value: 'Rect'
                },
                {
                    label: '圆形',
                    value: 'Circle'
                },
                {
                    label: '条码',
                    value: 'Barcode'
                },
                {
                    label: '二维码',
                    value: 'QrCode'
                }
            ],
            form: {},
            nodes: [],
            columns: Designer.columns,
            uploadImage,
            content: ''
        }
    },
    mounted() {
        beforeLoad.loadIcons()
        this.init()
    },
    methods: {
        init(content) {
            this.initDesigner();
            this.content = content;
        },
        initDesigner() {
            this.designer = new Designer({
                container: this.$refs.chart,
                width: 100,
                height: 50,
                unit: 'mm',
                uploadImage: uploadImage
            })

            this.designer.on('init', (data) => {
                this.form = {
                    ...data,
                }

                if (this.content) {
                    this.designer?.fromJSON(JSON.parse(this.content));
                }
            })

            this.designer.on('select:node', (node, data) => {
                this.columns = node.columns || []
                this.form = { ...data }
            })

            this.designer.on('delete:node', (node) => {
                this.columns = Designer.columns
                this.form = this.designer.getStageAttrs()
                console.log('delete:node', node);
            })

            this.designer.on('clear:selection', () => {
                this.columns = Designer.columns
                this.form = this.designer.getStageAttrs()
            })

            this.designer.on('update:node', (node, data) => {
                this.form = { ...data }
            })
        },
        handleCommand(type) {
            if (type === 'save') {
                this.handleSave()
            } else if (type === 'print') {
                this.handlePrint()
            } else {
                this.designer.command({ type })
            }
        },
        handleInputEvent(attrs) {
            this.designer.updateNodeAttrs(attrs)
        },
        getNodes() {
            return this.designer.getStageNodes();
        },
        handleSave() {
            return JSON.stringify(this.getNodes());
        },
        clear() {
            this.designer.destroy();
            this.designer = null;
        },
    }
}
</script>

<style lang="scss" scoped>
$boxshadow: 0 0 24px #ddd;

.designer-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
    height: 100%;
    width: 100%;
    box-shadow: $boxshadow;
}

.design-content {
    display: flex;
    justify-content: center;
    align-items: center;
    position: relative;
    height: 100%;
    width: calc(100% - 200px - 200px);
    padding: 20px;
    background-color: #fff;
    overflow: hidden;

    .design-stage {
        width: 100%;
        height: 100%;
    }
}


.menu-container {
    position: absolute;
    width: 350px;
    bottom: 0;
    left: 50%;
    z-index: 10;
    transform: translate(-50%, 0);
    background-color: #fff;
    padding: 10px 15px;
    border-bottom-left-radius: 4px;
    border-bottom-right-radius: 4px;
    box-shadow: $boxshadow;

    display: flex;
    justify-content: flex-end;

    .icon-item {
        width: 24px;
        height: 24px;
        cursor: pointer;
        margin-right: 10px;

        img {
            width: 100%;
            height: 100%;
        }
    }
}

.desiger-wrapper {
    height: 100%;
    align-items: center;
    justify-content: center;
    box-shadow: $boxshadow;
}

.konvajs-content {
    background-color: #fff;
    box-shadow: 0 0 4px #ddd;
}

.action-btns {
    height: 100%;
    width: 200px;
    background-color: #fff;
    padding: 20px;
    flex-shrink: 0;
    border-right: 1px solid #ddd;

    ul,
    li {
        list-style: none;
        padding: 0;
    }

    ul {
        display: flex;
        flex-wrap: wrap;
    }

    li {
        width: 70px;
        margin-bottom: 10px;
        margin-right: 8px;
        // height: 45px;
        padding: 5px 0;
        border-radius: 4px;
        // line-height: 45px;
        background-color: #f3f3f3;
        text-align: center;
        // border-bottom: 1px solid #ccc;
        cursor: pointer;

        &:hover {
            color: #fff;
            background-color: rgba($color: #999, $alpha: 0.7);
        }

        &:last-child {
            border-bottom: unset;
        }
    }
}

.right-side {
    width: 200px;
    height: 100%;

    background-color: #fff;
    box-shadow: $boxshadow;
    border-left: 1px solid #ddd;

    .el-tabs {
        height: 100%;
        display: flex;
        flex-direction: column;

        .el-tabs__content {
            // flex: 1;
            // overflow-y: auto;
            // padding-right: 10px;
            // .el-form {
            //     height: 100%;
            // }

            .el-form-item {
                margin-bottom: 2px;

                .el-form-item__label {
                    padding-bottom: 0px;
                }
            }
        }
    }
}
</style>