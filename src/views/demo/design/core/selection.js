export default class Selection {
  constructor(designer) {
    this.designer = designer;
    this.activeNode = null;
  }

  select(node) {
    if (this.activeNode) {
      this.unselect();
    }
    this.activeNode = node;
    node.select();
  }

  unselect() {
    let node = this.activeNode;
    this.activeNode = null;
    node.unselect();
  }

  clearSelection() {
    this.activeNode = null;
    this.designer.nodes.forEach((node) => {
      if (node.selected) {
        node.unselect();
      }
    });
    this.designer.emit('clear:selection');
  }

  get selectedNode() {
    return this.activeNode;
  }
}
