import VTable from "@/components/VTable";
import VTableColumn from "@/components/VTableColumn";
import VDialog from "@/components/VDialog";
import VResizer from "@/components/VResizer";
let components = [VTable, VTableColumn, VDialog, VResizer];

export default function installComponents(app) {
  components.forEach((component) => {
    app.component(component.name, component);
  });
}
