<script>
import { ElTableV2, ElAutoResizer } from 'element-plus';
import { defineComponent, h } from 'vue';
let methodsMap = (instance) => {
  let tableRef = instance.$refs.tableRef;
  for (let key in tableRef) {
    if (typeof tableRef[key] === 'function') {
      instance[key] = (...args) => tableRef[key].call(null, ...args);
    }
  }
};
export default defineComponent({
  name: 'VTableV2',
  props: {
    columns: {
      type: Array,
      default: () => [],
    },
    data: {
      type: Array,
      default: () => [],
    },
  },
  mounted() {
    methodsMap(this);
  },
  computed: {
    tableRef() {
      return this.$refs.tableRef;
    },
    events() {
      return {
        ...this.$listeners,
        'current-change': this.iCurrentChange,
        'selection-change': this.iHandleSelectionChange,
        'row-click': this.iRowClick,
      };
    },
  },
  render() {
    return h(
      'div',
      {
        style: {
          width: this.$attrs.width + 'px',
          height: this.$attrs.height + 'px',
        },
      },
      [
        h(
          ElAutoResizer,
          null,
          h(
            ElTableV2,
            {
              border: true,
              stripe: true,
              class: 'ht-table',
              ref: 'tableRef',
              columns: this.columns,
              headerHeight: 40,
              data: this.data,
              width: this.$attrs.width,
              height: this.$attrs.height,
              fixed: true,
              cache: 10,
            },
            this.$slots.default,
          ),
        ),
      ],
    );
  },
});
</script>
<style lang="scss" scoped>
.ht-table {
  --el-table-row-hover-bg-color: #cce0f7;
  --el-table-current-row-bg-color: #c2d4ea;
  --el-fill-color-lighter: #eef5fe;
  color: #3d5973;

  &.el-table--default {
    :deep(.el-table__cell) {
      padding: 6px 0;
    }

    :deep(.current-row) {
      color: #000;
    }
  }

  :deep(.el-table-v2__row) {
    &:nth-child(odd) {
      background-color: var(--el-fill-color-lighter);
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

  :deep(.el-table__header-wrapper),
  :deep(.el-table-v2__header-wrapper) {

    th,
    .el-table-v2__header-cell {
      background-color: var(--el-color-primary) !important;
      color: #fff;
    }

    .el-checkbox__input.is-indeterminate .el-checkbox__inner,
    .el-checkbox__input.is-checked .el-checkbox__inner {
      border-color: #fff;
    }
  }
}
</style>
