import Konva from "konva";
import { AbstractNode } from "./AbstractNode";

export class LineNode extends AbstractNode {
  type = "Line";

  static get commonAttrs() {
    return {
      x: 10,
      y: 10,
      points: [0, 0, 100, 0],
      stroke: "black",
      strokeWidth: 2,
    };
  }

  static create(designer) {
    let text = new Konva.Line({
      ...this.commonAttrs,
      draggable: true,
      resize: true,
      draggable: true,
      className: "Line",
    });

    let node = new LineNode(text, designer);
    node.tr.visible(false);
    node.draw(designer.layer);

    return node;
  }

  updateAttrs(attrs) {
    this.node.setAttrs({
      x: this.designer.transformUnitToPx(attrs.x),
      y: this.designer.transformUnitToPx(attrs.y),
      points: attrs.points ? attrs.points : [0, 0, this.designer.transformUnitToPx(attrs.width), 0],
      stroke: attrs.stroke,
    });
  }

  getNodeAttrs() {
    let attrs = this.node.getAttrs();
    attrs.width = this.designer.transformUnit(attrs.points[2] - attrs.points[0]);
    return attrs;
  }

  get columns() {
    return [
      {
        label: "Width",
        prop: "width",
        type: "inputNumber",
      },
      {
        label: "X",
        prop: "x",
        type: "inputNumber",
      },
      {
        label: "Y",
        prop: "y",
        type: "inputNumber",
      },
      {
        label: "ScaleX",
        prop: "scaleX",
        type: "inputNumber",
      },
      {
        label: "ScaleY",
        prop: "scaleY",
        type: "inputNumber",
      },
      {
        label: "Rotate",
        prop: "rotation",
        type: "inputNumber",
      },
      {
        label: "颜色",
        prop: "stroke",
        type: "color",
      },
      {
        label: "宽度",
        prop: "strokeWidth",
        type: "inputNumber",
      },
    ];
  }
}
