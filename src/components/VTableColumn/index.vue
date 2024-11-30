<script>
import { defineComponent, h } from "vue";
import { ElTableColumn } from "element-plus";

export default defineComponent({
  name: "VTableColumn",
  props: {
    columns: {
      type: Array,
      default: () => [],
    },
  },
  methods: {
    renderColumns() {
      return this.columns.map((column) => this.getColumnAttrs(column));
    },
    getColumnAttrs(column) {
      let slot = null;
      if (this.$slots[column.prop]) {
        slot = this.$slots[column.prop];
      }

      return h(ElTableColumn, {
        ...column,
        showOverflowTooltip: true,
      }, {
        default: slot,
      });
    },
  },
  render() {
    return this.renderColumns();
  },
});
</script>
<style scoped lang="scss">
.v-table {
  --el-table-header-bg-color: var(--el-color-primary);
  --el-table-header-text-color: #d4ebf8;

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
