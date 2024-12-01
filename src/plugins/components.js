import VTable from "@/components/VTable";
import VTableColumn from "@/components/VTableColumn";
import VDialog from "@/components/VDialog";

export default function installComponents(app) {
  app.component("VTable", VTable);
  app.component("VTableColumn", VTableColumn);
  app.component("VDialog", VDialog);
}
