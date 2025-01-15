<template>
    <el-form :model="form" label-position="top" labelWidth="100">
        <el-form-item v-for="col in columns" :label="col.label" :prop="col.prop">
            <template v-if="col.type == 'input'">
                <el-input v-model="form[col.prop]" clearable @input="handleInputEvent"
                    :disabled="col.disabled"></el-input>
            </template>
            <template v-if="col.type == 'inputNumber'">
                <el-input v-model.number="form[col.prop]" clearable @input="handleInputEvent" type="number"
                    :disabled="col.disabled"></el-input>
            </template>
            <template v-if="col.type == 'image'">
                <div class="image-container" v-if="form.imageUrl">
                    <el-image fit="cover" :src="concatPreviewSrc(form.imageUrl)" class="avatar"
                        :preview-src-list="[concatPreviewSrc(form.imageUrl)]" />
                    <div class="image-mask" @click="handleImageMaskClick">
                        <i class="el-icon-edit"></i>
                    </div>
                </div>

                <template v-else>
                    <el-upload class="avatar-uploader" :show-file-list="false" :on-success="handleAvatarSuccess"
                        :before-upload="beforeAvatarUpload">
                        <i class="el-icon-plus avatar-uploader-icon"></i>
                    </el-upload>
                </template>
            </template>
            <template v-if="col.type == 'select'">
                <el-select v-model="form[col.prop]" :options="col.options" @change="handleSelectChange"
                    :disabled="col.disabled">
                    <el-option v-for="option in col.options" :key="option.value" :label="option.label"
                        :value="option.value"></el-option>
                </el-select>
            </template>
        </el-form-item>
    </el-form>
</template>

<script>
export default {
    name: 'DesignForm',
    model: {
        prop: 'modelValue',
        event: 'update:modelValue'
    },
    props: {
        columns: {
            type: Array,
            default: () => []
        },
        modelValue: {
            type: Object,
            default: () => { }
        },
        uploader: {
            type: Function,
            default: () => { }
        }
    },
    emits: ['update:modelValue'],
    watch: {
        modelValue: {
            handler(val) {
                this.form = val
            },
            immediate: true,
            deep: true
        }
    },
    data() {
        return {
            form: {},
        }
    },
    methods: {
        concatPreviewSrc(url) {
            return process.env.VUE_APP_BASE_API + url;
        },
        handleInputEvent() {
            this.$emit('update:modelValue', this.form)
        },

        beforeAvatarUpload(file) {
            if (!file) return false;
            let formData = new FormData();
            formData.append('imageFile', file);
            this.uploader(formData).then(res => {
                if (this.$set) {
                    this.$set(this.form, 'imageUrl', res.url);
                } else {
                    this.form.imageUrl = res.url;
                }
                this.$emit('update:modelValue', this.form)
            })
            return false;
        },
        handleAvatarSuccess() {
            // Handle file upload success logic
        },
        handleImageMaskClick() {
            console.log('handleImageMaskClick');
            let input = document.createElement('input');
            input.type = 'file';
            input.accept = 'image/*';
            input.onchange = (e) => {
                console.log(e.target.files[0]);
                this.beforeAvatarUpload(e.target.files[0]);
                input.remove();
            }
            input.click();
            document.body.appendChild(input);
        },
        handleSelectChange(val) {
            console.log(val);
        }
    }
}
</script>

<style lang="scss" scoped>
.avatar-uploader {
    width: 100px;
    height: 100px;
    border: 1px dashed #d9d9d9;
    border-radius: 4px;


    ::v-deep .el-upload {
        width: 100%;
        height: 100%;
        border-radius: 4px;
        overflow: hidden;
        display: flex;
        align-items: center;
        justify-content: center;
    }
}

.image-container {
    width: 100px;
    height: 100px;
    border: 1px dashed #d9d9d9;
    border-radius: 4px;
    overflow: hidden;
    position: relative;

    .image-mask {
        width: 100%;
        height: 100%;
        background-color: rgba(0, 0, 0, 0.5);
        color: #fff;
        position: absolute;
        top: 0;
        left: 0;
        display: flex;
        align-items: center;
        justify-content: center;
        display: none;
    }

    &:hover {
        .image-mask {
            display: flex;
        }
    }
}
</style>