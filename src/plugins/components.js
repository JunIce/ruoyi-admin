import VTable from "@/components/VTable";
import VTableColumn from "@/components/VTableColumn";
import VDialog from "@/components/VDialog";
import VResizer from "@/components/VResizer";
import VPageWrapper from "@/components/VPageWrapper";
import VPageSearch from "@/components/VPageSearch";

let components = [
  VTable,
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
