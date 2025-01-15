import Konva from 'konva';
import { AbstractNode } from './AbstractNode';

export class CircleNode extends AbstractNode {
  type = 'Circle';

  static get commonAttrs() {
    return {
      x: 60,
      y: 60,
      radius: 50,
      stroke: 'black',
      strokeWidth: 1,
    };
  }

  static create(designer) {
    let circle = new Konva.Circle({
      ...this.commonAttrs,
      draggable: true,
      resize: true,
      draggable: true,
      className: "Circle",
    });

    let node = new CircleNode(circle, designer);
    node.draw(designer.layer);
    
    return node;
  }

  updateAttrs(attrs) {
    super.updateAttrs(attrs);
    this.node.radius(this.designer.transformUnitToPx(attrs.radius));
  }

  getNodeAttrs() {
    let attrs = super.getNodeAttrs();
    attrs.radius = this.designer.transformUnit(attrs.radius?.toFixed(2));
    return attrs;
  }

  get columns() {
    return [
      {
        label: 'X',
        prop: 'x',
        type: 'inputNumber',
      },
      {
        label: 'Y',
        prop: 'y',
        type: 'inputNumber',
      },
      {
        label: '半径',
        prop: 'radius',
        type: 'inputNumber',
      },
      {
        label: 'ScaleX',
        prop: 'scaleX',
        type: 'inputNumber',
      },
      {
        label: 'ScaleY',
        prop: 'scaleY',
        type: 'inputNumber',
      },
      {
        label: 'Rotate',
        prop: 'rotation',
        type: 'inputNumber',
      },
      {
        label: '颜色',
        prop: 'stroke',
        type: 'color',
      },
      {
        label: '宽度',
        prop: 'strokeWidth',
        type: 'inputNumber',
      },
    ];
  }
}
