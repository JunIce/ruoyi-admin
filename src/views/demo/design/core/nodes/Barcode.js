import Konva from 'konva';
import { AbstractNode } from './AbstractNode';

export class BarcodeNode extends AbstractNode {
  type = 'Barcode';

  static get commonAttrs() {
    return {
      x: 10,
      y: 10,
      width: 80,
      height: 80,
      image: '',
    };
  }

  static create(designer, image) {
    let rect = new Konva.Image({
      ...this.commonAttrs,
      className: 'Barcode',
      draggable: true,
      resize: true,
      draggable: true,
      image: image,
    });


    let node = new BarcodeNode(rect, designer);
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
        label: 'Height',
        prop: 'height',
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
        label: '绑定属性',
        prop: 'bindProp',
        type: 'input',
      },
    ];
  }
}
