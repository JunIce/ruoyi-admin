import Konva from 'konva';
import { AbstractNode } from './AbstractNode';

export class RectNode extends AbstractNode {
  type = 'Rect';

  static get commonAttrs() {
    return {
      x: 10,
      y: 10,
      width: 100,
      height: 100,
      stroke: 'black',
      strokeWidth: 1,
    };
  }

  static create(designer) {
    let rect = new Konva.Rect({
      ...this.commonAttrs,
      draggable: true,
      resize: true,
      draggable: true,
      className: "Rect",
    });

    let node = new RectNode(rect, designer);
    node.draw(designer.layer);
    
    return node;
  }

  get columns() {
    return [
      {
        label: 'Width',
        prop: 'width',
        type: 'inputNumber',
      },
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
