<template>
  <div class="page-search">
    <el-form :model="modelValue" ref="queryForm" size="default" :inline="true" :labelWidth="labelWidth">
      <slot name="default"></slot>
      <template v-if="advancedSearchVisible">
        <slot name="options"></slot>
      </template>
      <el-form-item>
        <el-button type="primary" icon="Search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="Refresh" size="mini" @click="resetQuery">重置</el-button>
        <el-button type="primary" link @click="advancedSearchVisible = !advancedSearchVisible">
          <el-icon>
            <ArrowDownBold v-if="!advancedSearchVisible" />
            <ArrowUpBold v-else />
          </el-icon>
          <span>{{ !advancedSearchVisible ? '高级筛选' : '收起筛选' }}</span>
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  name: 'VPageSearch',
  emits: ['search', 'reset'],
  props: {
    labelWidth: {
      type: String,
      default: '70'
    },
    modelValue: {
      type: Object,
      default: () => ({})
    }
  },
  data() {
    return {
      advancedSearchVisible: false
    }
  },
  mounted() {
    // console.log(this.$slots.options());
  },
  methods: {
    handleCloseAdvancedSearch() {
      this.advancedSearchVisible = false
    },
    handleQuery() {
      this.$emit('search')
    },
    resetQuery() {
      this.$refs.queryForm.resetFields()
      this.$emit('reset')
    }
  }
}
</script>
<style lang="scss" scoped>
.page-search {
  width: 100%;
  position: relative;

  .advanced-search {
    position: absolute;
    right: 0;
    top: 0;

    &.is-link:not(.is-disabled):hover {
      color: var(--el-color-primary);
    }
  }

  .advanced-search-content {
    min-height: 110px;
    padding-bottom: 50px;

    .advanced-search-footer {
      padding: 10px;
      width: 100%;
      text-align: center;
      position: absolute;
      bottom: 0;
      left: 0;
    }

    :deep(.el-form-item) {
      margin-right: 20px;
    }
  }
}
</style>
