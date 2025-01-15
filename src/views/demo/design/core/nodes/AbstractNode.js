import Konva from "konva";
import { formatValueToNumber } from "../util";

const ANCHORS = [
  "top-left",
  "top-center",
  "top-right",
  "middle-left",
  "middle-right",
  "bottom-left",
  "bottom-center",
  "bottom-right",
];

export class AbstractNode extends Konva.Group {
  static get commonAttrs() {
    return {
      x: 10,
      y: 10,
      width: 30,
      height: 10,
    };
  }

  formatValueToNumber = formatValueToNumber;

  constructor(node, parent) {
    super();
    this.id = this._generateRandomString(6);
    this.node = node;

    this._trVisible = true;
    this.designer = parent;

    // 创建 Transformer
    this.tr = new Konva.Transformer({
      node: this.node, // 指定要应用 Transformer 的节点
      enabledAnchors: ANCHORS,
      visible: this._trVisible,
    });
    this.add(this.node);
    this.add(this.tr);

    this.node.on("transform", () => {
      // 获取当前变换后的尺寸
      const scaleX = this.node.scaleX();
      const scaleY = this.node.scaleY();

      this.node.scaleX(1);
      this.node.scaleY(1);

      this.node.width(this.node.width() * scaleX);
      this.node.height(this.node.height() * scaleY);

      parent.emit("update:node", this, this.getNodeAttrs());
    });

    this.node.on("dragmove", () => {
      parent.emit("update:node", this, this.getNodeAttrs());
    });

    this.on("click", (e) => {
      if (parent.selection.selectedNode !== this) {
        parent.selection.select(this);
        parent.emit("select:node", this, this.getNodeAttrs());
      } else {
        parent.selection.unselect();
      }
    });
  }

  draw(layer) {
    layer.add(this);
  }

  hideTr() {
    this._trVisible = false;
    this.tr.visible(this._trVisible);
  }

  showTr() {
    this._trVisible = true;
    this.tr.visible(this._trVisible);
  }

  select() {
    this.selected = true;
    this.node.draggable(true);
    this.showTr();
  }

  unselect() {
    this.selected = false;
    this.node.draggable(false);
    this.hideTr();
  }

  getNodeAttrs() {
    let attrs = this.node.getAttrs();
    return {
      ...attrs,
      x: this.designer.transformUnit(attrs.x?.toFixed(2)),
      y: this.designer.transformUnit(attrs.y?.toFixed(2)),
      width: this.designer.transformUnit(attrs.width?.toFixed(2) || 0),
      height: this.designer.transformUnit(attrs.height?.toFixed(2) || 0),
      scaleX: attrs.scaleX || 1,
      scaleY: attrs.scaleY || 1,
    };
  }

  updateAttrs(attrs) {
    formatValueToNumber(attrs, [
      "width",
      "height",
      "x",
      "y",
      "scaleX",
      "scaleY",
    ]);

    this.node.setAttrs({
      ...attrs,
      x: this.designer.transformUnitToPx(attrs.x),
      y: this.designer.transformUnitToPx(attrs.y),
      width: this.designer.transformUnitToPx(attrs.width || 0),
      height: this.designer.transformUnitToPx(attrs.height || 0),
    });
  }

  _generateRandomString(length) {
    const characters =
      "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    let result = "";
    for (let i = 0; i < length; i++) {
      result += characters.charAt(
        Math.floor(Math.random() * characters.length)
      );
    }
    return result;
  }
}
