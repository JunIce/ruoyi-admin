
import VTable from '@/components/VTable'
import VTableColumn from '@/components/VTableColumn'

export default function installComponents(app) {
  app.component("VTable", VTable);
  app.component("VTableColumn", VTableColumn);
}
