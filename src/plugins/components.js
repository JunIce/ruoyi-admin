import VTable from "@/components/VTable";
import VTableV2 from "@/components/VTable/indexV2.vue";
import VTableColumn from "@/components/VTableColumn";
import VDialog from "@/components/VDialog";
import VResizer from "@/components/VResizer";
import VPageWrapper from "@/components/VPageWrapper";
import VPageSearch from "@/components/VPageSearch";

let components = [
  VTable,
  VTableV2,
  VTableColumn,
  VDialog,
  VResizer,
  VPageWrapper,
  VPageSearch,
];

export default function installComponents(app) {
  components.forEach((component) => {
    app.component(component.name, component);
  });
}
