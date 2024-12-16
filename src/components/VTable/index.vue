<template>
  <el-table v-bind="tableAttrs" v-on="$listeners" class="v-table">
    <slot></slot>
  </el-table>
</template>
<script>
export default {
  name: "VTable",
  inheritAttrs: false,

  computed: {
    tableAttrs() {
      return {
        ...this.$attrs,
        border: true,
        highlightCurrentRow: true,
        stripe: true,
      };
    },
  },
};
</script>
<style scoped lang="scss">
.v-table {
  --el-table-header-bg-color: var(--el-color-primary);
  --el-table-header-text-color: #d4ebf8;

  --el-table-row-hover-bg-color: #cce0f7;
  --el-table-current-row-bg-color: #c2d4ea;
  --el-fill-color-lighter: #eef5fe;
  color: #3d5973;
  --el-table-tr-bg-color: var(--current-bg-color);




  &.el-table--default {
    :deep(.el-table__cell) {
      padding: 6px 0;
    }

    :deep(.current-row) {
      color: #000;
    }
  }

  // 选中行样式
  :deep(.selected-row) {
    background-color: #86b8f0 !important;

    &.el-table__row--striped .el-table__cell {
      background-color: #86b8f0;
    }
  }

  :deep(.el-table__column-resize-proxy) {
    border-color: var(--el-color-primary);
    z-index: 1;
  }

  :deep(.el-table__header-wrapper) {
    th {
      background-color: var(--el-color-primary) !important;
      color: #fff;
    }

    .el-checkbox__input.is-indeterminate .el-checkbox__inner,
    .el-checkbox__input.is-checked .el-checkbox__inner {
      border-color: #fff;
    }
  }

  :deep(.el-table__body-wrapper) {
    .el-table__cell {
      padding: 6px 0;
    }
  }

  :deep(.el-table__header) {
    .el-table__cell {
      color: var(--el-table-header-text-color) !important;
      background-color: var(--el-table-header-bg-color) !important;
    }
  }

  :deep(.el-table__column-resize-proxy) {
    border-color: rgba(var(--el-color-primary), 0.5) !important;
  }
}
</style>
