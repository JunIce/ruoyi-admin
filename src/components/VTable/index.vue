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
  --el-table-header-text-color: #D4EBF8;

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
}
</style>
